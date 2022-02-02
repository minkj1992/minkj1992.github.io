# Hardhat


`Hardhat`를 통해 local에서 스마트 컨트랙트를 작성해본 뒤, contract를 upgrade시켜봅니다. 이후 `rinkeby` 테스트넷에 proxy 컨트랙트를 배포해보고, `verify`시켜봅니다.

<!--more-->
<br />

환경에 사용된 레포는 [github link](https://github.com/minkj1992/hardhat_demo)입니다.

- `hardhat`과 `openzeppelin` 환경에서 간단한 Upgradable contract를 생성해봅니다.

## setup

먼저 프로젝트 환경을 로컬에서 세팅해줍니다.

### init

```shell
$ yarn init -y
$ yarn add hardhat --dev
$ yarn add @openzeppelin/hardhat-upgrades --dev
```

테스트는 다음과 같이 할 수 있습니다.

### test

```shell
$ npx hardhat test
```

### deploy

로컬에서 `ganache`같은 내부 노드를 실행할 수 있습니다.

```shell
$ npx hardhat node
```

로컬에서 이제 배포를 해봅시다.

- deploy

```shell
$ npx hardhat run --network localhost ./scripts/SimpleStorageUpgrade.deploy.js
SimpleStorageUpgrade deployed to: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
```

런타임에 콘솔에서 붙어서 제대로 배포되었는지 검증해봅니다.

- check

```shell
$ npx hardhat console --network localhost

Welcome to Node.js v14.15.1.
Type ".help" for more information.
> const f = await ethers.getContractFactory("SimpleStorageUpgrade")
undefined
> const ssu = await f.attach("0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0")
undefined
> ssu.address
'0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0'
> (await ssu.get()).toString()
'500'
> let tx = await ssu.set(1000)
undefined
> (await ssu.get()).toString()
'1000'
```

**다음과 같이 활용하면 스마트 컨트랙트를 upgrade 시켜줄 수 있습니다.**

- upgrade contract version

```shell
$ npx hardhat run --network localhost ./scripts/SimpleStorageUpgradeV2.deploy.js
Compiling 1 file with 0.8.4
Solidity compilation finished successfully
SimpleStorageUpgrade version 2 deployed to: 0x9fE46736679d2D9a65F0992F2272dE9f3c7fa6e0
```

### deploy to remote network

> [hardhat docs](https://hardhat.org/config/)

이제 실제 remote 환경에서 배포해보겠습니다.

#### hardhat.config.js

```js
require("dotenv").config(); // yarn add dotenv

// ... 중략 ...

module.exports = {
  networks: {
    rinkeby: {
      url: `https://eth-rinkeby.alchemyapi.io/v2/${process.env.ALCHEMY_API_KEY}`,
      // 0x20CE8B2190949f48F5D32d5BbbfE7E3760811F61
      accounts: [process.env.TEST_ACCOUNT_PRIVATE_KEY],
    },
  },
  solidity: "0.8.4",

  etherscan: {
    apiKey: process.env.ETHERSCAN_API_KEY,
  },
};
```

- `ALCHEMY_API_KEY`: [alchemy](https://www.alchemyapi.io)에서 demo app을 생성하게 될 경우 view key를 하면 확인가능합니다.

![](/images/hardhat/1.png)

- `TEST_ACCOUNT_PRIVATE_KEY`: metamask에서 `rinkeby` 네트워크에 계정을 생성한 뒤 아래와 같이 비공개키 export를 누르면 확인 가능합니다.

![](/images/hardhat/6.png)
![](/images/hardhat/7.png)

- `ETHERSCAN_API_KEY`
  1. [register etherscan](https://etherscan.io/register)에서 회원 가입을 한 뒤
  2. [create etherscan api key](https://etherscan.io/myapikey)에서 `My API Keys` > `+ Add`해주어서 얻어줍니다.

![](/images/hardhat/5.png)

추가로 [rinkeby faucet](https://faucets.chain.link/)에 들어가게되면 address 기반으로 rinkeby 계정에 이더를 넣어줄 수 있습니다.

![](/images/hardhat/2.png)
![](/images/hardhat/3.png)
![](/images/hardhat/4.png)

## deploy remote blockchain network

> `rinkeby` network (test-net)

```shell
$ npx hardhat run --network rinkeby ./scripts/SimpleStorageUpgrade.deploy.js
Downloading compiler 0.8.4
Compiling 3 files with 0.8.4
Solidity compilation finished successfully

SimpleStorageUpgrade deployed to: 0xCe93de7572e3346F1f91Ad39ce06e8F6c6312b69
```

테스트넷이기 때문에 약간의 시간이 소요됩니다. (약 30초) 이후 deploy된 address는 [rinkeby etherscan](https://rinkeby.etherscan.io/)에서 확인가능합니다.

- https://rinkeby.etherscan.io/address/0xCe93de7572e3346F1f91Ad39ce06e8F6c6312b69

[proxy contract check](https://rinkeby.etherscan.io/proxyContractChecker)

![](/images/hardhat/8.png)

> The implementation contract at 0x92a949706c10fd221b9a073f4284b4bdbc47e6d7 does not seem to be verified.

아직 implementation이 검증되지 않았다고 뜬다. 이렇게 proxy가 아닌 implementation contract를 검증해주기 위해서는 아래와 같이 추가 작업해주면 됩니다.

## verify implementation

> [docs](https://hardhat.org/plugins/nomiclabs-hardhat-etherscan.html)

먼저 hardhat에서 제공해주는 etherscan 의존성을 설치해줍니다.

```shell
$ yarn add @nomiclabs/hardhat-etherscan --dev
```

이후 아래 명령어를 통해서, **로컬의 deploy된 컨트랙트와 실제 rinkeby에 배포된 컨트랙트를 비교해서 검증해주는 로직을 타줍니다.**

```shell
//npx hardhat verify --network rinkeby "<리모트 배포된 implementation address>"
$ npx hardhat verify --network rinkeby "0x92a949706c10fd221b9a073f4284b4bdbc47e6d7"
Compiling 3 files with 0.8.4
Solidity compilation finished successfully
Compiling 1 file with 0.8.4
Successfully submitted source code for contract
contracts/SimpleStorageUpgrade.sol:SimpleStorageUpgrade at 0x92a949706c10fd221b9a073f4284b4bdbc47e6d7
for verification on the block explorer. Waiting for verification result...

Successfully verified contract SimpleStorageUpgrade on Etherscan.
https://rinkeby.etherscan.io/address/0x92a949706c10fd221b9a073f4284b4bdbc47e6d7#code
```

![](/images/hardhat/9.png)

이상으로 [배포된 컨트랙트](https://rinkeby.etherscan.io/address/0x92a949706c10fd221b9a073f4284b4bdbc47e6d7#code)에서 정상적으로 배포된 컨트랙트를 확인할 수 있습니다.

