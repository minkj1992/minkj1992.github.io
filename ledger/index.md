# Ledger


How to setup my nano product.
<!--more-->

원래 거래소를 믿지 않기도 했고, 2022년 FTX 사태로 특히 더욱 Cold wallet의 중요성이 부각된 것 같아, 미국에 지내는 동안 $79.00에 `Nano S plus`를 구입하여 세팅하였다. 

앞으로 거래 시 `DeX`(Decentralized exchange, P2P방식의 분산형 암호화폐 거래소)의 특성을 가진 거래소에서 거래를 할 생각이며, 편의를 위할 때는 Ledger 인프라를 사용할 것 같다.

## TL;DR

1. PIN CODE
2. Recovery letters (24 words)
3. Install Ledger Live
4. (opt) Update firmware.
5. Install coin apps to nano device.
6. Add your accounts

## 1. `Cold Wallet` Background
> What you should basically know about.

- [x] My crypto assets are stored on the blockchain (p2p, network).
- [x] I need a private key to access and manage them.
- [x] My private key is stored within my Nano
- [x] Your Nano works as a "cold storage" wallet. This means that **it never exposes my private key online even when using the app**.
- [x] Validate transactions: `Ledger live` allows you to buy, sell, manage, exchange and earn crypto while remaining protected. You will validate every crypto transaction with your Nano.

---

## 2. How to set Nano S plus in Desktop?
> With Ledger live desktop app.

### 1. PIN CODE

- Nano S plus에서 4-8자리의 숫자 PIN코드를 설정합니다.
- PIN code를 세번 틀리게 될 경우, nano device는 reset됩니다.  

{{< admonition note "PIN code" >}}
_Your PIN code is the first layer of security. It physically secures access to your private key and your Nano. Your PIN code must be 4 to 8 digits long._

1. Don't share it.
2. You can change your PIN code if needed.
3. **Three wrong PIN code entries in a row will reset the device**.
4. Never store your PIN code on a computer or phone.
{{< /admonition  >}}

<center>

![](/images/nano_pincode.png)

</center>

### 2. Recovery letters (24 words)
> 준비물: 안전하게 24 영어단어들을 적어둘 공간.

- `Nano S plus`에서 보여주는 24개의 단어 리스트 (recovery letters)를 확인하며, 이를 안전한 곳에 적어둡니다.

<center>

![](/images/nano_recovery.png)

</center>


{{< admonition note "Recovery phrase" >}}
_Your recovery phrase is a secret list of 24 words that backs up your private keys. Your Nano generate a unique recovery phrase. Ledger does not keep a copy of it._

- **If you lose this recovery phrase, You will not be able to access your crypto in case You lose access to your Nano.**
{{< /admonition  >}}

{{< admonition note "How does recovery phrase work?" >}}
_Recovery phrase works like a unique master key. Your ledger device uses it to calculate private keys for every crypto asset you own._

_To restore access to your crypto, any wallet can calculate the same private keys from your recovery phrase._
{{< /admonition  >}}

{{< admonition note "What Happens if I lose Access To My Nano?" >}}
_To restore access to your crypto, any wallet can calculate the same private keys from your recovery phrase._

_1. Get a new hardware wallet._

_2. Select "Restore recovery phrase on a new device" in the Ledger app._

_3. Enter your recovery phrase on your new device to restore access to your crypto._
{{< /admonition  >}}


{{< admonition warning "가장 중요하게 기억할 것" >}}
_**When I connect my Nano to the ledger app, my private key is STILL OFFLINE!**_
{{< /admonition  >}}




### 3. Install Ledger Live

Ledger Live desktop앱을 다운로드합니다.

- [Download ledger live link](https://www.ledger.com/start)

<center>

![](/images/nano_download.png)

</center>



### 4. (opt) Update firmware.

- `Nano S plus`의 firmware를 업데이트 합니다.

{{< admonition tip "firmware?" >}}
_하드웨어 장치에 들어가는 소프트웨어의 일종. Ledger의 경우 버그 픽스 그리고 UI변경 때문에 하드웨어 펌웨어를 진행할 것으로 보인다._
{{< /admonition  >}}

<center>

![](/images/nano_update_firmware.png)

</center>


### 5. Install coin apps to nano device.

- 사용하고 싶은 종류의 코인들의 sw 어플리케이션을 다운로드합니다.

<center>

![](/images/nano_install_apps.png)

</center>


### 6. Add your accounts

This section will be continued with my 2023 investment scenario.

---

## 3. How to setup Nano S plus in Phone?

> my phone: Galaxy (Android)

1. Install Ledger Live app from google playstore [link](https://play.google.com/store/apps/details?id=com.ledger.live&hl=en_US&gl=US&pli=1).
2. Physically connect Nano Device and Phone with C type usb cable.
3. Type Pin code.
4. Done

<center>

![](/images/nano_phone.jpeg)

</center>
