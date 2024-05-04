# Robotics Coding Challenge


> This document records the series of experiences that a back-end server developer has experienced when moving to robotics in the second half of the year 2021. (2021.12.25 - 2022.01)

<!--more-->

<Admonition type="tip" icon="💡" title="Did you know...">
  <p>
    Use plugins to introduce shorter syntax for the most commonly used JSX
    elements in your project.
  </p>
</Admonition>

**어느 2021 11월** `Rust`를 공부하고 싶다는 생각이 들어서 할 만한 사이드 프로젝트 주제를 찾던 중 `ROS`가 재밌겠다는 생각에 로보틱스를 공부하게 되었다. (그런데 지금와서 돌이켜 보니 이 시장은 `C++`, `Pyhton`뿐..)

그렇게 로보틱스 공부와 자료 조사를 하는 과정에서 `boston dynamics`, `naver labs`, `bear`등의 기업을 찾아보게 되었고, 백엔드 개발자의 경력이 통할지 궁금해서 이력서를 제출해 보았다.

영어 resume가 익숙하지 않아서 미국에 살고 있는 친형의 코치를 받으며, 어찌어찌 이력서를 제출하게 되었고 지원 한 곳 중 한 곳에서 연락이 와서 코딩테스트를 보게 되었다. (~~사실 로보틱스 회사가 지원할 곳이 국내에는 2곳 밖에 없다~~)

{{< admonition note FYI >}}
연락이 온 회사의 시험 일정은 다음과 같다.

resume > online coding challenge(5 days) > 1st interview > cto interview.  
{{< /admonition >}}

## 1. Online Coding Challenge (Algorithm)

알고리즘 문제 3문제와 특정한 프로그램을 만들어야 하는 과제가 주어졌다. 이메일로 링크를 클릭하면 시험이 시작되며 총 5일정도 데드라인이 주어졌는데, 체감상 이틀(10시간)정도 집중하면 충분한 시간인 듯 하다.

알고리즘의 문제 난이도는 프로그래머스 2~3레벨 정도라서, 면접에 물어볼 걸 대비해 최대한 깔끔하게 코드를 짜려고 노력했는데 그걸 감안하더라도 많아야 3시간 정도면 3문제 모두 풀 수 있었던 것 같다. (5일 주길래 괜히 쫄았..)

- string 관련

```python
"""
time: O(n)
space: O(1)
"""
from typing import Optional


NUMBERS = list(range(10))
WORDS = ["zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]
OPERATORS = {"minus": "-", "plus": "+"}
N2W = {n: w for n, w in zip(NUMBERS, WORDS)}
W2N = {w: n for w, n in zip(WORDS, NUMBERS)}


def string_challenge(raw_input: str) -> str:
    n = len(raw_input)
    start = 0
    count = 0
    word_len = (3, 4, 5)
    expr = ""
    for i in range(n + 1):
        if count in word_len:
            word = decode(raw_input[start:i])
            if word:
                expr += word
                count = 0
                start = i
        count += 1
    num = eval(expr)
    return encode(num)


def decode(string: str) -> Optional[str]:
    if string in OPERATORS:
        return OPERATORS[string]
    if string in W2N:
        return str(W2N[string])
    return None


def encode(num: int):
    string_num = ""
    for n in str(num):
        if n == "-":
            string_num += "negative"
            continue
        string_num += N2W[int(n)]
    return string_num


# keep this function call here
print(string_challenge(input()))

```

- array 관련 (`binary tree`)

