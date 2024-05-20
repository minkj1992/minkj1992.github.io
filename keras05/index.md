# Fundamentals of machine learning


케라스 창시자에게 배우는 딥러닝, Fundamentals of machine learning
<!--more-->

## TL;DR

1. Understanding the tension between generalization and optimization, the fundamental issue in machine learning
2. Evaluation methods for machine learning models
3. Best practices to improve model fitting
4. Best practices to achieve better generalization

## 5.1 일반화: 머신러닝의 목표

머신러닝의 목표는 좋은 일반화 성능을 얻는 것입니다. 하지만 모든 머신러닝 모델은 과대적합의 문제가 발생합니다. 

- 과대적합: 훈련 데이터의 성능에 overfitting되어 평가 성능과의 차이가 커지는 것을 통해 발견할 수 있으며, 정확하게는 모델이 훈련 데이터에 overfit하게 훈련하여 일반화를 점차 잃어가는 현상

머신 러닝은 최적화와 일반화 사이의 줄다리기입니다.

- 최적화 (Optimization): 훈련 데이터에서 최고의 성능을 얻으려고 모델을 train하는 과정
- 일반화 (Generalization): 훈련된 모델이 이전에 본 적 없는 데이터에서 얼마나 잘 수행되는지

과대적합의 원인은 무엇일까? 어떻게 하면 좋은 일반화 성능을 달성할 수 있을까?

---

### 5.1.1 과소적합과 과대적합

validation과 train의 loss가 훈련이 진행되면서 같이 감소하게 되는 경우들은 “과소적합”이며, 이 과소적합이 발생한다는 것은 모델의 성능이 계속 발전될 여지가 있다는 뜻입니다. 

- Holdout 검증: Train / Test로 분리하여 모델의 성능 평가 (1번만 나눔)
- 과소적합 Underfitting:  훈련 데이터의 loss가 낮아질 수록 테스트 데이터의 loss도 낮아진다.

보통 모델은 훈련이 진행됨에 따라 validation loss가 낮아지다가 잠시 후 train보다 높아집니다.  이 높아지는 구간은 overfitting에 의해서, 낮아지는 구간은 모델이 아직 underfitting 되었기 때문입니다.

Overfitting은 모델이 훈련 데이터에 특화된 패턴을 학습하기 시작했다는 의미이며, 이 패턴은 새로운 데이터와 관련성이 적어 Generalization을 해치게됩니다.

과대적합은 아래의 경우 특히 발생할 가능성이 높습니다.

1. 데이터에 noise
2. 불확실한 특성
3. 드문 특성과 가짜 상관관계

**데이터에 노이즈**가 낀경우 mnist 경우, 이미지 사진의 형태가 정확히 보이지 않는 경우, 이를 무리하게 학습하려 시도하여 과적합 발생가능합니다. 또한 실제 모델 이미지에 대한 라벨링이 잘못된 경우에도 발생합니다.

데이터 잡음은 문제 정의에 **불확실성과 모호성**이 존재할 때에도 잡음이 발생할 수 있습니다. 예를 들면 바나나의 익은 정도 (덜익음, 썩음, 익음)를 판별하는 모델의 경우 라벨링하는 값에 주관이 개입될 수 있습니다. 또한 비슷하게 일정한 확률로 비가 오는 수치가 있을때, 랜덤하게 비가 오지 않는 경우도 존재합니다. 이 경우에도 무작위성에 의해서 잡음이 부여될 수 있습니다.

이렇듯 모델이 특성 공간의 모호한 영역에 너무 큰 확신을 가지면, 이런 확률적인 / 모호한 데이터에 과대적합 될 수 있습니다. 

마지막으로 **드문 특성과 가짜 상관관계**가 존재합니다.

잡음 feature는 모델이 해당 feature 패턴을 학습하기에, 필연적으로 overfitting을 유발시킵니다. 따라서 특성이 모델에 유익한지 또는 혼란스럽게 만드는지 확실하지 않다면 feature selection을 훈련전에 수행하는 것이 일반적입니다. 

**이를 위해서는 일반적으로 가용한 각 특성에 대해 유용성, 즉 특성과 레이블 사이의 상호 의존 정보(mutual information) 처럼 작업에 특성이 얼마나 유익한지 측정해야 합니다. 그 후 일정 threshold를 넘긴 특성만 사용합니다.** 

