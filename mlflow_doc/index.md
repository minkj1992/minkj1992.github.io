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

Deployments Server simplifies interactions with multiple llm providers. It has a lot of benefits such as following below.

1. **Unified endpoint**: Don't have to juggle between multiple provider APIs.
2. Simplified intefrations
3. Secure credential Management
    - Manges API keys in centralized storage
    - No more hard-coded cert keys
4. **Seamless provider swapping**
    - Swap providers without change codes.
    - Zero downtime provider, model or route swapping.

#### 2.2.1. Confiquring and Starting the Deployments Server

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
mlflow deployments start-server --config-path config.yaml --port {port} --host {host} --workers {worker count}

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

#### 2.2.2. Concepts of Deployments Server

- Provider
- Models
- Routes


### LLM Tracking





## 3. Evaluate
