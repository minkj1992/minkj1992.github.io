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


```sh
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


```sh
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

```jsthon
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

```jsthon
# create and run pod with finance namespace
# 생각해보니까 apply, create으로 pod 직접적으로 만들지 않았던 것 같네. 곧바로 run 했던 것 같은데, run = create + run like docker
k run redis -n=finance --image=redis

```

```jsthon
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

```jsthon
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

```js
floe@floe-QEMU-Virtual-Machine:~$ k get po -o wide
NAME    READY   STATUS    RESTARTS   AGE   IP            NODE       NOMINATED NODE   READINESS GATES
nginx   1/1     Running   0          40s   10.244.0.14   minikube   <none>           <none>
floe@floe-QEMU-Virtual-Machine:~$ k get nodes
NAME       STATUS   ROLES           AGE     VERSION
minikube   Ready    control-plane   5d19h   v1.28.3
```

- When there is no scheduler, there would be empty Node value on pod description.

```js
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

```
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

```js
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

```js
# Open terminal output with vim to easily find `/` N/n
> k describe no node01 | vim -
```


```js
k get no --no-headers | wc -l

# set label to node
k label no node01 color=blue

k create deploy blue --image=nginx --replicas=3

k describe no controlplane | grep -i taints
Taints:             <none>
k describe no node01 | grep -i taints
Taints:             <none>
```


```js
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
```js

controlplane ~ ➜  ps -aux | grep kubelet | grep -i config
root        4351  0.0  0.0 4519680 100556 ?      Ssl  02:23   0:36 /usr/bin/kubelet --bootstrap-kubeconfig=/etc/kubernetes/bootstrap-kubelet.conf --kubeconfig=/etc/kubernetes/kubelet.conf --config=/var/lib/kubelet/config.yaml --container-runtime-endpoint=unix:///var/run/containerd/containerd.sock --pod-infra-container-image=registry.k8s.io/pause:3.9

cat /var/lib/kubelet/config.yaml | grep staticPodPath
```

### How to create staticPod
> Create a static pod named static-busybox that uses the busybox image and the command sleep 1000

```js
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

```js
root@controlplane:~# kubectl get pods --all-namespaces -o wide  | grep static-greenbox
default       static-greenbox-node01                 1/1     Running   0          19s     10.244.1.2   node01       <none>           <none>
root@controlplane:~#
```

From the result of this command, we can see that the pod is running on node01.

2. Next, SSH to node01 and identify the path configured for static pods in this node.

  - Important: The path need not be /etc/kubernetes/manifests. Make sure to check the path configured in the kubelet configuration file.

```js
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

```js
root@node01:/etc/just-to-mess-with-you# ls
greenbox.yaml
root@node01:/etc/just-to-mess-with-you# rm -rf greenbox.yaml 
root@node01:/etc/just-to-mess-with-you#
```

4. Exit out of node01 using CTRL + D or type exit. You should return to the controlplane node. Check if the static-greenbox pod has been deleted:

```js
root@controlplane:~# kubectl get pods --all-namespaces -o wide  | grep static-greenbox
root@controlplane:~# 
```


# Chapter 5: Application Lifecycle Management


## Configmap

- Note that not to use `--from-file`, this is only handle single key like `--from-literal`
- Instead use `k create cm <NAME> --from-env-file=`

```js
controlplane ~ ➜  vim webapp.env 

controlplane ~ ➜  k create cm webapp-config-map --from-env-file=./webapp.env
configmap/webapp-config-map created
```

## Secret

```js
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

```js
// Move every resources from node-1 to others
k drain node-1
// after node upgrade
// cordon: block node from scheduling
// uncordon: enable scheduling back
k uncordon node-1
```

Running the uncordon command on a node will not automatically schedule pods on the node. When new pods are created, they will be placed on node01.

> We will be upgrading the controlplane node first. Drain the controlplane node of workloads and mark it UnSchedulable

```js
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

```js
$ k drain node01 --ignore-daemonsets --force
node/node01 already cordoned
Warning: deleting Pods that declare no controller: default/hr-app; ignoring DaemonSet-managed Pods: kube-flannel/kube-flannel-ds-rp464, kube-system/kube-proxy-8gmv5
evicting pod default/hr-app
```


> Question... I'm just curious that why there's still pod on drained node
>> controlplane ~ ➜  k drain controlplane --ignore-daemonsets 
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

## Kubernetes Software Versions


```js
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

```js
> pager /etc/apt/sources.list.d/kubernetes.list

deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyri
ng.gpg] https://pkgs.k8s.io/core:/stable:/v1.28/deb/ 
/
```

Switching to another Kubernetes package repository 

```js
vim /etc/apt/sources.list.d/kubernetes.list

# change version v1.28 -> v1.29
:s/v1.28/v1.29/g
```

2. Determine which version to upgrade to

```js
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
```js
> sudo apt-mark unhold kubeadm && \
sudo apt-get update && sudo apt-get install -y kubeadm=1.29.0-1.1 && \
sudo apt-mark hold kubeadm
```

- verify the upgrade plan

```js
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
```js

> sudo kubeadm upgrade apply $target_version

```

Now, upgrade the version and restart Kubelet. Also, mark the node (in this case, the "controlplane" node) as schedulable.

```js
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

```js
sudo apt-mark unhold kubeadm && \
sudo apt-get update && sudo apt-get install -y kubeadm='1.29.0-1.1' && \
sudo apt-mark hold kubeadm
```

3. kubeadm upgrade node (instead apply)

```js
> sudo kubeadm upgrade node
```

4. updgrade kubelet and kubectl

```js
sudo apt-mark unhold kubelet kubectl && \
sudo apt-get update && sudo apt-get install -y kubelet='1.29.0-1.1' kubectl='1.29.0-1.1' && \
sudo apt-mark hold kubelet kubectl

sudo systemctl daemon-reload
sudo systemctl restart kubelet


exit (back to master)
kubectl uncordon node01
```




