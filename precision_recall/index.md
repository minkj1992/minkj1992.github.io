# Precision과 Recall 절대 까먹지 않는방법


분명 어느정도 이해하고 넘어갔던 precision, recall인데 계속 헷갈린다. 이번 기회에 정확히 이해하고 외워보자.
<!--more-->

## TL;DR

- 기억방법: <ins>**리암씨팸**</ins>(리암 니슨씨가 사람 패는 모습을 상상)
    - **Re**call, **암**진단, Pre**ci**sion, 스**팸**분류
- 이해방법
    - Re-call(회상하다) -> 절대로 놓치면 안되는 케이스 -> 다시 상기해서 검사해봄
    - Precision(정밀) -> 함부로 버리면 안되는 케이스 -> 정밀하게 분류해야함
- 수식
    - Precision과 Recall은 분자/분모 모두 **TP**가 깔린다.
    - **P**recision: 모두 P로 구성됨
    $$\text{Precision} = \frac{\text{True Positives (TP)}}{\text{True Positives (TP)} + \text{False Positives (FP)}} $$
    - Recall: precision이 P로 구성되니 남는 FP 케이스가 분모에 존재
    $$\text{Recall} = \frac{\text{True Positives (TP)}}{\text{True Positives (TP)} + \text{False Negatives (FN)}}$$


## 고장난 Accuracy

흔히 ML 모델을 평가(evaluation)하는 지표로 정확도(Accuracy)를 사용하지만, Accuracy만으로는 불충분할 경우가 존재합니다. 이를 설명할 때 가장 흔히 드는 예시로 아래와 같은 문제가 존재합니다. 

> 암을 진단하는 병원에 내원한 환자가 암인지 아닌지 구분하는 예측 모델을 만들었습니다. 이때 100명의 환자에 대한 예측 결과는 아래와 같습니다.
> - TP(실제로 암이면서, 암으로 예측한 결과) = 1
> - TN(실제로 정상이고, 정상으로 예측한 결과) = 82
> - FN(실제로 암이지만, 정상으로 예측한 결과) = 16
> - FP(실제로 정상이지만, 암으로 예측한 결과) = 1
> 이진 분류일 때, 정확도 공식에 따르면 정확도는 아래와 같습니다.
> $$\text{Accuracy} = \frac{\text{TP} + \text{TN}}{\text{TP} + \text{TN} + \text{FP} + \text{FN}} = \frac{1 + 82}{1 + 82 + 16 + 1} = \frac{83}{100} = 0.83$$

즉 암을 진단하는 모델의 정확도는 83%입니다. 척보기에는 완전쓰래기 모델이라고 부르기에는 높은 정확도입니다. 하지만 잠깐 생각해보면 여기에 큰 문제가 있다는 것을 쉽게 생각해볼 수 있스빈다.

<ins>방문한 100명중에서, 실제 암환자는 17명 밖에 없었는데... 그 중에서 16명의 암환자를 정상이라고 오진한 모델의 정확도가 83%라고?</ins>

수식을 보면 정확도가 높게 평가된 이유는 82명이나되는 TN 때문입니다. 이는 다른 라벨과 비교해서 매우 많은 수가 포진되어있다는 것을 확인할 수 있습니다.
**이처럼 클래스(라벨)이 불균형하게 분포되어있으면, 정확도는 평가 지표로의 신뢰를 잃어버리게 됩니다.**(Accuracy Fails for Imbalanced Classification). 

클래스 분포가 약간 편향된 경우에 정확도가 여전히 유용한 측정항목이 될 수 있지만, <ins>클래스 분포의 불균형(skewed)이 심각한 경우 정확도는 모델 성능에 대한 신뢰할 수 없는 척도가 될 수 있습니다.<ins>

이런 불균형한 상황에 유용하게 쓸 수 있는 지표가 Precision, Recall입니다.

## 항상 헷갈리는 Precision, Recall

흔히 precision, recall을 설명할 때 등장하는 `Confusion Matrix`는 아래와 같습니다. 

<center>

![](/images/confusion_matrix.png)

source by [wandb](https://wandb.ai/mostafaibrahim17/ml-articles/reports/Precision-vs-Recall-Understanding-How-to-Classify-with-Clarity--Vmlldzo1MTk1MDY5)

</center>

- True Positive (TP) - 모델이 양성(Positive)을 양성으로 맞혔을 때
- True Negative (TN) - 모델이 음성(Negative)을 음성으로 맞혔을 때
- False Positive (FP) - 모델이 음성(Negative)을 양성(Positive)으로 잘못 예측했을 때
- False Negative (FN) - 모델이 양성(Positive)을 음성(Negative)으로 잘못 예측했을 때


`Confusion`이라는 말 그래돌 항상 헷갈리고, 이게 뭘 나타내는 지 모르겠습니다. 이를 위해 제가 이해한 방식을 설명드리겠습니다. 

우선 `Precision`(정밀도)의 수식은 아래와 같습니다.

$$\text{Precision} = \frac{\text{True Positives (TP)}}{\text{True Positives (TP)} + \text{False Positives (FP)}} $$

`Precision`은 수식의 분모를 보게되면 둘 모두 Positive라고 모델이 평가한 집합에서, 분자(실제로 True)인 비율을 의미합니다.

핵심은 분모에 존재하는 <ins>FP가 작아지는 것에 집중</ins>하는 지표로, <ins>**함부러 버리면 안됨**</ins> 케이스에서 유용한 지표입니다. 아래의 예시를 보겠습니다.

#### Precision이 사용되는 경우
- `탈모 흰머리 뽑기 로봇`: 정상 모근을 흰머리라고 오판하고 뽑으면 큰일남!
- `삼산 농장 잡초 제거 로봇` : 삼산을 잡초라고 오판하고 버리면 큰일남!
- `스팸 메일 필터링` : 중요한 메일을 스팸이라고 오판하고 버리면 큰일남!

이처럼 `Precision`은 위 예시와 같이 <ins>정상인것을 비정상이라고 판단하고 함부러 취급하면 큰일나는 경우에 유용</ins>한 지표입니다. **영어 단어를 보더라도 Precise(정밀한) 정밀하게 분류를 해야하는 뉘양스를 풍깁니다.**

---

다음으로 `Recall`(재현율)을 살펴보겠습니다. 사실 한국어 번역 재현율?? 뭔말이지 이해가 하나도 가지 않지만, Recall(상기,회상)이라는 단어처럼 소가 되새김질하는 것 같은 이미지가 떠오릅니다.

![](https://www.sciencetimes.co.kr/wp-content/uploads/2021/01/n-theheritagefarmme.jpg)

$$\text{Recall} = \frac{\text{True Positives (TP)}}{\text{True Positives (TP)} + \text{False Negatives (FN)}}$$

#### Recall이 중요한 경우 (Recall이 낮으면 안되는 경우)

- 의료 진단 (암 진단) : 암인데 암이 아니라고 판단하는 경우가 가장 크리티컬
- 자율 주행 사람 검출 : 자율 주행 차량의 경우 사람을 미검출하는 경우 인명 사고로 이어질 수 있음


## 전쟁으로 보는 Precision vs Recall

전쟁을 위해 살인 로봇을 만든다고 한다면 아군을 쏘지 않는 것이 중요하기 때문에, Precision이 중요합니다.
하지만 만약 스파이를 검출하는 로봇을 만든다고 하면, 스파이를 놓치지 않는 것이 중요하기 때문에 Recall이 중요해집니다.





