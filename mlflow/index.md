# MLflow code analysis


Let's analize mlflow source code
<!--more-->

## 1. Initialize

1. [mlflow/__init__.py](https://github.com/mlflow/mlflow/blob/master/mlflow/__init__.py)
2. [LazyLoader](https://github.com/mlflow/mlflow/blob/master/mlflow/utils/lazy_load.py#L8)

Mlflow uses so many large sized 3rd party libs, so it imports lazily when module is called. (singleton of global context)


```py
# __init__
# Lazily load mlflow flavors to avoid excessive dependencies.
catboost = LazyLoader("mlflow.catboost", globals(), "mlflow.catboost")
diviner = LazyLoader("mlflow.diviner", globals(), "mlflow.diviner")
fastai = LazyLoader("mlflow.fastai", globals(), "mlflow.fastai")
gluon = LazyLoader("mlflow.gluon", globals(), "mlflow.gluon")
```

`LazyLoader` inherits `types.ModuleType` to initialize instance as module type.

```py
class LazyLoader(types.ModuleType):
    def __init__(self, local_name, parent_module_globals, name):
        self._local_name = local_name
        self._parent_module_globals = parent_module_globals

        self._module = None
        super().__init__(str(name))

    def _load(self):
        """Load the module and insert it into the parent's globals."""
        if self._module:
            # If already loaded, return the loaded module.
            return self._module

        # Import the target module and insert it into the parent's namespace
        module = importlib.import_module(self.__name__)
        self._parent_module_globals[self._local_name] = module
        sys.modules[self._local_name] = module

        # Update this object's dict so that if someone keeps a reference to the `LazyLoader`,
        # lookups are efficient (`__getattr__` is only called on lookups that fail).
        self.__dict__.update(module.__dict__)

        return module

    def __getattr__(self, item):
        module = self._load()
        return getattr(module, item)

    def __dir__(self):
        module = self._load()
        return dir(module)

    def __repr__(self):
        if not self._module:
            return f"<module '{self.__name__} (Not loaded yet)'>"
        return repr(self._module)

```

- __init__에서 global() namespace를 받아 저장하고, module.__name__을 init하게 되고 LazyLoader 타입으로 부모 context(global context)에 모듈이 등록됩니다.
- __getattr__ 시점에, _load를 호출하여, 모듈이 import되지 않았을 경우, dynamic하게 import하여, global context에 이 module을 등록합니다.
- 이때 부모의 global context에 LazyLoader로 등록되어있던 모듈을 실제 import한 모듈로 overwrite합니다. 이를 통해서 최초 __getattr__ 이후로는 실제 import된 module이 LazyLoader를 대체하게 됩니다.
- `sys.modules`는 Python이 모든 로드된 모듈을 추적하는 데 사용하는 내부 캐시입니다. 모듈이 로드될 때, 이 캐시에 모듈을 등록하면, Python은 동일한 모듈을 재로드하는 대신 이미 로드된 모듈을 재사용합니다.


**문득 코드를 읽다, self._module를 저장하는 이유가 궁금해서 (내가 생각할 때는 불필요한 것 같은데) discussion을 남겨두었다.**

- [MLFLOW: Question on LazyLoader Implementation](https://github.com/mlflow/mlflow/discussions/10962)


## 2. Run

- [mlflow/__main__.py](https://github.com/mlflow/mlflow/blob/master/mlflow/__main__.py)
- [mlflow/cli.py](https://github.com/mlflow/mlflow/blob/master/mlflow/cli.py)

init 이후, `__main__.py`을 통해서 호출된 cli 모듈을 통해서 명령어에 대한 처리가 시작됩니다. cli.py에서는 크게 4가지의 명령어가 존재하며, 다음과 같습니다.

1. `run()`: Run an MLflow project from the given URI.
2. `server()`: Run the MLflow tracking server.
3. `gc()`: Permanently delete runs in the `deleted` lifecycle stage.
4. `doctor()`: Prints out useful information for debugging issues with MLflow.

기능 적인 측면에서 보면 크게 run, server 2가지만 파악하면 될 것같습니다.
이외에도 import를 통해서 나머지 명령어들을 불러옵니다. 최종적인 mlflow의 cli는 아래와 같습니다.

- load additional cli
```py

cli.add_command(mlflow.deployments.cli.commands)
cli.add_command(mlflow.experiments.commands)
cli.add_command(mlflow.store.artifact.cli.commands)
cli.add_command(mlflow.runs.commands)
cli.add_command(mlflow.db.commands)

# We are conditional loading these commands since the skinny client does
# not support them due to the pandas and numpy dependencies of MLflow Models
try:
    import mlflow.models.cli

    cli.add_command(mlflow.models.cli.commands)
except ImportError:
    pass

try:
    import mlflow.recipes.cli

    cli.add_command(mlflow.recipes.cli.commands)
except ImportError:
    pass

try:
    import mlflow.sagemaker.cli

    cli.add_command(mlflow.sagemaker.cli.commands)
except ImportError:
    pass


with contextlib.suppress(ImportError):
    import mlflow.gateway.cli

    cli.add_command(mlflow.gateway.cli.commands)

```

- FYI, suppress는 에러 ignore를 한줄로 사용하기 위해 사용됩니다.

```sh
> mlflow --help

Commands:
  artifacts    Upload, list, and download artifacts from an MLflow...
  db           Commands for managing an MLflow tracking database.
  deployments  Deploy MLflow models to custom targets.
  doctor       Prints out useful information for debugging issues with MLflow.
  experiments  Manage experiments.
  gc           Permanently delete runs in the `deleted` lifecycle stage.
  models       Deploy MLflow models locally.
  recipes      Run MLflow Recipes and inspect recipe results.
  run          Run an MLflow project from the given URI.
  runs         Manage runs.
  sagemaker    Serve models on SageMaker.
  server       Run the MLflow tracking server.
```



