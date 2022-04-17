# All Basic Computer Science


> Let's prepare basic computer science interview questions.

<!--more-->

## Operating System

...

## Network

### 웹 통신의 큰 흐름: https://www.google.com/ 을 접속할 때 일어나는 일

> 키워드: `dhcp`, `dns`, `nat`, `isp`, `3-way / 4-way handshake`, `ssl (ssl handshake)`

1. 가장 먼저 브라우저가 url에 적힌 값을 파싱해서 `HTTP Request Message`를 만들고, OS에 전송 요청을 합니다.

2. **OS는 `DNS Lookup`을 수행합니다.**
   룩업 과정은 etc/hosts > DNS Cache > Cache가 없을 경우 dns server로 ip를 얻어옵니다.

3. **DNS server로 ip request**
   이때 DNS server IP는 1차적으로 `isp`(internet service provider, ex kt, skt...)가 제공하는 정보들이 `dhcp`에 의해 컴퓨터에 세팅됩니다.

`dhcp`는 wifi를 쓸 경우, 공유기에 연결되어있는 `gateway ip`와 `router`의 `NAT`을 통해 `사설 ip`(private ip)를 할당 받으며, 외부 통신을 할 경우 router의 `Public ip`을 사용합니다.

ISP에 의해 세팅되어 있는 dns server로 아래 형식의 요청을 보내어, 도메인에 매핑된 ip를 받아옵니다.

```
- from: router ip(nat ip)
- to: 받아온 ip
- 게이트웨이 ip : wifi이면 공유기 연결 게이트웨이 ip / 스마트폰이면 자체 ip
```

4. **루트 도메인서버에서부터 서브도메인 서버순으로 dns query**
   이제 DNS Server로 DNS Query를 요청하게 되면 DNS 서버는 `Root name server`에 해당 도메인을 질의하고, `.com` `name server`의 ip를 받아오게 됩니다.

그 후 `.com 네임 서버`에 도메인 Query하게되면 `google.com`의 ip주소를 받고 최종적으로 `www.google.com`의 ip를 받아오게 됩니다.

5. pc는 최종 서버 ip로 HTTP Request를 보낸다.

- 3-way handshake

  > syn > ack, syn > ack

- 4-way handshake

  > fin > ack(close wait) > fin(last_ack) > ack