---

### 5.1.2 딥러닝에서 일반화의 본질

### 매니폴드 가설 예시

MNIST 28 * 28 unit8 벡터가 표현 가능한 가짓수는 784^256 로,  우주에 있는 원자 갯수 10^80보다 훨씬 큽니다. 하지만 이런 입력 중 매우 적은 수만 유효한 손글씨 데이터일 것이며, 이는 해당 공간에서 아주 작은 부분 공간만을 차지한다는 뜻입니다. (**연속적 (Continuous)**: 손글씨 숫자 이미지의 공간이 매끄럽고, 불연속적인 점이나 갑작스런 변화 없이 연결되어 있다는 것을 의미합니다.)

연속적인 형태에 약간의 수정이 존재하더라도 이는 여전히 handwritten으로 인식될 수 있으며, handwritten 공간은 784**256 표현 공간에서 작은 부분공간에 밀집되어있기 때문에, 약간의 수정된 이미지 또한 이 공간에 분포할 것이라고 가정합니다. 또한 endpoint(3과 8 각각의 최종 handwritting image)들은 중간과정에 변형이 일어나더라도 이 공간에 포함될 것을 가정합니다.

그럴 경우 다음과 같이 말할 수 있습니다. **“In technical terms, you would say that handwritten digits form a manifold within the space of possible 28 × 28 uint8 arrays.”**

즉 756차원  unit8에서 손글씨 숫자는 매니폴드를 형성합니다. 이를 일반화 하면

### 매니폴드 가설

1. 머신러닝 모델은 가능한 입력 공간안에서 비교적 간단하고, 저차원이며, 매우 구조적인(highly structured) 부분 공간(latent manifold)만 학습하면된다. 
    1. latent manifold = highly structured subspace
    2. **이는 다시말해 데이터가 고차원 공간에서 무작위로 분포되어 있는 것이 아니라, 더 낮은 차원의 매끄러운 구조를 따른다는 것입니다.**
2. 이 가설은 또한 고차원 공간에서의 두 데이터 점(예: 두 손글씨 숫자 이미지) 사이에 연속적인 경로가 존재하며, 이 경로 상의 모든 점들이 유효한 데이터 점임을 암시합니다. **그러므로 두 입력 데이터 사이를 보간(interpolate)할 수 있으며, 이 보간 과정에서 생성된 모든 중간 점들이 매니폴드 상에 위치한다고 주장합니다.**
    
    ![Screenshot 2024-05-20 at 2.46.52 PM.png](/images/keras05/Screenshot_2024-05-20_at_2.46.52_PM.png)
    

샘플 사이를 보간(interpolate)하는 능력은 딥러닝에서 일반화를 이해하는 열쇠입니다. **딥러닝은 latent manifold에서 sample들을 interpolate해서 continuous하게 빈곳을 채워서 해당 공간을 이해한다. (local generalization)**

### 딥러닝이 작동하는 이유?

매우 충분한 파라미터를 통해서 크고 복잡한 곡선(매니폴드)을 선택하여 훈련 데이터에 맞을 때까지 파라미터를 점진적으로 조정 + 학습 데이터들은 매니폴드 가설에 의해, 희소하게 분산된 독립 포인트가 아닌, 매니폴드 안에서 포인트들이 연속적인 경로를 따라 한 입력에서 다른 입력으로 변형될 수 있으며 중간과정들 한 매니폴드 공간안에 포함됨. 

![Screenshot 2024-05-20 at 2.46.11 PM.png](/images/keras05/Screenshot_2024-05-20_at_2.46.11_PM.png)

***이를 통해 부모 공간을 highly structued된 latent manifold를 train data를 통해서 찾아내고, 매니폴드 가설의 continuous에 의해서 각 point들은 연결되어 공간을 표현할 수 있다. 이를 통해 train data가 아니더라도  이전에 본 적 없는 입력을 이해할 수 있다. (generalization)*** 그러므로 딥러닝 모델이 훈련 샘플 사이를 단순히 보간하는 것 이상을 수행하리라고 기대해서는 안됩니다.

---

## 5.2 머신 러닝 모델 평가

