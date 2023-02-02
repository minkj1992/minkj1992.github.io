# Internet Computer development environment


How to set up Internet Computer development environment.
<!--more-->
## TL;DR
1. Install the Canister SDK
2. Build and deploy a dapp locally
3. Collect free cycles to power your dapp
4. Create a "cycles wallet" from which you can transfer cycles to any other dapps you want to power
5. Deploy a dapp on-chain


## 1. Set up infrastructure

### 1-1. Get Cycles Faucet
> [Get Cycles Faucet](https://anv4y-qiaaa-aaaal-qaqxq-cai.ic0.app/)

<center>

![](images/icp1.png)

</center>

- Request 20T cycles from [ICP Discord](https://discord.gg/4scZ5j2UJz)

<center>

![](images/icp3.png)

![](images/icp2.png)
</center>



### 1-2. Setup SDK
> [Setup SDK](https://internetcomputer.org/docs/current/developer-docs/build/install-upgrade-remove)

`dfx`: **D**i**f**inity e**x**ecution command-line interface

```js
$ sh -ci "$(curl -fsSL https://internetcomputer.org/install.sh)"
$ dfs --version
dfx 0.12.1
```



### 1-3. Claim cycles

```js
> dfx identity list
Creating the "default" identity.
WARNING: The "default" identity is not stored securely. Do not use it to control a lot of cycles/ICP.
To create a more secure identity, create and use an identity that is protected by a password using the following commands:
    dfx identity create <my-secure-identity-name> # creates a password protected identity
    dfx identity use <my-secure-identity-name> # uses this identity by default

Error: Failed to load identity manager.
Caused by: Failed to load identity manager.
    Cannot create identity directory at '...': Permission denied (os error 13)
```

1. Generate [Identity](https://support.dfinity.org/hc/en-us/articles/7453712440084-What-are-identities-) and set it default

```js
> sudo rm -rf leoo
> dfx identity list
anonymous
default *

# dfx identity new <my-secure-identity-name>
> sudo dfx identity new leoo.j
> dfx identity list
anonymous
default *
leoo.j

# dfx identity use <my-secure-identity-name> # uses this identity by default
> sudo dfx identity use leoo.j
Using identity: "leoo.j".
```

2. Claim Cycles

```js
> sudo dfx wallet --network ic redeem-faucet-coupon <COUPON_NUMBER>
Please enter the passphrase for your identity: [hidden]
Decryption complete.

> sudo dfx wallet --network=ic balance
Please enter the passphrase for your identity: [hidden]
Decryption complete.
20.099 TC (trillion cycles).
```

<center>

![](images/icp4.png)

</center>

## 2. Hello world locally


<center>

![](images/icp6.png)

</center>


I faced some permission error while doing `Hello_World` canister. So I posted below issue on ICP forum.

- [Issue that I faced](https://forum.dfinity.org/t/permissions-dfx-cli-on-osx/18220?u=leoo.j)

```js
> dfx new hello
> cd hello
```

- Terminal A
```js
// base root is hello
> dfx start
Running dfx start for version 0.12.1
Using the default definition for the 'local' shared network because /Users/minwook/.config/dfx/networks.json does not exist.
Dashboard: http://localhost:56958/_/dashboard
```

- Terminal B

```js
// base root is hello
> yarn install // or npm install
> dfx deploy

...

Deployed canisters.
URLs:
  Frontend canister via browser
    hello_frontend: http://127.0.0.1:4943/?canisterId=rkp4c-7iaaa-aaaaa-aaaca-cai
  Backend canister via Candid interface:
    hello_backend: http://127.0.0.1:4943/?canisterId=rno2w-sqaaa-aaaaa-aaacq-cai&id=r7inp-6aaaa-aaaaa-aaabq-cai
```

{{< admonition fix "if you faced frontend build hangs" >}}
[How to fix dfx deploy infinite hang](https://forum.dfinity.org/t/dfx-deploy-locally-with-a-new-dfx-identity/16470/24?u=leoo.j)

1. Open package.json file.
2. edit scripts like below.

- as-is
```json
  "scripts": {
    ...
    "generate": "dfx generate hello_backend"
  },
```

- to-be
```json
  "scripts": {
    ...
    "generate": "dfx --identity anonymous generate hello_backend"
  },
```
{{< /admonition  >}}


Finally Open your browser and navigate to the url output by the `dfx deploy`. 
In my case, it is `http://127.0.0.1:4943/?canisterId=rkp4c-7iaaa-aaaaa-aaaca-cai`.


<center>

![](images/icp5.png)

</center>

## 3. (OPT) Check my public wallet dashboard

You can check your wallet dashboard by below command.

```js
> dfx identity --network ic get-wallet
Please enter the passphrase for your identity: [hidden]
Decryption complete.
zze7b-pqaaa-aaaam-abciq-cai
```

With given wallet ID, you can browse your public dashboard site. 
Type **https://<wallet-id>.ic0.app/** on your web browser.


<center>

![](images/icp7.png)

[My public ICP Dashboard](https://zze7b-pqaaa-aaaam-abciq-cai.ic0.app/)

<p>- fin -</p>
</center>
