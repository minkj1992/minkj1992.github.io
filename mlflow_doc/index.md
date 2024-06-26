# Mlflow Document



<!--more-->

# Core Components

1. Tracking, tracking server
    - Log training statics
    - Log retrieval
    - register a model, to enable deployment
    
2. Model Registry, versioning
3. LLM deployment
4. Evaluate
5. Prompt engineering UI
6. Recipes, guided scenarios
7. Projects, packagin ml code, workflow, artifact


## 1. MLflow Tracking
> https://mlflow.org/docs/latest/tracking.html


The MLflow Tracking is an API and UI for logging parameters, code versions, metrics, and output files when running your machine learning code and for later visualizing the results. MLflow Tracking provides Python, REST, R, and Java APIs.

### Concepts

Tracking에는 Experiments와 Run 총 2가지의 개념이 존재합니다.

- `Runs`: script execution에 대한 tracking 단위
    - `with mlflow.start_run():`을 통해서 `mlflow.start_run()` + `mlflow.end_run()`를 자동으로 실행합니다.
    - `mlflow.autolog()`를 통해서 세팅할 수도 있습니다.
- `Experiments`: Group of `Runs`

```py
import mlflow

with mlflow.start_run():
    mlflow.log_param("lr", 0.001)
    # Your ml code
    ...
    mlflow.log_metric("val_loss", val_loss)
```