관측할 수 있는 것만 제어할 수 있습니다. 우리는 모델의 일반화 성능을 신뢰 있게 측정할 수 있어야 합니다.

- information leak: 하이퍼 파라미터를 모델 train과정에서 변경하는 행위도 크게 보면 trainning일종이다. 그러므로 validation data set또한 overfitting 가능하다.

Train / Validation / Test set 평가 방법(3)

1. `Hold-out validation`
    1. fixed validation
    2. 단점: validation, testset이 전체 데이터를 통계적으로 대표하지 못할수도 있다.
2. `K-fold cross-validation`
    1. k개로 train + validation fold를 나누고 0 ~ k-1 index를 loop돌면서 validation을 선택해서 train
    2. ***when the performance of your model shows significant variance based on your train-test split***
    3. O(k)
3. `iterated K-fold cross-validation` 
    1. O(P * K), when p == random 횟수

모델 평가 유의할 점

1. 대표성: 훈련 / 테스트 세트가 데이터 대표성이 있는가
2. 시간의 방향: 함부로 random no
3. 중복: 데이터에 중복이 있어서, train과 validation에 각각 들어가게 되면 **훈련 데이터의 일부로 테스트하는 최악의 경우가 발생**

---

## 5.3 훈련 성능 향상하기

모델 훈련은 3(사실 2)가지 단계로 진행됩니다.

1. 약간의 일반화 능력을 보이고 과대적합할 수 있는 모델을 얻기
2. (Overfitting 경계 찾기)
3. 과대적합과 싸워 일반화 성능 개선

1번 과정(과대적합 모델 얻기)에서 일반적으로 세가지 문제가 발생가능합니다.

1. 훈련이 되지 않음
    - 시간이 지나도 loss 줄어들지 않음
    - 너무 일찍 중단
2. 훈련은 되지만, 의미있는 일반화 달성 못함
3. 여전히 과소적합(underfitting) 상태

![일반적인 모델 트레이닝](/images/keras05/Untitled.png)

일반적인 모델 트레이닝

### 5.3.1 훈련이 되지 않을 경우

- 시간이 지나도 loss가 줄지 않음
- 너무 일찍 중단 될 경우

![RMSprop(1.)](/images/keras05/Untitled%201.png)

RMSprop(1.)

우선 이런 문제는 항상 극복가능합니다. 왜냐면 딥러닝 모델은 랜덤한 trainning 데이터에서도 모델을 훈련할 수 있기 때문입니다. 이런 **상황은 항상 경사 하강법 과정에 대한 설정에 문제가 있습니다.**

- 옵티마이저 선택
- 모델 가중치의 초깃값 분포
- **학습률**
- **배치 크기**

일반적으로 학습률과 배치 크기 튜닝으로 해결합니다.

학습률

- 너무 높은 학습률: 최적접합(proper fit)을 크게 뛰어넘는 업데이트 가능
- 너무 낮은 학습률: 훈련이 너무 느려, 멈춰보이는 것처럼 보일 수 있음

배치 크기

- 배치 크기 증가
    - 유익하고 잡음이 적은(분산이 낮은) 그레디언트가 만들어짐
    - local minimum
    - 단,  과적합 위험 존재

![RMSprop(1e-2)](/images/keras05/Untitled%202.png)

RMSprop(1e-2)

### c.f. 배치 크기(batch size)에 따른 기대 효과

### 요약

- **큰 배치**:
    - **장점**: 빠른 학습, 안정적 수렴
    - **단점**: 과적합 위험 증가
- **작은 배치**:
    - **장점**: Regularization 효과, 더 세밀한 학습
    - **단점**: 학습 시간 증가, local minima 위험

배치 크기(batch size)에 따른 기대 효과와 그 이유를 정리하면 다음과 같습니다.

