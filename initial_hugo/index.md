# Let's create blog with Hugo


> This article describes a series of technical steps to building this `serious` blog using the [hugo framework](https://gohugo.io/).
<!--more-->



사실 `Hugo`는 20년도에 혼자 제주도 여행을 가면서, 생각을 정리하고 싶어서 눈여겨 봤던 프레임워크이다.
당시 `go`를 공부하고 있었기 때문에 go로 만들어진 프레임워크라는 점에서 호감 +99점을 받았고, 이름이 다른 프레임워크들에 비해서 짧다는 장점이 있다. 

> 🤔 `Gatsby`는 무슨 왁스 이름같고, `Jekyll`는 스펠링도 어렵고 사실 어떻게 발음해야 될지도 잘 모를정도로 이름이 못생김

이 블로그는 `hugo`와 [LoveIt](https://github.com/dillonzq/LoveIt)을 사용해 만들었다. 

{{< admonition note keywords>}}
- Hugo
- LoveIt
- git submodule
- github workflows
- shell script
{{< /admonition >}}

## 1. Pre-Init Blog
> MacBook Pro (16-inch, 2019) Big Sur

### 1.1. Install hugo
- [install hugo docs](https://gohugo.io/getting-started/installing/)
```shell
# Install brew 
$ ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"
# download hugo
$ brew install hugo
# check hugo is successfully downloaded
$ hugo version
```

### 1.2. Select hugo templaate
1. github에서 hugo를 검색하고, star많은 순으로 정렬
   - 겉보기 이쁘다고 무턱대고 템플릿 가져다 쓰면, 아마 빠른시일 안에 블로그를 포기할 것임..
2. 1k 넘는 레포 중에서 이쁜거 찾는다. 
   1. 검색
   2. 카테고리 기능
   3. dark/light mode
   4. 댓글 기능

추천하는 테마들은 다음과 같다.
- ✅ https://github.com/dillonzq/LoveIt
- 👍 https://github.com/zzossig/hugo-theme-zzo
- https://github.com/luizdepra/hugo-coder
- https://github.com/adityatelange/hugo-PaperMod

### 1.3. Generate your profile image

- [프로필 이미지 생성한 곳](https://socialbook.io/cartoonize)
- [이말년 스타일 프로필 이미지](https://github.com/bryandlee/FreezeG)

필자는 프로필 사진에 `cartoonize`를 써보고 싶어서 github에서 ML 모델 위주로 검색했는데
원하는 곳을 찾지 못해서 대안으로 [여기](https://socialbook.io/cartoonize)를 사용 하게 되었다.
~~이말년 그림체로 프로필 만들고 싶은데 뭔가 해줘야 할게 많아서 포기... 누가 online 서버 만들어주면 좋겠다...🥺~~

![me](/images/profile.jpeg)


<details>
<summary>기타 작품들</summary>
<p>열정적인 스터디원들</p>
<span>
<img height="280" width="180" src="/images/profile2.jpeg" />
<img height="280" width="180" src="/images/y_good.png" />
<img height="280" width="180" src="/images/d_good_censored.png" />
</span>
</details>  


## 2. Init Blog
> 자 본격적으로 블로그 만들어보자.

### 2.1. Generate Blog

먼저 블로그용 폴더를 만들고 템플릿을 `submodule`로 추가해보자.

```shell
# hugo new site <YOUR FOLDER NAME>
$ hugo new site love
$ cd love
$ git init
$ git branch -M main
# git remote add origin <YOUR ROOT REPOSITORY>
$ git remote add origin https://github.com/minkj1992/love.git
# git submodule add <THEME REPOSITORY> themes/<THEME_NAME>
$ git submodule add https://github.com/dillonzq/LoveIt.git themes/LoveIt

# + 사내 계정이라 config 변경
$ git config user.email minkj1992@gmail.com
$ git config user.name "minkj1992"
```

필자의 경우 `github page`를 사용해서 블로그를 운영할 것이기 때문에, 미리 생성했던 `YOUR_ID.github.io`레포지토리를 `submodule`로 등록한다.


```shell
# git submodule add <YOUR_ID.github.io> public
$ git submodule add https://github.com/minkj1992/minkj1992.github.io public
```
{{< admonition warning >}}
$ git submodule add <YOUR_ID.github.io> public 명령어를 칠 때, **꼭 public을 디렉토리로 넣어주어야 한다!** 

hugo는 스태틱 파일들을 public/ 디렉토리로 빌드해주고, 우리의 `*.github.io`는 블로그의 스태틱 파일들을 가지고 있어야 하니까 :)
{{< /admonition >}}


[LoveIt config.toml](https://github.com/dillonzq/LoveIt/blob/master/exampleSite/config.toml)의 example config.toml을 참조해서, 본인의 블로그 루트 디렉토리에 config.toml을 생성하고 필요한 설정들을 추가해주자.


```shell
$ cp ./themes/LoveIt/exampleSite/config.toml ./config.toml

# if you want to add image to site home
$ mkdir -p assets/images
# after this command, paste your profile & log image to images and change config.toml

# (OPTIONAL) If you want to change css font-famiully and size, customize scss file.
$ touch assets/_override.scss
```

자 이제 첫 글을 작성해보자. 주의할점은 draft:false가 되어있어야 hugo를 github page 배포했을 때, 깨지지 않고 배포 된다.


```shell
# *.md draft must be falsed!
$ hugo new posts/initial_post.md
```

글을 작성했다면, 로컬에서 실행시켜보자 (hot-reload적용 됨)
참고로 commit이 안되서, 실행이 안된다고 하니 이쯤에서 우선 커밋 먼저 해준다. (`fatal: your current branch 'main' does not have any commits yet`)

```shell
$ git add . && git commit -m"Initial commit"

# hot reload debug run server
$ hugo server -D
```

정상적으로 블로그가 동작하는걸 확인했다면 배포를 해보자. 
배포는 아래 과정을 거쳐 진행된다.

1. hugo build (`$ hugo`)
2. ./public commit & push (submodule e.g minkj1992.github.io)
3. root repository push

먼저 hugo를 빌드하면 public/ 디렉토리에 파일들이 추가 된다. 추가된 파이들을 public의 remote로 push해주고, root 레포지토리로 돌아가서 push 해주면 된다.

필자는 아래의 스크립트를 사용해서 해당 과정을 진행해주고 있다.

#### 2.1.1. **`git-push.sh`**
```shell
#!/bin/sh

# If a command fails then the deploy stops
set -e
printf "\033[0;32m I Love Leoo.j \033[0m\n"
printf "\033[0;32mDeploying updates to GitHub...\033[0m\n"

printf "\033[0;32mBuild the project.\033[0m\n"
hugo -D
# hugo -t timeline # if using a theme, replace with `hugo -t <YOURTHEME>`


printf "\033[0;32m  Go To Public folder \033[0m\n"
cd public


printf "\033[0;32m  Setting for submodule commit \033[0m\n"
git config --local user.name "minkj1992"
git config --local user.email "minkj1992@gmail.com"
git submodule update --init --recursive


printf "\033[0;32m  Add changes to git. \033[0m\n"
git add .

printf "\033[0;32m  Commit changes.. \033[0m\n"
msg="rebuilding site $(date)"
if [ -n "$*" ]; then
	msg="$*"
fi
git commit -m "$msg"

printf "\033[0;32m  Push blog(presentation) source and build repos. \033[0m\n"
git push origin main


printf "\033[0;32m  Come Back up to the Project Root \033[0m\n"
cd ..
echo $pwd

printf "\033[0;32m  root repository Commit & Push. \033[0m\n"
git add .

msg="rebuilding site `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi

git commit -m "$msg"

git push origin main
```
shell을 만들었다면, 이제 배포 해보자.

```shell
$ sh git-push.sh <COMMIT_MSG>
```

{{< admonition tip >}}
만약 github action을 사용하고 싶다면, https://github.com/minkj1992/love/blob/main/.github/samples/gh-pages.sample 를 사용 해보라. (단 secrets.PERSONAL_TOKEN은 github setting에서 ENV등록해 주어야 함)

git hook을 쓰면 커밋이 편하긴 하지만, 개인적으로는 hook을 쓰면 로컬의 public/ 디렉토리의 git 버저닝이 관리되지 않고 있는게 눈에 거슬려서 shell을 사용 중이다.
{{< /admonition >}}


## 3. Conclusion
`Nexters`에서 2021년 회고글 작성하는 모임에 참석해서, 많은 개발자 분들이 notion으로 글을 정리하는 것에 자극 받아서, 바로 블로그를 만들게 되었는데 개인적으로 만족스럽다.

생각보다 에러 잡는데 시간을 많이 지체한 것 같고, `github hook`부분 기능을 잘 몰라 커스터마이징 하고 싶어 이것 저것 만져보다가 토요일 하루가 꼬박 걸렸는데 이 글을 읽는 여러분은 제가 했던 삽질을 경험하지 않길 바란다.


<center> - 끝 - </center>





