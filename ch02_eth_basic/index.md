# [마스터링 이더리움 CH02] Ethereum Basics


이번장에서는 [마스터링 이더리움 CH02] 이더리움 기초를 정리합니다. 인용글들은 [원문](https://github.com/ethereumbook/ethereumbook/blob/develop/02intro.asciidoc)을 참조하였습니다.

<!--more-->
<br />

## tl;dr

이번 챕터에서는 실제 solidity코드를 작성한 뒤, 실제 eth를 전송하고 회수하는 contract를 테스트넷에 배포해보겠습니다.

이를 위해 크게 3가지를 다루게 됩니다.

1. 가장 먼저 Metamask wallet을 생성하고, Ropsten test network기반의 faucet으로 부터 ether를 받습니다.
2. 그 뒤 faucet contract코드를 solidity로 작성한 뒤 Remix를 활용해 EVM-bytecode로 compile한 뒤 Faucet contract on Ropsten 네트워크로 등록합니다. (이때 withdraw() 함수를 추가)
3. 마지막으로 Faucet contract address로 ether를 보낸 뒤, withdraw()를 실행해 봅니다.

## Ether Currency Units

> 이더 화폐 단위

이더리움의 화폐 단위는 `ether(이더)`라고 불리며, ETH 심볼을 사용합니다. 또한 이더의 최소 단위는 `wei`(웨이)라고 불립니다.

$$ 10^{18} wei = 1 ETH $$

- 이더리움 내부에서는 항상 웨이를 부호 없는 정수를 사용합니다.

> _Ethereum’s currency unit is called ether, identified also as "ETH" or with the symbols Ξ (from the Greek letter "Xi" that looks like a stylized capital E) or, less often, ♦_

## Choosing an Ethereum Wallet

`이더리움 지갑`이란 이더리움 계정을 관리하는 데 사용되는 소프트웨어 애플리케이션이라고 생각하시면 됩니다.

> _In short, an Ethereum wallet is your gateway to the Ethereum system. It holds your keys and can create and broadcast transactions on your behalf._

- 사용자의 개인키를 보관
- 사용자를 대신하여 트랜잭션 생성 및 브로드캐스트 시행

[이더리움 공식 홈페이지](https://ethereum.org/ko/wallets/)의 정의에 따르면 지갑을 아래와 같이 정의 내리고 있습니다.

{{< admonition quote >}}
지갑은 이더를 전송할 수 있고 보유할 수도 있는 애플리케이션입니다. 또한 이더리움 디앱에서도 사용할 수 있습니다.
{{< /admonition >}}

또한 현재(2022.01.28)기준으로 아래 지갑들을 추천합니다.

- 메타마스크(MetaMask) iOS와 Android용 브라우저 확장 프로그램 및 모바일 지갑
- 마이크립토(MyCrypto) 웹 기반 지갑
- 트러스트월렛(TrustWallet) iOS와 Android용 모바일 지갑
- 마이이더월렛(MyEtherWallet) 클라이언트 측 지갑
- 오페라(Opera) 지갑이 통합된 주요 브라우저

지갑의 대표주자인 `Metamask`와 `TrustWallet`를 [trust wallet vs metamask](https://viraltalky.com/trust-wallet-vs-metamask-comparison/)글을 토대로 요약 정리 해보겠습니다.

<center>

|                                  |               Trust Wallet               |     Metamask      |
| :------------------------------: | :--------------------------------------: | :---------------: |
|               Cost               |                   Free                   |       Free        |
|         Desktop Software         |                   Yes                    |        Yes        |
|            Mobile App            |                   Yes                    |        No         |
|        Built-in exchange         | Yes (`Kyber Network` and `Web3 browser)` |        No         |
|           NFT Support            |                   Yes                    |        No         |
|         Staking Options          |                   Yes                    |        No         |
|    Available Cryptocurrencies    |    Bitcoin, BNB, and all ERC20 Tokens    | All ERC-20 Tokens |
|             Security             |                  Medium                  |      Medium       |
| Compatible with hardware wallets |                    No                    |        Yes        |

</center>

아직 어떤게 더 좋아보이는지는 모르겠지만, 개인적으로 `metamask`에 개발 레퍼런스가 더 많이 있는 것 같아, **메타마스크를 사용해볼 예정입니다.**

## Wallet

> _Metamask with Ropsten test network_

![](/images/metamask/1.png)

`Chrome extension`에 `MetaMask`를 치면 메타마스크 애플리케이션을 크롬에 추가할 수 있습니다. 메타마스크 익스텐션 설치 및 가입이 끝났다는 전제하에 설명을 진행하도록 하겠습니다.

![](/images/metamask/2.png)

테스트를 하기 위해 `Ropsten`테스트 네트워크로 설정해줍니다.

![](/images/metamask/3.png)
"구매"를 누르고 "포시트(수도꼭지)테스트"에서 [Ether 얻기]를 눌러줍니다. 이 경우 아래와 같이 새로운 웹페이지가 열립니다.

> `Faucet` 파우셋(Faucet)이란 수도꼭지란 뜻으로 코인 무료 지급 하는곳 으로 사용되고 있다. 이더리움 생태계에서는 대표적으로 `Ropsten`, `Kovan`이 존재합니다.

![](/images/metamask/4.png)
![](/images/metamask/5.png)

새 웹페이지에서 `request 1 ether from faucet`(초록 버튼)을 클릭하게 되면 자신의 address로 1eth가 들어오게 됩니다.

![](/images/metamask/6.png)

몇 초간의 대기 시간이 지나면, 이더 지급이 완료된 `transaction`을 확인할 수 있습니다.

![](/images/metamask/7.png)

이 때 트랜잭션 링크를 누르게 되면 아래와 같은 EtherScan 링크로 이동 되고, 트랜잭션 상세내용을 확인하실 수 있습니다.

![](/images/metamask/8.png)
![](/images/metamask/9.png)

다음과 같이 faucet으로 돌려보내기 또한 가능합니다.

참고로 test 네트워크도 마찬가지로 `gas`비를 받는데요, 이는 real 이더 메인넷과 동일한 환경을 제공하기 위해서 입니다.

## Introducing the World Computer

지금까지 이더리움 지갑에 대해서 살펴보았습니다. 앞서 이야기 하였듯 이더리움은 cryptocurrency 기능외에도 튜링 complete한 하나의 컴퓨터입니다.

`Ether`는 `smart contract`를 사용하기 위해 소모되는 payment의 개념이며, 이런 `smart contract`프로그램은 emplated computer called `Ethereum Virtual Machine`(`EVM`)위에서 동작합니다.

> _The EVM is a global singleton, meaning that it operates as if it were a global, single-instance computer, running everywhere. Each node on the Ethereum network runs a local copy of the EVM to validate contract execution, while the Ethereum blockchain records the changing state of this world computer as it processes transactions and smart contracts._

EVM 요약

- Global singleton
- Each node validate broadcasted contracts execution
- Ethereum blockchain records changing state. (tx, smart contract)

## Externally Owned Accounts (EOAs) and Contracts

이더리움에는 2가지 타입의 account가 존재합니다.

- EOA
- Contract Account

> _The type of account you created in the MetaMask wallet is called an externally owned account (EOA)._

먼저 `EOA`란 사용자를 대변하는 account입니다.

> _Externally owned accounts are those that have a private key; having the private key means control over access to funds or contracts._

EOA는 사용자의 private key를 소유하고 있으며, 이는 contract 또는 account의 코인에 접근 권한이 있다는 뜻입니다.

- Has private key
- Has address
- Simple EOA can't have smart contract code

> _That other type of account is a contract account. A contract account has smart contract code, which a simple EOA can’t have. Furthermore, a contract account does not have a private key. Instead, it is owned (and controlled) by the logic of its smart contract code: the software program recorded on the Ethereum blockchain at the contract account’s creation and executed by the EVM_

`Contract Account`는 `스마트 컨트랙트의 주소`에 해당되며, 스마트 컨트랙트가 블록에 포함되어 배포될때 해당 스마트 컨트랙트에 대한 주소가 생성이 되며, 이 주소를 통해서 메세지 전송이나 특정함수를 실행 할 수 있습니다.

- Has smart Contract
- Has address
- Does not have private key
- Owned by smart contract itself

> _However, when a transaction destination is a contract address, it causes that contract to run in the EVM, using the transaction, and the transaction’s data, as its input. In addition to ether, transactions can contain data indicating which specific function in the contract to run and what parameters to pass to that function. In this way, transactions can call functions within contracts._

트랜잭션의 destination이 `contract address`일 경우, 이는 트랜잭션을 통해 컨트랙트가 EVM에서 실행되도록 트리거 합니다.

트랜잭션 안에는 특정 contract function안에 어떤 parameter가 전달되는지를 기록함으로 써 특정 위치의 contract의 함수가 어떤 인자를 가지고 실행해야 할지를 지정할 수 있습니다.

{{< admonition note "Private key가 없다는 의미" >}}
Note that because a contract account does not have a private key, it cannot initiate a transaction. Only EOAs can initiate transactions, but contracts can react to transactions by calling other contracts, building complex execution paths.
{{< /admonition >}}

## A Simple Contract: A Test Ether Faucet

```sol
// SPDX-License-Identifier: CC-BY-SA-4.0

// Version of Solidity compiler this program was written for
pragma solidity 0.6.4;

// Our first contract is a faucet!
contract Faucet {
    // Accept any incoming amount
    receive() external payable {} // receive는 키워드

    // Give out ether to anyone who asks
    function withdraw(uint withdraw_amount) public {
        // Limit withdrawal amount
        require(withdraw_amount <= 100000000000000000);

        // Send the amount to the address that requested it
        msg.sender.transfer(withdraw_amount);
    }
}
```

```sol
receive() external payable {} // receive는 키워드
```

`solidity 0.6` 버전 이후 `fallback`기능은 2가지로 나눠지게 되었습니다.

- `receive() external payable` — for empty calldata (and any value)
- `fallback() external payable` — when no other function matches (not even the receive function). Optionally payable.

{{< admonition tip >}}
**fallback**
a.k.a default function이라고도 불리며, 이름 그대로 대비책 함수입니다.

특징

1. 먼저 무기명 함수, 이름이 없는 함수입니다.
2. external 필수
3. payable 필수

왜 쓰는가 ?

1. 스마트 컨트랙이 이더를 받을 수 있게 한다.
2. 이더 받고 난 후 어떠한 행동을 취하게 할 수 있다.
3. call함수로 없는 함수가 불려질때, 어떠한 행동을 취하게 할 수 있다.
   {{< /admonition >}}

```sol
        msg.sender.transfer(withdraw_amount);
```

`msg` object는 one of the inputs로 모든 contracts가 접근 가능한 객체입니다. transaction이 실행되도록 trigger 시킨 주체를 의미합니다. 또한 attribute인 `sender`는 `sender address of the transaction`를 의미합니다. 마지막으로 `transfer()`는 built-in 함수로 `ether`를 current contract -> `누군가`.transfer()의 `누군가`에게 전달하는 것을 의미하며 이 코드에서 `누군가`는 `address of the sender`입니다. 즉 코드를 한줄로 설명하면 `contract --eth--> msg.sender`로 작동해라는 명령어 입니다.

> _This meas transfer ether from current contract to the sender of the msg that triggered this contract execution_

## Compiling the Faucet Contract

자 이제 우리가 처음 작성한 스마트 컨트랙트 코드를 Solidity Compiler를 통해 EVM bytecode로 변환을 하여 EVM에서 실행 될 수 있도록 만들어보겠습니다.

Solidity Compiler로는 대표적으로 아래의 것들이 있습니다. 저희는 이 중 solidity 공식문서에서 권유하는 `Remix IDE`를 사용해보겠습니다.

> We recommend Remix for small contracts and for quickly learning Solidity.

참고로 대안으로 급부상하고 있는 [`Hardhat`](https://hardhat.org/hardhat-network/#how-does-it-work)이라는 개발환경 또한 존재합니다.

- web3.js 대신 ([ethers.js](https://docs.ethers.io/v5/))를 default로 사용함.

Remix에서 코드를 작성한 뒤, Remix 좌측 2번째 탭을 클릭한 뒤, 적절한 compiler 버전(이번 예제는 0.6.4)를 설정해주고 compile Faucet.sol 버튼을 클릭해주면
아래와 같은 화면을 볼 수 있습니다.

![](/images/metamask/10.png)

## Creating the Contract on the Blockchain

> _Now, we need to “register” the contract on the Ethereum blockchain._

이제 robsten test 네트워크에 생성해준 contract를 등록해보겠습니다.

> _Registering a contract on the blockchain involves creating a special transaction whose destination is the address 0x0000000000000000000000000000000000000000, also known as the zero address. The zero address is a special address that tells the Ether‐ eum blockchain that you want to register a contract. Fortunately, the Remix IDE will handle all of that for you and send the transaction to MetaMask._

자 Remix의 3번째 탭(DEPLOY & RUN TRANSACTIONS)을 클릭하여 아래와 같이 세팅해줍니다. Account는 앞서 metamask에서 계정을 생성해주었다면 remix에서 metamask 요청페이지를 열어주어, 계정을 연결시켜줄 것입니다.

여기에서 "Deploy"버튼을 누르게 되면

<center>

![](/images/metamask/11.png)

</center>
이렇게 Deployed Contract가 등록요청하는 metamask 창이 열리고 확인을 누릅니다.

<center>

![](/images/metamask/12.png)

</center>

확인을 눌러주면 Remix상에서 contract가 생성된 것을 보실 수 있습니다.

우측의 copy버튼을 눌러 Contract address를 복사해서 etherscan에서 확인해보겠습니다.

<center>

![](/images/metamask/13.png)

</center>

[생성한 컨트렉트](https://ropsten.etherscan.io/address/0x8726C3D2F253332767Abd7268a11291df4A2f40d)에서 보여지듯이 잘 생성된 것을 확인하실 수 있습니다.
자 그럼 1eth를 해당 컨트랙트로 보내보겠습니다.

<center>

![](/images/metamask/15.png)
![](/images/metamask/16.png)
![](/images/metamask/17.png)
![](/images/metamask/18.png)
![](/images/metamask/19.png)

</center>

metamask를 통해서 이더를 보낸뒤, etherscan을 통해서 확인해보면 정상적으로 value 1eth가 전송된 것을 확인할 수 있습니다. 자 이제 튜토리얼의 마지막 단계인 0.1eth를 회수 해보겠습니다.

Remix의 버튼에 "100000000000000000" (10\*17 wei = 0.1eth)를 기입하고 withdraw버튼을 클릭해줍니다.

<center>

![](/images/metamask/20.png)

</center>

etherscan을 통해서 보면 다음과 같이 0.1 eth를 전송한 트랜잭션을 확인할 수 있습니다.

<center>

![](/images/metamask/21.png)

</center>

## Conclusion

이상으로 마스터링 이더리움 ch02인 기초적인 이더리움에 대해서 정리해보았습니다. 개인적으로 스마트 컨트랙트가 어떻게 배포되는 지 궁금했었는데
staging 개념으로 테스트를 해볼 수 있는 테스트넷이 있다는 점과, 실제 컨트랙트를 배포해서 metamask 계정과 연동해서 동작시켜볼 수 있었던 점이 재밌었던 것 같습니다.

<center> - 끝 - </center>

