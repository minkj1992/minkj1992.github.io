# Openai Rate Limit


How to manage openai rate limit on production?
<!--more-->

## TL;DR

아래는 rate limit을 도식화 해본 그림입니다.

<center>

![](/images/openai_rate_limit_excali.png)

</center>

---

## Rate Limits

Openai는 rate limit을 통해 organization별로 api access를 관리합니다. 이때 openai의 rate limit은 총 5가지 방법으로 measure됩니다.

1. **RPM** (Requests per minute, `requests / 1min`)
2. **RPD** (`requests / 1day`)
3. **TPM** (`tokens / 1min`)
4. **TPD** (`tokens / 1day`)
5. **IPM** (`images / 1min`)


**추가로 5개중 1개만 할당량에 걸려도 rate limit가 발생합니다.**


<center>

![openai important](/images/openai_important.png)

</center>

1. **Rate limit은 user가 아닌 organization 기준으로 관리됩니다.**
2. 위의 5가지 quota는 model별로 다르게 관리됩니다.
    1. 예를들어 하나의 API key를 통해 서로 다른 model을 사용할 때, 둘 중 하나의 model이 rate limit에 걸리더라도, 다른 model은 영향을 받지 않습니다.
    2. 물론 이 또한 organization의 max usage limit을 넘어가지 말아야 합니다.

![openai usage limit](/images/openai_usage_limit.png) 


3. Openai의 rate limit은 organization level로 관리됩니다.
    - 이 때문에 하나의 organization에 속해있는 여러 api_key를 사용하더라도 특정 model이 rate limit에 걸리게 되면 다른 api key에서도 에러가 생깁니다.
    - 이를 방지하기 위해서, **organization을 여러개 생성해서 api_token을 관리하는 방법이 있습니다.**

- FYI, openai는 하나의 account당 하나의 organization만 만들 수 있습니다.

![openai organization](/images/openai_orga.png)

---

## Tiers

Rate limit 정책은 앞서 말했듯, organization과 model에 따라 달라집니다.

1. Organization (Tier)
2. Models (Rate Limit)

먼저 orgranization은 tier에 따라서 또는 usage limit (bucket per month)설정에 따라서 다른 limit을 가집니다.

![](/images/openai_usage_tier.png) 


마찬가지로 2번째의 model에 따라서도 tier별로 그리고 model별로 다른 rate limit 정책을 따릅니다. 아래는 tier3의 예시입니다.

![](/images/openai_model_rate_limit.png) 

**실제 production에서 rate limit을 관리한다면 현재 tier에 대한 RPM, RPD들을 티어별로 저장해야 할 것 같습니다.**

---

## Rate Limit Headers

개발에서 현재 Rate limit을 확인하는 방법은 HTTP 요청을 통해 받은 response의 header를 통해서만 확인 가능합니다. (물론 UI를 통해서도 가능)

~~세계적인 기업에서 token, model별로 현재 rate limit metric을 알려주는 api endpoint가 하나도 없다는게 충격적입니다.~~


![](/images/openai_limit_header.png)

1. `x-ratelimit-limit-requests`: 남은 최대 요청 수
2. `x-ratelimit-limit-tokens`: 남은 최대 토큰 수
3. `x-ratelimit-remaining-requests`: 남은 허용 요청 수
4. `x-ratelimit-remaining-tokens`: 남은 허용 요청 수
5. `x-ratelimit-reset-requests`: request rate limit이 초기화 되기 까지 남은 시간
6. `x-ratelimit-reset-tokens`: token rate limit이 초기화 되기 까지 남은 시간

사실 `remaing`과 `limit`의 차이점은 현재 진행하고 있는 request를 포함시키는지 여부입니다.

즉 만약 `x-ratelimit-limit-requests`가 60이고 `x-ratelimit-remaining-requests`가 59라면, 후자는 현재 진행되고 있는 request를 포함시켰기 때문에 1차이가 나는 것 같습니다.


## TODO;

production managing architecture.

- source codes
    - [How to handle rate limits.ipynb](https://github.com/openai/openai-cookbook/blob/297c53430cad2d05ba763ab9dca64309cb5091e9/examples/How_to_handle_rate_limits.ipynb)
    - [openai/api_requests_parallel_preprocessor.py](https://github.com/openai/openai-cookbook/blob/297c53430cad2d05ba763ab9dca64309cb5091e9/examples/api_request_parallel_processor.py)


## Refs

- [Rate limit advice openai](https://platform.openai.com/docs/guides/rate-limits?context=tier-free)
- [Q&A openai rate limit strategy](https://community.openai.com/t/rate-limiting-strategies/7128/3)
- [openai docs](https://platform.openai.com/account/limits)