- Batch Size가 클때
    1. **학습 속도**:
        - **효과**: 학습이 빠르게 진행되며, 수렴 속도가 빠릅니다.
        - **이유**: 큰 배치는 병렬 연산의 효율을 극대화하여 하드웨어(특히 GPU)의 성능을 최대한 활용할 수 있습니다. 한번의 업데이트로 많은 데이터를 학습하므로 수렴 속도가 빨라집니다.
    2. **수렴 안정성**:
        - **효과**: 그래디언트의 분산이 작아지므로, 매번의 업데이트가 안정적입니다. 로컬 옵티마(local optima)에 빠질 확률이 작습니다.
        - **이유**: 배치 크기가 크면 개별 샘플의 노이즈가 평균화되어 그래디언트의 추정이 더 정확해집니다.
    3. **과적합**:
        - **효과**: 과적합의 위험이 있습니다.
        - **이유**: 큰 배치는 모델이 데이터의 일반적인 패턴을 학습하기 때문에, 작은 배치보다 정교한 regularization 효과가 덜할 수 있습니다. 데이터의 작은 변동에 대한 민감도가 낮아질 수 있습니다.
- Batch Size가 작을 때
    1. **학습 횟수 증가**:
        - **효과**: 1 epoch 당 iteration이 많아져서 학습이 더 세밀하게 이루어집니다.
        - **이유**: 작은 배치는 더 많은 업데이트를 필요로 하므로, 한 epoch 동안 여러 번의 파라미터 조정이 이루어집니다.
    2. **Regularization 효과**:
        - **효과**: Loss의 분산이 커져서 regularization 효과가 있습니다.
        - **이유**: 작은 배치는 한 번에 계산하는 데이터 양이 적기 때문에 그래디언트에 더 많은 노이즈가 포함됩니다. 이는 모델이 과적합하지 않고, 더 일반적인 패턴을 학습하게 돕습니다.
    3. **로컬 미니마**:
        - **효과**: local minima로 빠질 가능성이 높습니다.
        - **이유**: 노이즈가 많아 불안정한 경사 하강을 겪을 수 있으며, 이로 인해 매끄럽지 않은 경로를 따르게 될 수 있습니다.
    4. **학습 시간**:
        - **효과**: 학습 시간이 오래 걸립니다.
        - **이유**: 작은 배치로 인한 여러 번의 업데이트가 필요하므로, 전체 학습 과정이 오래 걸릴 수 있습니다. 하드웨어 효율이 떨어지기도 합니다.
- Q. 배치샘플을 늘리면 더 유익하고 noise가 적은(분산이 낮은) 그래디언트가 만들어지는 이유? (201p)
    
    ![Screenshot 2024-05-20 at 3.35.00 PM.png](/images/keras05/Screenshot_2024-05-20_at_3.35.00_PM.png)
    

### 5.3.2 훈련은 되지만, 의미있는 일반화 달성 못함

모델이 훈련되지만 어떤 이유에서인지 검증 지표가 전혀 나이지지 않는 경우, 즉 모델이 훈련되지만 일반화되지 않습니다. (랜덤 분류기가 달성 할 수 있는 것과 크게 다르지 않는 성능)

1. 단순하게 입력 데이터에 타깃 예측을 위한 정보가 충분하지 않는 경우
2. 현재 사용하는 모델의 종류가 문제에 적합하지 않는 경우

### 5.3.3 여전히 Underfitting하는 경우

![Screenshot 2024-05-20 at 7.14.34 PM.png](/images/keras05/Screenshot_2024-05-20_at_7.14.34_PM.png)

모델이 훈련되고 검증 지표가 향상되며 최소한 어느 정도의 일반화 능력을 달성하고 있지만, validation loss가 역전되지 않고 멈추어 있거나 매우 느리게 좋아지는 것 같을 때, **항상 overfitting이 가능하다는 것을 기억해야 합니다.**

훈련 손실이 줄어들지 않는 문제와 마찬가지로 이런 문제도 항상 해결할 수 있습니다. Overfitting할 수 없는 것처럼 보인다면 이는 모델의 표현능력(representational power)가 부족한 것입니다. 즉 용량이 더 큰 모델이 필요합니다.

- layer를 추가한다. (더 많은 가중치를 가지도록)
- 또는 층의 크기(more parameters)를 늘린다.
- 더 적합한 종류의 층 사용(5.3.2 구조에 대한 더 나은 가정)

## 5.4 일반화 성능 향상하기

모델이 어느정도 Overfitting을 할 수 있다면, 이제 Generalization을 극대화하는 데 초점을 맞출 차례입니다.

매니폴드 가설을 살펴보면서 우리는 딥러닝의 일반화가 데이터의 latent space(잠재 구조)에서 비롯된다는 것을 배웠습니다. 데이터를 사용해 샘플 사이를 부드럽게 보간(Interpolation)할 수 있다면 Generalization 성능을 가진 딥러닝 모델을 훈련 할 수 있습니다.

