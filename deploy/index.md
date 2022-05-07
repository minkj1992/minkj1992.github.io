# Deploy Golang Web server on google cloud run


`Google Cloud Run`을 사용해서, `golang`으로 만들어진 `gin server`를 배포하는 모든 과정을 작성합니다.

<!--more-->
<br />

## tl;dr

## Why Google Cloud?

[Cloud Run vs App Engine](https://dev.to/pcraig3/cloud-run-vs-app-engine-a-head-to-head-comparison-using-facts-and-science-1225)

위 글내용을 정리하면 cloud run은 req당 가격을 측정하며, 컨테이너가 req가 종료되면 동작하지 않기 때문에 가격적인 장점이 있지만, app-engine과 비교하여 성능적인 차이가 있긴합니다. (ping test 기준 56ms)

제가 빌드할 서비스는 트래픽이 많이 필요한 서비스는 아니기 때문에 가격적으로 유리한 cloud run을 사용하기로 했습니다. (심지어 $0.09/month로, heroku $7 hobby plan보다 유리함)

- [google cloud run 가격](https://cloud.google.com/run/pricing#tables)

아래는 cloud run에 대한 장점입니다.

1. 인스턴스 자동확장 (max 1k request per container)
2. https 제공
3. 가격 저렴 / req당 가격을 측정
4. 200만회 request / month 항상 무료

## Spec

최종적으로 적용할 서비스 리스트는 아래와 같습니다.

- `Cloud run rest APIs`
- `Cloud SQL`: sql db
  - mysql의 json field 제공 (Google Cloud SQL is more than MySQL v5.7.11.)
- `Cloud Code` & `Cloud Build`: CI/CD
- `Cloud Storage FUSE`: image / audio server
- devops
  - `Cloud Monitoring`
  - `Cloud Logging`
  - `Cloud Trace`

## Go 앱을 Cloud Run에 배포

- [Quick Start go with Cloud Run](https://cloud.google.com/run/docs/quickstarts/build-and-deploy#clean-up)
- [컨테이너 이미지의 Go 앱을 Cloud Run에 배포](https://cloud.google.com/run/docs/quickstarts/build-and-deploy/deploy-go-service#clean-up)

실제 프로젝트 진행은 qwiklabs.com을 통해서 가상 컴퓨팅환경에서 진행하였습니다.

- [Qwiklabs의 데모 계정](https://www.qwiklabs.com/focuses/5162?parent=catalog)
  - $8를 결제하였다.

## Cloud SQL

> [Cloud SQL 빠른 시작](https://cloud.google.com/sql/docs/mysql/connect-instance-cloud-shell?hl=ko)

- cloud shell에서 mysql 연결

```sh
gcloud sql connect voda --user=root

(SQL) create database voda;
```

다음으로 cloud build를 적용해서 깃헙 소스코드에서 cloud run을 지속적으로 배포해보겠습니다.

## CI/CD

- [Cloud Build를 사용하여 Git에서 지속적 배포](https://cloud.google.com/run/docs/continuous-deployment-with-cloud-build)

1. exchange-diary 깃헙 레포에 cloud build 앱 다운
2. build configuration:
   1. 현재는 모든 브랜치에 적용(`^develop$`)
   2. build type: `/Dockerfile`

## Cloud Run

1. 인그레스: 모든 트래픽 허용 (추후 변경)
2. cloud run build를 통해서 앱 배포
3. Cloud SQL을 어떻게 업데이트 해줄 수 있을까?

