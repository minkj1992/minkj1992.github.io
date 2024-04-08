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



## 3. Evaluate
