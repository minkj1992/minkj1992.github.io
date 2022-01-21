# [Cryptozombies02] Zombies Attack Their Victims



{{< admonition quote>}}
좀비에게 먹이를 주어서 조합이 가능하게 해보자.
{{< /admonition >}}

<!--more-->
<br/>

## [ch02] Zombies Attack Their Victims

이번 장을 마치면 다음과 같은 고양이를 먹은 좀비를 생성할 수 있다.


![](images/cat_zombie2.png)

### `Mappings` and `Addresses`

이더리움 블록체인은 은행계좌와 같은 `account`를 사용해서 유저를 식별합니다. 
이때 각 `account`들은 이더리움 블록체인상의 coin인 `ether`를 단위로 `balance`를 가지게 됩니다.
이를 `통화`를 통해 각 계정은 송금/인출 등의 은행과 같은 기능들을 할 수 있습니다. 이를 위해 이더리움에서 각 계정은 은행 계좌 번호와 같은 `address`를 가지고 있으며, 

여기서 말하는 `address`는 `EOA`(Extenally Owned Account)의 address입니다. 보통은 `EOA`간의 메세지는 이더를 보내지만, EOA는 컨트랙트 어카운트에 메세지를 보내 해당 코드를 실행 시킬 수 도 있습니다. 

`Mapping`은 기본적으로 python의 `dict`와 같은 key-value 저장소입니다.

```sol
contract ZombieFactory {

    ...
    mapping (uint => address) public zombieToOwner;
    mapping (address => uint) ownerZombieCount;

    ...
}
```

### `Msg.sender`

`solidity`에는 모든 함수에서 이용 가능한 특정 전역 변수들이 있는데, 그 중의 하나가 **현재 함수를 호출한 사람 (혹은 스마트 컨트랙트)의 주소**를 가리키는 `msg.sender`이다.


{{< admonition tip>}}
`solidity`에서 함수 실행은 항상 `external caller`(외부 호출자)가 시작하며, 컨트랙트는 외부에서 함수를 호출  하기 전까지 블록체인 상에서 아무것도 하지 않는다.

즉 스마트 컨트랙트는 `msg.sender`(호출자)가 항상 존재합니다.
{{< /admonition >}}

```sol
mapping (address => uint) favoriteNumber;

function setMyNumber(uint _myNumber) public {
  favoriteNumber[msg.sender] = _myNumber;
}

function getMyNumber() public view returns (uint) {
  return favoriteNumber[msg.sender];
}
```

### `Require` 
특정 조건이 True가 아닐 경우, 에러를 발생시키고 함수를 벗어나게 됩니다. 

```sol
function sayHiToLeoo(string _name) public returns (string) {
  // solidity는 고유의 스트링 비교 기능이 없다. 그러므로 keccak256 해시값을 
  // 비교해 스트링이 같은 값인지 판단하는 코드
  require(keccak256(_name) == keccak256("Leoo.j"));
  
  return "Hi";
}
```

### `Inheritance`

```sol
contract Animal {
  function cry() public returns (string) {
    return "Default cry";
  }
}

contract Dog is Animal {
  function cry() public returns (string) {
    return "Bark";
  }
}
```

### `Import`

파일들로 코드를 분리하고, 다른 파일에 있는 코드를 불러오고 싶을 때, 솔리디티는 `import`라는 keyword를 사용합니다.

```sol
import "./someothercontract.sol"; // SomeOtherContract

contract newContract is SomeOtherContract {

}
```

### Storage vs Memory

`solidity`가 변수를 저장할 수 있는 공간에는 2가지 종류가 있습니다.

- `storage`
- `memory`

`Storage`는 블록체인 상에 영구적으로 저장되는 변수들입니다. `state variable`(함수 외부에 선언된 변수)인 경우 초기 설정상 `Storage`로 관리되어 블록체인 상에 영구적으로 저장됩니다.

이와 반대로 함수 내부에 선언된 변수는 `memory`로 자동 선언되어 함수 호출 종료시 사라지게 됩니다.

단 명시적으로 `storage`, `memory` 키워드들을 사용해주어야 하는 상황이 존재하는데, 바로 함수 내에서 `struct`, `배열`을 처리할 때 입니다.