```python
"""
time: O(n)
space: O(1)
"""

from collections import defaultdict
from typing import List, Tuple, Dict


def is_binary_tree(arr: List[str]) -> bool:
    """
    1. #child > 2
    2. child's parent > 1
    3. multiple tree
    """
    p2c: dict = defaultdict(list)  # parent : [c1,c2..]
    c2p = {}  # child:parent

    for str_tuple in arr:
        child, parent = parse(str_tuple)

        if len(p2c[parent]) == 2:
            return False
        p2c[parent].append(child)

        if child in c2p:
            return False
        c2p[child] = parent

    return is_single_tree(p2c, c2p)


def parse(str_tuple: str) -> Tuple[int, int]:
    str_pair = str_tuple.split(",")
    child = int(str_pair[0][1:])
    parent = int(str_pair[1][:-1])
    return child, parent


def is_single_tree(p2c, c2p) -> bool:
    root = None
    for parent in p2c:
        if root not in c2p:
            if root:
                return False
            root = parent
    return True


# keep this function call here
answer = "true" if is_binary_tree(input()) else "false"
print(answer)
```

- math 관련 (`postfix calculate`)

```python
"""

"""
OPERATOR = ("+", "-", "*", "/")


def calculate(postfix_notation: str) -> int:
    stack = []
    for s in postfix_notation.split():
        value = None
        if not is_operator(s):
            value = int(s)
        else:
            r_value = stack.pop()
            l_value = stack.pop()
            expression = f"{l_value}{s}{r_value}"
            value = int(eval(expression))
        stack.append(value)

    return stack[0]


def is_operator(s: str) -> bool:
    return s in OPERATOR


# keep this function call here
```

## 2. Online Coding Challenge (Assignment)

Online coding 테스트는 총 알고리즘 3문제와 과제가 하나 주어지는데, 5일이나 시간을 주었던게 이 `assignment` 때문이 아닐까 생각된다.
문제는 `Simple xxx controller`를 구현하는 것인데, 문제 요구 사항은 간단했지만 `Employer`들께 잘 보이려면 간단하게 제출하면 안 될 것 같은 직감이 빡와서 알고리즘을 풀고 해당 sw 컨트롤러 관련 정보를 모으는데 5시간 넘게 소모한 것 같다.

그 뒤 세운 거창했던 계획..

- [x] Design pattern 한개 이상 사용하자
- [x] Unit 테스트는 Coverage 80% 넘기자.
- [-] Integration 테스트는 pytest-bdd를 활용해 시나리오들을 추가하자.
- [x] 트랜잭션을 구현해서 에러가 생긱 경우 rollback 시키는 기능을 추가하자
- [x] README는 꼼꼼히!

서치 결과 해당 과제에 어울리는 패턴은 총 3가지로 `state pattern`, `command pattern`, `chain of responsibility` 정도가 있었던 것 같다. controller가 여러가지 상태를 가지고, 그에 따라서 실행되어야 할 stage에 따라서 메서드들이 다르게 접근 되어야 된다는 점에서 `state-pattern`이 좋아 보였다. 총 10가지 정도 state를 클래스로 구현하였는데, 만들고 보니까 state 클래스별로 **실행되지 말아야할 method**들의 에러 메시지를 일일히 구현해주는 부담이 있었고, 무엇보다 내가 만드는 controller가 10개 정도의 state에 비즈니스로직이 퍼져있어 가독성이 떨어졌다.

그래서 깔끔하게 포기하고, 선택한 `command pattern`.
트랜잭션을 구현하고 싶었기 때문에, `undo`, `redo`기능을 편하게 구현해야 했었는데 `Command Pattern`을 사용하니 배치로 여러 작업을 묶어서 실행하면 되어서 편하게 구현이 되었던 것 같다. 거기다 구현해야 하는 controller가 추후에는 더 많은 command를 받아들여야 하는 요구사항이 있었기 때문에, 명령들을 분리해서 관리할 수 있어 좋았던 것 같다.

평소에 디자인 패턴을 현업에서 많이 써보지 못해서 익숙하지 않았는데, 이번 기회에 공부하면서 구현해보니 재밌었던 것 같다.
마지막으로 `chain of responsibility`를 적용해 보려고 했는데, 사실 이건 이름부터 너무 길어서 마음에 들지 않았기도 했고 테스트 코드 짜려니 시간도 부족한 것 같아서 바로 손절했다. (카리스마🔪.. 찢었다...)

