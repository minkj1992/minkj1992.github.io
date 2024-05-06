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

> edit, practice에 대한 시간이 빠져있었다. 조금 일정을 미뤄 + 2일뒤인 24.4.28에 시험을 볼 예정이다.

24.04.19 ~


# Chapter 2: Core concepts

![](/images/k8s-full.png)

![](/images/k8s-full2.png)

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

The kubelet is the primary "node agent" that runs on each node. It can register the node with the `apiserver`. The lifecycle of the kubeadm CLI tool is decoupled from the kubelet, which is a daemon that runs on each node within the Kubernetes cluster. It means you have to install kubelet and kubeadm when you init cluster by kubeadm. 

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

- get node

```
k get po -o wide
```



- get yaml from running and edit on runtime

```
# if already deployed
# $ k get po redis -o yaml 
$ kubectl run redis --image=redis123 --dry-run=client -o yaml > redis.yaml
$ k create -f redis.yaml
$ k edit or vim redis.yaml

# :%s/redis123/redis
```

- Q. k edit 으로 containerStatus를 edit하면 어떻게 되는거지?

## Replicaset

- It is often used to guarantee the availability of a specified number of identical pods
- Pods created from ReplicaSets can be distributed and executed on multiple nodes based on schduling, topologySpreadConstraints(affinity, maxSkew, labelSelector ..)


```py
- scale
- replace
```

## Deployments

A Deployment provides declarative updates for Pods and ReplicaSets.

The following are typical use cases for deployments

1. rollout a ReplicaSet
2. Declare the new state of the Pods
3. Rollback to earlier deployment revision
4. Scale out deployment
5. Pause the rollout of a deployment


> ReplicaSet-A for controlling your pods, then You wish to update your pods to a newer version, now you should create Replicaset-B, scale down ReplicaSet-A and scale up ReplicaSet-B by one step repeatedly **(This process is known as rolling update).**


```py
k api-resources | grep deployment
k create deployment --image=nginx nginx --replicas=4 --dry-run=client -o yaml > nginx-deployment.yaml
```

## Services

Kubernetes Services enables communication between various components within and outside of the application.

- NodePort: Where the service makes an internal port accessible on a port on the NODE.
- ClusterIP
- LoadBalancer

#### NodePort

- NodePort uses node machine's port and Node's IP.

![](/images/k8s-nodeport1.png)

Kubernetes sets up a cluster IP address, the same as if you had requested a Service of `type: ClusterIP` (10.106.1.12)

- To connect the service to the pod, use selector


```yaml

---
apiVersion: v1
kind: Service
metadata:
  name: myapp-svc
spec:
  type: NodePort
  ports:
  - targetPort: 80
    port: 80
    nodePort: 30008
  selector:
    app: myapp
    type: front-end

---
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  labels:
    app: myapp
    type: front-end
spec:
  containers:
  - name: nginx-container
    image: nginx

```


**service.spec.selector must be equal to pod.metadata.labels** to connect each other.

#### How do I verify if the NodePort service and the pod are properly connected?
> To confirm whether the service and the pod are properly connected, we can check the endpoints via the service describe as shown below, and then compare them with the IP of the pod.

![](/images/nodeport-with-pod.png)


#### A service with multipe pods with single service

- Random algorithm is used to balance the load of traffic
- Session Affinity: yes in this case

![](/images/multi_pod_nodeport.png)

#### When pods are distributed across multiple nodes

![](/images/multi_node_nodeport.png)

Let's look at what happens when the Pods are distributed across multiple nodes. In this case, we have the web application on Pods on separate nodes in the cluster, When we create a service, without having to do any additional configuration.

Kubernetes automatically creates a service that spans **across all the nodes in the cluster** and **maps the target port to the same node port on all the nodes in the cluster.** 

**This way you can access your application using the IP of any node in the cluster and using the same port number which in this case is 30,008.** As you can see, using the IP of any of these nodes, and I'm trying to curl to the same port, and the same port is made available on all the nodes part of the cluster.

#### ClusterIP

- The service creates a **Virtual IP** inside the cluster to enable communication between different services such as a set of frontend servers to a set of backend servers.
- **A kubernetes service can help us group the pods together and provide a single interface to access the pod in a group.**

![](/images/k8s_clusterip.png)


#### LoadBalancer

Where the service provisions a loadbalancer for our application in supported cloud providers.


## Namespaces

In k8s, namespaces provide a mechanism for **isolating groups of resources** within a single cluster. Names of resoures need to be unique within a namespace, but not across namespaces.

> This means that when we have namespaces such as (dev, sandbox, prod), then we can generate golang backend pods for each environment respectively(accordingly).

![](/images/k8s-namespace-isolation.png)

#### namespace cli

```py
$ k get ns
NAME              STATUS   AGE
kube-system       Active   9m56s
kube-public       Active   9m56s
kube-node-lease   Active   9m56s
default           Active   9m56s
finance           Active   22s
marketing         Active   22s
dev               Active   21s
prod              Active   21s
manufacturing     Active   21s
research          Active   21s

$ k get ns --no-headers | wc -l
10

$ k get po -n=research --no-headers | wc -l
2
```

```py
# create and run pod with finance namespace
# 생각해보니까 apply, create으로 pod 직접적으로 만들지 않았던 것 같네. 곧바로 run 했던 것 같은데, run = create + run like docker
k run redis -n=finance --image=redis

```

```py
# swich ns
$ kubectl config set-context $(kubectl config current-context) --namespace=dev

# view pods in all namespace
$ kubectl get pods --all-namespaces
```

#### kubernetes DNS rule

![](/images/k8s_dns_rule.png)

- <Service_Name>.<Namespace>.svc.cluster.local

- same namespace: just use service name
- another namespace: db-service.dev.svc.cluster.local

#### deterministic namespace

- If you want to make sure that this pod gets you created in the dev env all the time, even if you don't specify in the command line, you can move the --namespace definition into the pod-definition file.

```yaml
# or $ kubectl create -f pod-definition.yaml --namespace=dev

apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
  namespace: dev
  labels:
     app: myapp
     type: front-end
spec:
  containers:
  - name: nginx-container
    image: nginx
```


#### ResourceQuota

- To limit resources in a namespace, create a resource quota. To create one start with ResourceQuota definition file.

![](/images/k8s_ns_resource_quota.png)

```yaml
apiVersion: v1
kind: ResourceQuota
metadata:
  name: compute-quota
  namespace: dev
spec:
  hard:
    pods: "10"
    requests.cpu: "4"
    requests.memory: 5Gi
    limits.cpu: "10"
    limits.memory: 10Gi
```

## Imperative

```py
$ k run nginx-pod --image=nginx:alpine
$ k run redis --image=redis:alpine --labels="tier=db"
$ k expose po redis --port=6379 --name=redis-service
$ k create deploy webapp --image=kodekloud/webapp-color --replicas=3
$ k run custom-nginx --image=nginx --port=8080
$ k create ns dev-ns
$ k create deploy redis-deploy -n dev-ns --image=redis --replicas=2
$ k run httpd --image=httpd:alpine && k expose po httpd --port=80 --name=httpd
```


# Chapter 3: Schedule

## Manual schduling

- Schduler bind pod to nodes
- If there is no scheduler, pod's status would be 'Pending'


```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    name: nginx
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
    - containerPort: 8080
```

- when there is scheduler

```py
floe@floe-QEMU-Virtual-Machine:~$ k get po -o wide
NAME    READY   STATUS    RESTARTS   AGE   IP            NODE       NOMINATED NODE   READINESS GATES
nginx   1/1     Running   0          40s   10.244.0.14   minikube   <none>           <none>
floe@floe-QEMU-Virtual-Machine:~$ k get nodes
NAME       STATUS   ROLES           AGE     VERSION
minikube   Ready    control-plane   5d19h   v1.28.3
```

- When there is no scheduler, there would be empty Node value on pod description.

```py
# There is no aligned node to the pod.
$ k describe po nginx | grep Node

# There is no scheduler.
$ k get po -n kube-system | grep scheduler

# get nodes
$ k get nodes
```

- Then, if we want to manually scheule our pod, write nodeName to pod yaml

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx
spec:
  ...
  nodeName: node02
