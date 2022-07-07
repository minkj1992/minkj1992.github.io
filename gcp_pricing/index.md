# Google Cloud Computing 가격 정책 비교


사이드 프로젝트를 하면서 GCP로 배포할 때, 서버 비용이 걱정되었다면!
<!--more-->
<br />

## tl;dr
사이드 프로젝트 성격인 서비스에, 대부분의 사용자가 한국인이라고 가정하면 아래와 같이 설정하면 됩니다.

- [x] Just use `Taiwan` region for every instance
- [x] Check cloud run minimum instance number(set 0)
- [x] Set cloud sql vCPU 1

## Intro
2022년 6/25일(현재 날짜 7/7일) GDG 해커톤을 통해 정말 간단한 golang 서버를 배포했었다.
갑자기 문득 궁금한 생각이 들어, cloud console에 가서 금액을 확인해보니,,

- ₩268,826,, 말이되나?
![](/images/gcp_price_total)

트래픽 없는 해커톤 서비스가, 약 12일 정도 사용했는데 26만원 찍혔다는건 말도 안된다. 
이에 대한 이유를 분석해보고 다음 `myply` 사이드 프로젝트를 할 떄, 적용해보기 위해서 정리를 해봅니다.


- 대부분 cloud sql이 잡아먹고 있었다. (원인: vCPU 4인 default instance가 살아있어서)
![](/images/gcp_price_cloud_sql)

- cloud run은 100만건까지는 공짜인줄 알았는데, 500원씩 사용되었다. (원인: min instance count 1로 되어있었다.)
![](/images/gcp_price_cloud_run)



## GCP network
> https://cloud.google.com/vpc/network-pricing

- Ingress ( users -> instance ) is free

![Screen Shot 2022-07-07 at 11 07 13 AM](https://user-images.githubusercontent.com/37536298/177675005-4ad16f45-dbc9-48a1-ab55-1cf7efeccfe1.png)


- egress ( kr -> taiwan $0.05 )

![Screen Shot 2022-07-07 at 11 08 19 AM](https://user-images.githubusercontent.com/37536298/177674992-11590792-af5d-454c-b0f5-4cde9032c33a.png)



## Cloud run
> https://cloud.google.com/run/pricing

1. Set cpu less than 1, memory 128MiB
2. Check CPU is only allocated during request processing
3. Execution env to set `First generation`
4. Set autoscaling  min number of instances field to 0. (max: 4~5?)

![Screen Shot 2022-07-07 at 10 17 45 AM](https://user-images.githubusercontent.com/37536298/177668849-966c3246-f473-49cd-a11f-46c078c21fcb.png)

- cloud run pricing by region, `FYI, seoul is 2nd grade.`
![Screen Shot 2022-07-07 at 10 42 21 AM](https://user-images.githubusercontent.com/37536298/177672190-eaa1e8e5-99f0-447f-8d05-30e6dc914486.png)


## Cloud SQL
> https://cloud.google.com/sql/pricing

- Set number of vCPU 1 (4로 설정하면 대략 하루 18,000원)
- Seoul이 vCPU가격이 다른 region에 비해서 훨씬 높다.  -> Select Taiwan

- Taiwan의 경우 us-central들과 가격 정책이 같다.
![Screen Shot 2022-07-07 at 10 28 07 AM](https://user-images.githubusercontent.com/37536298/177670809-bced7463-906c-4f8b-b70e-cee1c9d197bb.png)

- lowa
![Screen Shot 2022-07-07 at 10 29 18 AM](https://user-images.githubusercontent.com/37536298/177670842-0eeda4fb-d18a-4b07-b660-2523ed505276.png)

- seoul
![Screen Shot 2022-07-07 at 10 29 47 AM](https://user-images.githubusercontent.com/37536298/177670883-006897d1-1b58-4691-9d82-23a270b28e6f.png)

- **Data Egress(outboud) pricing**
  - Set `cloud run` region same as `cloud sql` 

![Screen Shot 2022-07-07 at 10 32 37 AM](https://user-images.githubusercontent.com/37536298/177671199-c2760c24-b0ab-4871-a9fc-69b6e34342ba.png)

## Conclusion

저와 비슷한 환경에서 사이드 프로젝트를 한다면 cloud run과 cloud sql을 지금까지 분석했던 방식으로 설정하면 될 것 같습니다. 단, 트래픽이 늘어나서 taiwan <-> korea에 대한 이그래스 가격이 합리적이지 못하다고 느낀다면, region을 seoul로 바꾸는 것도 좋아보입니다. 

간단하게 계산해보면 `18GB`(9/0.5)이상의 트래픽이 한달 간 발생한다면, region을 다시 seoul로 바꿔서 서비스해야 할 것 같습니다. 단 트래픽이 늘어나면 sql vCPU도 늘어나야 하기 때문에 다시 계산해야 될 수도 있습니다. (*여기서 $9는 cloud sql single vCPU 한달 사용 차액 $9(taiwan - seoul, per vCPU)*)

정리하면, youtube의 홈화면(최적화가 잘되어있다는 것도 감안 필요, 또한 스트리밍 시청하게 되면 영상 사이즈로 고려해야 함)은 약 613kb 네트워크 리소스를 사용하며, 이를 18GB만큼 사용하려면 `29363.7846656`으로 해당 서비스가 약 한달에 3만번 이상 hit를 받는 서비스라고 생각한다면, 네트워크 이그레스를 고려해야 한다.


