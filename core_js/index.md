# Core Javscript


[모던 Javascript 튜토리얼](https://ko.javascript.info/?map)
<!--more-->


## 1. JS 기본

### 변수

`typeof 연산자`는 값의 자료형을 반환해줍니다. 그런데 두 가지 예외 사항이 있습니다.

```js
typeof null == "object" // 언어 자체의 오류
typeof function(){} == "function" // 함수는 특별하게 취급됩니다.
```


### 함수

함수 default 파라미터로 expression(표현식)을 넘겨줄 수도 있습니다.

```js
function showMessage(from, text = anotherFunction()) {
  // anotherFunction()은 text값이 없을 때만 호출됨
  // anotherFunction()의 반환 값이 text의 값이 됨
}
```

또한 python과 달리, default가 나오고, default를 주지 않더라도 에러가 나지 않습니다. 

```js
> function F(a="yes", b) {return `${a}, ${b}`}
undefined
> F()
'yes, undefined'
```

이런 점에서는 a는 default를 쓰면서, b에 인자를 주고 싶다면 `F(a=undefined, b="yes")`처럼 해야하는 불편한 상황이 생길 수 있을 것 같아서 python처럼 강제해서 관리하는게 더 좋아보이네요.


> **함수 parameter(매개변수) 기본값 평가 시점**
>> 자바스크립트에선 함수를 호출 때마다, argument가 없을 경우에만 default parameter를 평가합니다. 만약 해당 argument가 전달된다면 default의 expression은 호출되지 않습니다.


> **return문이 없거나 return 지시자만 있는 함수는 undefined를 반환합니다.**

### 함수 선언문 vs 함수 표현식

```js
funciton sum(a, b) {
    return a + b;
}

const sum = function(a,b) {
    return a + b;
};
```

둘의 큰 차이점은

1. 함수 선언문: 자바 스크립트는 스크립트 실행 전, 준비단계에서 선언된 함수 선언문들을 모두 찾아, 해당 함수를 생성합니다. 그렇기 때문에 **함수 선언문은 함수 선언문이 정의되기 전에도 호출 가능합니다.**
2. 함수 표현식: **실제 코드 흐름이 해당 함수에 도달했을 때 함수를 생성합니다.**

또 한가지 중요한 점은, 함수 선언문은 선언된 블록 내 어디서든 접근할 수 있지만, 블록 밖에서는 함수에 접근하지 못합니다. 다시 말해서

```js
const booooooolean = true;

switch (booooooolean) {
  case "unreachable":
    function greet() {
      console.log("unreachable");
    }
    greet();
    break;
}

greet();
// TypeError: greet is not a function
```

greet이 undefined이기 때문에, is not a function 에러가 발생합니다.

## 3. 객체:기본

### 3.1. 객체

- **상수 객체는 수정될 수 있습니다.**

```js
const user = {
    name: "John",
};

user.name = "Pete"; // ok
```

단, property flag를 사용하면 immutable하게 처리할 수 있습니다. 물론 이 또한 user가 let인지 const인지와는 상관없습니다.

```js
use strict; // 엄격 모드에서만 가능합니다.

let user = {
  name: "John"
};

Object.defineProperty(user, "name", {
  writable: false
});

user.name = "Pete"; // Error: Cannot assign to read only property 'name'
```

- **계산된 프로퍼티**(computed property)

대괄호로 쌓여진 property는 computed property를 나타냅니다.

```js
const fruit = 'apple';
const bag = {
  [fruit + 'Computers']: 5 // bag.appleComputers = 5
};
```

- **문자형, 심볼형이 아닌 key값은 문자열로 자동형변환 됩니다.**

```js
// 선언시 7은 문자열로 "7"로 자동-형변환이 일어납니다.
const o = {
  7: "hi",
};

console.log(o[7]); // 문자 또는 심볼이 아니기 때문에 "7"로 접근합니다.
console.log(o["7"]);
```

- 자바스크립트 객체의 중요한 특징: **존재하지 않는 프로퍼티에 접근하려 해도 에러가 발생하지 않고 undefined를 반환합니다.**

```js
let user = {}
console.log(user.alkdjfklcvjixc === undefined) //true
```

이를 해결하기 위해서 **in 연산자를 사용해서 property가 들어있는지 확인합니다.**

```js
let user = { name: "John", age: 30 };

alert( "age" in user ); // 1. true
alert( "blabla" in user ); // 2. false

let key = "age"
alert( key in user) // 3. true, 문제 상황
```

위의 3번쨰 문제 상황 발생 가능하니, 그냥 ""를 사용해서 key를 체크하는게 좋습니다.


또한 property를 검사할 때는, 아래와 같은 이유로 **=== undefined로 체크하는 것 대신, in을 사용해야 합니다.**


- property를 undefined로 선언한 경우
```js
let o = {
    a: undefined
}; 


console.log(o.a); // undefined, a라는 property key가 있든 없든, undefined가 나오기 때문에 비교문이 의미없다.
console.log("a" in o) // true
```

#### for ... in 반복문
> object의 모든 property key를 순회가능합니다. 

이건 array를 loop도는 `for(;;)`와는 본질적으로 다른 기능입니다. (많이 헷갈렸었음.)

```js
let user = {
  name: "John",
  age: 30,
  isAdmin: true,
};

for (let k in user) {
  console.log(k); // name, age, isAdmin
  console.log(user[k]); // John, 30, true
}
```

참고로 `for (let k in user)`에서 let을 지워도 동작하는데, 이는 js가 변수 선언 키워드(var, const, let)이 없으면 전역변수로 사용하기 때문이다. 그러니 변수에 키워드 무조건 넣는게 좋다고 생각합니다. 또한 **use strict에서는 이를 허용하지 않습니다.**

#### property 정렬 순서

- 정수 프로퍼티(integer property)는 자동으로 정렬
- 그 외의 프로퍼티는 객체에 추가한 순서 그대로 정렬됩니다.

따라서 아래 코드에서, 선언된 나라 순서가 중요하다면 주의해야합니다.

```js
let codes = {
  "49": "독일",
  "41": "스위스",
  "44": "영국",
  // ..,
  "1": "미국"
};

for (let code in codes) {
  alert(code); // 1, 41, 44, 49
}

// 해결방법
let codes = {
  "+49": "독일",
  "+41": "스위스",
  "+44": "영국",
  // ..,
  "+1": "미국"
};

for (let code in codes) {
  alert( +code ); // 49, 41, 44, 1
}
```
### 3.2. by reference
> 참조에 의한 객체 복사

- 객체와 primitive type의 근본적인 차이 중 하나는 **객체는 ‘참조에 의해(by reference)’ 저장되고 복사된다**는 것입니다.

- reference copy대신 Shallow copy를 하고 싶다면 Object.assin()를 사용하면 됩니다.

```js
let user = {
  name: "John",
  sizes: {
    height: 182,
    width: 50
  }
};

let clone = Object.assign({}, user);
```

shallow copy이기 때문에 user.sizes와 clone.sizes는 동일한 객체를 가리킵니다.



@TODO CONTINUE

- https://ko.javascript.info/object-methods
