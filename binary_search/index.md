# Binary Search and BST


> Binary Search와 BST에 대해서 정리합니다.
<!--more-->


`Binary Search`는 `divide conquer`의 일종으로 검색 범위를 binary하게 줄여나가면서 원하는 데이터를 검색하는 알고리즘입니다.

```python
def binary_search(sorted_arr, target):
    n = len(sorted_arr)
    if n == 0:
        return -1
    
    low, high = 0, n-1
    while low <=high:
        mid = (low + high) // 2
        if sorted_arr[mid] == target:
            return mid
        
        if sorted_arr[mid] > target:
            high = mid -1
        else:
            low = mid + 1
    return -1 # low == high + 1 == mid
```

눈여겨 봐야할 포인트는 다음 2가지이다.

1. 검색 대상이 되는 arr가 `sorted`되어있다.
2. while의 조건으로 low <= high 등호가 들어있다.
   - 검색의 범위 element가 2개로 좁혀졌을 때 `// 2` 연산에 의해서 왼쪽만 탐색이 될 텐데, 찾아야하는 값이 우측 값에 존재한다면 low와 high가 같아야만(low == high == mid) 검색이 가능하다.

## Binary Search 특징
- `retrieve` Time complexity: `O(log N)`
- `retrieve` Space complexity: `O(1)`
- 삽입 / 삭제 불가


## Binary Search Tree (`BST`)
> n= number of elements, h = tree height

