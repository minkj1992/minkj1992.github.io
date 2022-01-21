# [Cryptozombies03] dAdvanced Solidity Concepts



{{< admonition quote>}}
나만의 좀d비 덱을 만들어보자.
{{< /admonition >}}

<!--more-->
<br/>



## [ch03] Advanced Solidity Concepts
> 챕터3를 통과하게 되면 [나만의 좀비 덱](https://share.cryptozombies.io/en/lesson/3/share/leoo?id=Y3p8MTcwMTU4)을 가지게 됩니다.

### `Ownable Contracts`

external function으로 setter를 열어두게 되면, 아무나 내 컨트랙트 안의 state variable을 수정할 수 있게 되는 보안적인 이슈가 생기게 된다. 이를 대처하기 위해 주로 사용하는 방식은 `contract`를 `ownable`하게 만들어 **특별한 권리를 가지는 특정 소유자가 있음을 지정할 수 있다.**

아래는 
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


{{< admonition tip Gas가 필요한 이유>}}
이더리움 진영에서는 이더림움을 `World Computer`라고 소개합니다. 

전세계에 퍼져있는 개별 노드들이 누군가가 만든 함수를 실행할 때 네트워크 상의 모든 노드 각각이 함수의 output을 검증하기 위해 그 함수를 실행해야 합니다. 

이더리움은 `Turing complete`하기 때문에 무한 루프와 같이 컴퓨팅 자원을 많이 소모되는 코드가 악의적으로 생성된다면 이더리움이라는 하나의 컴퓨터에 악영향을 끼칠것입니다. 이런 이유로 이더리움 개발자들은 연산 처리에 각각 비용을 할당했으며 사용자들은 space / time에 비례하여 gas를 지불해야 합니다.

+ 크립토 좀비에 따르면 `side-chain`에서는 반드시 gas를 지불하지는 않는다고 하네요, `Loom Network`를 사용하는 크립토 좀비가 대표적인 예시라고 합니다. 이더리움 메인넷에서 월드 오브 워크래프트 같은 게임을 직접적으로 돌리는 것은 절대 말이 되지 않기 때문이다. (엄청 높은 가스 비용) 하지만 다른 합의 알고리즘을 가진 사이드체인에서는 가능할 수 있지. 
{{< /admonition >}}


### `Gas`비 절약법

기본적으로 `uint256`이 아닌 `uint8`과 같은하위 타입들로 저장소를 절약하는 것은 아무런 이득이 없다고 합니다. 왜냐면 솔리디티에서 uint의 크기에 상관없이 `256bit` 저장공간을 미리 잡아두기 때문입니다.

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

// `mini`는 구조체 압축을 했기 때문에 `normal`보다 가스를 조금 사용하게 된다.
NormalStruct normal = NormalStruct(10, 20, 30);
MiniMe mini = MiniMe(10, 20, 30); 
```

이런 이유로, 구**조체 안에서는 가능한 작은 크기의 정수 타입을 쓰는 것이 좋다**고 할 수 있다.

또한 **동일한 데이터 타입은 하나로 묶어놓는 것이 좋다**. 즉, 구조체에서 서로 옆에 있도록 선언하면 솔리디티에서 사용하는 저장 공간을 최소화해준다고 합니다. 예를 들면, `uint c; uint32 a; uint32 b;`라는 필드로 구성된 구조체가 `uint32 a; uint c; uint32 b;` 필드로 구성된 구조체보다 가스를 덜 소모합니다. uint32 필드가 묶여있기 때문이지.

### `Time Units`

{{< admonition tip 2038년 문제 >}}
![year 2038 problem](https://ko.wikipedia.org/wiki/2038%EB%85%84_%EB%AC%B8%EC%A0%9C)  POSIX 시간 표기법은 시간을 1970년 1월 1일 자정 UTC 이후 경과된 초 시간을 이용하여 표현하는데,대부분의 32비트 시스템에서 초 시간을 저장하는 데 이용되는 time_t 자료 형식은 부호 있는 32비트 정수형이다. 즉 이 형식을 이용하여 나타낼 수 있는 **최후의 시각은 1970년 1월 1일 자정에서 정확히 2147483647초가 지난 2038년 1월 19일 화요일 03:14:07 UTC이다.** 이 시각 이후의 시각은 범위를 초과하여 내부적으로 음수로 표현되며, 프로그램의 이상 작동을 유발하는데, 왜냐하면 이러한 값은 2038년 대신 프로그램의 구현 방법에 따라 1970년 또는 1901년을 가리키기 때문이다
{{< /admonition >}}
