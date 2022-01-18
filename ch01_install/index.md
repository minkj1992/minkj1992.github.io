# [ch01] Ros2 Foxy install



`ros2 Foxy`를 `mac pro` 로컬 환경에 설치해봅니다.
<!--more-->
<br />

## 1. tl;dr
- 실행환경: virtualbox(ubuntu 20.04)
- mount: 현재 디렉토리
- 코드작성
  - vm vs code에서 실행
  - host에서 코드 작성을 하려하였으나, lint와 링크 기능이 제대로 동작하지 않음

## 2. setup on `virtual box` (✅)
> [setup mac like keyboard on virtual box](https://bradwhittington.wordpress.com/2011/04/08/copy-paste-with-cmd-c-cmd-v-virtualbox-ubuntu-os/)

최종적으로 성공한 방식은 `vm`에 `ubuntu`를 설치해서 `foxy`를 실행하는 방식입니다. 아래 절차를 거쳐 세팅을 진행하였습니다.

### 2.1. ubuntu iso download

가장 먼저 ubuntu iso 이미지를 다운 받습니다.
https://mirror.kakao.com/ubuntu-releases/focal/ 에서 ubuntu 20.04 lts download 

### 2.2. vritualbox download

다음으로 [virtualbox](https://www.virtualbox.org/)를 최신버전으로 다운 받습니다.

### 2.3. virtual box 셋업


가장 먼저 설치한 iso를 연결해줍니다.

#### 2.3.1. iso 등록
![](/images/ros/0.png)

다음으로 아래 스펙으로 virtualbox를 세팅해줍니다. 
1. mem: 8192MB
2. HDD: 30 GB
3. 키보드 / 드래그앤 드롭 Bidirectional 설정

#### 2.3.2. 일반 > 기본
![](/images/ros/1_1.png)

#### 2.3.3. 일반 > 고급
![](/images/ros/1_2.png)

#### 2.3.4. 시스템 > 마더보드
![](/images/ros/2_1.png)

#### 2.3.5. 시스템 > 프로세서
![](/images/ros/2_2.png)


#### 2.3.6. 디스플레이
![](/images/ros/3_1.png)

#### 2.3.7. 공유폴더 설정
![](/images/ros/4.png)

#### 2.3.8. 전체 설정
![](/images/ros/5.png)

여기까지 기본적인 ubuntu 설정이 완료되었습니다. 다음으로 ubuntu 자체 설정을 해주겠습니다.

### 2.4. ubuntu install

ubuntu는 한글설정 / 맥북 키보드 shortcut 설정 / zsh 등, 처음 세팅하게 되면 설정해주어야 할 것들이 존재합니다.

가장 먼저는 apt 저장소를 [카카오 미러](https://memostack.tistory.com/217)로 변경해주세요. (이렇게 하면 더 빠르게 apt 설치가 가능합니다.)


1. [virtual box cmd right로 변경](https://superuser.com/a/829588)
2. [키보드 세팅(mac like)]

```bash
    $ sudo apt-get install keyboard-configuration
    $ sudo dpkg-reconfigure keyboard-configuration

    # select macbook pro(intel)
    # MacIntosh
    # English
    # English (Macintosh)
    # Both Alt keys
    # No compose key
    # Terminal preference에서 copy & paste 설정
```
3. mount host folder(소스코드) to ubuntu

```bash
# 이건 임시이며, 항상 mount시키고 싶다면, vmware에 공유 폴더 설정해주어야 한다.
$ sudo mount -t vboxsf ros2-sandbox /home/leoo/shared/ros2-sandbox/
```

### 2.5. install ros dependencies
```bash
$ sudo apt update
$ sudo apt install build-essential gcc make perl dkms
$ sudo apt update && sudo apt upgrade
$ sudo apt install terminator
$ sudo add-apt-repository universe
$ sudo apt-get update
$ sudo apt-get install python3-pip
```

### 2.6. install ros2 binary
```bash
$ locale  # check for UTF-8

$ sudo apt update && sudo apt install locales
$ sudo locale-gen en_US en_US.UTF-8
$ sudo update-locale LC_ALL=en_US.UTF-8 LANG=en_US.UTF-8
$ export LANG=en_US.UTF-8
$ locale  # verify settings

$ sudo apt update && sudo apt install curl gnupg2 lsb-release
$ sudo curl -sSL https://raw.githubusercontent.com/ros/rosdistro/master/ros.key  -o /usr/share/keyrings/ros-archive-keyring.gpg
$ echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/ros-archive-keyring.gpg] http://packages.ros.org/ros2/ubuntu $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/ros2.list > /dev/null

$ sudo apt update
$ sudo apt install ros-foxy-desktop
$ echo 'source /opt/ros/foxy/setup.bash' >> ~/.bashrc
$ sudo apt install python3-argcomplete
```

만약 이렇게 세팅하였을 때 자신의 vm이 너무 느리다고 판단 된다면 [vm 설정방법](https://mkyong.com/mac/virtualbox-running-slow-and-lag-on-macos-macbook-pro/)을 참고해주세요.


지금까지 세팅을 완료하였으면 예쁘게 동작하는 vm을 보실 수 있습니다. :)


#### 2.6.1. vm에서 동작하는 화면
![](/images/ros/virtualbox.png)


## 3. setup with `Docker` (🚫)
> [docker 설치 refs](https://roomedia.tistory.com/entry/1%EC%9D%BC%EC%B0%A8-macOS-Catalina-10155%EC%97%90-ros2-foxy-%EC%84%A4%EC%B9%98%ED%95%98%EA%B8%B0)

Host로 설정하는 방식이 실패하여, 그 다음으로 선택한 방식입니다. 도커를 사용한 방식은 문제없이 동작하였지만, 추후 터미널을 여러개 띄워 네트워크 통신이 많아지는 걸 고려하면, 추가적으로 설정해야 할 부분들이 많아 보였습니다. 추가로 문서 또한 리눅스 문서들이 많으므로 vm ware 사용

```bash
$ docker pull osrf/ros:noetic-desktop-full-buster
$ brew install socat

# https://www.cyberciti.biz/faq/apple-osx-mountain-lion-mavericks-install-xquartz-server/
$ brew install --cask xquartz
$ sudo reboot
# xquartz 보안 설정 모두 열어주기
```
- host 터미널에서 아래 명령어 실행

```bash
# ip 확인 후 xhost에 추가
$ ip=$(ifconfig en0 | grep inet | awk '$1=="inet" {print $2}')
$ xhost + $ip

# 컨테이너 생성
$ docker run -it -e DISPLAY=$ip:0 --name ros osrf/ros:noetic-desktop-full-buster
```


## 4. setup on `Host` (🚫)
> **This was failed**

이 방식은 맥북에 바로 설치하는 방식으로 `mac os`에 ros에 필요한 dependencies들을 바로 추가해주었습니다.

### 4.1. Pre-Install ROS2 on mac
```bash
$ brew doctor

$ softwareupdate --all --install --force
$ sudo rm -rf /Library/Developer/CommandLineTools
$ sudo xcode-select --install
$ brew link kubernetes-cli
$ brew link python@3.9
$ echo 'export PATH="/usr/local/sbin:$PATH"' >> ~/.zshrc
```

### 4.2. Install ROS2 on mac

```bash
brew install python@3.8
brew unlink python && brew link --force python@3.8
echo 'export PATH="/usr/local/opt/python@3.8/bin:$PATH"' >> ~/.zshrc
export LDFLAGS="-L/usr/local/opt/python@3.8/lib"
export PKG_CONFIG_PATH="/usr/local/opt/python@3.8/lib/pkgconfig"

brew install asio tinyxml2 tinyxml eigen pcre poco
brew install openssl && echo "export OPENSSL_ROOT_DIR=$(brew --prefix openssl)" >> ~/.zshrc
brew install qt freetype assimp sip pyqt5
brew install console_bridge log4cxx spdlog cunit graphviz

python3 -m pip install pygraphviz pydot catkin_pkg empy ifcfg lark-parser lxml netifaces numpy pyparsing pyyaml setuptools argcomplete

pip3 install -U colcon-common-extensions

# OpenCV는 필수는 아닙니다. 설치시 시간이 엄청 오래 걸리니 고민해보세요.
brew install opencv
```

### 4.3. Download ROS Foxy Binary
```bash
# https://github.com/ros2/ros2/releases

mkdir -p ~/ros2_foxy
cd ~/ros2_foxy
tar xf ~/Downloads/ros2-foxy-20211013-macos-amd64.tar.bz2
```

### 4.4. Check installed
```bash
$ . ~/ros2_foxy/ros2-osx/local_setup.zsh
```

### 4.5. Error
```
 $ . ~/ros2_foxy/ros2-osx/local_setup.zsh
[connext_cmake_module] Warning: The location at which Connext was found when the workspace was built [[/Applications/rti_connext_dds-5.3.1]] does not point to a valid directory, and the NDDSHOME environment variable has not been set. Support for Connext will not be available.
```

- `csrutil disable`로 시도해보았지만 실패
- ros2 바이너리 버전을 낮춰서 시도해보았지만 역시 실패
- big sur과 맞지 않는 source code에러가 있는 듯하다.


아래는 host에 설치하기 위해서 참조했던 문서들입니다.
- [공식](https://docs.ros.org/en/foxy/Installation/macOS-Install-Binary.html)
- [how-to-install-ros2-foxy-on-macos](https://snowdeer.github.io/ros2/2020/09/15/how-to-install-ros2-foxy-on-macos/)
- [building-ros2-on-macos-big-sur-m1](http://mamykin.com/posts/building-ros2-on-macos-big-sur-m1/)



