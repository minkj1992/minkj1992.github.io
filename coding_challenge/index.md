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

```
들어가기 앞서 이른 아침에도 이렇게 면접 볼 수 있도록 기회 주셔서 감사합니다.

안녕하세요 열정 비타민 제민욱입니다.
사내에서는 열정을 상징하는 레오라는 네이밍을 사용하고 있고요, 개발과 인생 모두 열정적으로 살고 싶어서 이렇게 지었습니다.

저는 현재 2년 경력을 가진 주니어 개발자 입니다.
처음 카카오에 들어갔을 때는 블록체인 tf에 들어가 python, django로 개발을 하였습니다.
이후 함께한 팀이 'krust'라는 카카오 블록체인 법인회사를 만들게 되면서, 따라가지 않고 지갑 서비스라는 새로운팀으로 이동하였습니다.
그곳에서 기존 블록체인 tf에서 관리하던 서버들을 운영하고 kotlin/spring 기반으로 프로젝트들을 마이그레이션 하는 작업을 하고 있습니다.
```

- motivation to change career

```
크게 2가지 이유로 커리어 전환을 생각하게 되었습니다.

1. 로보틱스로 커리어 전환을 하고 싶다.
언젠가 총 4가지 분야중에서 웹/앱 백엔드 개발자가 커리어 전환이 필요할 시점이 올 것이라 생각하는데요. 블록체인, 딥러닝, 그래픽스, 로보틱스입니다.

그래서 막연하게 커리어 전환을 언젠가 해야지 생각하고 있다가, 크러스트로 이직 권유를 받았을 때, 결정을 해야하는게 현실로 닥쳐오면서 진지하게 고민했던 것 같아요.

당시 진지하게 고민해보니 블록체인 개발자로 커리어를 전환하는 것이 저의 주관이 아닌, 환경에 더 크게 영향을 받은 느낌을 받았고, 여러 면담 끝에 거절하게 되었습니다.

그러다 보니 이후 부터는 더 간절하게 나머지 분야들에 대해 리서치하게 되었습니다. 그리고 그 중 로보틱스가 그리고 있는 미래상에 제 가슴이 뛴다는 느낌을 받아서 로보틱스로 커리어 전환을 목표하게 되었습니다.

비록 새 팀에 적응하느라 잦은 야근으로 많이 학습하진 못했지만, 주말마다 udemy ros2강의를 들으면서 흥미를 더 키워갔던 것 같습니다.

2. 가능하다면 스타트업 환경에서 근무하고 싶다.
카카오의 경우 인프라팀, db팀, 모니터링 툴등 편리한 기능들이 존재해서 빠르게 서비스들을 만들어가는 장점이 있는데요. 이게 한편으로는 단점이라고 생각합니다.

이런 환경이 개발자 입장에서는 오히려, 설계나 운영에서 고민해야 하는 많은 기회들을 빼앗아 가기 때문입니다. 그래서 넥스터즈 같은 대외활동을 하기도 했지만, 제대로 성장하기 위해서는 스타트업 환경에서 근무 하고 싶다고 생각하고 있습니다.
```

- 산인공 자격증

```
해당 서비스는 산업인력공단측에서 사용자들이 보유한 자격증을 NFT 카드로 발급해주고 싶은 니즈에서 출발한 프로젝트입니다.

프로젝트는 크게 4가지 역할로 나눠지는데요.
1. 프레젠테이션
    - 타 부서 서비스와의 dependency
    - 카카오 클라 버전 validation
    - 약관 등
    - api 요청
2. api 서버
    - 카카오 accountId -> 지갑 서비스에서 사용자 개인정보 획득
    - 산인공 api 자격증 리스트
    - 블록체인이기 때문에 속도 이슈로 queue에 NFT 발급 이벤트 enqueue
3. worker / queue
    - queue consume해서 트랜잭션 트래킹 및 실제 블록체인 네트워크망으로 erc721 nft 발행 요청
4. private klaytn nodes
    - 유저 address로 실제 nft 발급

발행 요청이 완료되면 유저들은 tms로 발급된 자격증 카드들을 카카오 지갑 서비스에서 확인 가능합니다.
```

- strength and weakness

```
- 강점: 좋아하는 분야에 꾸준합니다.

일례로 5년전에 복싱을 시작해서 재미를 느끼고 매일 복싱을 하고 있습니다. 그리고 개발을 좋아하는데요, 조그만 거라도 매일매일 새로운 지식을 배우고 성장하는 것이 즐거운 것 같아요. 또 최근 블로그를 만들었는데, 그날 배웠던 지식 하나하나 잘 정리하는 게 더 재미를 주는 것 같습니다.

- 약점: 어릴때 종종 준비물들을 까먹었던 기억이 있습니다. 이런 습관들이 나이가 들 수록 크리티컬해지는 것을 느꼈습니다. 그러다 보니 자연스럽게 구글 캘린더와 일기를 작성하는 습관을 가져가게 된 것 같습니다.
```

- 대외활동

```
Nexters는 현직 개발자와 디자이너들이 팀을 꾸려서 단기간에 하나의 웹/앱 서비스를 만드는 대외활동입니다. 열정있는 사람들을 더 알고 싶은 마음과 서비스의 처음과 끝을 모두 경험해보고 싶어서 지원하게 되었습니다.

활동은 약 2개월동안 아이디어 >  MVP > 프로토타입-피드백 > 배포를 타이트하게 진행하며, 개발자들은 현직에서 사용해보지 못했던 기술들을 정해 스터디하며 프로젝트를 진행합니다.

* 운영진: 홍보, 인터뷰, 프로젝트 일정, 연사 초청, 스터디 진행.
```

- 관심있는 기술

```
graphql
- query
- mutation
- resolver

pytest

ros2 (ros foxy)
- node: 주체
- topic: pub/sub grouping
- service (rpc): rpc
- action: 목적 지향형, long-running behaviors (서비스보다 더 오래걸리는 작업)
    - goal
    - result
    - feedback
```

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
- What was good in using Jenkins Pipeline?
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

<center>-끝-</center>