위와 같은 코드를 실행시키면 1번의 run이 실행생성되며, 이때 별다른 [Backend Stores](https://mlflow.org/docs/latest/tracking/backend-stores.html#backend-stores)와 [Artifact Stores](https://mlflow.org/docs/latest/tracking/artifacts-stores.html)를 설정하지 않았으면 local의 **./mlruns**로 backend stores 그리고 artifact stores가 지정됩니다.

좀더 상세히 설명하자면, Tracking run을 실행하게되면 그 결과로 artifact들과 run의 metadata들이 생성됩니다. 이를 저장하는 방식에 따라서 mlflow에서는 총 2가지로 구분합니다.

- `Backend Store`: persist metadata for each `run`, `experiment`
    - i.e run ID, start ~ end time
    - parameter, metrics
- `Artifact Store`: run을 실행하고 생성된 artifact
    - 일반적으로 model, model weight, images, data files(parquet)와 같이 large file을 관리합니다.




### Tracking server

생성된 metadata(backend store)와 artifact(artifact store)들을 UI 상에서 보여주고 싶으면 mlflow tracking server를 실행시켜야 합니다. 

```shell
# 1. default 127.0.0.1:5000
mlflow server
# 2. ui is alias of server
$ mlflow ui
# 3. explictly notate host and post
$ mlflow server --host 127.0.0.1 --port 8080
```

FYI, 아래 코드 참조하면 `mlflow ui`는 `mlflow server`의 alias입니다. 


```py
# mlflow/cli.py
class AliasedGroup(click.Group):
    def get_command(self, ctx, cmd_name):
        # `mlflow ui` is an alias for `mlflow server`
        cmd_name = "server" if cmd_name == "ui" else cmd_name
        return super().get_command(ctx, cmd_name)
```


tracking server를 실행하게 되면, ./mlruns의 artifact들을 읽어들입니다. 이를 도식화해보면 아래와 같습니다. 


![](/images/mf_tracking.png)


**만약 python api를 통해서 tracking server를 실행시키고 싶다면 아래와 같이 할 수 있습니다.**

```py
# search for runs that has the best validation loss among all runs in the experiment.

client = mlflow.tracking.MlflowClient()
experiment_id = "0"
best_run = client.search_runs(
    experiment_id, order_by=["metrics.val_loss ASC"], max_results=1
)[0]
print(best_run.info)
# {'run_id': '...', 'metrics': {'val_loss': 0.123}, ...}

```

## 2. LLMs
> https://mlflow.org/docs/latest/llms/index.html


MLflow also support for LLMs aims to abstract(with unified interface) inticating processes while building and deploying llm products.

### 2.1. Concepts

MLflow가 현재(2.11.3) LLM관련 제공하는 feture들을 group화 시키면 아래와 같습니다.

- Deployment Server
- LLM Evaluate
- Prompt Engineering UI
- LLM Tracking System


### 2.2. Deployments Server
> previously known as AI Gateway, [learn more](https://mlflow.org/docs/latest/llms/deployments/index.html)

#### Benefit
Deployments Server simplifies interactions with multiple llm providers. It has a lot of benefits such as following below.

1. **Unified endpoint**: Don't have to juggle between multiple provider APIs.
2. Simplified intefrations
3. Secure credential Management
    - Manges API keys in centralized storage
    - No more hard-coded cert keys
4. **Seamless provider swapping**
    - Swap providers without change codes.
    - Zero downtime provider, model or route swapping.


#### Deep dive to deployments server
```sh
$ poetry add 'mlflow[genai]'

export OPENAI_API_KEY=your_api_key_here
```

위의 세팅을 한 뒤, 아래 config.yaml을 생성하여 LLM deployment server의 스펙을 정의해줍니다.

```yaml
# config.yaml

endpoints:
  - name: completions
    endpoint_type: llm/v1/completions
    model:
      provider: openai
      name: gpt-3.5-turbo
      config:
        openai_api_key: $OPENAI_API_KEY
    limit:
      renewal_period: minute
      calls: 10

  - name: chat
    endpoint_type: llm/v1/chat
    model:
      provider: openai
      name: gpt-3.5-turbo
      config:
        openai_api_key: $OPENAI_API_KEY

  - name: embeddings
    endpoint_type: llm/v1/embeddings
    model:
      provider: openai
      name: text-embedding-ada-002
      config:
        openai_api_key: $OPENAI_API_KEY

```

이후 mlflow deployments cli를 실행시켜줍니다.

```sh
mlflow deployments start-server --config-path config.yaml --port 8080 --host localhost --workers 2

```

cli 명령어가 실행되면 아래 코드가 동작하게 되며 worker를 지정해 준 만큼 uvicorn 객체가 실행됩니다. 

```py
# mlflow/deployments/server/runner.py
    def start(self) -> None:
        self.process = subprocess.Popen(
            [
                sys.executable,
                "-m",
                "gunicorn",
                "--bind",
                f"{self.host}:{self.port}",
                "--workers",
                str(self.workers),
                "--worker-class",
                "uvicorn.workers.UvicornWorker",
                f"{app.__name__}:create_app_from_env()",
            ],
            env={
                **os.environ,
                MLFLOW_DEPLOYMENTS_CONFIG.name: self.config_path,
            },
        )

```

> uvicorn은 ASGI를 구현한 구현체이며, 내부적으로 uvloop를 사용하고 있습니다. uvloop는 Cython과 libuv(v8 engine event loop)를 사용해 구현되어있습니다. 보통 uvicorn을 multi process (core, parallel)에서 사용할 때는 gunicorn을 사용하며, uvicorn 내부적으로 `uvicorn.workers.UvicornWorker`라는 gunicorn과 compatible한 worker 구현체를 가지고 있어 이를 사용합니다. FYI, pyhton 3.12 부터는 sub interpreters가 도입되어, startup에대한 속도를 끌어올렸다고 하는데 더 자세한 내용은 아래를 참고하시면 됩니다. [Running Python Parallel Applications with Sub Interpreters](https://tonybaloney.github.io/posts/sub-interpreter-web-workers.html). 이걸 보다보면 Ray core나 dask parallel 같은 프레임워크는 GIL을 어떤 방식으로 우회하는지 궁금하네요.

이제 http://127.0.0.1:8080/docs로 이동하게되면 아래와 같은 swagger를 확인할 수 있습니다.

![](/images/mf_deployment_server.png)

만약 config.yaml을 수정하고 save 한다면, deployemnt server의 swagger 또한 변경 된것을 확인할 수 있을 겁니다.


```py
# mlflow/deployments/server/runner.py

def run_app(config_path: str, host: str, port: int, workers: int) -> None:
    config_path = os.path.abspath(os.path.normpath(os.path.expanduser(config_path)))
    with Runner(
        config_path=config_path,
        host=host,
        port=port,
        workers=workers,
    ) as runner:
        for _ in monitor_config(config_path):
            _logger.info("Configuration updated, reloading workers")
            runner.reload()
```

이는 deployment server를 실행하게 되면, 아래 gunicorn으로 실행되는 runner 객체가 생성되고 난 뒤, monitor_config(config_path)를 호출 하면서, config file을 계속 watch하고 있기 때문입니다.

> 만약 아래와 같은 에러가 도중에 발생한다면, known issue(psutil을 따로 관리하는 게 의도된)이므로, 그냥 pip install psutil 또는 poetry add psutil을 통해서 다운로드 받으시면 됩니다.

    ```
      File "/Users/minwook/code/personal/mlflow-demo/.venv/lib/python3.11/site-packages/mlflow/gateway/utils.py", line 71, in kill_child_processes
        import psutil
    ModuleNotFoundError: No module named 'psutil'
    ```


```py
class Runner:
    ...
        def reload(self) -> None:
        kill_child_processes(self.process.pid)

```

만약 file에 변경이 일어나게되면 runner.reload()를 통해서 process의 child들인 uvicorn.workers.UvicornWorker들이 종료되게 됩니다. 하지만 parent인 gunicorn은 종료되지 않았기 때문에, 새로운 config를 읽어서 uvicorn worker들이 재실행됩니다.

```py
from watchfiles import watch
def monitor_config(config_path: str) -> Generator[None, None, None]:
    with open(config_path) as f:
        prev_config = f.read()

    for changes in watch(os.path.dirname(config_path)):
        if not any((path == config_path) for _, path in changes):
            continue

        if not os.path.exists(config_path):
            _logger.warning(f"{config_path} deleted")
            continue

        with open(config_path) as f:
            config = f.read()
        if config == prev_config:
            continue

        try:
            _load_route_config(config_path)
        except Exception as e:
            _logger.warning("Invalid configuration: %s", e)
            continue
        else:
            prev_config = config

        yield

```

문득 궁금해지는 건, polling 하면서 file을 읽는 부분의 성능적인 이슈를 어떻게 해결했는지인데, 조금만 생각해봐도 file을 polling하면서 계속 watch가 실행된다면 부하가 있을 것 같아서 watchfiles라는 오픈소스를 확인해보니, rust를 호출해서 성능적인 퍼포먼스를 끌어올렸습니다.


> watchfile 소스코드를 읽다 보니, rust로 파일을 읽어들이는 것을 확인했다. 글들을 찾아보니 [py03](https://github.com/PyO3/pyo3)를 통해서 python에 대한 rust binding이라고 하는데, 이를 통해서 python을 통해서 rust코드를 호출 할 수 있는 것 처럼 읽혀졌다. 좀 더 찾아보니 numpy 또한 rust를 호출해서 성능을 향상 시키는 시도가 이뤄지고 있다. [PyO3 + rust-numpy](https://terencezl.github.io/blog/2023/06/06/a-week-of-pyo3-rust-numpy/) 아직도 어떻게 python이 rust binary를 호출하는지 내부구조가 잘 그려지지 않지만 여기까지만 파악하기로 한다.


마지막으로 현재 mlflow는 config에서 endpoint를 설정할 때 아래 3가지 type만 제공하고 있다.

- “llm/v1/completions”
- “llm/v1/chat”
- “llm/v1/embeddings”

이 endpoint는 내부적으로 model provider에 따른 api를 wrapping하고 있다. see follwing url if you want to learn more in detail.


[supported-provider-models](https://mlflow.org/docs/latest/llms/deployments/index.html#supported-provider-models)


#### with python code

만약 swagger가 아닌 소스코드 상에서, deployment에 접근하고 싶다면 [MlflowDeploymentClient](https://mlflow.org/docs/latest/python_api/mlflow.deployments.html#mlflow.deployments.MlflowDeploymentClient)를 사용하면 됩니다.




### LLM Tracking





## 3. Evaluate