```

- After that delete and replace our pod resource to schedule on node02

```py
# kill pod and replace resource
k replace --force -f nginx.yaml
```

## Labels and Selectors

- Labels are key/value pairs that are attached to objects such as Pods
  - Unlike names and UIDS, labels do not provide uniqueness. In general, we expect many objects to carry the same labels.
- Via a label selector, the client/user can identify a set of objects. The label selector is the core grouping primitive in Kubernetes.



## Taints and Tolerations

- Node affinity: a property of Pods that attracts them to a set of nodes (either as a preference or a hard requirement).
- Taints: Taints are the opposite -- they allow a **node** to repel(격퇴하다) a set of pods.
- Tolerations: **Tolerations are applied to pods**. Tolerations allow the scheduler to schedule pods with matching taints. 

```py
kubectl taint nodes node1 key1=value1:NoExecute
kubectl taint nodes node1 key1=value1:NoSchedule
kubectl taint nodes node1 key1=value1:PreferNoSchedule
```

Taint Effect fields

- `NoExecute`
  - Pods that do not tolerate the taint are evicted immediately
  - Pods that tolerate the taint without specifying tolerationSeconds in their toleration specification remain bound forever
- `NoSchedule`
  - No new pods will be scheduled unless matching toleration (key1=value1)
  - Pods currently running on the node are not evicted.
- `PreferNoSchedule`
  - soft version of NoSchedule. The control plane will try to avoid but not guaranteed.

```
# Create a taint on node01 with key of spray, value of mortein and effect of NoSchedule
k taint nodes node01 spray=mortein:NoSchedule
```

## Node affinity

he primary feature of Node Affinity is to ensure that the pods are hosted on particular nodes.

- With Node Selectors we cannot provide the advance expressions.
  - e.g. A OR B, NOT A

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: myapp-pod
spec:
  containers:
  - name: data-processor
    image: data-processor
  affinity:
    nodeAffinity:
      requireDuringScedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: size
            opeator: In
            values:
            - Large
            - Medium
  
```


### Node Affinity Types
> https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/#types-of-inter-pod-affinity-and-anti-affinity

- Available
  - requiredDuringSchedulingIgnoredDuringExecution
  - preferredDuringSchedulingIgnoredDuringExecution

- Future plan
  - double require: requiredDuringSchedulingRequiredDuringExecution
  - prefer require: preferredDuringSchedulingRequiredDuringExecution


Wrap up the available node affinity types states

- **DuringScheduling: Required | Preferred**
- **DuringExecution: Ignored**



## Taints and tolerations and Node Affinity