- [tls/ssl](https://www.cloudflare.com/ko-kr/learning/ssl/what-happens-in-a-tls-handshake/)
- [좀 더 자세한 과정 설명](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)

![](/images/rsa_ssl.jpeg)

```
# RSA 키 교환 알고리즘
1. client hello (protocol version, 암호 알고리즘, 압축 방식, 클라 난수)
2. server hello (세션 ID, ca 인증서, 서버난수)
3. verify ca and get public key
4. 클라는 난수(pre master secret) 생성 후 public key로 암호화 후 서버 전달
5. 클라 세션키 생성 및 서버는 난수를 private key로 복호화 하여 대칭키(세션 키) 생성
6. 클라는 세션키(대칭키)로 암호화한 fin message를 서버로 전달
7. 서버 또한 세션키로 암호화한 fin message를 전달
8. 이후 세션키를 통해 통신 계속 진행
```

### TCP vs UDP

|                |        TCP         |              UDP               |
| :------------: | :----------------: | :----------------------------: |
|    연결방식    |    연결형서비스    |        비 연결형 서비스        |
| 패킷 교환 방식 |   가상 회선 방식   |        데이터그램 방식         |
|   전송 순서    |   전송 순서 보장   |    전송 순서가 바뀔 수 있음    |
| 수신 여부 확인 | 수신 여부를 확인함 |   수신 여부를 확인하지 않음    |
|   통신 방식    |  1:1 통신만 가능   | 1:1 / 1:N / N:N 통신 모두 가능 |
|     신뢰성     |        높음        |              낮음              |
|      속도      |        느림        |              빠름              |

### Web Socket Handshake

[web socket mdn](https://developer.mozilla.org/ko/docs/Web/API/WebSockets_API/Writing_WebSocket_servers)

![](/images/websocket.png)

클라와 서버가 서로 TCP/IP 4계층 레이어에서 통신한다. 즉 conneciton을 들고 있다.

- http 요청 이후, upgrade요청 한다.
- ping을 지속적으로 쏴서, health-check

### 로드 밸런싱(Load Balancing)

로드 밸런싱이란 여러 서버에게 균등하게 트래픽을 분산 시켜주는 것이다.

nginx의 경우 기본적으로 라운드 로빈 방식으로 동작합니다.

- scale-out
- scale-up

### Nginx가 10k problem을 해결한 방식

기존 방식은 request당 하나의 process 또는 thread를 사용해서 요청들을 처리했습니다. 이에 반해 nginx는 worker pool을 두고 request가 들어올 때 마다, async하게 worker(default cpu 당 1)에게 task를 위임합니다.

이렇게 하게 될 경우 process/thread에 비해, pcb/tcb를 만드는데 들어가는 비용을 줄일 수 있으며 또한 사용자의 요청이 많아질 경우, 상대적으로 context switching에 사용되는 비용을 줄일 수 있습니다. 마지막으로 os가 스케쥴링에 들어가는 비용이 줄어듭니다.

즉 지정된 갯수의 미리생성된 process(thread) 워커를 사용함으로써, 기존의 request가 늘어날 때마다, os 리소스가 급격히 늘어나는 것을 방지하여, 이에 대한 side effect(스케쥴링, context-switching등에 대한 오버헤드를 막아줍니다.) 또한 워커에 필요한 리소스들을 미리 생성해두기 때문에 Process 생성에 들어가는 오버헤드를 줄여줍니다.

## Database

### DB 트랜잭션이란?

> 트랜잭션은 데이터베이스의 데이터를 조작하는 논리적인 작업의 단위(unit of work)입니다.

- 트랜잭션은 ACID
  - Atomicity
    - all or nothing (rollback)
    - db transaction, rollback으로 보장
  - Consistency
    - transaction이 commit 되어도 DB의 여러 제약 조건에 맞는 상태를 보장하는 성질이다. 송금하는 사람의 계좌 잔고가 0보다 작아지면 안 된다.
  - Isolation
    - transaction이 진행되는 중간 상태의 데이터를 다른 transaction이 볼 수 없도록 보장하는 성질이다. 송금하는 사람의 계좌에서 돈은 빠져나갔는데 받는 사람의 계좌에 돈이 아직 들어가지 않은 DB 상황을 다른 transaction이 읽으면 안 된다.
    - lock으로 구현
  - Durability
    - transaction이 Commit했을 경우 해당 결과가 영구적으로 적용됨을 보장하는 성질이다

### 트랜잭션과 lock에 대해서 isolation과 연결 지어 설명해주세요

DB엔진은 ACID 원칙을 희생하여 동시성을 얻을 수 있는 방법을 제공합니다.

- Row level lock
  - shared lock: read lock
  - exclusive lock: write lock
- Record lock
  - s lock: read index lock
  - x lock: write index lock
- Gap lock: db index record의 gap에 걸리는 lock (gap = db에 실제 record가 없는 부분)

lock은 모두 transaction이 commit 되거나 rollback 될 때 함께 unlock

- `Consistent read`
  - Isolation
- https://s1107.tistory.com/45
- http://labs.brandi.co.kr/2019/06/19/hansj.html
- https://suhwan.dev/2019/06/09/transaction-isolation-level-and-lock/

- index
  - https://idea-sketch.tistory.com/43?category=547413
  - https://idea-sketch.tistory.com/45

### DB index에 대해 설명해주세요

> https://idea-sketch.tistory.com/43?category=547413

### todo

- dirty read, Non-Repeatable Read, Phantom Read
- optimistic lock, pessimistic lock
- slow query
- Index

## Software Engineering

## Design Pattern

## Language

## Computer Architecture

