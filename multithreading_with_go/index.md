# Mastering Multithreading Programming with golang


Learn about Multithreading, Concurrency & Parallel programming with practical and fun examples in Google's Go Lang.

<!--more-->
</br>

## 1. Introduction

Concurrency의 세계에서는 thread 수를 아무리 높이더라도, 일정 수준이 지나면 병목현상이 생긴다.

![](/images/parallel/amdal.png)

### Amdal's law
> Focus on latency (speed up)


![](/images/amdal.png)

`암달의 법칙`은 컴퓨터 시스템의 일부를 개선할 때 전체적으로 얼마만큼의 최대 성능 향상이 있는지 계산하는 데 사용된다.

> 1 / ((1-p) + (p/s))

1. 작업의 portion이 낮은 작업을 아무리 개선시키더라도, 시스템 전체에 미치는 영향은 미미하다. 즉 전체 작업의 효율을 최대한 증가시키고 싶다면 그 중에 가장 비중이 큰 작업부터 초점을 맞추는 것이 좋다.
2. (도표와 같이) 아무리 병렬작업이 늘어나더라도, 병렬화가 불가능한 작업들에 의해 병목현상이 발생하여 `speedup`의 한계가 정해지게 된다.

```python
from pprint import pprint

def amdahl(p,s):
    """
    Amdahl's law
    
    p is the proportion of execution time that the part benefiting from improved resources originally occupied.
    
    s is the speedup of the part of the task that benefits from improved system resources.
    """
    return 1 / ((1-p) + (p/s))
    

def simulate():    
    # represents p in amdahl's law
    portion = [95/100, 90/100, 75/100, 50/100]
    # represents s in amdahl's law
    number_of_processors = [1,2,4,8,16,32,64,128,256, 512, 1024, 2048, 4096, 8192, 16384, 32768, 65536]
    
    for s in number_of_processors:
        print(f'##################number_of_processors:{s}#############')
        for p in portion:
            pprint(f'parallel portion:{p}, speedup: {amdahl(p,s)}')
        

simulate()
"""
##################number_of_processors:1#############
'parallel portion:0.95, speedup: 1.0'
'parallel portion:0.9, speedup: 1.0'
'parallel portion:0.75, speedup: 1.0'
'parallel portion:0.5, speedup: 1.0'
##################number_of_processors:2#############
'parallel portion:0.95, speedup: 1.9047619047619047'
'parallel portion:0.9, speedup: 1.8181818181818181'
'parallel portion:0.75, speedup: 1.6'
'parallel portion:0.5, speedup: 1.3333333333333333'
##################number_of_processors:4#############
'parallel portion:0.95, speedup: 3.478260869565217'
'parallel portion:0.9, speedup: 3.0769230769230775'
'parallel portion:0.75, speedup: 2.2857142857142856'
'parallel portion:0.5, speedup: 1.6'
##################number_of_processors:8#############
'parallel portion:0.95, speedup: 5.925925925925925'
'parallel portion:0.9, speedup: 4.7058823529411775'
'parallel portion:0.75, speedup: 2.909090909090909'
'parallel portion:0.5, speedup: 1.7777777777777777'
##################number_of_processors:16#############
'parallel portion:0.95, speedup: 9.142857142857139'
'parallel portion:0.9, speedup: 6.400000000000001'
'parallel portion:0.75, speedup: 3.3684210526315788'
'parallel portion:0.5, speedup: 1.8823529411764706'
##################number_of_processors:32#############
'parallel portion:0.95, speedup: 12.54901960784313'
'parallel portion:0.9, speedup: 7.8048780487804885'
'parallel portion:0.75, speedup: 3.657142857142857'
'parallel portion:0.5, speedup: 1.9393939393939394'
##################number_of_processors:64#############
'parallel portion:0.95, speedup: 15.421686746987941'
'parallel portion:0.9, speedup: 8.767123287671234'
'parallel portion:0.75, speedup: 3.8208955223880596'
'parallel portion:0.5, speedup: 1.9692307692307693'
##################number_of_processors:128#############
'parallel portion:0.95, speedup: 17.414965986394545'
'parallel portion:0.9, speedup: 9.343065693430658'
'parallel portion:0.75, speedup: 3.9083969465648853'
'parallel portion:0.5, speedup: 1.9844961240310077'
##################number_of_processors:256#############
'parallel portion:0.95, speedup: 18.618181818181803'
'parallel portion:0.9, speedup: 9.66037735849057'
'parallel portion:0.75, speedup: 3.9536679536679538'
'parallel portion:0.5, speedup: 1.9922178988326849'
##################number_of_processors:512#############
'parallel portion:0.95, speedup: 19.284369114877574'
'parallel portion:0.9, speedup: 9.827255278310943'
'parallel portion:0.75, speedup: 3.9766990291262134'
'parallel portion:0.5, speedup: 1.996101364522417'
##################number_of_processors:1024#############
'parallel portion:0.95, speedup: 19.635666347075723'
'parallel portion:0.9, speedup: 9.912875121006778'
'parallel portion:0.75, speedup: 3.988315481986368'
'parallel portion:0.5, speedup: 1.9980487804878049'
##################number_of_processors:2048#############
'parallel portion:0.95, speedup: 19.816158684083195'
'parallel portion:0.9, speedup: 9.956246961594557'
'parallel portion:0.75, speedup: 3.994149195514383'
'parallel portion:0.5, speedup: 1.9990239141044412'
##################number_of_processors:4096#############
'parallel portion:0.95, speedup: 19.90765492102064'
'parallel portion:0.9, speedup: 9.978075517661392'
'parallel portion:0.75, speedup: 3.9970724566967553'
'parallel portion:0.5, speedup: 1.9995118379301928'
##################number_of_processors:8192#############
'parallel portion:0.95, speedup: 19.953720618682237'
'parallel portion:0.9, speedup: 9.98902572856969'
'parallel portion:0.75, speedup: 3.998535692495424'
'parallel portion:0.5, speedup: 1.9997558891736849'
##################number_of_processors:16384#############
'parallel portion:0.95, speedup: 19.976833506065944'
'parallel portion:0.9, speedup: 9.994509851765999'
'parallel portion:0.75, speedup: 3.999267712210899'
'parallel portion:0.5, speedup: 1.9998779371376258'
##################number_of_processors:32768#############
'parallel portion:0.95, speedup: 19.98841004056484'
'parallel portion:0.9, speedup: 9.9972541721329'
'parallel portion:0.75, speedup: 3.9996338225870436'
'parallel portion:0.5, speedup: 1.9999389667063383'
##################number_of_processors:65536#############
'parallel portion:0.95, speedup: 19.99420334070626'
'parallel portion:0.9, speedup: 9.998626897551304'
'parallel portion:0.75, speedup: 3.9998169029127695'
'parallel portion:0.5, speedup: 1.9999694828875292'
"""
```


