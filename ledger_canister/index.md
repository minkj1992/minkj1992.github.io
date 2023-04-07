# The Ledger canister


[The Ledger canister from internet computer docs](https://internetcomputer.org/docs/current/references/ledger)
<!--more-->

## TL;DR

`Ledger canister`는 `IC principals`에 속한 accounts의 집합입니다.

각 account들은 token balance와 관련되어있으며, account owner들은 다른 account로 토큰 transfer를 진행할 수 있습니다.

모든 transfer 실행은 `append-only transaction ledger`로 기록됩니다.

`Ledger canister`의 interface는 token의 `minting`과 `burning`을 지원합니다.


## 1. Accounts

### 1.1 용어
- `IC principal`: `account` = 1: n
- `account` : `owner` = 1:1 (no joint accounts)

모든 account들은 한개의 principal에 속하며, principal은 1개이상의 account를 관리할 수 있습니다. `subaccount_identifier`를 통해 principal내에서 account를 구별합니다.

그러므로 논리적으로 각 `ledger account`는 `(account_owner`, `subaccount_identifier)` pair값과 상응합니다.

`account identifier`는 32-byte string으로 저장되며, 계산식은 아래와 값습니다.

```sh
account_identifier(principal, subaccount_identifier) = CRC32(h) || h
```

```sh
h = sha224(“\x0Aaccount-id” || principal || subaccount_identifier)
```

### Type

```rs
type Tokens = record {
     e8s : nat64;
};



// Account identifier  is a 32-byte array.
// The first 4 bytes is big-endian encoding of a CRC32 checksum of the last 28 bytes
type AccountIdentifier = blob;


//There are three types of operations: minting tokens, burning tokens & transferring tokens
type Transfer = variant {
    Mint: record {
        to: AccountIdentifier;
        amount: Tokens;
    };
    Burn: record {
         from: AccountIdentifier;
         amount: Tokens;
   };
    Send: record {
        from: AccountIdentifier;
        to: AccountIdentifier;
        amount: Tokens;
    };
};

type Memo = u64;

// Timestamps are represented as nanoseconds from the UNIX epoch in UTC timezone
type TimeStamp = record {
    timestamp_nanos: nat64;
};

Transaction = record {
    transfer: Transfer;
    memo: Memo;
    created_at_time: Timestamp;
};

Block = record {
    parent_hash: Hash;
    transaction: Transaction;
    timestamp: Timestamp;
};

type BlockIndex = nat64;

//The ledger is a list of blocks
type Ledger = vec Block
```

- `amount`: 전송할 토큰의 양
- `fee`: 송금 시 지불해야 하는 수수료
- `from_subaccount`: ICP가 발생해야 하는 호출자의 계정을 지정하는 하위 계정 식별자입니다. 이 매개변수는 선택 사항입니다. --- 호출자가 지정하지 않으면 모두 0 벡터로 설정됩니다.
- `to`: 토큰을 전송해야 하는 계정 식별자
- **`memo`: 발신자가 선택한 64비트 숫자입니다. 예를 들어 특정 전송을 식별하기 위해 다양한 방법으로 사용할 수 있습니다.**

### Burning token

`minting account`로 토큰을 transfer하는 것은, 토큰을 간단히 remove시키는 방식입니다. 즉 token을 `burn`하는 것입니다. 

Burn transaction은 ledger에 `(Burn(from, amout))`형태로 저장되며, burn transaction fee는 0입니다. 대신 burn되어져야 할 token의 amount는 `standard_fee`를 초과해야 transaction이 동작합니다.

{{< admonition note "burning transaction" >}}
_A burning transaction is the process of "burning" ICP, whereby a certain amount of ICP are destroyed._

**The main use case is that of purchasing cycles, through which ICP are destroyed while at the same time a corresponding amount of cycles is created, using the current exchange rate between ICP and ( SDR), in such a way that one SDR corresponds to one trillion (10E12) cycles.** 

_It is represented as a transaction from the source account to the ICP supply account._
{{< /admonition  >}}


## 2. Ledger block

**확장성을 위해 ledger canister는 entire ledger transactions들을 저장하지 않습니다.** 대신 ledger canister는 가장 최근 블록으로 구성된 suffix of the ledger를 저장하고 있습니다. 그리고 나머지 블록들은 모두 `archive canisters`에 저장됩니다.

## 3. Conclusion

Ledger canister는 ICP token관련된 block들을 저장하며, 아래 3가지 기능을 interface로 제공합니다.

- `transfer()`
- `burn()` / `mint()`
- `balance()`

또한 확장성을 위해 `archive canisters`에 나머지 블록들을 저장해두고, 최신 블록의 suffix만을 저장해두어 light weight합니다. 아마 이게 가능한 이유는 icp만의 특별한 encryption 덕분이며 더 자세한 암호화 방식은 추후 공부해야 할 것 같습니다.

최근에 완성된 `invoice canister`는 이 `ledger canister`위에서 동작하는 high level canister이며, BTC / ETH / ICP 각 canister들의 payment를 추상화시켜줘 더 편리하게 처리할 수 있습니다.

`plug`, `infinityWallet`, `stoic`등의 `wallet`들은 이런 ledger의 `transfer`기능을 추상화시켜서 처리해주는 것 같습니다.