```sol
contract SandwichFactory {
  struct Sandwich {
    string name;
    string status;
  }

  Sandwich[] sandwiches; // state variable (storage)

  function eat(uint _idx) public {
    string defaultState = "NOT EATEN"; // implicit memory
    Sandwich storage mySandwich = sandwiches[_idx]; // arr should explict

    Sandwich memory anotherSandwich = sandwiches[_idx + 1];
    sandwiches[_idx + 1] = anotherSandwich;
  }
}
```

```sol
pragma solidity ^0.4.19;

import "./zombiefactory.sol";

contract ZombieFeeding is ZombieFactory {

  function feedAndMultiply(uint _zombieId, uint _targetDna) public {
      require(msg.sender == zombieToOwner[_zombieId]);
      Zombie storage myZombie = zombies[_zombieId];
  }

}
```

### Extra Function Visibility

`solidity`에는 public과 private 이외에도 `internal`과 `external`이라는 함수 접근 제어자가 있다.

- `internal`
  - 상속하는 컨트랙트에서도 접근 가능 (java protected와 비슷해 보임?)
  - 나머지는 private과 동의
- `external`
  - 컨트랙트 바깥에서만 호출 될 수 있음
  - 컨트랙트 내의 다른 함수에 의해 호출될 수 없다.
  - 나머지는 public과 동의

`internal`은 상속하는 컨트랙트에서도 접근 가능하다는 점을 제외하면 private과 같다. 느낌 상 java의 `protected`와 유사해 보이며, `state variable`은 default로 internal 접근자를 가진다.

`external`은 **함수가 컨트랙트 바깥에서만 호출** 될 수 있고 **컨트랙트 내의 다른 함수에 의해서 호출 될 수 없다**는 부분만 제외하면 public과 같다.

```sol
contract Sandwich {
  uint private sandwichesEaten = 0;

  function eat() internal {
    sandwichesEaten++;
  }
}

contract BLT is Sandwich {
  uint private baconSandwichesEaten = 0;

  function eatWithBacon() public returns (string) {
    baconSandwichesEaten++;
    // eat 함수가 internal로 선언되었기 때문에 여기서 호출이 가능하다 
    eat();
  }
}
```

### interface

블록체인 상에서, 다른 컨트랙트와 상호작용을 하고 싶다면 `Interface`를 정의해야 합니다.

```sol
contract LuckyNumber {
  mapping(address => uint) numbers;

  function setNum(uint _num) public {
    numbers[msg.sender] = _num;
  }
  function getNum(address _myAddress) public view returns (uint) {
    return numbers[_myAddress];
  }
}
```
예를 들어 다음과 같은 외부 컨트랙트가 있다고 가정 할 때, 우리는 다음과 같은 interface를 만들 수 있습니다.

```sol
// 예시에서는 contract NumberInterface {}를 사용한다.
interface NumberInterface {
  function getNum(address _myAddress) public view returns (uint);
}
```

크립토 좀비에 제공된 예시에서는 contract NumberInterface {}를 사용하고 있는데, `interface`키워드가 추가 된 것인지 아니면 contract보다 interface가 제약조건이 많기 때문에 간단하게 contract로 구현했는지 모르겠지만, 좀 더 명확한 표현이 좋아서 예제를 변경하였습니다. 실제 interface 사용은 다음과 같습니다.

```sol
contract MyContract {
  address NumberInterfaceAddress = 0xab38...
  // ^ 이더리움상의 FavoriteNumber 컨트랙트 주소이다
  NumberInterface numberContract = NumberInterface(NumberInterfaceAddress)
  // 이제 `numberContract`는 다른 컨트랙트를 가리키고 있다.

  function someFunction() public {
    // 이제 `numberContract`가 가리키고 있는 컨트랙트에서 `getNum` 함수를 호출할 수 있다:
    uint num = numberContract.getNum(msg.sender);
    // ...그리고 여기서 `num`으로 무언가를 할 수 있다
  }
}
```

`interface`라는 키워드는 아래와 같은 제약조건이 있습니다.