intro에서의 한계는 `Gustafson's law`로도 설명이 가능하다.
+ 추가로 intro에서의 한계는 문제의 사이즈를 늘리는 방법으로 해결할 수 있다. (E.g FPS 게임에서 60fps 까지만 프레임을 지원할 수 있지만, 나머지 사운드에 대한 연산이나 타 유저와 통신하는 채팅에 대한 연산을 활성화 시킬 수 있다.)

![](/images/parallel/gustaf.png)

### Gustafson's law
> Focus on Throughput


> S(P)=P−a(P−1)

- P: 프로세서의 개수
- a: 병렬화되지 않는 부분의 비율
- S(P): 이론상 성능 향상 비율
  

성능 향상은 **같은 시간 동안 처리하는 데이터량의 비율**을 의미한다.

![](/images/gustafsons_law.png)


## 2. Creating and using thread

### process vs thread vs green thread

먼저 process는 memory space를 isolate하게 관리하며, fork()를 통해서 복사된다.

thread는 memory space를 공유하기 때문에 isolate하지 않으며, thread간 context switch에 따른 overhead가 발생한다. 

아래는 single processor에서 multi thread를 표현한 그림이다.

![](/images/parallel/thread.png)

context switch의 오버헤드는 스레드 수가 많지 않다면 큰 비중을 차지 하지 않지만, 스레드 수가 커짐에 따라 문제가 발생한다.

![](/images/parallel/context_switch.png)

![](/images/parallel/context_switch2.png)

이에 대한 대안으로 `green thread`가 사용되는데, `green thread`는 쉽게 말해 user level thread로 kernel level thread(흔히 우리가 말하는 thread)과 1:n관계를 가진다.

![](/images/parallel/green_thread.png)

하나의 kernel thread안에서 여러 `green thread`가 존재하기 때문에, 연산중이 kernel thread안에 존재하는 특정 `green thread`에서 `IO` 연산이 필요하여 Interrupt가 발생할 경우, io 연산이 불필요한 green thread들 까지 불필요하게 cpu연산을 하지 못하게 된다.

![](/images/parallel/green_thread2.png)

이를 해결하기 위해 golang에서는 hybrid 방식의 green thread를 사용한다. 

![](/images/parallel/hybrid_groutine.png)

즉 io가 필요한 green trhead가 발생할 경우, 동일한 방식으로 io interrupt를 시키면서

![](/images/parallel/hybrid_groutine2.png)


새로운 kernel level thread를 만들어주어, io연산이 불필요한 green thread들을 새로운 스레드로 넣어주어 효과적으로 동작하도록 한다.

![](/images/parallel/hybrid_groutine3.png)
