# CKA certificate


MLflow ambassador 혜택으로 linux foundation의 무료 수강권을 얻을 수 있어서, k8s 자격증 시리즈를 구매했다.
이 참에 cka -> ckad -> cks를 빠르게 따볼 계획이다.
<!--more-->


![](/images/cka2.png)

## Prerequisite

> https://training.linuxfoundation.org/full-catalog/#

- [Important Insturctions: CKA and CKAD](https://docs.linuxfoundation.org/tc-docs/certification/tips-cka-and-ckad)
- [CKA Curriculum_v1.29](https://github.com/cncf/curriculum/blob/master/CKA_Curriculum_v1.29.pdf)



<center>

![](/images/CKA_Curriculum_v1.29-0.jpg)
![](/images/CKA_Curriculum_v1.29-1.jpg)
![](/images/CKA_Curriculum_v1.29-2.jpg)

</center>

## Planning


1. CNCF endorsed udemy lecture [Certified Kubernetes Administrator (CKA) with Practice Tests](https://www.udemy.com/course/certified-kubernetes-administrator-with-practice-tests/)를 듣는다.
2. [Course 필기](https://github.com/kodekloudhub/certified-kubernetes-administrator-course)를 참조해서 복습
3. 실습이 필요하다면 [kubernetes-the-hard-way](https://github.com/kelseyhightower/kubernetes-the-hard-way/tree/master) 셋업해서 로컬에서 연습해볼 예정


하루 약 3~4시간씩 강의를 들어보자 약 23시간 강의이기 때문에 다 듣는데 약 5~6일 소요되기 때문에, 일주일 뒤에 테스트볼 예정이다. (~ 24.4.26)


24.04.19 ~


# Chapter 2: Core concepts

## Cluster architecture

- Master: Manage, Plan, Schedule, Monitor nodes
    - etcd cluster: save, distributed
    - kube-apiserver: orch all within cluster
    - kube controller manager: run many controller (node, replica )
    - kube-scheduler
- Worker Nodes
    - kubelet: api client
    - kube-proxy: enable communication with other service
    - container runtime
        - docker, rkt, containerd, crio, podman...

## Docker vs containerd

- k8s(CRI <--- OCI)  <-> runtime (docker, rkt, podman...)

- containerd가 실제 k8s의 CRI( container runtime interface)와 호환되며, containerd는 최초의 docker의 runtime
- crictl is belongs to k8s, netctl and ctl is belongs to containerd

## ETCD

- key-value store
    - stores information in the form of document or pages 


> https://tech.kakao.com/2021/12/20/kubernetes-etcd/

[etcd kakao](https://tech.kakao.com/storage/2021/12/01-2.png)

- 높은 신뢰성을 제공하기 위해 ETCD는 RSM(Replicated state machine)이다.
- 이는 똑같은 데이터를 여러 서버에 계속 복제하는 것이고, 이 방법을 사용하는 머신을 RSM이라 칭합니다.
- 여러 서버에 복제하게 되면 발생하는 데이터 복제 과정에서 발생하는 여러 문제를 해결하기 위해 consensus를 확보하는 것이 핵심이며, 아래 4가지 속성을 만족한다는 것을 뜻합니다.
- etcd는 Raft알고리즘을 통해 이를 구현합니다.
    1. Safety
    2. Available
    3. Independent from timing
    4. reactivity

[etcd dive deep](https://medium.com/@extio/deep-dive-into-etcd-a-distributed-key-value-store-a6a7699d3abc)

- leader만 write 가능 이후 이를 follower에게 전파하여 append log
- follower는 client로 부터 read 요청을 처리 가능하다. (Q. timinig )

[etcd kv api](https://etcd.io/docs/v3.3/learning/api_guarantees/)

etcd tries to **ensure the strongest consistency** and durability guarantees for a distributed system. This specification enumerates the **KV API** guarantees made by etcd.

1. Atomicity: 모든 API request are atomic
2. Consistency: All Api calls ensure [sequential consistency](https://en.wikipedia.org/wiki/Consistency_model#Sequential_consistency), the strongest consistency guarantee ava from distributed systems


- 엄격한 일관성 모델보다 약한 메모리 모델입니다.
- 변수에 대한 쓰기는 즉시 표시될 필요는 없지만, 서로 다른 프로세서에 의한 변수에 대한 쓰기는 모든 프로세서에서 동일한 순서로 표시되어야 합니다.
- 모든 실행 결과가 데이터 저장소에 있는 모든 프로세스의 (읽기 및 쓰기) 작업이 순차적 순서로 실행된 것과 동일하고 각 개별 프로세서의 작업이 이 순서대로 나타나는 경우 순차적 일관성이 충족됩니다

For watch operations, etcd guarantees to return the same value for the same key across all members for the same revision.

> it is impossible for etcd to ensure strict consistency. etcd does not guarantee that it will return to a read the “most recent” value (as measured by a wall clock when a request is completed) available on any cluster member.

## etcd in k8s

`--advertise-client-urls`(internal_ip:3479): uri that etcd will listen, kubeapi가 여기로 접근

## kube-api server
> https://github.com/kodekloudhub/certified-kubernetes-administrator-course/blob/master/docs/02-Core-Concepts/06-Kube-API-Server.md


- We can trigger kube-apiserver by kubectl and kubeadm or directly we can send request to kube-apiserver by api(i.g. curl)
- Kube-apiserver is the only componenet that iteracts directly to the etcd datastore

```
curl -X POST /api/v1/namespaces/default/pods ... [other]
```


1. Authenticate User
2. Validate Request
3. Retrieve data
4. Update ETCD
5. Scheduler
    - kube-scheduler keep watching kube-apiserver so that it can know when etcd changes
    - etcd changed -> scheduler check and identifies right node to put pod on -> request to kube-apiserver -> send request to kubelet
6. kubelet

## Kube controller manager

Kube Controller manager manages various controllers in k8s. Then what is controller?

In k8s terms, a controller is a process that continuously monitors the state of the componenets within the system and works towards bringing the whole system to the desired functioning state

> In Kubernetes, controllers are control **loops** that watch the state of your cluster, then make or request changes where needed. Each controller tries to move the current cluster state closer to the desired state

```
--node-monitor-period duration     Default: 5s

--node-monitor-grace-period duration     Default: 40s

# https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/#taint-based-evictions
`pod-eviction-timeout flag is deprecated as per v1.26 . -> tolerationSeconds
```

## Kube scheduler

Kube-scheduler is responsible for scheduling pods on nodes. The kube-schduler is only responsible for deciding which pod goes on which node. It doesn't actually place the pod on the nodes, that's the job of the `kubelet`

Schedule pod <-> node

1. Filter Nodes
2. Rank Nodes
3. Post state which node to shcdule

## Kubelet

#### The kubelet is the primary "node agent" that runs on each node. It can register the node with the `apiserver`

#### The lifecycle of the kubeadm CLI tool is decoupled from the kubelet, which is a daemon that runs on each node within the Kubernetes cluster. It means you have to install kubelet and kubeadm when you init cluster by kubeadm. 

- The kubelet will create the pods on the Nodes
- Monitor Node & Pods

## Kube proxy

- kube-proxy is a network proxy that runs on each node in cluster, implementing part of the kubernetes `Service` concept.
- kube-proxy maintains network rules on nodes, allow network communication to pods from network sessions inside or outside of cluster.
- It uses the operating system packet filtering layer (OSI-L3) if can.

## Pod

- single container pod
- multi-container pod
    - pod <-localhost-> helper containers, also shares persist vol

```yml
apiVersion:
kind:
metadata:
    name:
    labels:
        app:
        type:
spec:
    containers:
```


