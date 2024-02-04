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


### 2.2. Start a Local Mlflow Server

Mlflow 코드를 분석하기 위해서는, 실제 UI상에서 experiment들이 어떻게 진행되어야 하는지 파악하는게 우선이라 생각되어 local에서 mlflow pull 받아서 확인해보았습니다.

- [Locally Run MLflow tracking server](https://mlflow.org/docs/latest/getting-started/tracking-server-overview/index.html#start-a-local-mlflow-server)



### 2.3. Getting Started with MLflow

대략 mlflow를 local에서 실행해보고 났으니, 구체적으로 mlflow의 구성요소들을 정리해 보겠습니다.

- [x] [Getting Started with Mlflow](https://mlflow.org/docs/latest/getting-started/index.html)
    - [x] [MLflow Tracking Quickstart](https://mlflow.org/docs/latest/getting-started/intro-quickstart/index.html)

![](/images/tracking-basics.png)


## 3. Concepts

- [Tracking Concepts](https://mlflow.org/docs/latest/tracking.html#concepts)

### 3.1. MLflow Tracking

- `Runs`: Executions of some piece of code
    - Each run records metrics, parameter, start ~ end times, artifacts(model weights, images, etc)
- `Experiments`: Group of runs, for a specific task

```bash
# UI browser
mlflow ui --port 5000

# mlflow server
mlflow server --host 127.0.0.1 --port 50
```

#### Tracking Components

- Tracking APIs: Tracking Server와 interact할 수 있는 인터페이스
    - 내 생각에는 이게, python, REST등의 client형식으로 관리 될 것 같다.(auth가 필요하니)
- Backend Store: metadata for each Run (i.g. run ID, metrics ..)
    - Default `/mlruns/**` (file based) 
    - Databaed-based (db ...)
- Artifact Store: 
    - input data files, model weight, images 따위
    - svn과 연동한다면 Artifact에서 작업이 되어야 할 것.
    - Parquet, S3 등으로 대체 가능
- Tracking Server
    - standalone HTTP server that provides REST API for accessing backend/artifact store.

![](/images/tracking-setup-overview.png)


Deployment 코드를 보다가 아래와 같은 패턴을 발견습니다. 이렇게 데코레이터에서 원본으로 wraps을 하는 이유는

1. `__name__`, `__module__` 함수 이름/모듈 보존: 로깅 / 디버깅 시 wrapper func가 원본 함수 이름을 물려 받을 수있음
2. `__doc__` repr 보존: 원본 함수의 주석이 그대로 보존 될 수 있어, wrapper가 되더라도 주석을 wrapper에 쓰는 것이 아니라, 실제 기능하는 코드에 주석을 넣어둘 수 있습니다.

오픈소스에서 `@functools.wraps(fn)`를 쓰는 이유는 딱 위의 2가지 이유 정도 있을 것 같습니다.


```py
@cache_return_value_per_process
def get_or_create_nfs_tmp_dir():
    """
    Get or create a temporary NFS directory which will be removed once python process exit.
    """
    from mlflow.utils.databricks_utils import get_repl_id, is_in_databricks_runtime
    from mlflow.utils.nfs_on_spark import get_nfs_cache_root_dir

    nfs_root_dir = get_nfs_cache_root_dir()
    ...

def cache_return_value_per_process(fn):
    """
    A decorator which globally caches the return value of the decorated function.
    But if current process forked out a new child process, in child process,
    old cache values are invalidated.

    Restrictions: The decorated function must be called with only positional arguments,
    and all the argument values must be hashable.
    """

    @functools.wraps(fn)
    def wrapped_fn(*args, **kwargs):
        if len(kwargs) > 0:
            raise ValueError(
                "The function decorated by `cache_return_value_per_process` is not allowed to be "
                "called with key-word style arguments."
            )
        if (fn, args) in _per_process_value_cache_map:
            prev_value, prev_pid = _per_process_value_cache_map.get((fn, args))
            if os.getpid() == prev_pid:
                return prev_value

        new_value = fn(*args)
        new_pid = os.getpid()
        _per_process_value_cache_map[(fn, args)] = (new_value, new_pid)
        return new_value

    return wrapped_fn
```


### 3.2. Mlflow Projects
> https://mlflow.org/docs/latest/projects.html

Mlflow Project란 소스 코드를 재사용 가능하게 패키징화 시킨 컴포넌트입니다.

- API, CLI entrypoint를 가지고 있습니다.
- chainning으로 multi-step workflows에 추가시킬 수 있습니다.
- Conda env나 MLproject (yaml file)를 통해서 properties들을 추가할 수 있습니다.

**Project는 Git URI 또는 local directory에서 `mlflow run` command를 사용해서 실행 가능합니다. 코드 상에서는 `mlflow.projects.run()`을 통해 python api로 실행가능합니다.** 이 api는 k8s나 databricks 환경에서 remote로 실행 가능합니다.

Project가 실행되는 환경을 구성하기 위해서는 아래 4가지 구성이 옵션들이 가능합니다.

1. Virtualenv (preferred)
2. Docker container
3. Conda
4. System environment

여기에서는 Viertualenv 세팅 위주로 살펴보겠습니다. 

#### MLflow Project (`Virtualenv`)

이 세팅을 구성한다면, mlflow는 pyenv를 활용해 isolated environment를 만들고, virtualenv를 통해 dependencies들을 포함시킵니다. 또한 mlflow 코드를 실행시키기 전에(prior to running the project code), 이 isolated env를 activate 시킵니다.

이 세팅을 활용하기 위해서는 2개의 파일 세팅이 필요합니다.

- `MLproject`: 시작 지점
- `python_env.yaml`: virtualenv environment description 


**MLProject example**

```
name: llm_summarization

python_env: python_env.yaml

entry_points:
  main:
    command: python summarization.py

```

**python_env example**

```yaml
python: "3.10"
build_dependencies:
  - pip
dependencies:
  - langchain>=0.0.244
  - openai>=0.27.2
  - evaluate>=0.4.0
  - mlflow>=2.4.0
```

이후 project를 실행하기 위해서는, MLproject 파일이 있는 디렉토리에서 아래 코들르 사용하면 됩니다.

```
> mlflow run .
```

이렇게 명령어를 `작성하게 되면 cli상에서 mlflow는 projects 패키지를 호출하며, 실제 코드에서는 아래 부분이 호출 됩니다.

<script src="https://gist.github.com/minkj1992/efbd5b3d78367829167c37dbe4460379.js"></script>


- **load_project**에서 python_env.yaml을 읽어, Project 객체를 생성합니다.
- Project는 Entrypoint와 Parameter를 init하고 이를 통해 source code 파일들을 link 또는 remote의 경우 storage_dir로 download합니다.
- 이후 projects.run을 통해 아래 코드가 호출되어 local 또는 databricks 같은 remote에 존재하는 코드들을 실행합니다. [code](https://github.com/mlflow/mlflow/blob/bf141c74ffb816e4bfce54177eef704c25849d0b/mlflow/projects/__init__.py#L103-L120) 
- 이후 projects.backend.local.py의 run이 호출되어 새로운 process가 동작합니다. [코드](https://github.com/mlflow/mlflow/blob/bf141c74ffb816e4bfce54177eef704c25849d0b/mlflow/projects/backend/local.py#L70)



> mlflow project를 실행해보다가, examples/llms/summarization을 실행해보는데, dep 에러가 발생해서 contribute했다. langchain이 워낙 빠르게 변화하는 open source다 보니까, dependency 관리가 쉽지 않나 보다. [FIx llm example tiktoken dependency error #10989](https://github.com/mlflow/mlflow/pull/10989)


### 3.3. MLflow LLMs
> https://mlflow.org/docs/latest/llms/index.html

Mlflow introduce `Deployment`(previously AI Gateway) that simplifies interactions with multiple LLM provides.

1. Support multiple llm providers
2. Integrate to mlflow model serving
3. Centralized scattered API keys
4. Seamless Provider swapping


Mlflow LLM는 아래와 같은 core concept들을 가지고 있습니다.

- Deployments
- Evaluation
- Prompt UI
- Flavors
    - OpenAI
    - Langchain
    - Transformer / BERT
- Tracking


- 드디어 MLflow의 dev dependency 관련된 세팅을 처리했고, bug fix를 contribute했다. https://github.com/mlflow/mlflow/pull/10998

