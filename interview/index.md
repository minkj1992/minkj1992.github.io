# All Basic Computer Science


> Let's prepare basic computer science interview questions.

<!--more-->


## Network

### HTTP의 GET과 POST 비교
> get과 post는 http 프로토콜의 가장 기본적인 메소드들 중 2개입니다.

#### GET
- 조회 기능
- http request message(1. request line, 2. header, 3. body)중 request line에 url (쿼리 포함)로 전송
- 웹 브라우저나 중간의 프록시 서버 (의도치 않게) cache 가능성 존재
- 길이 제한, 안정성

#### POST
- 서버로 데이터를 전송하여 리소스 생성 기능
- http reqeust message의 body에 데이터를 담아 전송하여, GET과 비교해 더 많은 양의 데이터 전송 가능

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


TCP
- 연결
- 신뢰성, 패킷 손실 또는 순서 다를 경우 -> 패킷 재전송 또는 순서 재정렬
- 순서 보장, 데이터는 송신한 순서대로 수신됨
- 흐름 제어, 슬라이딩 윈도우를 통해 받을 buffer 양 계산하여 전송

UDP
- 비연결, 데이터 바로 전송
- 신뢰성 x, 재전송 하지 않고, 패킷 순서 보장 x
- 흐름 제어 x
- 최소한의 오버헤드
- DNS, 온라인 게임, VoIP(voice call), 미디어 스트리밍 그리고 http3 (QUIC)