- 다른 Contract로 부터 상속받을 수 없습니다, 하지만 다른 interface로부터는 상속받을 수 있습니다.
- 모든 function들은 `public`, `external`이어야 합니다.
- `constructor`를 선언할 수 없습니다.
- `variable`를 선언할 수 없습니다.
- `struct`를 선언할 수 없습니다.
- `enum`를 선언할 수 없습니다.
- 내부에는 `추상함수`, 즉 함수 시그니처만 존재합니다.
  - Interfaces cannot have any functions implemented


### summary 

최종적으로 다음과 같은 코드가 만들어집니다.

```sol
pragma solidity ^0.4.19;
import "./zombiefactory.sol";
contract KittyInterface {
  function getKitty(uint256 _id) external view returns (
    bool isGestating,
    bool isReady,
    uint256 cooldownIndex,
    uint256 nextActionAt,
    uint256 siringWithId,
    uint256 birthTime,
    uint256 matronId,
    uint256 sireId,
    uint256 generation,
    uint256 genes
  );
}
contract ZombieFeeding is ZombieFactory {

  address ckAddress = 0x06012c8cf97BEaD5deAe237070F9587f8E7A266d;
  KittyInterface kittyContract = KittyInterface(ckAddress);

  function feedAndMultiply(uint _zombieId, uint _targetDna, string _species) public {
    require(msg.sender == zombieToOwner[_zombieId]);
    Zombie storage myZombie = zombies[_zombieId];
    _targetDna = _targetDna % dnaModulus;
    uint newDna = (myZombie.dna + _targetDna) / 2;
    if (keccak256(_species) == keccak256("kitty")) {
      newDna = newDna - newDna % 100 + 99;  // 끝에 2자리를 99로 변경한다.
    }
    _createZombie("NoName", newDna);
  }

  function feedOnKitty(uint _zombieId, uint _kittyId) public {
    uint kittyDna;
    (,,,,,,,,,kittyDna) = kittyContract.getKitty(_kittyId);
    feedAndMultiply(_zombieId, kittyDna, "kitty");
  }
}
```

- 요구사항
  - 고양이 좀비(`kitty zombie`)는 DNA 마지막 2자리로 99를 갖는다고 가정한다. 그러면 우리 코드에서는 만약(if) 좀비가 고양이에서 생성되면 좀비 DNA의 마지막 2자리를 99로 설정한다.

#### 자바스크립트와 web3.js를 활용하여 우리의 컨트랙트와 상호작용하는 예시
```js
var abi = /* abi generated by the compiler */
var ZombieFeedingContract = web3.eth.contract(abi)
var contractAddress = /* our contract address on Ethereum after deploying */
var ZombieFeeding = ZombieFeedingContract.at(contractAddress)

// 우리 좀비의 ID와 타겟 고양이 ID를 가지고 있다고 가정하면 
let zombieId = 1;
let kittyId = 1;

// 크립토키티의 이미지를 얻기 위해 웹 API에 쿼리를 할 필요가 있다. 
// 이 정보는 블록체인이 아닌 크립토키티 웹 서버에 저장되어 있다.
// 모든 것이 블록체인에 저장되어 있으면 서버가 다운되거나 크립토키티 API가 바뀌는 것이나 
// 크립토키티 회사가 크립토좀비를 싫어해서 고양이 이미지를 로딩하는 걸 막는 등을 걱정할 필요가 없다 ;) 
let apiUrl = "https://api.cryptokitties.co/kitties/" + kittyId
$.get(apiUrl, function(data) {
  let imgUrl = data.image_url
  // 이미지를 제시하기 위해 무언가를 한다 
})

// 유저가 고양이를 클릭할 때:
$(".kittyImage").click(function(e) {
  // 우리 컨트랙트의 `feedOnKitty` 메소드를 호출한다 
  ZombieFeeding.feedOnKitty(zombieId, kittyId)
})

// 우리의 컨트랙트에서 발생 가능한 NewZombie 이벤트에 귀를 기울여서 이벤트 발생 시 이벤트를 제시할 수 있도록 한다: 
ZombieFactory.NewZombie(function(error, result) {
  if (error) return
  // 이 함수는 레슨 1에서와 같이 좀비를 제시한다: 
  generateZombie(result.zombieId, result.name, result.dna)
})
```