### 5.4.1 Dataset curation

1. 입력 → 출력 매핑하는 공간을 조밀하게 샘플링해야하니, 데이터가 충분한지 확인할 것
2. 레이블 할당 에러 최소화 (이상치확인)
3. 누락 값 처리, 데이터 정제
4. 많은 특성 중, 확실하지 않는 특성이 있다면 feature engineering
    1. 잠재 매니폴드를 더 매끄럽고 간단하고 구조적으로 만듭니다.
    2. 좋은 특성은 적은 자원을 사용해 문제를 더 멋지게 풉니다.(시계 cnn 딥러닝 모델을 통한 현재시각 확인 → 시계 각도와 시간을 매핑한 함수)
    3. 좋은 특성은 더 적은 데이터로 문제를 풀 수 있습니다. (샘플 갯수가 적다면 feature에 있는 정보가 매우 중요해집니다.)

### 5.4.3 Early stopping

validation이 역전되는 overfitting 구간을 찾으면 `EarlyStopping` 콜백을 사용해 이를 처리가능

### 5.4.4 모델 Regulation (규제)

**Regulation은 훈련 데이터에 완벽하게 맞추려는(overfitting) 모델의 능력을 적극적으로 방해하여, overfitting에 의한 모델의 validation loss를 줄이는 것이 목적입니다. → 모델 Generalization 상승,** 

Regulation을 통해 모델은 더 간단하고 더 평범하게, 곡선을 부드럽게, 더 일반적으로 만드는 경향을 가집니다. 

- 너무 작은 모델은 Overfitting 되지 않는다.
    - 모델의 기억 용량에 제한이 있어, 훈련 데이터를 단순 기억도 못할 정도의 사이즈
    
    ![Screenshot 2024-05-20 at 7.46.47 PM.png](/images/keras05/Screenshot_2024-05-20_at_7.46.47_PM.png)
    
- 너무 큰 모델은 바로 Overfitting된다.
    - 모델이 바로 overfitting된다.
    - validation loss 곡선이 고르지 않고 분산이 크다면 모델이 너무 큰 것
    - 이는 신뢰할 수 있는 검증 과정을 사용하지 않는다는 징후로도 해석가능, 예를들면 validation set이 너무 작은 경우

![Screenshot 2024-05-20 at 7.46.32 PM.png](/images/keras05/Screenshot_2024-05-20_at_7.46.32_PM.png)

### L1-norm vs L2-norm

[https://laid.delanover.com/difference-between-l1-and-l2-regularization-implementation-and-visualization-in-tensorflow/](https://laid.delanover.com/difference-between-l1-and-l2-regularization-implementation-and-visualization-in-tensorflow/)

[https://seongyun-dev.tistory.com/52#google_vignette](https://seongyun-dev.tistory.com/52#google_vignette)

![[https://scott.fortmann-roe.com/docs/BiasVariance.html](https://scott.fortmann-roe.com/docs/BiasVariance.html)](/images/keras05/5aa4ddcd-4a37-4078-ba00-86858aeac5b8.png)

[https://scott.fortmann-roe.com/docs/BiasVariance.html](https://scott.fortmann-roe.com/docs/BiasVariance.html)

![Screenshot 2024-05-20 at 10.25.19 AM.png](/images/keras05/e1fbf612-293f-44ee-a29d-c0e8372520ab.png)

![Screenshot 2024-05-20 at 8.16.40 PM.png](/images/keras05/Screenshot_2024-05-20_at_8.16.40_PM.png)

![Screenshot 2024-05-20 at 10.27.11 AM.png](/images/keras05/d471e465-a4a6-4346-85c1-374044c48820.png)

### Dropout

무작위로 층의 출력 특성을 일부 제외시키는 방식

혁펜하임에 따르면 연차 쓰게 해서, 남아있는 직원들이 서로 더 긴밀하게 일해보는 것을 반복. 이를 통해 overfitting시키는 node 배제하고 학습도 가능 

테스트할 때(대통령이 오면) 연착 쓰던 직원들 모두 불러서 동작하며, 드롭아웃 비율로 출력을 낮춰줘야 한다.
