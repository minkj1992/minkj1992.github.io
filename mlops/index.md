# All about mlops


> [MLOps for ALL](https://mlops-for-all.github.io/en/docs/introduction/intro)

# 1. Introduction

## What is MLOps?

- Continuous Tranning pipeline
- Model CI / CD

## Levels of MLOps

#### 0단계: 수동 process

![](https://mlops-for-all.github.io/assets/images/level-0-85b288b20c458e64055199fc50b1fe86.png)


- ML -model-> Ops, 즉 model을 통해 ops팀과 ml팀이 소통하는 방식
- model
    - codes
    - 학습된 weight and bias
    - environment

#### 1단계: ML 파이프라인 자동화

![](https://mlops-for-all.github.io/assets/images/level-1-pipeline-b2979b34d4804546ef4005cdf0f6311a.png)

1. Trainning Pipeline
2. Continuous Training

#### 2단계: CI/CD 파이프라인 자동화


## Componenet of MLOps

![](https://mlops-for-all.github.io/assets/images/mlops-component-540cce1f22f97807b54c5e0dd1fec01e.png)

1. Experimentation, prototype (ml engineer)
    1. jupyter notebook
    2. Data, Hyper parameter, eval metrics
    3. Visualization
2. Data Processing
    1. 크게 3가지에서 사용됩니다.
        1. ML model develop phase
        2. COntinuous Training pipeline
        3. API deployment
    2. 다양한 데이터 소스와 서비스에 호환되는 데이터 connector feature
    3. Encoder / Decoder
    4. 데이터 변환과 Feature engineering
    5. 학습과 서빙을 위한 scale-out 가능한 Batch / Stream data feature
3. Model Training
    1. ML framework 실행을 위한 env
    2. GPU 분산 학습을 위한 환경 제공
    3. Hyper parameter tunning 그리고 최적화 기능 (Hyper parameter? 모델링시 사용자가 직접 세팅해주는 값들, top_p, top_k, token_length ...)
4. Model Evaluation
    1. 모델 performance 측정
    2. CT 결과의 성능 지속적인 추적
    3. Visualization
5. Model Serving
    1. low latency, high availability (HA)
    2. 다양한 ML 모델 프레임워크 지원
    3. 복잡한 형태의 모델간, 스텝별 flow 서빙 (preprocess, postprocess, 모델간 통신)
    4. autoscaling
    5. logging, 특히 llm을 직접 관리한다면, CT를 위해 agent의 logging들이 필수
6. Online Experimentation
    1. A/B testing
    2. 새로운 모델 생성 시, 해당 모델을 배포하면 어느 정도의 성능을 보일지 검증하는 기능
    3. Multi-armed bandit testing, 한정된 리소스(시간, 트래픽)에서 여러개의 테스트 그룹 중 가장 좋은 그룹 선택
7. Model Monitoring
    1. model 성능 측정
8. ML Pipeline
    1. 다양한 이벤트들을 통한 실행 기능
9. Model REgistry
    1. 모델 lifecycle관리하는 중앙 저장소
    2. versioning, metadata
10. Dataset and Feature Repository
    1.  Dataset sharing, search, versioning
    2.  Event streaming 및 온라인 추론 작업에 대한 실시간 처리 및 서빙 기능
    3.  사진 / 텍스트 / 테이블 등 다양한 형태의 데이터 지원
11. ML metadata and Artifact Tracking
    1.  ML 산출물
    2.  ML artifacts history관리 기능

## Why Kubernetes?

- Container를 통한 소통 편리성
- phase 분리 용이성 등
- node 리소스 효율적 관리 GPU 등.

# 2. Setup Kubernetes

- kustomize
- istio, service mesh
- CSI
- Argo
- Helm
- k3s



# 3. Kubeflow

## Kubeflow Concepts
1. Component
    - Component contents
    - Component wrapper, kubeflow로 component가 전달
2. Artifacts, Componenet를 통해 생산
3. Pipeline
4. Run


### 1.1. Component contents

컴포넌트 콘첸츠를 구성하는 것은 총 3가지로
1. Environment
2. Python code w\ Config
3. Generates Artifacts


```py
import dill
import pandas as pd

from sklearn.svm import SVC

train_data = pd.read_csv(train_data_path)
train_target= pd.read_csv(train_target_path)

clf= SVC(
    kernel=kernel
)
clf.fit(train_data)

with open(model_path, mode="wb") as file_writer:
     dill.dump(clf, file_writer)
```


![](/images/kubeflow_components.png)


### 1.2. Component Wrapper

컴포넌트 래퍼는 컴포넌트 콘텐츠에 필요한 config를 전달하고 실행시키는 작업을 합니다.


![](/images/component_wrapper.png)


### 2. Artifacts

- Model
    - 파이썬 코드
    - 학습된 weights
    - network 구조
    - 실행시키기 위한 환경
- Data
    - 전처리된 feature
    - 모델의 예측값
- Metric
    - Dynamic metric, train loss와 같이 epoch마다 계속 변화하는 값
    - Stataic Metric, 학습이 끝난 뒤 최종적으로 모델을 평가하는 정확도 등

### 3. Pipeline

파이프라인은 컴포넌트의 집합과 이를 실행시키는 순서도로 구성되어있습니다. 순서도는 DAG 이뤄져 있으며 조건문을 포함시킬 수 있습니다.

또한 컴포넌트를 실행시키기 위해서는 Config가 필요한데, Pipeline을 구성하는 컴포넌트의 Config들을 모아 둔 것이 파이프라인 Config 입니다.


![](/images/kubeflow_pipeline.png)

### 4. Run

Kubeflow에서는 실행된 파이프라인을 Run이라고 부릅니다.
**파이프라인이 실행되면, 각 컴포넌트들이 아티팩트들을 생성하고, Kubeflow pipeline에서는 Run하나당 고유한 ID를 생성한 뒤, Run에서 생성되는 모든 아티팩트들을 저장합니다.**

![](/images/kubeflow_run.png)


### 5. Experiment

Experiment란 Kubeflow 에서 실행되는 Run을 논리적으로 관리하는 단위입니다.

### 6. InputPath, OutputPath



