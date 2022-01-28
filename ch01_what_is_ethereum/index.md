# [마스터링 이더리움 CH01] 이더리움이란 무엇인가



이번장에서는 `Mastering Ethereum`과 [김혐남님의 세미나 내용](http://www.umlcert.com/mastering_ethereum-1/)을 기반으로 이더리움이란 무엇인가에 대해서 정리합니다.
<!--more-->


### tl;dr
- **Ethereum is a Blockchain Platform.**

이더리움을 블록체인 생태계에서 포지션을 생각한다면 플랫폼으로 구분할 수 있습니다.

- **Ethereum is "the world computer"**

이더리움 진영에서 내새우는 정의는 (합의를 기반으로 동작하는) `World Computer`입니다.



## 1. What is Ethereum

{{< admonition quote >}}
Ethereum is often described as "the world computer"
{{< /admonition >}}

이더리움 진영에서 정의하는 이더리움이란 (탈 중앙화된) 월드 컴퓨터입니다. `World`와 `Computer`에 집중해본다면, `World`라는 단어처럼 이더리움은 "전세계를 하나로 묶어줄 수 있는 connection"을 제공해주어야 할 것으로 보입니다. 묶어준 다는 점에서 p2p 네트워크가 필요해 보입니다. `Computer`라는 단어는 `program`, 그리고 이 프로그램을 개발할 수 있는 `language`가 필요해 보입니다. 또한 그 프로그램은 decentralized 성격을 가져야 하기 때문에, 일반 프로그램과 비교해 특별한 성질을 가질 것으로 보입니다.

> Ethereum is an open source with globally decentralized computing infrastructure that executes programs called `smart contracts`. It uses a blockchain to synchronize and store the system's state changes.

`smart contract`라는 특별한 program을 사용해 globally decentralized computing infrastructure 이면서도, 하나의 world computer 즉 state가 sync되는 시스템을 구현합니다. 

{{< admonition tip >}}

Ethereum is a deterministic but practically `unbounded state machine`, Which means consisting of a **globally accessible singleton state** and a virtual machine.

{{< /admonition >}}

p2p로 퍼져있으면서도, 하나의 컴퓨터로 동작하기 위해서 스마트 컨트랙트는 `globally accessible singleton state`와 virtual machine 개념이 존재합니다.

> It uses a blockchain to synchronize and store the system's state changes, along with a cryptocurrency called `ether` to meter and constrain execution resource costs.

## 2. Compared to `Bitcoin`

- `in common`
  - p2p network connecting participants.
  - to synchronize `PoW` they uses `Byzantine fault-tolerant consensus` algorithm.
  - uses `cryptographic primitives`
    - hashes
    - digital signatures
    - digital currency

{{< admonition tip >}}
`PoW`(proof of work) is a form of adding new blocks of transactions to a cryptocurrency's blockchain.
{{< /admonition >}}

- `in contrast`
  - Ethereum's purpose is not primarily to be a digital currency payment network. `ether` is intended as a `utility currency` to pay for use of the Ethereum platform as the world computer.
  - Bitcoin's Script language is intentionally constrained to simple true/false evaluation of spending conditions, but ethereum's language is `Turing complete`

{{< admonition tip >}}


A `Turing complete system` means a system in which a program can be written that will find an answer (although with no guarantees regarding runtime or memory).

So, if somebody says "my new thing is Turing Complete" that means in principle (although often not in practice) it could be used to solve any computation problem.

+ 프로그래밍 언어는 이와 비슷하게 `Turing complete`하다 왜냐하면 프로그램이 실행될 충분한 메모리와 시간이 주어진다면 특정 computational problem을 풀어낼 수 있기 때문이다.
{{< /admonition >}}

`김혐남`님의 말을 빌리자면, 이더리움을 간단히 **블록체인의 플랫폼**이라 소개합니다. 블록체인 플랫폼이 되려다보니 이더리움은 블록체인 튜링 완전 프로그래밍이 가능해야 했고, `Turing Complete`해지니 `application`을 만들 수 있게 되었고, 이 앱은 `block chain` 위에서 실행되니 블록체인의 특성을 지닌 앱이 될 수 있었습니다. 이더리움은 이런 블록체인의 기능을 플랫폼처럼 추상화 시켜, 참여하는 개발자들이 쉽게 블록체인의 특성을 지닌 애플리케이션을 개발할 수 있도록 도와줍니다.

## 3. Components of a Blockchain
The components of an open, public blockchain are:


1. A peer-to-peer (P2P) network connecting participants and propagating transactions and blocks of verified transactions, based on a standardized "gossip" protocol

2. Messages, in the form of transactions, representing state transitions

3. A set of consensus rules, governing what constitutes a transaction and what makes for a valid state transition

4. A state machine that processes transactions according to the consensus rules

5. A chain of cryptographically secured blocks that acts as a journal of all the verified and accepted state transitions

6. A consensus algorithm that decentralizes control over the blockchain, by forcing participants to cooperate in the enforcement of the consensus rules

7. A game-theoretically sound incentivization scheme (e.g., proof-of-work costs plus block rewards) to economically secure the state machine in an open environment

8. One or more open source software implementations of the above ("clients")

## 4. The Birth of Ethereum

> 이더리움 창립자들은 프로그래밍을 통해 다양한 애플리케이션을 지원할 수 있는 특정 목적에 국한되지 않는 블록체인에 대해 생각하고 있었다. 이 생각은 이더리움과 같은 범용 블록체인을 사용하여 개발자가 피어투피어 네트워크, 블록체인, 합의 알고리즘 등의 기본 메커니즘을 구현하지 않고도, 특정 애플리케이션을 프로그래밍할 수 있다는 것이다. 이더리움 플랫폼은 **세부사항을 추상화하고 탈중앙화 블록체인 애플리케이션을 위한** 결정적이고 안전한 프로그래밍 환경을 제공한다.

이더리움은 범용 블록체인(=블록체인 플랫폼)으로서 위치하고 있다는 것을 확인할 수 있습니다.

> 탈중앙화의 첫 번째 대상은 ‘가치의 이동’입니다. 디파이라 불리는 탈중앙 금융이 블록체인의 성공적인 킬러 앱이 될 수 있는 이유는 바로 가치의 이동에서 찾을 수 있습니다.

블록체인이 왜 탈중앙화할 수 밖에 없었는지를 고민해 보면 `정보의 이동`이 아닌 핵심은 `가치의 이동`이 필요했기 때문입니다.

> The original blockchain, namely Bitcoin’s blockchain, tracks the state of units of bitcoin and their ownership. You can think of Bitcoin as a distributed consensus state machine, where transactions cause a global state transition, altering the ownership of coins. The state transitions are constrained by the rules of consensus, allowing all participants to (eventually) converge on a common (consensus) state of the system, after several blocks are mined. 

> **Ethereum answers the question: "What if we could track any arbitrary state and program the state machine to create a world-wide computer operating under consensus?"**

