# A Day before CKA (합격)


Do Killer.sh and Mock test.
<!--more-->

## TL;DR


24.04.19 ~ 24.05.08

중간에 모두의 연구소와 아주대 멘토링 그리고 mlflow 엠버서더 활동때문에 일주일을 [다른 곳](https://github.com/minkj1992/llama3-langchain-mlflow)에 신경써야했지만, 그래도 다행히 아이펠 과정 본격적으로 시작하기 전에 빠르게 합격한 것 같다. (합격 여부만 나오고 아직 점수를 확인을 못해서 업데이트를 기다려야 할 것 같다.)

![](/images/cka.png)

개인적으로 CKA를 처음으로 돌아가 다시 시작한다면, 그리고 만약 정말 빠르게 자격증을 따야한다면

1. killer.sh 먼저 풀어보기 또는 https://killercoda.com/killer-shell-cka 문제 유형 풀기
2. 3일 반복
3. 1차 시험
4. 합격 결과에 따라서, 필요한 내용 공부

이렇게 하면 좋을 것 같다. 이제 ckad랑 cks가 남았는데 빠르게 처리해야겠다.





#### OS Upgrades
> https://beta.kodekloud.com/user/courses/udemy-labs-certified-kubernetes-administrator-with-practice-tests/module/85e5661f-74bc-4534-adc8-f51d138fdace/lesson/afa1849c-2679-44e5-9e17-4105bf40a914


> We need to take node01 out for maintenance. Empty the node of all applications and mark it unschedulable.

```py
k drain node01 --ignore-daemonsets
```

> drain problem

```
controlplane ~ ➜  kubectl drain node01 --ignore-daemonsets
node/node01 cordoned
error: unable to drain node "node01" due to error:cannot delete Pods declare no controller (use --force to override): default/hr-app, continuing command...
There are pending nodes to be drained:
 node01
cannot delete Pods declare no controller (use --force to override): default/hr-app
```

- If there are any pods that are neither mirror pods nor managed by a replication controller, replica set, daemon set, stateful set, or job, then drain will not delete any pods unless you use --force.

```py
k cordon node01
```


### Cluster Upgrade Process

1. Control Plane Node upgrade
    1. 패키지 저장소 버전 update
    2. kubeadm upgrade
    3. drain
    4. Upgrade kubelet and kubectl
    5. Uncordon
2. Worker Node upgrade


#### 1. Control Plane Node upgrade
> https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/#upgrading-control-plane-nodes

1. 패키지 저장소 버전 update
2. kubeadm upgrade
3. drain
4. Upgrade kubelet and kubectl
5. Uncordon


```py
# 1. 패키지 저장소 버전 update
vim /etc/apt/sources.list.d/kubernetes.list
sudo apt update
sudo apt-cache madison kubeadm

# 2. kubeadm upgrade
## 2.1. download
sudo apt-mark unhold kubeadm && \
sudo apt-get update && sudo apt-get install -y kubeadm='1.30.x-*' && \
sudo apt-mark hold kubeadm
## 2.2. verify
sudo kubeadm upgrade plan
## 2.3. apply package
sudo kubeadm upgrade apply v1.30.x
# 3. drain
kubectl drain <node-to-drain> --ignore-daemonsets
# 4. Upgrade kubelet and kubectl
sudo apt-mark unhold kubelet kubectl && \
sudo apt-get update && sudo apt-get install -y kubelet='1.30.x-*' kubectl='1.30.x-*' && \
sudo apt-mark hold kubelet kubectl

# 4.2. restart kubelet
sudo systemctl daemon-reload && sudo systemctl restart kubelet
# 5. Uncordon
kubectl uncordon <node-to-uncordon>
```

#### 2. Worker Node Upgrade (linux)
> https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/upgrading-linux-nodes/

Control plane과 다 똑같은데, 2번 kubeadm upgrade 명령어만 틀림 apply 대신 upgrade (`sudo kubeadm upgrade node`)

```py
# 1. 패키지 저장소 버전 update
vim /etc/apt/sources.list.d/kubernetes.list
sudo apt update
sudo apt-cache madison kubeadm

# 2. kubeadm upgrade
## 2.1. download
sudo apt-mark unhold kubeadm && \
sudo apt-get update && sudo apt-get install -y kubeadm='1.30.x-*' && \
sudo apt-mark hold kubeadm
## 2.2. upgrade
sudo kubeadm upgrade node
# 3. drain
kubectl drain <node-to-drain> --ignore-daemonsets
# 4. Upgrade kubelet and kubectl
sudo apt-mark unhold kubelet kubectl && \
sudo apt-get update && sudo apt-get install -y kubelet='1.30.x-*' kubectl='1.30.x-*' && \
sudo apt-mark hold kubelet kubectl
# 4.2. restart kubelet
sudo systemctl daemon-reload && sudo systemctl restart kubelet
# 5. Uncordon
kubectl uncordon <node-to-uncordon>
```

===

#### JsonPath
> 리소스 path를 어디서 찾아야되는거지
```
Print the names of all deployments in the admin2406 namespace in the following format:

DEPLOYMENT   CONTAINER_IMAGE   READY_REPLICAS   NAMESPACE

<deployment name>   <container image used>   <ready replica count>   <Namespace>
. The data should be sorted by the increasing order of the deployment name.


Example:

DEPLOYMENT   CONTAINER_IMAGE   READY_REPLICAS   NAMESPACE
deploy0   nginx:alpine   1   admin2406
Write the result to the file /opt/admin2406_data.
```

1. k edit을 통해서 structure 뽑기
2. output에 맞춰서 jsonpath 작성


```
kubectl -n admin2406 get deployment -o custom-columns=DEPLOYMENT:.metadata.name,CONTAINER_IMAGE:.spec.template.spec.containers[].image,READY_REPLICAS:.status.readyReplicas,NAMESPACE:.metadata.namespace --sort-by=.metadata.name > /opt/admin2406_data
```


#### network test
> Create a nginx pod called nginx-resolver using image nginx, expose it internally with a service called nginx-resolver-service. Test that you are able to look up the service and pod names from within the cluster. Use the image: busybox:1.28 for dns lookup. Record results in /root/CKA/nginx.svc and /root/CKA/nginx.pod

Use the command kubectl run and create a nginx pod and busybox pod. Resolve it, nginx service and its pod name from busybox pod.

To create a pod nginx-resolver and expose it internally:
kubectl run nginx-resolver --image=nginx
kubectl expose pod nginx-resolver --name=nginx-resolver-service --port=80 --target-port=80 --type=ClusterIP

To create a pod test-nslookup. Test that you are able to look up the service and pod names from within the cluster:

kubectl run test-nslookup --image=busybox:1.28 --rm -it --restart=Never -- nslookup nginx-resolver-service
kubectl run test-nslookup --image=busybox:1.28 --rm -it --restart=Never -- nslookup nginx-resolver-service > /root/CKA/nginx.svc

Get the IP of the nginx-resolver pod and replace the dots(.) with hyphon(-) which will be used below.
kubectl get pod nginx-resolver -o wide

kubectl run test-nslookup --image=busybox:1.28 --rm -it --restart=Never -- nslookup <P-O-D-I-P.default.pod> > /root/CKA/nginx.pod


#### service account
> Create a new service account with the name pvviewer. Grant this Service account access to list all PersistentVolumes in the cluster by creating an appropriate cluster role called pvviewer-role and ClusterRoleBinding called pvviewer-role-binding. Next, create a pod called pvviewer with the image: redis and serviceAccount: pvviewer in the default namespace.

Pods authenticate to the API Server using ServiceAccounts. If the serviceAccount name is not specified, the default service account for the namespace is used during a pod creation.

- https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/

Now, create a service account pvviewer:

`kubectl create serviceaccount pvviewer`
To create a clusterrole:

`kubectl create clusterrole pvviewer-role --resource=persistentvolumes --verb=list`
To create a clusterrolebinding:

`kubectl create clusterrolebinding pvviewer-role-binding --clusterrole=pvviewer-role --serviceaccount=default:pvviewer`
Solution manifest file to create a new pod called pvviewer as follows:

```yaml
---
apiVersion: v1
kind: Pod
metadata:
  labels:
    run: pvviewer
  name: pvviewer
spec:
  containers:
  - image: redis
    name: pvviewer
  # Add service account name
  serviceAccountName: pvviewer
```

#### NetworkPolicy

**test-* namespace인 pod들의 Ingress 요청을 허용하고 싶었고, namespace 필터링하면 될 거라고 생각했지만 아니었다. Selector는 label기반으로 만들어진다.**