![](https://github.com/kodekloudhub/certified-kubernetes-administrator-course/blob/master/images/tn-na.PNG?raw=true)

The combination of Taint + Tolearation can block other pod to be scheduled on tainted node, but cannot ensure that tolearated pod are being placed on the matching tainted node. so if that case we need affinity


![](https://github.com/kodekloudhub/certified-kubernetes-administrator-course/blob/master/images/tn-nsa.png?raw=true)

As such, a combination of taints and tolerations and node affinity rules can be used together to completely dedicate nodes for specific parts.

## cli

```py
# Open terminal output with vim to easily find `/` N/n
> k describe no node01 | vim -
```


```py
k get no --no-headers | wc -l

# set label to node
k label no node01 color=blue

k create deploy blue --image=nginx --replicas=3

k describe no controlplane | grep -i taints
Taints:             <none>
k describe no node01 | grep -i taints
Taints:             <none>
```


```py
k describe no controlplane 
Name:               controlplane
Roles:              control-plane
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=controlplane
                    kubernetes.io/os=linux
                    node-role.kubernetes.io/control-plane=
```

- label exists operator

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: red
  name: red
spec:
  replicas: 2
  selector:
    matchLabels:
      app: red
  template:
    metadata:
      labels:
        app: red
    spec:
      containers:
      - image: nginx
        name: nginx
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: node-role.kubernetes.io/control-plane
                operator: Exists
```


## Resource and Limit

- If the node where a Pod is running has enough of a resource available, it's possible (and allowed) for a container to use more resource than its request for that resource specifies. 
- However, a container is not allowed to use more than its resource limit.
- Kubelet and container runtime enforce the limit.

### Limit cpu vs Limit memory

- memory: oom kill 
- cpu: throttle

#### **memory**
> [when a process in the container tries to consume more than the allowed amount of memory, the system kernel terminates the process that attempted the allocation, **with an out of memory (OOM) error**.](https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/)

#### **cpu**

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: cpu-demo
  namespace: cpu-example
spec:
  containers:
  - name: cpu-demo-ctr
    image: vish/stress
    resources:
      limits:
        cpu: "1"
      requests:
        cpu: "0.5"
    args:
    - -cpus
    - "2"
```

> [Configured the Container to attempt to use 2 CPUs, but the Container is only being allowed to use about 1 CPU. The container's CPU use is being **throttled**, because the container is attempting to use more CPU resources than its limit.](https://kubernetes.io/docs/tasks/configure-pod-container/assign-cpu-resource/)


![](/images/resource_limit_cpu_memory.png)

- CPU: Request and No Limit is ideal

## Daemon Sets

A DaemonSet ensures that all or some Nodes run a copy of a Pod.

- As nodes are added to the cluster, Pods are added to them
- As nodes are removed from the cluster, those Pods are garbage collected

Some typical use of a DaemonSet are:

- running a cluster storage daemon on every node
- running a logs collection daemon on every node
- running a node monitoring daemon on every node

Also kube-proxy componenet can be deployed as DaemonSets

### How to create?

1. first create deployment --dry-run=client -o yaml > ds.yaml
2. delete status / replicas
3. and 

## Static pods
> https://kubernetes.io/docs/tasks/configure-pod-container/static-pod/

Static Pods are managed directly by the kubelt daemon on a specific node without the `kube-apiserver` observing them. Unlike Pods that are managed by the control plane(etcd, api, scheduler, controller manager ..); instead, the kubelet watches each static Pod.

- Static Pods are always bound to one Kubelt on a specific node.
- The kubelet automatically tries to create a mirror Pod on the kube-apiserver for each static Pod. 
- This means static pods running on a node are visible on the API server, but cannot be controlled from there.

> Mirror pod? A pod object that a kubelt uses to represent a static pod

- **Kubelet only can understand pod level**

#### Use Case

- kubeadm: Deploy control plane component as static Pods
  - kubeadm은 kubelet을 통해 `/etc/kubernetest/manifests` 안에 있는 control plane  component spec을 읽어 static pods들을 생성하여 관리한다.
- edge computing (iot)

#### Check wheter pod is static or not

There is two way

1. k get nodes && k get po -A
  - static pod naming: [POD NAME]-[NODE NAME]

2. k describe po [POD NAME] and check Owner: Node or other resource types


#### How to find staticPodPath
```py

controlplane ~ ➜  ps -aux | grep kubelet | grep -i config
root        4351  0.0  0.0 4519680 100556 ?      Ssl  02:23   0:36 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime-endpoint=unix:///var/run/containerd/containerd.sock --pod-infra-container-image=registry.k8s.io/pause:3.9

cat /var/lib/kubelet/config.yaml | grep staticPodPath
```

### How to create staticPod
> Create a static pod named static-busybox that uses the busybox image and the command sleep 1000

```py
controlplane ~ ➜  k run static-busybox --image=busybox --dry-run=client -o yaml > /etc/kubernetes/manifests/static-busybox.yaml

controlplane ~ ➜  vim /etc/kubernetes/manifests/static-busybox.yaml
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: static-busybox
  name: static-busybox
spec:
  containers:
  - image: busybox
    name: static-busybox
    command: ["sleep"]
    args:
    - "1000"
  dnsPolicy: ClusterFirst
  restartPolicy: Always
```


### How to find staticPod and delete it
> Question: We just created a new static pod named static-greenbox. Find it and delete it.
>> This question is a bit tricky. But if you use the knowledge you gained in the previous questions in this lab, you should be able to find the answer to it.


1. First, let's identify the node in which the pod called static-greenbox is created. To do this, run:

```py
root@controlplane:~# kubectl get pods --all-namespaces -o wide  | grep static-greenbox
default       static-greenbox-node01                 1/1     Running   0          19s     10.244.1.2   node01       <none>           <none>
root@controlplane:~#
```

From the result of this command, we can see that the pod is running on node01.

2. Next, SSH to node01 and identify the path configured for static pods in this node.

  - Important: The path need not be /etc/kubernetes/manifests. Make sure to check the path configured in the kubelet configuration file.

```py
root@controlplane:~# ssh node01 
root@node01:~# ps -ef |  grep /usr/bin/kubelet 
root        4147       1  0 14:05 ?        00:00:00 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime-endpoint=unix:///var/run/containerd/containerd.sock --pod-infra-container-image=registry.k8s.io/pause:3.9
root        4773    4733  0 14:05 pts/0    00:00:00 grep /usr/bin/kubelet

root@node01:~# grep -i staticpod /var/lib/kubelet/config.yaml
staticPodPath: /etc/just-to-mess-with-you

root@node01:~# 
```

Here the staticPodPath is /etc/just-to-mess-with-you


3. Navigate to this directory and delete the YAML file:

```py
root@node01:/etc/just-to-mess-with-you# ls
greenbox.yaml
root@node01:/etc/just-to-mess-with-you# rm -rf greenbox.yaml 
root@node01:/etc/just-to-mess-with-you#
```

4. Exit out of node01 using CTRL + D or type exit. You should return to the controlplane node. Check if the static-greenbox pod has been deleted:

```py
root@controlplane:~# kubectl get pods --all-namespaces -o wide  | grep static-greenbox
root@controlplane:~# 
```


# Chapter 5: Application Lifecycle Management


## Configmap

- Note that not to use `--from-file`, this is only handle single key like `--from-literal`
- Instead use `k create cm <NAME> --from-env-file=`

```py
controlplane ~ ➜  vim webapp.env 

controlplane ~ ➜  k create cm webapp-config-map --from-env-file=./webapp.env
configmap/webapp-config-map created
```

## Secret

```py
k create secret generic db-secret --from-env-file=./db.env
```

- `envFrom`

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: envfrom-secret
spec:
  containers:
  - name: envars-test-container
    image: nginx
    envFrom:
    - secretRef:
        name: test-secret
```

# Chapter 6: Cluster Maintenance


## Node upgrade (OS upgrade)

- drain: cordon + move resources
- uncordon: node ensable to be scheduled 
- cordon: node disable to be scheduled

```py
// Move every resources from node-1 to others
k drain node-1
// after node upgrade
// cordon: block node from scheduling
// uncordon: enable scheduling back
k uncordon node-1
```

Running the uncordon command on a node will not automatically schedule pods on the node. When new pods are created, they will be placed on node01.

> We will be upgrading the controlplane node first. Drain the controlplane node of workloads and mark it UnSchedulable

```py
> k drain node01 --ignore-daemonsets
node/node01 cordoned
Warning: ignoring DaemonSet-managed Pods: kube-flannel/kube-flannel-ds-rp464, kube-system/kube-proxy-8gmv5
evicting pod default/blue-667bf6b9f9-qm6x9
evicting pod default/blue-667bf6b9f9-hbzk9
pod/blue-667bf6b9f9-hbzk9 evicted
pod/blue-667bf6b9f9-qm6x9 evicted
node/node01 drained
```

There are daemonsets created in this cluster, especially in the kube-system namespace. To ignore these objects and drain the node, we can make use of the --ignore-daemonsets flag.

```py
$ k drain node01 --ignore-daemonsets --force
node/node01 already cordoned
Warning: deleting Pods that declare no controller: default/hr-app; ignoring DaemonSet-managed Pods: kube-flannel/kube-flannel-ds-rp464, kube-system/kube-proxy-8gmv5
evicting pod default/hr-app
```


> Question... I'm just curious that why there's still pod on drained node

```py
controlplane ~ ➜  k drain controlplane --ignore-daemonsets 
node/controlplane cordoned
Warning: ignoring DaemonSet-managed Pods: kube-flannel/kube-flannel-ds-9wfn6, kube-system/kube-proxy-x5qj8
evicting pod kube-system/coredns-5dd5756b68-l5w24
evicting pod kube-system/coredns-5dd5756b68-5nbck
evicting pod default/blue-667bf6b9f9-pxxm6
evicting pod default/blue-667bf6b9f9-m72qc
pod/blue-667bf6b9f9-m72qc evicted
pod/blue-667bf6b9f9-pxxm6 evicted
pod/coredns-5dd5756b68-5nbck evicted
pod/coredns-5dd5756b68-l5w24 evicted
node/controlplane drained

controlplane ~ ➜  k get no
NAME           STATUS                     ROLES           AGE   VERSION
controlplane   Ready,SchedulingDisabled   control-plane   29m   v1.28.0
node01         Ready                      <none>          29m   v1.28.0

controlplane ~ ➜  k get po -o wide
NAME                    READY   STATUS    RESTARTS   AGE     IP            NODE     NOMINATED NODE   READINESS GATES
blue-667bf6b9f9-987gj   1/1     Running   0          2m55s   10.244.1.4    node01   <none>           <none>
blue-667bf6b9f9-bcdtn   1/1     Running   0          2m55s   10.244.1.3    node01   <none>           <none>
blue-667bf6b9f9-gnlz5   1/1     Running   0          14s     10.244.1.10   node01   <none>           <none>
blue-667bf6b9f9-lgbg4   1/1     Running   0          14s     10.244.1.9    node01   <none>           <none>
blue-667bf6b9f9-tfcj2   1/1     Running   0          2m55s   10.244.1.2    node01   <none>           <none>

controlplane ~ ➜  k get po -A -o wide
NAMESPACE      NAME                                   READY   STATUS    RESTARTS   AGE    IP            NODE           NOMINATED NODE   READINESS GATES
default        blue-667bf6b9f9-987gj                  1/1     Running   0          3m2s   10.244.1.4    node01         <none>           <none>
default        blue-667bf6b9f9-bcdtn                  1/1     Running   0          3m2s   10.244.1.3    node01         <none>           <none>
default        blue-667bf6b9f9-gnlz5                  1/1     Running   0          21s    10.244.1.10   node01         <none>           <none>
default        blue-667bf6b9f9-lgbg4                  1/1     Running   0          21s    10.244.1.9    node01         <none>           <none>
default        blue-667bf6b9f9-tfcj2                  1/1     Running   0          3m2s   10.244.1.2    node01         <none>           <none>
kube-flannel   kube-flannel-ds-4krgp                  1/1     Running   0          29m    192.20.38.9   node01         <none>           <none>
kube-flannel   kube-flannel-ds-9wfn6                  1/1     Running   0          29m    192.20.38.6   controlplane   <none>           <none>
kube-system    coredns-5dd5756b68-7dv8x               1/1     Running   0          21s    10.244.1.8    node01         <none>           <none>
kube-system    coredns-5dd5756b68-ptnml               1/1     Running   0          21s    10.244.1.11   node01         <none>           <none>
kube-system    etcd-controlplane                      1/1     Running   0          29m    192.20.38.6   controlplane   <none>           <none>
kube-system    kube-apiserver-controlplane            1/1     Running   0          29m    192.20.38.6   controlplane   <none>           <none>
kube-system    kube-controller-manager-controlplane   1/1     Running   0          29m    192.20.38.6   controlplane   <none>           <none>
kube-system    kube-proxy-bwrl5                       1/1     Running   0          29m    192.20.38.9   node01         <none>           <none>
kube-system    kube-proxy-x5qj8                       1/1     Running   0          29m    192.20.38.6   controlplane   <none>           <none>
kube-system    kube-scheduler-controlplane            1/1     Running   0          29m    192.20.38.6   controlplane   <none>           <none>
```

## Kubernetes Software Versions


```py
# get kubectl version
k version

# get kubeadm upgrade plan
kubeadm upgrade plan
```

You can find all kubernetes releases at https://github.com/kubernetes/kubernetes/releases.
Downloaded package has all the kubernetes components in it except `ETCD cluster` and `CoreDNS` as they are seperate projects.


## Cluster Upgrade Introduction

> Q. Is it mandatory for all of the kubernetes components to have the same versions?

No, The components can be at different release versions. At any time, **kubernetes supports only up to the recent 3 minor versions**, and the recommended approach is to upgrade one minor version at a time, instead of upgrading all 3 steps at once.

![](https://github.com/kodekloudhub/certified-kubernetes-administrator-course/raw/master/images/up2.PNG)



## Upgrading kubeadm clusters
> https://v1-29.docs.kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/


### 1. Upgrade kubeadm master node

0. k drain <node> --ignore-daemonsets

1. (opt) [Update package repository](https://v1-29.docs.kubernetes.io/docs/tasks/administer-cluster/kubeadm/change-package-repository/)

```py
> pager /etc/apt/sources.list.d/kubernetes.list

deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyri
ng.gpg] https://pkgs.k8s.io/core:/stable:/v1.28/deb/ 
/
```

Switching to another Kubernetes package repository 

```py
vim /etc/apt/sources.list.d/kubernetes.list

# change version v1.28 -> v1.29
:s/v1.28/v1.29/g
```

2. Determine which version to upgrade to

```py
> sudo apt update
> sudo apt-cache madison kubeadm


   kubeadm | 1.29.4-2.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   kubeadm | 1.29.3-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   kubeadm | 1.29.2-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   kubeadm | 1.29.1-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
   kubeadm | 1.29.0-1.1 | https://pkgs.k8s.io/core:/stable:/v1.29/deb  Packages
```

3. Upgrading control plane nodes


- upgrade kubeadm
```py
> sudo apt-mark unhold kubeadm && \
sudo apt-get update && sudo apt-get install -y kubeadm=1.29.0-1.1 && \
sudo apt-mark hold kubeadm
```

- verify the upgrade plan

```py
target_version=v1.29.0

> sudo kubeadm upgrade plan $target_version

[upgrade/config] Making sure the configuration is correct:
[upgrade/config] Reading configuration from the cluster...
[upgrade/config] FYI: You can look at this config file with 'kubectl -n kube-system get cm kubeadm-config -o yaml'
[preflight] Running pre-flight checks.
[upgrade] Running cluster health checks
[upgrade] Fetching available versions to upgrade to
[upgrade/versions] Cluster version: v1.28.0
[upgrade/versions] kubeadm version: v1.29.0
[upgrade/versions] Target version: v1.29.0
[upgrade/versions] Latest version in the v1.28 series: v1.29.0

Components that must be upgraded manually after you have upgraded the control plane with 'kubeadm upgrade apply':
COMPONENT   CURRENT       TARGET
kubelet     2 x v1.28.0   v1.29.0

Upgrade to the latest version in the v1.28 series:

COMPONENT                 CURRENT   TARGET
kube-apiserver            v1.28.0   v1.29.0
kube-controller-manager   v1.28.0   v1.29.0
kube-scheduler            v1.28.0   v1.29.0
kube-proxy                v1.28.0   v1.29.0
CoreDNS                   v1.10.1   v1.11.1
etcd                      3.5.9-0   3.5.10-0

You can now apply the upgrade by executing the following command:

        kubeadm upgrade apply v1.29.0

_____________________________________________________________________


The table below shows the current state of component configs as understood by this version of kubeadm.
Configs that have a "yes" mark in the "MANUAL UPGRADE REQUIRED" column require manual config upgrade or
resetting to kubeadm defaults before a successful upgrade can be performed. The version to manually
upgrade to is denoted in the "PREFERRED VERSION" column.

API GROUP                 CURRENT VERSION   PREFERRED VERSION   MANUAL UPGRADE REQUIRED
kubeproxy.config.k8s.io   v1alpha1          v1alpha1            no
kubelet.config.k8s.io     v1beta1           v1beta1             no
_____________________________________________________________________

```

- choose a version to upgrade and apply
```py

> sudo kubeadm upgrade apply $target_version

```

Now, upgrade the version and restart Kubelet. Also, mark the node (in this case, the "controlplane" node) as schedulable.

```py
> sudo apt-mark unhold kubelet kubectl && \
sudo apt-get update && sudo apt-get install -y kubelet='1.29.0-1.1' kubectl='1.29.0-1.1' && \
sudo apt-mark hold kubelet kubectl

> sudo systemctl daemon-reload
> sudo systemctl restart kubelet
> sudo kubectl uncordon controlplane
```

### 2. Upgrade kubeadm worker node

1. mirror update


2. https://v1-29.docs.kubernetes.io/docs/tasks/administer-cluster/kubeadm/upgrading-linux-nodes/

```py
sudo apt-mark unhold kubeadm && \
sudo apt-get update && sudo apt-get install -y kubeadm='1.29.0-1.1' && \
sudo apt-mark hold kubeadm
```

3. kubeadm upgrade node (instead apply)

```py
> sudo kubeadm upgrade node
```

4. updgrade kubelet and kubectl

```py
sudo apt-mark unhold kubelet kubectl && \
sudo apt-get update && sudo apt-get install -y kubelet='1.29.0-1.1' kubectl='1.29.0-1.1' && \
sudo apt-mark hold kubelet kubectl

sudo systemctl daemon-reload
sudo systemctl restart kubelet


exit (back to master)
kubectl uncordon node01
```


### 3. Problem: backup and restore etcd
> Q. An ETCD backup for cluster2 is stored at /opt/cluster2.db. Use this snapshot file to carryout a restore on cluster2 to a new path /var/lib/etcd-data-new. Once the restore is complete, ensure that the controlplane components on cluster2 are running. The snapshot was taken when there were objects created in the critical namespace on cluster2. These objects should be available post restore.

### Solution
Step 1. Copy the snapshot file from the student-node to the etcd-server. In the example below, we are copying it to the /root directory:

```py
student-node ~  scp /opt/cluster2.db etcd-server:/root
cluster2.db                                                                                                        100% 1108KB 178.5MB/s   00:00    

student-node ~ ➜  
```

Step 2: Restore the snapshot on the cluster2. Since we are restoring directly on the etcd-server, we can use the endpoint https:/127.0.0.1. Use the same certificates that were identified earlier. Make sure to use the data-dir as /var/lib/etcd-data-new:

```py
etcd-server ~ ➜  ETCDCTL_API=3 etcdctl --endpoints=https://127.0.0.1:2379 --cacert=/etc/etcd/pki/ca.pem --cert=/etc/etcd/pki/etcd.pem --key=/etc/etcd/pki/etcd-key.pem snapshot restore /root/cluster2.db --data-dir /var/lib/etcd-data-new
{"level":"info","ts":1662004927.2399247,"caller":"snapshot/v3_snapshot.go:296","msg":"restoring snapshot","path":"/root/cluster2.db","wal-dir":"/var/lib/etcd-data-new/member/wal","data-dir":"/var/lib/etcd-data-new","snap-dir":"/var/lib/etcd-data-new/member/snap"}
{"level":"info","ts":1662004927.2584803,"caller":"membership/cluster.go:392","msg":"added member","cluster-id":"cdf818194e3a8c32","local-member-id":"0","added-peer-id":"8e9e05c52164694d","added-peer-peer-urls":["http://localhost:2380"]}
{"level":"info","ts":1662004927.264258,"caller":"snapshot/v3_snapshot.go:309","msg":"restored snapshot","path":"/root/cluster2.db","wal-dir":"/var/lib/etcd-data-new/member/wal","data-dir":"/var/lib/etcd-data-new","snap-dir":"/var/lib/etcd-data-new/member/snap"}

etcd-server ~ ➜  
```

Step 3: Update the systemd service unit file for etcdby running vi /etc/systemd/system/etcd.service and add the new value for data-dir:

```
[Unit]
Description=etcd key-value store
Documentation=https://github.com/etcd-io/etcd
After=network.target

[Service]
User=etcd
Type=notify
ExecStart=/usr/local/bin/etcd \
  --name etcd-server \
  --data-dir=/var/lib/etcd-data-new \
---End of Snippet---
```

Step 4: make sure the permissions on the new directory is correct (should be owned by etcd user):

```py
etcd-server /var/lib ➜  chown -R etcd:etcd /var/lib/etcd-data-new

etcd-server /var/lib ➜ 


etcd-server /var/lib ➜  ls -ld /var/lib/etcd-data-new/
drwx------ 3 etcd etcd 4096 Sep  1 02:41 /var/lib/etcd-data-new/
etcd-server /var/lib ➜ 
```

Step 5: Finally, reload and restart the etcd service.

```py
etcd-server ~/default.etcd ➜  systemctl daemon-reload 
etcd-server ~ ➜  systemctl restart etcd
etcd-server ~ ➜  
```

Step 6 (optional): It is recommended to restart controlplane components (e.g. kube-scheduler, kube-controller-manager, kubelet) to ensure that they don't rely on some stale data.

```py

student-node ~ ✖ k delete po kube-apiserver-cluster2-controlplane kube-controller-manager-cluster2-controlplane kube-scheduler-cluster2-controlplane -n kube-system 
pod "kube-apiserver-cluster2-controlplane" deleted
pod "kube-controller-manager-cluster2-controlplane" deleted
pod "kube-scheduler-cluster2-controlplane" deleted

ssh cluster2-controlplane

cluster2-controlplane ~ ✖ systemctl restart kubelet
cluster2-controlplane ~ ➜  systemctl status kubelet
● kubelet.service - kubelet: The Kubernetes Node Agent
   Loaded: loaded (/lib/systemd/system/kubelet.service; enabled; vendor preset: enabled)
  Drop-In: /etc/systemd/system/kubelet.service.d
           └─10-kubeadm.conf
   Active: active (running) since Sat 2024-04-27 13:51:11 UTC; 6s ago
```



# Chapter 7: Security
> https://kubernetes.io/docs/concepts/security/

## Authentication

- Account
- Service accuount

All the user authentication is managed by kube-apiserver, authenticate



## TLS Certificates

- key, pem = public, priviate
- certificate
- certificate authority (CA)
- Certificate Signning Request (CSR): with public key
- PKI (public Key infrastructure)

### Format

- Certificate (pulic key)
  - ***.crt, *.pem**
  - i.g. server.crt, server.pem, client.crt, client.pem
- Private key
  - ***.key, *-key.pem**
  - i.g. server.key, server-key.pem, client.key, client-key.pem


When server requests to CSR to certify server with CSR, CA will verify the request and if it passed, then CA encrypt request(+server's pub key) with CA's private key and return to server and finally return to client.

Client especially browser has CA's public key so that the browsers uses the public key of the Certificate Authority to validate the certificate was actually signed by the Verified Certificate Authority themselves.

> Q. But does the public key is only for encryption not decryption then how does browser validate ca's encrypted data is valid or not?

In public key cryptography, the public key is indeed primarily used for encryption, **but it also has a crucial role in verifying digital signatures**. Digital signatures are created by encrypting a hash (a unique fingerprint) of the data using the private key. The resulting encrypted hash, along with the data, forms the digital signature.

To verify the digital signature, the recipient uses the public key associated with the private key used to create the signature. This process works as follows:

1. The recipient uses the public key to decrypt the digital signature, resulting in the original hash value.
2. The recipient independently computes the hash of the received data.
3. If the decrypted hash matches the independently computed hash, the signature is valid. This indicates that the data hasn't been altered since it was signed and that the signature was indeed created with the private key associated with the public key used for decryption.

In the context of SSL/TLS certificates:

- The Certificate Authority (CA) signs the digital certificate with its private key, creating a digital signature.
- Your browser, possessing the CA's public key, decrypts the digital signature to obtain the hash of the certificate.
- The browser then independently computes the hash of the certificate data.
- If the decrypted hash matches the computed hash, the browser knows that the certificate is authentic and was indeed issued by the CA. (decrypt라는 용어가 광범위하게 잘못사용되는 것도 한 몫하는 것 같다.)

So, the CA's public key is used to validate the CA's digital signature on the certificate by decrypting the signature and verifying its integrity, ensuring that it was signed with the CA's private key.

> Q. Then how computed hash can ensures whether the signature is valid or not?

public key cryptography: A class of cryptographic techniques employing two-key ciphers. **Messages encrypted with the public key can only be decrypted with the associated private key. Conversely, messages signed with the private key can be verified with the public key.**

1. CA public key로 valid 하다고 판결하면, server를 신뢰하고, 서버의 public key를 신뢰해서, 이를 통해서 symm key를 encrypt해서 server에 보낸다. 중간에 이를 가로채는 것은 서버가 정상 서비스 업체라고 가정한다면, private key를 마음대로 사용하지 않을테니 중간에 가로챈 사람들을 symm key를 복호화 할 수 없다.

## TLS in kubernetes

- Root Certificates (CA)
- Server Certificates (server)
- Client Certificates (client)


![](/images/k8s-tls.png)

### Server Certificates for servers

- KUBE-API server
  - apiserver.crt
  - apiserver.key
- ETCD server
  - etcdserver.crt
  - etcdserver.key
- KUBELET server (node01, node02 ....)
  - node01.crt
  - node01.key

### Client Certificates for clients

- admin user: kubectl REST API to access kube-api server
  - admin.crt
  - admin.key
- KUBE SCHEDULER: to access kube-api server API
  - scheduler.crt
  - scheduler.key
- KUBE CONTROLLER-MANAGER: to access kube-api server api
  - controller-manager.crt
  - controller-manager.key
- KUBE-PROXY
  - kube-proxy.crt
  - kube-proxy.key

![](/images/k8s-tls2.png)

## Kube config

Kubeconfig is the file to organize information about clusters, users, namespaces and authentication mechanisms.

kubectl command line tool uses kubeconfig file to find the information it needs to choose a cluster and communicate with API server of a cluster. Which  measn kubeconfig contains

- kind: Config
- cluster
- contexts: mapping with cluster and users
  - namespace
- users
  - crt files

```py
> k config view
> k config use-context
```

Base directory is `$HOME/.kube/config`. 

`echo env | grep -i HOME` to figure out home

## Authorization

1. ABAC: user당 json형식으로 kind: policy를 관리하는 방식
2. RBAC: policy group을 rule로 묶고, user와 binding하는 방식
3. Webhook



```py
# kubeadm 

cat /etc/kubernetes/manifests/kube-apiserver.yaml | grep -i authorizaion

...elipsis
--authorization-mode=Node,RBAC,Webhook
```


Above setting means first Node auth try and RBAC try and Webhook try


## RBAC

- Role and Role Bindings are under the scope of namespaces
  - so withoud namespace = default namespace role
  - with namespace -> specific namespace


```py
# check Access
k auth can-i create deployments --as dev-user -n production

k auth can-i list nodes --as michelle
```


## Cluster Roles

Unlike Role, Cluster-Role's scope is not limited by namespace. In other words, it applies to all namespaces.

- Namespaced
  - pods, replicasets, jobs, deployments, services, secrets, roles, rolebindings configmaps, pvc
  - `k api-resources --namespaced=true`
- Cluster Scoped
  - nodes, PV, clusterrole, clusterrolebindings, certificatesigningrequests(csr), namespaces
  - `k api-resources --namespaced=false`

```py

$ k create clusterrole storage-admin --resource=persistentvolumes,storageclasses --verb=*
$ k get clusterrole storage-admin -o yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: "2024-05-04T06:22:03Z"
  name: storage-admin
  resourceVersion: "1696"
  uid: d73ce51f-6ea7-4dfa-9db9-2564e3fa277c
rules:
- apiGroups:
  - ""
  resources:
  - persistentvolumes
  verbs:
  - '*'
- apiGroups:
  - storage.k8s.io
  resources:
  - storageclasses
  verbs:
  - '*'
```

## Service account

![](/images/k8s-ua-vs-sa.png)

The user account is literally an account for users, and the service account is literally an account for services such as Prometheus, Grafana, Kubeflow..


```py
controlplane ~ ➜  k get sa -n default
NAME      SECRETS   AGE
default   0         11m
dev       0         52s

controlplane ~ ➜  k create sa test-sa
serviceaccount/test-sa created

controlplane ~ ➜  k get sa
NAME      SECRETS   AGE
default   0         11m
dev       0         85s
test-sa   0         3s

controlplane ~ ➜  k describe sa test-sa
Name:                test-sa
Namespace:           default
Labels:              <none>
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   <none>
Tokens:              <none>
Events:              <none>
```

- Every namespace has it's own default service account

```py
controlplane ~ ➜  k describe sa default
Name:                default
Namespace:           default
Labels:              <none>
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   <none>
Tokens:              <none>
Events:              <none>

controlplane ~ ➜  k get po
No resources found in default namespace.

```

- default sa is used when, create resources without sa

```py
controlplane ~ ➜  k run nginx --image=ngin
pod/nginx created


controlplane ~ ➜  k describe po nginx | grep -i default
Namespace:        default
Service Account:  default
  Normal   Scheduled  104s                default-scheduler  Successfully assigned default/nginx to controlplane


- Default volume and mount are automatically created at `/var/run/secrets/*`.
- Kubernetes automatically mounts the default service account token to inside pode.
- ServiceAccount token Secrets store credentials identifying a ServiceAccount for Pods.
- Legacy method providing long-lived credentials, but In Kubernetes v1.22+, recommended to obtain short-lived, rotating tokens using `TokenRequest API`.
- Methods to obtain short-lived tokens include direct API calls or through kubectl.



```py
controlplane ~ ➜  k describe po nginx

...elipsis

Containers:
  nginx:
    Container ID:   
    Image:          ngin
    Image ID:       
    Port:           <none>
    Host Port:      <none>
    State:          Waiting
      Reason:       ErrImagePull
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-2pp72 (ro)
Volumes:
  kube-api-access-2pp72:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true      
```

We can check it like this

```py
controlplane ~ ✖ k exec -it nginx -- ls /var/run/secrets/kubernetes.io/serviceaccount
ca.crt  namespace  token
controlplane ~ ➜  k exec -it nginx -- cat /var/run/secrets/kubernetes.io/serviceaccount/token
eyJhbGciOiJSUzI1NiIsImtpZCI6Im1VUi1aZllLak5Fc0MyV184aE1ycDVvcmN6Z1pnNWZlT1JGLXFxQ2NtTm8ifQ.eyJhdWQiOlsiaHR0cHM6Ly9rdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLmNsdXN0ZXIubG9jYWwiLCJrM3MiXSwiZXhwIjoxNzQ2MzQyNzYzLCJpYXQiOjE3MTQ4MDY3NjMsImlzcyI6Imh0dHBzOi8va3ViZXJuZXRlcy5kZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsIiwia3ViZXJuZXRlcy5pbyI6eyJuYW1lc3BhY2UiOiJkZWZhdWx0IiwicG9kIjp7Im5hbWUiOiJuZ2lueCIsInVpZCI6ImRjNjgzOTI0LTZmYjAtNGFlOC04YjNmLTRlZmU5NmEwMjI0YSJ9LCJzZXJ2aWNlYWNjb3VudCI6eyJuYW1lIjoiZGVmYXVsdCIsInVpZCI6IjgxY2M0MDU1LTBhNjAtNDNiNC05ODE3LTk3OWIwYTUwODA3MiJ9LCJ3YXJuYWZ0ZXIiOjE3MTQ4MTAzNzB9LCJuYmYiOjE3MTQ4MDY3NjMsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDpkZWZhdWx0OmRlZmF1bHQifQ.KOu1urbQ_aqLSDe0tOChy9ZbEmeUFkf-yDKU0TPeK_zKqi1tbZRGV4pVY_6ac90ZCeWTTu2hA1jsUYFTLVyfWfCy7jb7H7BR3gBVrMnUncSYIbjGeZNuJK_3JJ_xaSN3cuKJyJbK19cQG19pACkp3TvPxXfmdAKFcuGXdpvp9m4vXGGHV4zMKebStuk5guhKyDsVQycoLSTse4mUohARPRb8BFTNcSTwUHaQ2crTo-FBa46XbwQUkQt0JTIeCijXb9cRe0zvuiZggUEv2i8BBVb6G6OAVt9n5uzYAv4WLSSF8ovfygbQhivq5U-BJP7B85IaEKwuejLDzMthXa9ioA
```


![](/images/k8s-sa-policy.png)

Unlike the existing method on the left where a secret and a token were created when creating sa, from 1.24 onwards, you must explicitly create a token instead of creating a secret to confirm it. Also, if necessary, you must create a secret as shown below and connect to sa.

![](/images/k8s-sa-policy2.png)


> [But, Note: You should only create a ServiceAccount token Secret if you can't use the TokenRequest API to obtain a token, and the security exposure of persisting a non-expiring token credential in a readable API object is acceptable to you. For instructions, see Manually create a long-lived API token for a ServiceAccount.](https://kubernetes.io/docs/concepts/configuration/secret/#serviceaccount-token-secrets)


## Image security

```py
root@controlplane ~ ➜  k create secret -h
Create a secret with specified type.

 A docker-registry type secret is for accessing a container registry.

 A generic type secret indicate an Opaque secret type.

 A tls type secret holds TLS certificate and its associated key.

Available Commands:
  docker-registry   Create a secret for use with a Docker registry
  generic           Create a secret from a local file, directory, or literal value
  tls               Create a TLS secret
```


Create a secret object with the credentials required to access the registry.

```
Name: private-reg-cred
Username: dock_user
Password: dock_password
Server: myprivateregistry.com:5000
Email: dock_user@myprivateregistry.com
```

```py
root@controlplane ~ ➜  k create secret docker-registry private-reg-cred --docker-username=dock_user --docker_email=dock_user@myprivateregistry.com --docker-password=dock_password --docker-server=myprivateregistry.com:5000
secret/private-reg-cred created

root@controlplane ~ ➜  k get secret
NAME               TYPE                             DATA   AGE
private-reg-cred   kubernetes.io/dockerconfigjson   1      6s

root@controlplane ~ ✖ k describe secret private-reg-cred
Name:         private-reg-cred
Namespace:    default
Labels:       <none>
Annotations:  <none>

Type:  kubernetes.io/dockerconfigjson

Data
====
.dockerconfigjson:  176 bytes

```

- edit deployment

```py
> root@controlplane ~ ➜  k edit deployments.apps web 

    spec:
      containers:
      - image: myprivateregistry.com:5000/nginx:alpine
        imagePullPolicy: IfNotPresent
        imagePullSecrets: private-reg-cred


error: deployments.apps "web" is invalid
A copy of your changes has been stored to "/tmp/kubectl-edit-4288536569.yaml"
error: Edit cancelled, no valid changes were saved.

> vim /tmp/kubectl-edit-4288536569.yaml

    spec:
      imagePullSecrets:
      - name: private-reg-cred
      containers:
      - image: myprivateregistry.com:5000/nginx:alpine
        imagePullPolicy: IfNotPresent

> k apply -f /tmp/kubectl-edit-4288536569.yaml
```

## Security Context
> https://kubernetes.io/docs/tasks/configure-pod-container/security-context/

A security context defines privilege and access control settings for a Pod or Container.

- `pod.spec.securityContext`
- `pod.spec.containers.securityContext`

#### Attributes

- `runAsGroup`
- `runAsUser`
- `capabilities`: Adds and removes POSIX capabilities from running containers.

[The runAsGroup field specifies the primary group ID of <runAsGroup_value> for all processes within any containers of the Pod. If this field is omitted, the primary group ID of the containers will be root(0).](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/#set-the-security-context-for-a-pod)


```py
controlplane ~ ➜  ps -ef | grep -i sleep
 6747 root      0:00 sleep 4800
 7688 root      0:00 grep -i sleep

controlplane ~ ➜  k exec -it ubuntu-sleeper -- ps -ef | grep -i sleep
root           1       0  0 09:17 ?        00:00:00 sleep 4800
```


#### How to check capabilities

```py
controlplane ~ ➜  k exec -it ubuntu-sleeper -- sh
# ls
bin  boot  dev  etc  home  lib  lib64  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var
# ps -aux
USER         PID %CPU %MEM    VSZ   RSS TTY      STAT START   TIME COMMAND
root           1  0.0  0.0   2692  1120 ?        Ss   04:13   0:00 sleep 4800
root          56  0.0  0.0   2796  1096 pts/0    Ss   04:15   0:00 sh
root          63  0.0  0.0   7884  4112 pts/0    R+   04:15   0:00 ps -aux
> cd /proc/1
> cat status | grep -i Cap
CapInh: 0000000000000000
CapPrm: 00000000a80425fb
CapEff: 00000000a80425fb
CapBnd: 00000000a80425fb
CapAmb: 0000000000000000

# edit pod.spec.securityContext.capabilities
k edit po ubuntu-sleeper
```


## Network policies
> https://kubernetes.io/docs/concepts/services-networking/network-policies/

- By default, a pod is non-isolated for egress and ingress, which means default is all allowed
- Only traffic flows marked as "Ingress" or "Egress" in `spec.PolicyTypes` are impacted by network policy.
- If ingress is allowed, a response to the request is automatically delivered.


![](/images/k8s-network-policy1.png)


![](/images/k8s-network-policy2.png)


> [kubernetes-network-policy-recipes](https://github.com/ahmetb/kubernetes-network-policy-recipes/blob/master/09-allow-traffic-only-to-a-port.md)

```yaml
kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: api-allow-5000
spec:
  podSelector:
    matchLabels:
      app: apiserver
  ingress:
  - ports:
    - port: 5000
    from:
    - podSelector:
        matchLabels:
          role: monitoring
```

- Drop all non-whitelisted traffic to app=apiserver.
- **Allow traffic on port 5000 from pods with label role=monitoring in the same namespace.**


#### Problem
> Use the spec given below. You might want to enable ingress traffic to the pod to test your rules in the UI.Also, ensure that you allow egress traffic to DNS ports TCP and UDP (port 53) to enable DNS resolution from the internal pod.

```
Policy Name: internal-policy
Policy Type: Egress
Egress Allow: payroll
Payroll Port: 8080
Egress Allow: mysql
MySQL Port: 3306

controlplane ~ ➜  k get all
NAME           READY   STATUS    RESTARTS   AGE
pod/external   1/1     Running   0          23m
pod/internal   1/1     Running   0          23m
pod/mysql      1/1     Running   0          23m
pod/payroll    1/1     Running   0          23m

NAME                       TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/db-service         ClusterIP   10.97.124.55     <none>        3306/TCP         23m
service/external-service   NodePort    10.102.212.250   <none>        8080:30080/TCP   23m
service/internal-service   NodePort    10.107.186.112   <none>        8080:30082/TCP   23m
service/kubernetes         ClusterIP   10.96.0.1        <none>        443/TCP          105m
service/payroll-service    NodePort    10.103.183.130   <none>        8080:30083/TCP   23m
```

#### Solution

```py
k get netpol payroll-policy -o yaml > netpol.yaml
```

- netpol.yaml
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: internal-policy
spec:
  egress:
  - to:
    - podSelector:
        matchLabels:
          name: payroll
    ports:
    - port: 8080
      protocol: TCP
  - to:
    - podSelector:
        matchLabels:
          name: mysql
    ports:
    - port: 3306
      protocol: TCP
  - ports:
    - port: 53
      protocol: UDP
    - port: 53
      protocol: TCP
  podSelector:
    matchLabels:
      name: internal
  policyTypes:
  - Egress
```

```py
k apply -f netpol.yaml
```



# Chapter 8: Storage


- [Docker: Storage drivers versus Docker volumes](https://docs.docker.com/storage/storagedriver/#storage-drivers-versus-docker-volumes)

As we start learning about containerization, we come across two important parts: storage drivers and volumes. These are crucial for managing data in Docker setups. Storage drivers handle storing image layers and container data, aiming for efficiency while considering performance. Volumes, on the other hand, provide a way to store data persistently, connecting the temporary nature of containers with the need for lasting data storage and sharing. In this discussion, we'll explore the key differences between storage drivers and volumes, understand why memory efficiency and Copy-on-Write (CoW) mechanisms matter for storage drivers, and see how volume drivers help keep data stored across container lifetimes. Let's summarize Docker storage and volume together.

Containers are designed with a stateless assumption to facilitate reuse. Docker, not only for container reuse but also for efficient image building, divides images into layers. These layers enable reuse as the created image is read-only (RO).

However, containers created from images can write and modify files. This occurs through Copy-on-Write (COW) functionality, where changes made in the layer trigger a copy and write process in the background. Storage drivers handle this role.

Storage drivers need to be memory efficient and support COW, but they can impact performance, especially with large data IO operations like databases (db). To address this, volumes are used. Volumes ensure persistence beyond container lifespan and enable data sharing between containers. Depending on the driver, data can be stored in cloud or node hosts.

### ko
```ko
- container는 재사용하기 위해 stateless를 가정한다.
- docker는 container 재사용 뿐 아니라, image build 효율을 위해 image 각각에도 layer를 나눠서 관리한다.
- 만들어진 image는 RO(stateless)이기 때문에 재사용 가능하다.
- 하지만 image를 통해 만들어진 container는 파일을 쓰고 고칠수도 있다. 이는 COW, 실제 해당 layer에서 Write이 일어날때 copy 후 write하기 때문이며, 이 기능은 background에서 일어난다. 이 역할을 하는 것이 storage driver이다.
- storage driver는 기본적으로 memory efficiend하게 만들어져야 하며, cow를 제공해야 하기 때문에 큰 데이터 IO의 경우(db) 성능에 안좋은 영향을 받는다.
- 그렇기 때문에 Volume을 사용한다. volume은 persist를 위해 사용하며, container lifespan을 벗어난 데이터들을 관리할 수도 있고, container간에 share도 가능하다. driver에 따라서 cloud, node host에 저장할 수도 있다.
```


## CSI (container storage interface)

- CRI: container runtime interface
- CNI: container network interface
- [CSI](https://github.com/container-storage-interface/spec/blob/master/spec.md): container storage interface

## Volume

Kubernetes volumes preserve container data, preventing loss on crashes, enabling file sharing in Pods, and enhancing application management by providing directories accessible to containers with various types, including ephemeral and persistent, ensuring data persistence across restarts.

## Persistent Volume

Volume have to write spec on pod definition, but when system is bigger than It is hard to manage all pods to mapping each volumes. So persistent volume and Persistent volume claim concepts are invented.


> Configure a volume to store these logs at /var/log/webapp on the host.

- spec

```
Name: webapp

Image Name: kodekloud/event-simulator

Volume HostPath: /var/log/webapp

Volume Mount: /log
```

```py
controlplane ~ ➜  k get all
NAME         READY   STATUS    RESTARTS   AGE
pod/webapp   1/1     Running   0          5m26s

NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.96.0.1    <none>        443/TCP   11m
```

```py
> k get po webapp -o yaml > webapp.yaml

# edit
> k replace -f 
```

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: webapp
spec:
  containers:
  - name: event-simulator
    image: kodekloud/event-simulator
    env:
    - name: LOG_HANDLERS
      value: file
    volumeMounts:
    - mountPath: /log
      name: log-volume

  volumes:
  - name: log-volume
    hostPath:
      # directory location on host
      path: /var/log/webapp
      # this field is optional
      type: Directory
```

- Should volumeMounts.name == volumes.name
- not hostpath, hostPath


```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-log
spec:
  persistentVolumeReclaimPolicy: Retain
  accessModes:
    - ReadWriteMany
  capacity:
    storage: 100Mi
  hostPath:
    path: /pv/log
```

- [**Claims use the same conventions as volumes when requesting storage with specific access modes.**](https://kubernetes.io/docs/concepts/storage/persistent-volumes/#access-modes-1)


## Persistent Volume Claims

1. PVC created
2. k8s search for matching pv (pending state pvc)
3. Bind

## Storage Class

**It would be great if make pv then physical memory provisioned, which is called `Dynamic Provisioning`. And we can achieve it by `Storage class(sc)` instead PV. Actually pv is created when sc definition is called but we do not manually create pv, because we need to manage volume dynamically.**


## Wrap up

1. PV: provision된 physical volume을 관리하는 instance  (node처럼 cluster resource)
2. PVC: volume를 쓰기위한 요청. **pod는 pvc를 통해 pv에 리소스 사용**
3. StorageClass(sc)
  - PVC와 PV는 같은 storageClassName일 때 bind
  - 물론 default storageClass가 있으면 storageClassName omit가능

**PV를 프로비저닝 하는 방법은 총 2가지이다.**

1. Static Provisioning
  - kube-apisesrver를 통해서 (kubectl) admin(관리자)이 직접 pv를 생성하는 방식
2. Dynamic Provisioning
  - storage class를 기반으로 자동으로 pv를 생성하여 provision하는 방식
  - **PVC가 특정 spec을 정의한 storage class를 요청하면 pvc의 양만큼 sc에 정의된 스펙의 pv가 생성되고 프로비저닝 된다.**



# Chapter 9: Networking


## Switching Routing

#### Switching

```py
$ ip addr
$ ip address

controlplane ~ ➜  ip address
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
2: flannel.1: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1400 qdisc noqueue state UNKNOWN group default 
    link/ether 7e:f7:9a:1e:dc:79 brd ff:ff:ff:ff:ff:ff
    inet 10.244.0.0/32 scope global flannel.1
       valid_lft forever preferred_lft forever
3: cni0: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1400 qdisc noqueue state UP group default qlen 1000
    link/ether 8e:a7:28:ce:4f:34 brd ff:ff:ff:ff:ff:ff
    inet 10.244.0.1/24 brd 10.244.0.255 scope global cni0
       valid_lft forever preferred_lft forever
4: veth94e4cabd@if2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1400 qdisc noqueue master cni0 state UP group default 
    link/ether ca:df:3d:76:bb:96 brd ff:ff:ff:ff:ff:ff link-netns cni-cf9503cc-e140-55a0-5d1b-f0475dc95336
5: veth58ee9409@if2: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1400 qdisc noqueue master cni0 state UP group default 
    link/ether b2:e3:5b:bd:ca:14 brd ff:ff:ff:ff:ff:ff link-netns cni-908defd9-94bd-1a21-0282-6ad3f832bc73
13028: eth0@if13029: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1450 qdisc noqueue state UP group default 
    link/ether 02:42:c0:00:f3:09 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 192.0.243.9/24 brd 192.0.243.255 scope global eth0
       valid_lft forever preferred_lft forever
13030: eth1@if13031: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default 
    link/ether 02:42:ac:19:00:15 brd ff:ff:ff:ff:ff:ff link-netnsid 1
    inet 172.25.0.21/24 brd 172.25.0.255 scope global eth1
       valid_lft forever preferred_lft forever
```


- eth0: 과거 디폴트 랜카드 지정 번지 (이더넷 카드 0번)
- lo: loopback interface
- veth: virtual ethernet interface
- ens
- em1
- MULTICAST: Can handle multicast packet
- UP: NUC is working
- LOWER_UP: L1 layer, Physical device signal up
- mtu: Maximum Transmission Unit (1500 default)
- qdisc: Queuing Disciplines, NIC에 들어오기전 Queue에 저장되는 패킷들 우선순위 부여 알고리즘
- state: NIC 현재 작동 상태
- qlen: 전송큐 크기
- link/ether
  - L2 layer, Link layer protocol is ethernet
  - 바로 옆에 나오는 주소는 해당 NIC의 MAC 주소
  - brd는 브로드캐스트 시 사용되는 주소
- inet: L3, 바로 옆 주소는 ipv4 or ipv6에 따른 주소
- scope
  - **Global**: Indicates accessibility and validity of the interface from a global perspective, allowing access from external networks. This is often seen in instances hosted in the cloud.
  - **Link**: Specifies that the interface is only accessible and valid within the local LAN, restricting access to the local network.
  - **Host**: Indicates that the interface is only valid and accessible within the host itself, limiting access to the local host.
- valid_lft, preferred_lft
  - Valid Lifetime and Preferred Lifetime


```py
ip link add
```


#### subnet mask
> [https://www.cloudflare.com/ko-kr/learning/network-layer/what-is-a-subnet/](https://www.cloudflare.com/ko-kr/learning/network-layer/what-is-a-subnet/)

라우터의 정의="서브넷이 다른 네트워크와 연결하기 위해 최적의 경로를 찾아서 목적지까지 패킷을 전달하는 장비"

1. IP = 주소가 속한 네트워크 + 해당 네트워크 안의 장치 = subnet + device address
2. 라우터 uses subnet mask to route (cidr)


아래 예시는 cidr 적용안된 설명같긴 하다.

> 실제 예를 들어, IP 패킷의 주소가 IP 주소 192.0.2.15라고 가정해 보겠습니다. 이 IP 주소는 클래스 C 네트워크이므로 네트워크는 "192.0.2"(또는 기술적으로는 정확하게 192.0.2.0/24)로 식별됩니다. 네트워크 라우터는 패킷을 "192.0.2"라고 표시된 네트워크의 호스트로 전달합니다. 패킷이 해당 네트워크에 도착하면 네트워크 내의 라우터가 라우팅 테이블을 참조합니다.서브넷 마스크 255.255.255.0을 사용하여 이진법 계산을 합니다.장치 주소 "15"(나머지 IP 주소는 네트워크를 나타냄)를 확인하고 패킷이 이동해야 하는 서브넷을 계산합니다.패킷을 해당 서브넷 내에서 패킷을 전달하는 라우터 또는 스위치로 전달하고 패킷은 IP 주소 192.0.2.15에 도착합니다(라우터 및 스위치에 대해 자세히 알아보기).



#### Router / Gateway

```py
controlplane ~ ➜  route
Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         172.25.0.1      0.0.0.0         UG    0      0        0 eth1
10.244.0.0      0.0.0.0         255.255.255.0   U     0      0        0 cni0
10.244.1.0      10.244.1.0      255.255.255.0   UG    0      0        0 flannel.1
172.25.0.0      0.0.0.0         255.255.255.0   U     0      0        0 eth1
192.0.243.0     0.0.0.0         255.255.255.0   U     0      0        0 eth0


# when to add
ip route add 192.168.2.0/24 via 192.168.1.1
```


#### Default Destination

- same 0.0.0.0, which means any IP addr
- internet에 존재하는 다양한 ip들을 일일히 처리 불가, default를 두어 처리한다.

```
default         172.25.0.1      0.0.0.0         UG    0      0        0 eth1
0.0.0.0
```

#### What if Gateway 0.0.0.0?

```
10.244.0.0      0.0.0.0         255.255.255.0   U     0      0        0 cni0
```

같은 네트워크 안에 있는 디바이스는 router / Gateway를 타고 가지 않아도 되니, 0.0.0.0값이 gateway에 존재



![](/images/k8s-network.png)

- To send packets from A to C, two routes need to be set up: `A -> B -> C` and `C -> B -> A`
- **In Linux, packets received on eth1 are not automatically forwarded to eth0; this is due to security concerns.**

![](/images/k8s-network2.png)

To enable forwarding, you can either set it in `/etc/sysctl.conf` or temporarily change it by setting `/proc/sys/net/ipv4/ip_forward` to 1.


![](/image/k8s_network_wrap.png)


## DNS



```py
controlplane ~ ➜  cat /etc/hosts
127.0.0.1       localhost
::1     localhost ip6-localhost ip6-loopback
fe00::0 ip6-localnet
ff00::0 ip6-mcastprefix
ff02::1 ip6-allnodes
ff02::2 ip6-allrouters
192.2.61.9      controlplane
10.0.0.6 docker-registry-mirror.kodekloud.com

# ctrl + D to finish
controlplane ~ ➜  cat >> /etc/hosts
1.1.1.1         test
```

- `>>`: append
- `>`: change


- `/etc/hosts`: Domain name mapping
- `/etc/resolv.conf`: Domain nameserver ip
  - nameserver
  - **search** field
- `/etc/nsswitch.conf`: Order configuration

```
controlplane ~ ➜  cat /etc/nsswitch.conf 
# /etc/nsswitch.conf
#
# Example configuration of GNU Name Service Switch functionality.
# If you have the `glibc-doc-reference' and `info' packages installed, try:
# `info libc "Name Service Switch"' for information about this file.

passwd:         files
group:          files
shadow:         files
gshadow:        files

hosts:          files dns
networks:       files

protocols:      db files
services:       db files
ethers:         db files
rpc:            db files

netgroup:       nis
```

- hostsf를 보면 file -> dns 순으로 해독된다. 즉 /etc/hosts가 먼저 해석된다.



```py
$ ping
# etc/hosts를 처리하지 않음, dns서버만 처리함
$ nslookup
$ dig
```

## CoreDNS

## Namespace

![](/images/k8s-network3.png)

When the container is created, it has its own private routing table and ARP table, and a virtual Ethernet interface (veth0) is automatically created.

#### Create Netwok ns

```py
ip netns add red
ip netns add blue

ip netns
red
blue
```

#### Exec in network ns

```py
ip netns exec red ip link

# simpler
ip -n red link
```

#### ARP
> address resolution protocol

- 동일한 네트워크 내에 존재하는 호스트들의 IP주소와 ethernet 주소(mac 주소)를 확인하는 명령어
- 즉, 특정 네트워크 내에 어떤 호스트들이 존재하는지를 확인할 수 있는 것이 바로 arp 명령어

```py
arp 
ip netns exec red arp
```

![](/images/k8s-network4.png)

- Host는 container에 대해서 알 수 없다.



![](/images/k8s-network5.png)


```py
ip netns add red
ip netns add blue

ip netns
red
blue

# same
ip -n red link
ip netns exec red ip link

ip netns exec red arp
ip netns exec red route

# link
# i.g ip link add veth0 type veth peer name veth1 (veth0 <-> veth1)
ip link add veth-red type veth peer name veth-blue # (veth-red <-> veth-blue)
ip link set veth-red netns red && ip link set veth-blue netns blue # set namespace

# set ip
# dev means device
ip -n red addr add 192.168.15.1 dev veth-red
ip -n blue addr add 192.168.15.2 dev veth-blue

# up link (running)
ip  -n red link set veth-red up
ip  -n blue link set veth-blue up

# ping test
ip netns exec red ping 192.168.15.2
```

#### Switch

![](/images/k8s-network6.png)

만약 container들이 많아진다면? switch가 필요하다.( virtual switch )

- [x] LINUX BRIDGE
- Open vSwitch(ovs)

```py
debian@debian:~$ ip netns list
blue
red

debian@debian:~$ sudo ip netns add v-net-0 type bridge
debian@debian:~$ ip netns list
v-net-0
blue
red
ip link set dev v-net-0 up
```

- **veth 타입과 달리 bridge type을 생성하면 자동으로 network namespace도 생성됨**

```py
# delete past cable
# other pair automatically deleted
ip -n  red link del veth-red

# create new cable (veth-red <-> veth-red-br)
ip link add veth-red type veth peer name veth-red-br
ip link add veth-blue type veth peer name veth-blue-br

# set veth to namespace
ip link set veth-red netns red # red
ip link set veth-red-br master v-net-0 # bridge
ip link set veth-blue netns blue # blue
ip link set veth-blue-br master v-net-0 # bridge

# set ip
ip -n red addr add 192.168.15.1 dev veth-red
ip -n blue addr add 192.168.15.1 dev veth-blue

# up
ip -n red link set veth-red up
ip -n red link set veth-blue up
```

- Set v-net-0 ip address on host namespace

```py
$ ping 192.168.15.1
Not Reachable!
$ ip addr add 192.168.15.5/24 dev v-net-0
$ ping 192.168.15.1 # red ok
```

#### Route
> "How to ping outside of host network from isolated namespace container?"

![](/images/k8s-network7.png)

```py
# 1. allow request 
# open v-net-0 -> LAN
ip -n blue ip route add 192.168.1.0/24 via 192.168.15.5

# 2. allow response
iptables -t nat -A POSTROUTING -s 192.168.15.0/24 -j MASQUERADE
```

- Link to internet (route)

```py
$ ip -n blue ping 8.8.8.8
Connect: Network is unreachable

$ ip -n blue route
... nothing related 8.8.8.8

# default to v-net-0
$ ip -n blue ip route add default via 192.168.15.5

$ ip -n blue ping 8.8.8.8
success
```

> Q. Then how to send request from outside host(192.168.1.3) to host's container?


![](/images/k8s-network8.png)

A. Set port forwarding rule on Host.

```py
$ iptables -t nat -A PREROUTING --dport 80 --to-destination 192.168.15.2:80 -j DNAT
```


## c.f) Linux namespace

We need 2 things to isolate network

1. UTS namespace
2. Network namespace

> 리눅스 네임스페이스란?
[리눅스 네임스페이스는 프로세스를 실행할 때 시스템의 리소스를 분리해서 실행할 수 있도록 도와주는 기능입니다. 한 시스템의 프로세스들은 기본적으로 시스템의 리소스들을 공유해서 실행됩니다. 이를 단일 네임스페이스라고 생각해볼 수 있습니다. 실제로 리눅스에서는 1번 프로세스(init)에 할당되어있는 네임스페이스들을 자식 프로세스들이 모두 공유해서 사용하는 구조로 이루어져있습니다.](https://www.44bits.io/ko/keyword/linux-namespace)

```
$ ls -al /proc/1/ns
total 0
dr-x--x--x 2 root root 0 Jan 31 03:47 .
dr-xr-xr-x 9 root root 0 Jan 24 14:46 ..
lrwxrwxrwx 1 root root 0 Jan 31 03:47 cgroup -> 'cgroup:[4026531835]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 ipc -> 'ipc:[4026531839]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 mnt -> 'mnt:[4026531840]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 net -> 'net:[4026531993]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 pid -> 'pid:[4026531836]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 pid_for_children -> 'pid:[4026531836]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 user -> 'user:[4026531837]'
lrwxrwxrwx 1 root root 0 Jan 31 03:47 uts -> 'uts:[4026531838]'

$ ls -l /proc/1074/ns
total 0
lrwxrwxrwx 1 root root 0 Jan 31 04:08 cgroup -> 'cgroup:[4026531835]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 ipc -> 'ipc:[4026531839]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 mnt -> 'mnt:[4026531840]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 net -> 'net:[4026531993]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 pid -> 'pid:[4026531836]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 pid_for_children -> 'pid:[4026531836]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 user -> 'user:[4026531837]'
lrwxrwxrwx 1 root root 0 Jan 31 04:08 uts -> 'uts:[4026531838]'

diff <(ls -Al /proc/1/ns | awk '{ print $11 }')  \
       <(ls -Al /proc/1074/ns | awk '{ print $11 }')
```


### UTS namespace
> https://www.44bits.io/ko/post/container-network-1-uts-namespace

```py
debian@debian:~$ hostname
debian

debian@debian:~$ touch /tmp/utsns1
debian@debian:~$ sudo unshare --uts=/tmp/utsns1 hostname utsns1
debian@debian:~$ sudo nsenter --uts=/tmp/utsns1 hostname
utsns1

sudo nsenter --uts=/tmp/utsns1 bash
debian@debian:~$ sudo nsenter --uts=/tmp/utsns1 bash
root@utsns1:/home/debian# hostname
utsns1
```

So we can isolate hostname with these commands.

- `unshare`
- `nsenter`


## Docker Networking

![](/images/k8s-docker-net0.png)
![](/images/k8s-docker-net1.png)
![](/images/k8s-docker-net2.png)
![](/images/k8s-docker-net3png)

## CNI (Container Networking Interface)


![](/images/k8s-network-cmd.png)
