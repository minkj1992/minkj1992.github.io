# [마스터링 이더리움 CH03] Ethereum Client


이번장에서는 [마스터링 이더리움 CH03] 이더리움 기초를 정리합니다. 인용글들은 [원문](https://github.com/ethereumbook/ethereumbook/blob/develop/03clients.asciidoc)을 참조하였습니다.

<!--more-->
<br />

{{< admonition >}}
_An Ethereum client is a software application that implements the Ethereum specification and communicates over the peer-to-peer network with other Ethereum clients
... 중략 ...
Ethereum is defined by a formal specification called the "Yellow Paper" (see [references])._

{{< /admonition >}}

이더리움 클라이언트는 sw application으로 [yellowpaper](https://ethereum.github.io/yellowpaper/paper.pdf)에 명시된 ethereum spec들을 만족시키는 구현체입니다. p2p network상에서 서로 통신하며 이더리움 생태계를 이끌어갑니다.

> This is in contrast to, for example, Bitcoin, which is not defined in any formal way.

## Ethereum Networks

{{< admonition >}}

_There exist a variety of Ethereum-based networks that largely conform to the formal specification defined in the Ethereum Yellow Paper
... 중략 ...
Among these Ethereum-based networks are Ethereum, Ethereum Classic, Ella, Expanse, Ubiq, Musicoin, and many others._

{{</ admonition >}}

Yellow Paper에 기제된 스펙을 기반으로 구현된 Ethereum based 네트워크들은 많이 존재합니다. (Ethereum, Ethereum Classic, Ella, Expanse, Ubiq, Musicoin...)
대부분 프로토콜 수준에서는 호환되지만, 각 네트워크 마다 세세한 부분에서 다른 점들이 존재하기 때문에 이더리움 클라이언트 `maintainers`들이 각 네트워크를 지원하기 위해 약간씩의 코드 변경작업이 필요합니다. 이 때문에 모든 버전의 이더리움 클라이언트 소프트웨어가 모든 이더리움 기반 블록체인을 실행하는 것은 아닙니다.

책에서 소개하는 대표적인 `Ethereum protocol` 구현체는 다음과 같습니다.

- **Parity**, written in Rust
- **Geth**, written in Go
- cpp-ethereum, written in C++
  - 현재는 `aleth`라고 레포가 되어있으며 deprecated됨
- pyethereum, written in Python
  - 현재는 `py-evm`로 관리되고 있습니다.

## Should I Run a Full Node?

> _The health, resilience, and censorship resistance of blockchains depend on them having many independently operated and geographically dispersed full nodes. Each full node can help other new nodes obtain the block data to bootstrap their operation, as well as offering the operator an authoritative and independent verification of all transactions and contracts._

글을 읽고 보니, 이더리움에서 말하는 `Node`와 `Client`개념이 헷갈려 먼저 정리하고 들어가겠습니다.

[Client vs Node 공식문서](https://ethereum.org/en/developers/docs/nodes-and-clients/)에 따르면

- `Node`
  - 분산된 이더리움 네트워크 컴퓨터들에서 동작하는 소프트웨어를 지칭합니다.
  - "Node" refers to a running piece of client software.
  - features: verify blocks, transaction data
  - Types
    - full node: block 전부 copy
    - light node: header만 copy
    - archive node
- `Client`
  - 사용자들의 컴퓨터에서 node를 실행할 수 있도록 하는 application을 뜻합니다
  - node는 piece of client입니다.
  - A client is an implementation of Ethereum that verifies all transactions in each block, keeping the network secure and the data accurate.

{{< admonition >}}

1. **Full client**
   Full clients store the entire Ethereum blockchain; a process that can take several days to synchronize and requires a huge amount of disk space – over 1 Terabyte to be exact, according to the latest figures. Full clients allow connected nodes to perform all tasks on the network, including mining, transaction and block-header validation and running smart contracts.

2. **Light client**
   Ethereum clients may be implemented in full or in part. The above overview gives an explanation of how a “full” client works, however it is important to know that you don’t always need to run a full client. Typically when data storage and speed are at issue, developers will elect to use what are called “light clients.”

   Light clients offer a subset of the functionality of a full client. Light clients can provide faster speeds and free up data storage availability because, unlike the full clients, they do not store the full Ethereum blockchain.

   The scope of a light client’s functionality is tailored toward the goals of the Ethereum client. For example, light clients are frequently used for private keys and Ethereum address management within a wallet. Additionally, they tend to handle smart contract interactions and transaction broadcasts. Other uses for remote clients include web3 instances within JavaScript objects, dapp browsers and retrieving exchange rate data.

3. **Remote client**
   There is a third type of client called a remote client which is similar to a light client. The main difference being, a remote client does not store its own copy of the blockchain, nor does it validate transactions or block headers. Instead, remote clients fully rely on a full or light client to provide them with access to the Ethereum blockchain network. These types of clients are predominantly used as a wallet for sending and receiving transactions.

{{< /admonition >}}

> The terms "remote client" and "wallet" are used interchangeably, though there are some differences. Usually, a remote client offers an API (such as the web3.js API) in addition to the transaction functionality of a wallet

참고로 `remote client`라는 용어와 `wallet`은 interchangeably하게 사용되는데, 둘의 미묘한 차이점으로는 `remote client`는 api(such as web3.js)를 제공한다는 점이다.

<center>

**remote client = wallet + api**

</center>

> Ethereum remote clients do not validate block headers or transactions. They entirely trust a full client to give them access to the blockchain, and hence lose significant security and anonymity guarantees.

이더리움 `remote clients`는 `light client` 처럼 block header 검증하지 않습니다. 이 덕분에 local hw 스펙을 줄일 수 있으나, 외부 full client들에게 depend한다는 특징이 있습니다.

## The JSON-RPC interface

> [왜 JSON-RPC를 사용할까?](https://www.getoutsidedoor.com/2019/08/10/%EC%99%9C-json-rpc%EB%A5%BC-%EC%82%AC%EC%9A%A9%ED%95%A0%EA%B9%8C/)

글을 읽다가 왜 이더리움은 `json-rpc`를 사용하는지 궁금해서 찾아보게 되었습니다. REST와 json-rpc의 차이점에 대해서 정리하겠습니다.

- `JSON-RPC`: tcp base로 원격/로컬 프로세스 procedure(함수)에 직접 접근하는 방식

  - over tcp
  - only single endpoint
  - crud외 표현 가능

- `REST`
  - over http(s)
  - crud(http method)를 벗어난 표현에 제한적이다.

rpc는 소스코드 > idl(interface definition language) > rpcgen> stub 코드 생성 > `rpc runtime`을 통한 packet 통신 (tcp L4) 과정을 통해서 서버와 클라이언트간 통신을 하게 만듭니다.
http(L7)보다 더 낮은 레벨(L4)에서 동작하기 때문에 기능이 덜 할 수 있지만, 제약이 덜합니다.

## Conclusion

뭔가 3장은 내심 겉할기 식의 설명들이 많았던 것 같습니다. `rpc-json`을 굳이 쓰는 이유도 없었고 client와 node에 대한 설명도 크게 와닿지 않은 것 같아 답답하네요. 예를 들면 full client를 실행하는 방법을 알려주면서, full client를 사용자가 왜 운용해야하는지를 light client와 비교하는 핵심이 빠진 느낌이 듭니다. (full client를 사용해야지 reward가 들어오는 걸까요..? 좀 더 공부가 필요하네요)

<center> - 끝 - </center>

