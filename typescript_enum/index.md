# Typescript Better Enum



Typescript Better Enum
<!--more-->

## TL;DR
- TS버전 5.*.*미만일 경우에는 Enum을 사용하지 말자, 대신 as const 사용하자.
- v5 부터는 enum 사용

```ts
const PostState = {
    Draft: "DRAFT",
    Scheduled: "SCHEDULED",
    Published: "PUBLISHED"
} as const;

type PostStateType = typeof PostState[keyof typeof PostState];

// Usage
const x: PostStateType = PostState.Draft
const y: PostStateType = "SCHEDULED"
```

다만 TS version5부터는, [All enums Are Union enums](https://devblogs.microsoft.com/typescript/announcing-typescript-5-0/#all-enums-are-union-enums) enum의 기존 문제점들이 해결되었습니다.


## Enum의 문제점과 v5의 변화

기존의 타입스크립트에서 enum은 기본적으로 숫자 기반 enum이었습니다. 이는 각 enum 멤버가 숫자 값을 가지며, 이 숫자들은 컴파일 타임에 할당됩니다. 그러나 타입스크립트의 타입 체킹 시스템은 enum 타입으로 선언된 변수에 어떠한 숫자도 할당할 수 있게 허용했습니다. 이는 enum이 숫자의 집합으로 간주되었기 때문입니다.

예를 들어, 다음과 같은 enum이 있다고 가정해보겠습니다:

```typescript
enum PostState {
    Draft, // 0
    Published, // 1
    Private, // 2
}
```

여기서 `PostState` 타입의 함수에 숫자를 전달하는 것은 허용되었습니다:

```typescript
function handlePost(state: PostState) {}

handlePost(0); // Draft를 의미합니다.
handlePost(1); // Published를 의미합니다.
handlePost(9999); // 타입 체크를 통과합니다.
```

`handlePost(9999)`가 허용되는 이유는 타입스크립트가 enum 타입을 숫자의 서브타입으로 간주했기 때문입니다. 즉, `PostState` 타입은 실제로 숫자 타입에 할당 가능한 모든 값을 포함하고 있었습니다. (모든 숫자값을 허용하는 문제) 이는 타입 안전성을 저해하는 문제로 인식되었고, 코드의 의도를 명확하게 표현하는 데에도 문제가 있었습니다.

타입스크립트 2.0에서 도입된 enum 리터럴 타입은 이 문제를 부분적으로 해결했습니다. 각 enum 멤버는 리터럴 타입을 가지게 되어, `PostState` 타입은 실제로 `0 | 1 | 2`와 같은 유니언 타입이 되었습니다. 그러나 계산된 멤버가 있는 경우에는 여전히 문제가 발생할 수 있었습니다.

타입스크립트 5.0에서는 모든 enum 멤버가 고유한 타입을 가지게 되어, 이러한 문제를 완전히 해결했습니다. 이제 `handlePost(9999)`와 같은 호출은 타입 에러를 발생시키게 되어, 타입 안전성이 향상되었습니다.


타입스크립트 5.0에서의 변화:

타입스크립트 5.0에서는 모든 enum 멤버가 고유한 타입을 가지게 되어, enum은 멤버 타입들의 유니언이 됩니다. 이로 인해 각각의 enum 멤버를 더 명확하게 다룰 수 있게 되며, 타입 안전성이 향상됩니다.

```typescript
// 타입스크립트 version 5.*.*
enum PostState {
    Draft, // 이제 'PostState.Draft' 타입을 가짐
    Published, // 'PostState.Published' 타입
    Private, // 'PostState.Private' 타입
}

// 이 함수는 이제 'PostState'의 특정 멤버만을 받을 수 있습니다.
function handlePost(state: PostState.Draft | PostState.Published) {}

handlePost(PostState.Draft); // ✅ 
handlePost(PostState.Private); // ❌ 타입 체크 해줌
handlePost(999); // ❌ 값 체크 해줌
```

이제 `handlePost` 함수는 `PostState`의 모든 멤버를 받는 것이 아니라, 오직 `Draft` 또는 `Published` 상태의 포스트만을 처리할 수 있습니다. 이는 타입스크립트가 더 엄격한 타입 체크를 할 수 있게 해주어, 잘못된 값이 함수에 전달되는 것을 방지합니다.

또한, 계산된 값을 가진 enum 멤버에 대해서도 고유한 타입을 생성하여, 이전에는 타입스크립트가 처리하지 못했던 경우에도 타입 안전성을 제공합니다.

```typescript
enum PostState {
    Draft = "DRAFT",
    Published = "PUBLISHED",
    Private = Math.random() // 계산된 값
}

// 이전 버전에서는 'Private' 멤버의 타입이 계산될 수 없었지만,
// 타입스크립트 5.0에서는 'Private'도 고유한 타입을 가집니다.
```

이러한 변화는 코드의 명확성과 타입 안전성을 크게 향상시키며, 타입스크립트를 사용하는 개발자들에게 더 나은 개발 경험을 제공합니다.