들어가기 앞서, [ratsgo](https://ratsgo.github.io/data%20structure&algorithm/2017/10/22/bst/)를 참조하여 정리했음을 알려드립니다.

<center>

![](/images/bst_origin.png)

</center>

`이진 탐색 트리`란 `Binary Search`와 `Linked list`를 결합한 자료구조 입니다. 특히 `Binary Search`의 **탐색 속도**(`O(log n)`) 와 링크드리스트의 **삽입/삭제** `O(1)`의 장점을 결합했다는 특징이 있습니다.

참고로 binary search는 삽입/삭제가 불가하며, 링크드리스트는 탐색 속도가 `O(n)`이라는 단점들이 있습니다. BST는 서로의 장점을 사용해 각각의 단점을 `O(h)`로 보완합니다.

### 주요 특징
- left.val < root < right.val
- `inorder traverse`(중위 순회)시 결과가 정렬된 리스트가 주어진다.
  - left -> node -> right
- 구성하는 노드에서 중복된 노드가 없어야 한다. (**unique 보장**)
- 노드 끼리 `우선순위 대소 비교`가 가능해야 한다.
- retrieve, insert, delete의 계산복잡성은 모두 `𝑂(ℎ)`

### 기본 데이터 형태

```python
class Node:
    def __init__(self, val):
        self.val = val
        self.left = None
        self.right = None

class BinarySearchTree:
    def __init__(self):
        self.root: Optional[Node] = None

    def set_root(self, val):
        self.root = Node(val)
```


### retrieve / find
> Time Complexity: `O(h)`

탐색 대상과 root를 비교하여 left / right를 찾아나간다. 이 경우 Binary Search와 비슷하게 `O(h)` 시간 복잡도를 가진다. (아래와 같은 극단적 불균형 트리인 경우이면서, min/max값을 탐색한다면 `O(n)`)

<center>

![](/images/unbalanced_bst.png)

</center>

```python
    def find(self, val):
        node = self.find_node(self.root, val):
        return True if node else False

    def find_node(self, node, val) -> Optional[Node]:
        if not node:
            return None
        elif val == node.val:
            return node
        elif val < node.val:
            return self.find_node(node.left, val)
        else:
            return self.find_node(node.right, val)
```

### insert
> Time Complexity: `O(h)`

`O(logn)`이 아닌 이유는 비대칭(`Unbalanced Binary Tree`)인 경우 tree의 높이가 n까지도 가능하기 때문이다. (sorted arr를 차례대로 insert 시킬경우)

이를 해결하기 위해서는 `BF`(balance factor)를 사용해 balance를 맞추는 `AVL` 또는 `B-`같은 트리를 사용해야 한다.

```python
    def insert(self, val):
        if not self.root:
            self.set_root(val)
        else:
            self.insert_node(self.root, val)

    def insert_node(self, node, val):
        if val <= node.val:
            if node.left:
                self.insert_node(node.left, val)
            else:
                node.left = Node(val)
        elif val > node.val:
            if node.right:
                self.insert_node(node.right, val)
            else:
                node.right = Node(val)
```


{{< admonition tip >}}
AVL 트리는 rotation을 사용해 tree의 insert / delete 시 balance를 맞춘다. 그러므로 검색의 경우 항상 `O(log n)`을 보장한다.

특별한 점은 `single rotation`, `double rotation`을 통해서 tree의 balance를 맞추어 주는데 [자세한 설명](https://ratsgo.github.io/data%20structure&algorithm/2017/10/27/avltree/)을 참조하세요
{{< /admonition >}}


### delete
> Time Complexity: `O(h)`
삭제는 총 3가지 경우가 존재합니다.

1. leaf node (자식노드가 없는 경우) -> 그냥 제거

<center>

![](/images/bst_nochild.png)

</center>

1. 자식노드가 하나 존재하는 경우 -> 제거 후, 자식 노드를 삭제된 노드의 부모로 연결

<center>

![](/images/bst_onechild.png)

</center>

1. 자식노드가 둘 존재하는 경우


이 경우에는 predecessor 또는 successor를 삭제할 노드와 위치를 뒤 바꾼 다음, 
1와 2의 삭제 방법을 사용하면 됩니다. (참고로 successor와 predecessor는 자식노드가 1개 또는 없는 경우 밖에 존재하지 않습니다.)

**predecessor로 제거, successor로 제거 둘다 가능 합니다.**

<center>

![](/images/bst_twochild.png)

</center>

{{< admonition tip >}}
predecessor: 삭제 대상 노드의 왼쪽 서브트리 가운데 최대값
successor: 삭제 대상 노드의 오른쪽 서브트리 가운데 최소값

그림 기준으로 **16**을 inorder traverse를 해보면 다음과 같습니다.

> 4, 10, 13, **16**, 20, 22, 25, 28, 30, 42

이때, predecessor(13), successor(20)가 됩니다. 
{{< /admonition >}}

```python
# delete 방법 (d = 삭제 대상 노드의 레벨)
1. 삭제 대상 노드의 오른쪽 서브트리를 찾는다.
2. successor(1에서 찾은 서브트리의 최소값) 노드를 찾는다.
3. 2에서 찾은 successor의 값을 삭제 대상 노드에 복사한다.
4. successor 노드를 삭제한다.
```

가장 계산이 복잡한 자식 노드가 둘 모두 존재하는 경우의 시간 복잡도를 분서해보겠습니다. 

1에서 d레벨(트리 높이) 만큼 이동을 해주어야 하며, 2에서 최대 h-d 레벨(트리높이)만큼 이동해주어야 합니다. 3과 4의 연산은 계산에서 제외한다면 `O(d + h -d)` => `O(h)`가 만들어집니다.

e.g) 간단히 가장 복잡할 것 같은 `root`를 지운다 가정하였을 때, d = 1, h = h이므로 `O(1 + h - 1)`이 됩니다.

### traverse (inorder)
> Time Complexity: O(n)

<center>

![](/images/bst.png)

</center>

위의 그림의 경우 1 -> 3 -> 5 -> 7 -> 8 -> 10로 순회 가능하다.

```python
    def traverse(self):
        return self.traverse_node(self.root)

    def traverse_node(self, node):
        result = []
        if node.left:
            result.extend(self.traverse_node(node.left))
        if node:
            result.extend([node.val])
        if node.right:
            result.extend(self.traverse_node(node.right))
        return result
```

## Conclusion
이상으로 binary search와 bst에 대하여 알아 보았습니다. 관련해서 leetcode문제는 

- [binary search](https://github.com/minkj1992/algorithm/tree/main/practice/leetcode/binarySearch)
에 정리를 해두었습니다.


<center> - 끝 - </center>