#### TCP 흐름제어 (sliding window)
> [패킷의 흐름과 오류를 제어하는 TCP](https://evan-moon.github.io/2019/11/22/tcp-flow-control-error-control/)

클라와 서버는 각각 데이터를 담을 수 있는 버퍼를 가지고 있고, `window`라는 일종의 마스킹 도구를 가지고 있다. 서버측의 윈도우 크기는 3way handshake를 통해, 마지막 클라이언트 측에서 보내준 버퍼 크기를 사용하여 서버측 윈도우 크기를 정하게 된다.

```py
localhost.initiator > localhost.receiver: Flags [S], seq 1487079775, win 65535
localhost.receiver > localhost.initiator: Flags [S.], seq 3886578796, ack 1487079776, win 65535
localhost.initiator > localhost.receiver: Flags [.], ack 1, win 6379
```

최초 SYN과 SYN+ACK에서 버퍼 크기를 `65535`로 이야기 한 뒤, 1/10 크기인 `6379`로 윈도우 사이즈를 정했다. 그 뒤 서버가 클라이언트에게 데이터를 전송할 때 마다 응답값인 ack에 window size와 ack number를 주어서 **수신 가능한 버퍼 공간과 다음에 받기 원하는 바이트의 번호를 요청할 수 있습니다.**

또한 데이터가 전송되는 서버 측 버퍼는 아래와 같은 3가지 구조로 존재하게 될 것입니다.

```py
전송완료 & ACK 응답 받음| 전송완료 & ACK 응답 받지 못함| 전송 대기 중
```

![](https://evan-moon.github.io/static/a04bfa93161e2a1c8ef37a6f19e1b0dc/21b4d/sw-3.png)

서버는 클라이언트의 window size에 따라 window를 이동시키고, 해당 패킷들을 전송하게 됩니다. 만약 이때 서버에서 전송한 패킷들 중 특정 데이터가 유실되거나 잘못 보낼 경우들을 처리하기 위해서 총 2가지 방식이 존재합니다.

1. Go Back N
2. Selective Repeat

![](https://evan-moon.github.io/static/1bc040f7f4090114bb5b65db8e769d5b/e9d87/go-back-n.png)

![](https://evan-moon.github.io/static/393dda4bf3efb487525f9567f2d08f12/21b4d/go-back-n-example.png)

`Go Back N`은 오류가 발생한 패킷 넘버부터 새롭게 패킷을 보내는 것으로, 기존의 보내졌던 패킷들 중 오류 패킷 이후의 정상적으로 수신받은 모든 데이터들도 패기합니다. 

![](https://evan-moon.github.io/static/a16fa4ea93d682445ee5ce78e4673ab5/e9d87/selective-repeat.png)

`Selective Repeat`은 에러가 난 패킷만 재전송해주는 방식이며, 이 경우에는 클라 측 버퍼의 패킷 순서가 꼬일 수 있기 때문에, 손실된 세그먼트들을 저장하는 버퍼(reordering buffer)를 따로 두어, 패킷들을 담아두고 기존의 버퍼가 해당 reordering buffer에 들어있는 데이터들이 기존 receiver buffer에 정렬되도록 합니다.

```python
$ sysctl net.inet.tcp | grep sack:
net.inet.tcp.sack: 1
```
### DNS lookup

![](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FbuJeCx%2FbtrRtVNCPHr%2FwP0QDJYq9fI9cwos0HkSRK%2Fimg.png)

1. 브라우저에서 naver.com을 입력하면
2. /hosts 파일에서 dns: ip에 대한 매핑 확인
3. (없을 경우) local pc dns cache 확인
4. (없을 경우) /resolv.conf에서 local dns server ip 확인

> Local DNS Server란?
>> ISP provider(skt, kt, 등) 또는 Public DNS(google, cloudflare)등이 역할 가능하며, 주요 역할은 DNS name을 IP 주소로 변환하는 것

![](https://img1.daumcdn.net/thumb/R1280x0/?scode=mtistory2&fname=https%3A%2F%2Fblog.kakaocdn.net%2Fdn%2FbIx8eY%2FbtrRtU2gjnv%2F20sXtl8sBhwb70IXcJqgQ0%2Fimg.png)

5. local dns server의 cache에서 확인
6. (없을 경우) local dns server - (recursive query) -> ROOT DNS SERVER
  1. TLD DNS server IP 획득
  2. TLD (.com, .kr 등)
7. local dns server -> TLD DNS SERVER
  1. Second domain dns의 authoritative server ip 획득
8. local dns server --> authoritative server
9. Authoritative dns -> subdomain(3rd level domain) 체크
  1. 예를들면 ftp.naver.com, blog.naver.com에서 ftp.* 또는 blog.* 와 매핑된 Ip 주소
10. ip 획득, cache 저장



### 웹 통신의 큰 흐름: https://www.google.com/ 을 접속할 때 일어나는 일

> 키워드: `dhcp`, `dns`, `nat`, `isp`, `3-way / 4-way handshake`, `ssl (ssl handshake)`


(pre-step) 노트북 기준 wifi 연결시, IP주소를 얻기 위해 DHCP 요청하며 이를 통해 라우터(또는 wifi 공유기)를 통해 사설 IP주소, 서브넷 마스크, 게이트웨이 주소등을 전달받아둔 상태입니다. 

1. 먼저 브라우저가 url에 적힌 값을 파싱해서 `HTTP Request Message`를 만들고, OS에 전송 요청을 합니다.

2. **DNS lookup**, /hosts -> cache 를 확인해서 IP 주소가 없을 경우 DNS lookup을 실시합니다.
3. **NAT**, 노트북에 할당된 사설 IP주소를 라우터/wifi공유기가 실행하여 public IP주소로 변환합니다.
4. 라우터의 라우팅 테이블을 통해 패킷 다음 목적지 선택
5. ISP의 라우터로 전달 후 다음 목적지 선택
6. google.com의 데이터 센터로 전달
7. 3way handshake
8. 패킷 전달
9. (keep-alive 이후) 4way handshake


### DNS round robin 방식
> DNS round robin이란 부하 분산 기술로, Authoritative Nameserver는 여러 IP 주소를 순차적으로 반환하여 RR방식으로 부하 분산하는 것입니다.

FYI, DNS RR은 (DNS lookup에도 영향)을 주는 IP 주소에 대한 load balancing이다. 

#### DNS RR 문제점


1. cache에 의해 균등하게 분산되지 못함
2. health check이 존재하지 않음



### handshakes
#### 3-way handshake

  > syn > syn-ack > ack
  
#### 4-way handshake

  > fin > ack(close wait) > fin(last_ack) > ack

#### TLS handshake

![](https://cf-assets.www.cloudflare.com/slt3lc6tev37/5aYOr5erfyNBq20X5djTco/3c859532c91f25d961b2884bf521c1eb/tls-ssl-handshake.png)



- [tls/ssl](https://www.cloudflare.com/ko-kr/learning/ssl/what-happens-in-a-tls-handshake/)
- [좀 더 자세한 과정 설명](https://blog.cloudflare.com/keyless-ssl-the-nitty-gritty-technical-details/)

![](/images/rsa_ssl.jpeg)

```py
# RSA 키 교환 알고리즘
1. client -> server: client hello (protocol version, 암호 알고리즘, 압축 방식, 클라 난수)
2. server -> client: server hello (세션 ID, ca 인증서, 서버난수)
3. client -> client: verify ca and get public key
4. client -> server: 클라는 난수(pre master secret) 생성 후 public key로 암호화 후 서버 전달
5. both: 클라 세션키 생성 및 서버는 난수를 private key로 복호화 하여 대칭키(세션 키) 생성
6. client -> server: 클라는 세션키(대칭키)로 암호화한 fin message를 서버로 전달
7. server -> client: 서버 또한 세션키로 암호화한 fin message를 전달
8. 이후 세션키(master key, 대칭키)를 통해 통신 계속 진행
```

### Web Socket Handshake

[web socket mdn](https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers)

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
    - lock 또는 mvcc(snapshot in postgre)로 구현
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

