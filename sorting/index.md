# Sorting algorithms


> Let's summary list of Sorting Algorithms
<!--more-->


## tl;dr

|index|sorting name|time|space|description|
|:---:|:--:|:--:|:---:|:---:|
|1|Bubble|O(n^2)|O(1)|루프당 max가 가장 뒤, swap O(n^2)|
|2|Selection|O(n^2)|O(1)|루프당 min 맨 앞, swap O(n)|
|3|Insertion|O(n^2)|O(1)|`I`까지 sort 보장, 정렬이 어느정도 되어있다면 사용할 것|
|4|Merge|O(nlogn)|O(n)|nlogn 알고리즘 중 유일한 stable|
|5|Heap|O(nlogn)|O(1)|insert(O(logn) * n개 원소, space가 1이 포인트|
|6|Quick|O(nlogn)|O(n)|piv기준 작으면 left 크거나 같으면 right,balanced partition을 위해 random piv를 해준다.|

## Bubble Sort
## Selection Sort
## Insertion Sort

- 이미 정렬된 상태라면 `O(n)`의 빠른 속도를 보인다.
- 정렬된 상태에서 빠른 이유는 각 insert마다 1번의 비교만 하면 되기 때문이다.
- reversed를 사용하면 insert시 arr re-arrange를 방지 가능하다.
- 단점: 삽입을 하게 되면 데이터가 하나씩 뒤로 밀려야 되기 때문에 배열이 길어질수록 효율이 떨어진다.
- 개인적으로 input()받을 때 insertionSort를 사용하면 입력과 정렬을 동시에 할 수 있어서 더욱 효율적인 것 같다.

```python
def insertion_sort(arr):
    n = len(arr)

    for i in range(1, n):
        val = arr[i]
        j = i - 1
        if arr[j] < val:
            continue
        while j >= 0 and val < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = val
    return arr 
```


## Merge Sort

```python
def merge_sort(arr):
    def merge(left, right):
        l = r = 0
        result = []
        while l < len(left) and r < len(right):
            if left[l] < right[r]:
                result.append(left[l])
                l+=1
            else:
                result.append(right[r])
                r+=1

        return result + left[l:] + right[r:] # 나머지 (left over)
    if len(arr) <= 1:
        return arr
    mid = len(arr) // 2
    return merge(merge_sort(arr[:mid]), merge_sort(arr[mid:]))
```
## Heap Sort

## Quick Sort
- 성능을 생각하면 piv를 random으로 찾아야 한다.
```python
from __future__ import annotations


def quick_sort(arr: list[int]) -> list[int]:
    if len(arr) <= 1:
        return arr
    piv = arr[0]
    others = arr[1:]

    left = [v for v in others if v <= piv]
    right = [v for v in others if v > piv]
    return quick_sort(left) + [piv] + quick_sort(right)
```
