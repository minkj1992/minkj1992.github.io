# [Cryptozombies03] Advanced Solidity Concepts



{{< admonition quote>}}
나만의 좀비 덱을 만들어보자.
{{< /admonition >}}

<!--more-->
<br/>



## [ch03] Advanced Solidity Concepts
> 챕터3를 통과하게 되면 [나만의 좀비 덱](https://share.cryptozombies.io/en/lesson/3/share/leoo?id=Y3p8MTcwMTU4)을 가지게 됩니다.

### `Ownable Contracts`

external function으로 setter를 열어두게 되면, 아무나 내 컨트랙트 안의 state variable을 수정할 수 있게 되는 보안적인 이슈가 생기게 된다. 이를 대처하기 위해 주로 사용하는 방식은 `contract`를 `ownable`하게 만들어 **특별한 권리를 가지는 특정 소유자가 있음을 지정할 수 있다.**
 
```sol
/**
 * @title Ownable
 * @dev The Ownable contract has an owner address, and provides basic authorization control
 * functions, this simplifies the implementation of "user permissions".
 */
contract Ownable {
  address public owner;
  event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

  /**
   * @dev The Ownable constructor sets the original `owner` of the contract to the sender
   * account.
   */
  function Ownable() public {
    owner = msg.sender;
  }

  /**
   * @dev Throws if called by any account other than the owner.
   */
  modifier onlyOwner() {
    require(msg.sender == owner);
    _;
  }

  /**
   * @dev Allows the current owner to transfer control of the contract to a newOwner.
   * @param newOwner The address to transfer ownership to.
   */
  function transferOwnership(address newOwner) public onlyOwner {
    require(newOwner != address(0));
    OwnershipTransferred(owner, newOwner);
    owner = newOwner;
  }
}
```

`function Ownable()`는 `Constructor`(생성자)입니다. 컨트랙트와 동일한 이름을 가졌으며 default로 제공되어 특별한 작업을 할 게 아니라면 생략가능합니다. 생성자는 컨트랙트 생성시 단 한번만 실행됩니다.

`modifier onlyOwner()`에서 modifier는 `function modifier`(함수 제어자)입니다. 함수에 대한 접근을 제어하기 위해 사용되는 함수의 일종으로, **보통 함수 실행 전 요구사항 충족여부를 확인하는 데 사용됩니다.**

예시의 `onlyOwner()` 함수는 컨트랙트의 소유자에 한해서만 해당 함수를 실행할 수 있도록 하기 위해 제어해주는 기능을 해줍니다. 즉 `transferOwnership`(소유권 이전) 함수는 onlyOwner 조건을 만족시킬 때만 실행됩니다.

`_` 키워드는 쉽게 modifier 검사를 마친 뒤, 실행 될 함수가 들어가게 된다 생각하면 됩니다.

`indexed` 키워드에 대해서는 추후에 더 알아보겠습니다.

### `Gas`
> 이더리움 DApp이 사용하는 연료

솔리디티에서는 사용자들이 만든 DApp의 함수를 실행할 때마다 `Gas`라 불리는 화폐(ETH, 이더)를 지불해야합니다. 엄밀히 말해서는 사용자가 `ETH`(이더)를 이용해 `Gas`를 구매한다.

`Gas`비는 연산비용에 따라 다릅니다. 즉 함수의 로직이 얼마나 복잡한지에 따라 연산이 소모되는 gas cost가 상승합니다.

이런 시스템이기 때문에 코드 최적화가 암묵적으로 강제된다 할 수 있습니다. 가스는 함수를 실행하는 사용자들이 실제 돈을 쓰기 때문에 코드 최적화가 되지 않았다면 당연히 많은 사용자들이 생성한 코드를 사용하지 않게 됩니다.


{{< admonition tip Gas가필요한이유>}}
이더리움 진영에서는 이더림움을 `World Computer`라고 소개합니다. 

전세계에 퍼져있는 개별 노드들이 누군가가 만든 함수를 실행할 때 네트워크 상의 모든 노드 각각이 함수의 output을 검증하기 위해 그 함수를 실행해야 합니다. 

이더리움은 `Turing complete`하기 때문에 무한 루프와 같이 컴퓨팅 자원을 많이 소모되는 코드가 악의적으로 생성된다면 이더리움이라는 하나의 컴퓨터에 악영향을 끼칠것입니다. 이런 이유로 이더리움 개발자들은 연산 처리에 각각 비용을 할당했으며 사용자들은 space / time complexity에 비례하여 gas를 지불해야 합니다.

추가로 크립토 좀비에 따르면 `side-chain`에서는 반드시 gas를 지불하지는 않는다고 하네요, `Loom Network`를 사용하는 크립토 좀비가 대표적인 예시라고 합니다. 
이더리움 메인넷에서 롤 같은 게임을 직접 돌리게 되면 말도 안되게 엄청 높은 가스 비용이 들테니까요. 하지만 다른 합의 알고리즘을 가진 사이드체인에서는 가능하다고 합니다. 
{{< /admonition >}}


### `Gas`비 절약법

기본적으로 `uint256`이 아닌 `uint8`과 같은하위 타입들로 저장소를 절약하는 것은 아무런 이득이 없다고 합니다. 
왜냐면 솔리디티에서 uint의 크기에 상관없이 `256bit` 저장공간을 미리 잡아두기 때문입니다.

단 `struct` 안에서 `uint`를 사용한다면 더 작은 크기를 사용할 때, storage 절약이 가능하다고 합니다.

```sol
struct NormalStruct {
  uint a;
  uint b;
  uint c;
}

struct MiniMe {
  uint32 a;
  uint32 b;
  uint c;
}

// `mini`는 구조체 압축을 했기 때문에 `normal`보다 가스를 조금 사용하게 됩니다.
NormalStruct normal = NormalStruct(10, 20, 30);
MiniMe mini = MiniMe(10, 20, 30); 
```

이런 이유로, **구조체 안에서는 가능한 작은 크기의 정수 타입을 쓰는 것이 좋다**고 할 수 있습니다.또한 **동일한 데이터 타입은 하나로 묶어놓는 것이 좋습니다**. 

즉 구조체에서 서로 가까이 있도록 선언하면 솔리디티에서 사용하는 저장 공간을 최소화해줍니다. 
예를 들면, `uint c; uint32 a; uint32 b;`라는 필드로 구성된 구조체가 `uint32 a; uint c; uint32 b;` 필드로 구성된 구조체보다 uint32 필드들이 묶여있기 때문에 가스를 덜 소모합니다. 

### `Time Units`

Solidity provides some native units for dealing with time.
- `now`

`now`를 사용하게 되면 unix timestamp(1970년 1월 1일부터 지금까지의 초 단위 합)을 `uint256`타입으로 얻을 수 있습니다.
참고로 `unix time`은 전통적으로 32bit로 저장되는데 이 경우 `Year 2038` 문제가 발생할 것입니다. 만약 우리 `DApp`이 2038년까지 운영되길 원한다면 어쩔 수 없이 `64bit`를 써야하지만, trade of로 유저들은 저장하는데 더 많은 gas를 소모하게 됩니다.

{{< admonition tip 2038년_문제 >}}
[year 2038 problem](https://ko.wikipedia.org/wiki/2038%EB%85%84_%EB%AC%B8%EC%A0%9C)란? 

POSIX 시간 표기법은 시간을 1970년 1월 1일 자정 UTC 이후 경과된 초 시간을 이용하여 표현하는데,대부분의 32비트 시스템에서 초 시간을 저장하는 데 이용되는 time_t 자료 형식은 부호 있는 32비트 정수형이다. 즉 이 형식을 이용하여 나타낼 수 있는 **최후의 시각은 1970년 1월 1일 자정에서 정확히 2147483647초가 지난 2038년 1월 19일 화요일 03:14:07 UTC이다.** 이 시각 이후의 시각은 범위를 초과하여 내부적으로 음수로 표현되며, 프로그램의 이상 작동을 유발하는데, 왜냐하면 이러한 값은 2038년 대신 프로그램의 구현 방법에 따라 1970년 또는 1901년을 가리키기 때문입니다.
{{< /admonition >}}

```sol
uint lastUpdated;

// `lastUpdated`를 `now`로 설정
function updateTimestamp() public {
  lastUpdated = now;
}

// 마지막으로 `updateTimestamp`가 호출된 뒤 5분이 지났으면 `true`를, 5분이 아직 지나지 않았으면 `false`를 반환
function fiveMinutesHavePassed() public view returns (bool) {
  return (now >= (lastUpdated + 5 minutes));
}
```

### Passing structs as arguments
**솔리디티에서는 `private` 또는 `internal` 함수의 인자로서 구조체의 `storage 포인터`를 전달할 수 있습니다.**

이때 구조체는 포인터타입이며, 솔리디티에서는 이를 `storage pointer`라고 부르고 있습니다. 문득 `memory pointer`명칭도 존재하는지 찾아보니 서치하지 못한걸 보면 `storage pointer`라는 명칭만 있는 것 같습니다.

`storage pointer`라는 명칭을 처음 접해서 개념을 정리하기 위해서 이곳저곳을 찾다, [Storage Pointers in Solidity](https://blog.b9lab.com/storage-pointers-in-solidity-7dcfaa536089)라는 글을 읽었습니다. 이해한 부분까지 정리해보면 `struct`타입은 기본적으로 pointer 타입인 것고 이를 function에서 local variable로 참조해서 사용하면 `storage`형태로 저장되는 것 같습니다. 

아래의 코드를 보면

```sol
contract FirstSurprise {
 
 struct Camper {
   bool isHappy;
 }
 
 mapping(uint => Camper) public campers;
 
 function setHappy(uint index) public {
   campers[index].isHappy = true;
 }
 function surpriseOne(uint index) public {
   Camper c = campers[index];
   c.isHappy = false;
 }
}

```

`setHappy`를 통하지 않고도, `surpriseOne()`의 `Camper c = campers[index]` c가 `storage pointer`타입이기 때문에 side-effect가 생길 수 있다는 점이 핵심인 듯합니다.

최근에 이런 목소리를 반영해서 solidity compiler는 이런 상황일 때 아래와 같은 warning을 띄워준다고 합니다.

> Variable is declared as storage pointer. Use explicit “storage” keyword to silence this warning.

이런 맥락에서 크립토 좀비에서는 **구조체를 함수의 인자로 전달하면 storage pointer**라고 설명하고 있으며, 아래와 같이 함수 인자에 **명시적으로 storage를 쓰도록 하고 있습니다.**

```sol
function _doStuff(Zombie storage _zombie) internal {
  // _zombie로 할 수 있는 것들을 처리
}
```
그럼 이런 배경지식을 기억한채 다시 좀비로 넘어와보겠습니다.
우리는 좀비들이 끊임없이 kitty를 먹고 증식하는 것을 막기 위해서 `feedAndMultiply`에 다음 제약조건을 추가해보려고 합니다.

- 먹이를 먹으면 좀비가 재사용 대기에 들어간다.
- 좀비는 재사용 대기 시간이 지날 때까지 고양이들을 먹을 수 없다.

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

  KittyInterface kittyContract;

  function setKittyContractAddress(address _address) external onlyOwner {
    kittyContract = KittyInterface(_address);
  }

  function _triggerCooldown(Zombie storage _zombie) internal {
    _zombie.readyTime = uint32(now + cooldownTime);
  }

  function _isReady(Zombie storage _zombie) internal view returns (bool) {
      return (_zombie.readyTime <= now);
  }

  function feedAndMultiply(uint _zombieId, uint _targetDna, string _species) internal {
    require(msg.sender == zombieToOwner[_zombieId]);
    Zombie storage myZombie = zombies[_zombieId];
    require(_isReady(myZombie)); // 새로 추가 된정보
    _targetDna = _targetDna % dnaModulus;
    uint newDna = (myZombie.dna + _targetDna) / 2;
    if (keccak256(_species) == keccak256("kitty")) {
      newDna = newDna - newDna % 100 + 99;
    }
    _createZombie("NoName", newDna);
    _triggerCooldown(myZombie); // 새로 추가 된정보
  }

  function feedOnKitty(uint _zombieId, uint _kittyId) public {
    uint kittyDna;
    (,,,,,,,,,kittyDna) = kittyContract.getKitty(_kittyId);
    feedAndMultiply(_zombieId, kittyDna, "kitty");
  }
}
```

우선 `Zombie` storage pointer를 인자로 받는 `_isReady()`, `_triggerCooldown()`함수를 만듭니다.

-  `_isReady()`: 좀비가 재사용 대기시간을 넘겼는지 확인
-  `_triggerCooldown()`: 좀비가 kitty와 조합(eat) 되었다면, 좀비의 readyTimed을 now + cooldownTime(1일)로 업데이트 해줍니다.

이후 좀비에게 먹이를 공급하는 `feedAndMultiply()`함수에 아무나 접근하지 못하도록 `internal`로 함수를 지정해줍니다.

### Function modifiers with arguments
> 앞서 `modifier onlyOwner`같은 커스텀 function modifier를 보았는데, 이에 더해 function modifier에 argument를 넣어주는 법을 배워봅시다.

```sol
// usrId => age mapping
mapping (uint => uint) public age;

modifier olderThan(uint _age, uint _userId) {
  require(age[_userId] >= _age);
  _;
}

function buyCigarette(uint _userId) public olderthan(19, _userId) {
  🚬()
}
```

위의 코드는 담배를 판매하는 간단한 contract입니다. `functio nmodifier`의 인자로 나이와 userId를 제공하여 나이를 검사를 구현해주었습니다.

이 기능을 활용하여 우리의 `zombie`에게 level 속성을 부여해보고, 속성에 따라서 아래와 같은 능력치 제한을 두는 modifier를 만들어보겠습니다.

- 레벨 2 이상인 좀비인 경우, 사용자들은 그 좀비의 이름을 바꿀 수 있네.
- 레벨 20 이상인 좀비인 경우, 사용자들은 그 좀비에게 임의의 DNA를 줄 수 있네.


#### zombieHelper.sol
```sol
pragma solidity ^0.4.19;

import "./zombiefeeding.sol";

contract ZombieHelper is ZombieFeeding {

  modifier aboveLevel(uint _level, uint _zombieId) {
    require(zombies[_zombieId].level >= _level);
    _;
  }

  function changeName(uint _zombieId, string _newName) external aboveLevel(2, _zombieId) {
    require(msg.sender == zombieToOwner[_zombieId]);
    zombies[_zombieId].name = _newName;
  }

  function changeDna(uint _zombieId, uint _newDna) external aboveLevel(20, _zombieId) {
    require(msg.sender == zombieToOwner[_zombieId]);
    zombies[_zombieId].dna = _newDna;
  }
}
```

### Saving Gas With 'View' Functions
> View functions don't cost gas

**view 함수는 사용자에 의해 외부에서 호출되었을 때 가스를 전혀 소모하지 않는다.**

블록체인에 상태를 기록한다는 것은, 모든 `single node`들에게 트랜잭션이 추가되어야 한다는 것을 의미합니다. 하지만 반대로 view / pure function의 경우 블록체인 상에 어떤 것도 수정하지 않기 때문에 gas 소모가 없습니다.

만약 web3.js에게 view function를 호출해달라 요청하는 것은 실제로는 로컬 이더리움 노드에 query만 날리면 되기 때문에 가스 소모가 없게 됩니다.



{{< admonition warning >}}
앞부분 설명을 보다보니 문득 view function gas가 들지 않는다면, view function을 infinite 호출하게되면 이더리움 망가뜨릴수 있지 않을까하는 생각에 검색하게 되었고 확인해보니 pure / view function은 internally call 해주게 되면 gas비가 든다고 한다. 즉 크립토좀비가 이번 세션에서 설명하는 것은 blockchain 외부(i.g web3.js)에서 호출하면 free gas cost라는 의미이다.

**Pure and view functions still cost gas if they are called internally** from another function. They are only free if they are called externally, from outside of the blockchain.

This [View/Pure Gas usage - Cost gas if called internally by another function?](https://ethereum.stackexchange.com/questions/52885/view-pure-gas-usage-cost-gas-if-called-internally-by-another-function/52887#52887) goes into greater depth on this topic.

{{< /admonition >}}

자세히 보니 크립토 좀비의 **참고**에도 아래와 같은 hint가 작성되어있네요. (데헷 😧)

{{< admonition tip >}}
만약 view 함수가 동일 컨트랙트 내에 있는, view 함수가 아닌 다른 함수에서 내부적으로 호출될 경우, 여전히 가스를 소모할 것이네. 이것은 다른 함수가 이더리움에 트랜잭션을 생성하고, 이는 모든 개별 노드에서 검증되어야 하기 때문이네. 그러니 view 함수는 외부에서 호출됐을 때에만 무료라네.
{{< /admonition >}}

이제 우리의 좀비 DApp에 사용자의 전체 좀비 군대를 볼 수 있는 메소드를 추가해보자. `getZombiesByOwner()`라는 네이밍에 `external view function`으로 만들어 보겠습니다.



#### Declaring arrays in memory

솔리디티에서 `storage`에 write하는 것은 비싼 연산 중 하나입니다. 이더리움은 `World computer`이기 때문에 main-net기준으로 storage를 사용할 경우, 연결되어 있는 전세계 수많은 node들에 update를 시키게 되기 때문이죠. 이러다 보니 대부분의 프로그래밍 언어가 크기가 상당한 collection에 각각 접근( `O(N)` )하는 것을 지양하는 것과 달리, 솔리디티는 그 접근이 `external view`함수라면 storage를 쓰는 것보다 `memoery`를 써서 각각 element에 접근하는 것이 더 저렴한 방법입니다. (이는 gas비 때문인데, 훗날 이더리움 가격이 떨어진다면 달라질지도)

{{< admonition tip >}}
생각해보니 실제로 storage에 write하는 것은 O(N) * per_gas_cost는 아닌것 같네요.

만약 1만명이 사용하는 contract라고 가정했을 떄 N=element갯수라면, `gas_cost = (O(N) * 하나의_write_연산에_사용되는_gas_cost) * 10000`이 되기 때문에 만약 1만명이 아닌 사용하고자 하는 사람의 숫자가 많아진다면, 즉 `if 10000 >= storage's element size`라면 `N`의 정의가 달라지게 될 것 같습니다.
{{< /admonition >}}

`Storage`에 아무것도 쓰지 않고도 함수 안에서 새로운 배열을 만들기 위해서는 `memory`키워드를 사용하면 됩니다. 이는 `storage`배열을 직접 업데이트하는 것보다 gas_cost 측면에서 훨씬 (크립토 좀비에 따르면)  저렴하다고 합니다. 그러므로 **collection을 storage로 관리하지 말고 memory로 전환하여 관리합시다.**

```sol
pragma solidity ^0.4.19;

import "./zombiefeeding.sol";

contract ZombieHelper is ZombieFeeding {

  ... 중략 ...

  function getZombiesByOwner(address _owner) external view returns(uint[]) {
    uint[] memory result = new uint[](ownerZombieCount[_owner]);
    return result;
  }

}

```