~~글을 쓰다 보니 갑자기 귀찮아 져서 급하게 마무리 하자면~~ 그렇게 중간에 약속들도 있고 해서 테스트는 원래 계획했던 bdd는 시나리오만 짜서 제출하게 되었다. 사실 워낙 simple한 controller이다 보니 integration 테스트가 의미 없긴 했지만 bdd는 짜보고 싶었는데 조금 아쉽다. [🤫..](https://github.com/minkj1992/atm_controller)

## 3. Prepare Interview (4 round)

{{< admonition tip >}}
면접은 총 4시간 동안 이뤄지며, 매 시간마다 담당자와 1:1 면접이 진행됩니다. 미국과 한국 면접관들이 조율하여 들어오시며, zoom으로 치뤄집니다.
{{< /admonition >}}

### 공통 질문 준비

- self-introduction
- motivation to change career
- 대표 프로젝트 설명
- strength and weakness
- 대외활동
- 관심있는 기술

### 이력서 기반 질문 준비

- icp

### 일반적인 cs 질문 준비

- DB
  - What is the main difference of delete and truncate in SQL?
    - `delete`: 데이터만 삭제 되며 테이블 용량은 줄어 들지 않는다.
      - 테이블 유지 / 데이터 삭제 / 롤백 o (저장공간 유지)
    - `trucanate`: 테이블의 최초 상태로 되돌림
      - 테이블 유지 / 데이터 삭제 / 롤백 x
    - `drop`: 테이블의 정의 자체를 완전히 삭제한다.
      - 테이블 삭제 / 데이터 삭제 / 롤백 x
  - [uuid](https://github.com/aragorn/home/wiki/Database-Key-Design)
  - Why do we need cherry-pick command in git?
    - 현재 브랜치가 아닌 다른 브랜치의 특정 commit을 현재 브랜치로 가져오고 싶을 때
- Web
- Hash collision
  - 확률적으로 1퍼 미만으로 하려면?

### Coding 과제 리뷰 준비

- How to get handled the security threat in getting access to the docker volume?
- what is the path normalization?

### 화면 공유 온라인 코딩 테스트 준비 (2번)

- leetcode medium to hard
- bfs / dfs
- heap
- dp(interleaving)

```python
# https://leetcode.com/problems/interleaving-string

import pprint

def is_interleaved(s1: str, s2: str, s3: str) -> bool:
    n, m = len(s1), len(s2)
    if (n + m) != len(s3):
        return False
    board = [[False] * (m + 1) for _ in range(n + 1)]
    for i in range(n + 1):
        for j in range(m + 1):
            # check edge
            if i == j == 0:
                board[i][j] = True
                continue
            if i == 0:
                if s2[j - 1] == s3[i + j - 1]:
                    board[i][j] = board[i][j - 1]
                continue
            if j == 0:
                if s1[i - 1] == s3[i + j - 1]:
                    board[i][j] = board[i - 1][j]
                continue
            # check except edge
            if s1[i - 1] == s3[i + j - 1]:
                if s2[j - 1] == s3[i + j - 1]:
                    board[i][j] = board[i - 1][j] or board[i][j - 1]
                else:
                    board[i][j] = board[i - 1][j]
            else:
                if s2[j - 1] == s3[i + j - 1]:
                    board[i][j] = board[i][j - 1]
    pprint.pprint(board)
    return board[n][m]


s1 = "aabcc"
s2 = "dbbca"
s3 = "aadbbcbcac"
is_interleaved(s1, s2, s3)

s1 = "abcde"
s2 = "12345"
s3 = "12345abcde"
is_interleaved(s1, s2, s3)
```

## 4. Do Interview (4 round)

> success

- [x] 5 round interview (5-6hours)
- [x] skipped english interview
## 5. payment

- [x] salary: 카카오 대비 45% 상승
- [x] stock: 당시 가치 전체 연봉의 30%의 4년 토탈 1년 cliff.

## 6. apartment

- [x] 성수동 원룸 2년 계약 1000/80

<center>-끝-</center>

