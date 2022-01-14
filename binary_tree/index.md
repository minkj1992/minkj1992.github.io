# Binary Tree


> Binary Tree관련 알고리즘들을 학습하고 정리합니다.
<!--more-->


## Binary Tree
1. child > 2면 안된다.
2. parent > 1이면 안된다.
3. root(부모가 없는 노드)는 한개만 존재해야 한다.
4. array로 구현하면 편의를 위해 0인덱스를 비워둔다.
   1. parent = child % 2
   2. lchild = parent * 2
   3. rchild = parent * 2 + 1

## Heap

1. max heap, min heap (등호도 고려된다.)
2. 대소 관계는 부모-자녀 간에만 고려된다.
3. left child 먼저 삽입된다. (즉 leaf 중에 left 없이 right가 있는 경우는 없다.)


### Heap insert
1. 인덱스 마지막에 새로운 요소 append
2. (if parent is exist) 부모와 대소 비교 하여 exchange. (아래 -> 위 heapify)

### Heap pop
1. root pop
2. 힙의 마지막 element를 root로 이동
3. 힙 재구성 (= 위 -> 아래 heapify)
   1. (if child exist) l, r 비교하여 현재 노드가 작다면 exchange (max heap 기준)
   2. 재귀적으로 반복   



### Heap sort
- `O(n + n*logn)` => `O(nlogn)`
  - Max heap 구성(O(n))
  - 루트와 말단 노드 교체 후 heapify (`O(nlogn)`)
    - O(logn) = 트리 최대 높이 = heapify시 depth
    - n = 모든 노드들에 대하여 검사

