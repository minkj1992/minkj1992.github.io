# DFS and BFS


> 간단하게 bfs와 dfs를 `python`으로 구현해보고, 상황별 장단점을 분석해본다.

<!--more-->

## 1. Graph Traverse
### 1.1. BFS
```python
from collections import deque
graph = {
  '5' : ['3','7'],
  '3' : ['2', '4'],
  '7' : ['8'],
  '2' : [],
  '4' : ['8'],
  '8' : []
}


def bfs(start_node="5"):
    queue = deque([start_node,])
    visited = set()
    while queue:
        node = queue.pop()
        if node in visited:continue
        
        visited.add(node)
        for neighbor in graph[node]:
            queue.appendleft(neighbor)
```
### 1.2. DFS
```python
graph = {
  '5' : ['3','7'],
  '3' : ['2', '4'],
  '7' : ['8'],
  '2' : [],
  '4' : ['8'],
  '8' : []
}

# 5 8 7 3 4 2
def dfs(start_node="5"):
    stack = [start_node,]
    visited = set()

    while stack:
        node = stack.pop()
        if node in visited:continue

        visited.add(node)
        for neighbor in graph[node]:
            stack.append(neighbor)
```

```python
def recursive_dfs(node="5", visited = set()):
    visited.add(node)

    for neighbor in graph[node]:
        if neighbor in visited: continue
        recursive_dfs(neighbor, visited)
```
### 1.3. BFS vs DFS 
> 문제 유형별로 어떤 알고리즘이 더 유리한지 서술합니다.

- :(far fa-times-circle fa-fw): : impossible
- :(far fa-thumbs-up fa-fw): : good and possible
- :(far fa-thumbs-down fa-fw): : bad but possible

<center>

|index|Problem|BFS|DFS
|:---:|:--:|:--:|:---:|
|1|그래프의 **모든 정점을 방문** 하는 문제| :(far fa-thumbs-up fa-fw): | :(far fa-thumbs-up fa-fw):|
|2|각 경로 마다 특징을 저장해둬야 하는 문제| :(far fa-thumbs-down fa-fw): | :(far fa-thumbs-up fa-fw):|
|3|최단 거리 문제| :(far fa-thumbs-up fa-fw): | :(far fa-thumbs-down fa-fw):|
|4|문제의 그래프가 매우 클 경우| :(far fa-times-circle fa-fw): | :(far fa-thumbs-up fa-fw):|
|5|검색 시작 지점과 원하는 대상이 가까이 있을 경우| :(far fa-thumbs-up fa-fw): | :(far fa-thumbs-down fa-fw):|

</center>

1. 단순히 모든 node 방문이라면 둘 모두 사용가능합니다.



2. a->b로 가는 경로를 구할 때, 경로 안에서 `같은 숫자`가 x번 이상 없어야 하는 경우, `dfs`는 함수의 인자에 local 변수들을 좀 더 손 쉽게 줄 수 있는 반면, `queue`를 활용하는 `bfs`는 상태를 기억하기 좀 더 까다롭다. 물론 queue에 node를 넣을 때, local state를 같이 넣어주면 되긴 하지만, 이는 명시적이지 못하므로 `dfs`가 더 유리하다 생각된다.



3. `bfs`의 경우 level(e.g 이동 count, tree의 level, ) 단위로 확장되기 때문에 목적지에 도착하는 순간 return한 값이 최소 이동거리가 되지만, dfs의 경우에는 깊이 있게 하나씩 파니까 상대적으로 `bfs`보다 오래 걸린다.



4. `Python`의 경우 하나의 리스트가 가질 수 있는 `Py_ssize_t`는 `536870912`인 반면, 함수에 대한 메모리 제한은 두지 않고 있다. 그러므로 queue를 이용해서 구현하는 `bfs`의 경우 그래프가 크다면, 제한이 있는 반면 `recursive function`으로 구현한 `dfs`한정(stack으로 구현한 경우 동일한 문제발생)해서 `function call stack`을 사용해서 메모리 관리를 하기 때문에 상대적으로 더 큰 그래프 탐색에 사용될 수 있다.



5. `3.`의 경우와 같은 원리이다.


## Conclusion
필자는 왠만하면 `bfs`를 좀 더 선호하는 경향이 있는데, 과거의 경험을 떠올려 보면, (정확하게 기억은 나지 않지만) path가 연속적으로 이어져야하는 로직에서 bfs가 불편했던 것 같다. `DFS`의 경우에는 처음부터 끝까지 연속적으로 탐색하기 때문에 비교적 쉽게 구현이 되었는데, `BFS`는 spread 하면서 이동하기 때문에 해당 상황에서 불리 했던 것 같다.

dfs, bfs를 언제 써야하는지 개인적으로 헷갈렸었는데 표로 정리하고 나니 나중에도 유용하게 볼 것 같다.

<center> - 끝 - </center>

