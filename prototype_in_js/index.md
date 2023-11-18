# Prototype in Javascript


What is difference between **`__proto__`**, **`prototype`**, **`[[Prototype]]`**
<!--more-->

## TL;DR

- `.__proto__`: getter of [[Prototype]]
- `[[Prototype]]`: Internal Property
- `.prototype`
    1. **function 정의**: js는 자동으로 함수 정의할때,  `함수.prototype`에 `{ constructor: 함수 }`인 object를 넣어준다. 
    2. **Constructor**: js는 자동으로 constructor로 instance를 생성(new operator)할 때, `const f = new F()` 라고 한다면, `f.__proto__ = F.prototype`를 시행한다.

> FYI, `__proto__`를 직접적으로 활용하는 방식은 (deprecated)되었습니다.


```js
function F() {}

// 1. Function 정의 시
F.prototype = {
  constructor: F,
}

// 이때 js는 위와 같이 prototype property를 추가하고, 이 property안에 constructor라는 property를 지닌, object를 추가하고, 
// constructor property는 함수 그자체를 가리킵니다.


// 2. Constructor (new 사용) 시
const f = new F();

f.__proto__ = F.prototype

// instance의 [[Prototype]]에 constructor.prototype에 대한 reference를 추가합니다.
```


## [`__proto__` vs prototype vs [[Prototype]]](https://stackoverflow.com/a/62077007)


### Function

Javscript에서 Function은 사실 Object입니다. 더 정확히 말하자면 Function이란 키워드는 Object에 `[[Call]]`(ECMA-262) internal property를 추가한 object입니다. 

<center>

![](/images/prototype_in_js3.png)

</center>

여기서 말하는 Function은 `function`이라는 함수를 생성할 때, 사용하는 예약어가 아닌 위 콘솔에서 확인할 수 있는 미리 생성되어있는(Built-in) object를 뜻합니다.

**즉, 모든 function들은 Function(Built-in Function Object)의 instance입니다.** 쉽게 설명해서 **`function = new Function()`** 라고 생각할 수 있습니다.

또한 JS에서 `function F() {}`처럼 function을 정의할 때, 내부적으로 `prototype`이라는 property를 추가하고, 이 값으로 `{ constructor: self }`인 object를 추가합니다.

> JS의 Built-in들은 native source code로 정의되어있어, 매우 빠릅니다. (c++)

```js
function F() {    
}

// js는 아래코드를, 자동으로 실행합니다.
F.prototype = {
    constructor: F
}
```


### Constructor

모든 function들은 `new`라는 operator를 통해서 `constructor`의 기능을 할 수 있습니다.

new키워드를 통해서 instance가 생성될 때, js에서는 추가적으로 constructor.property를 `instance.__proto__`안에 넣어줍니다.

```js
function F() {}


const f = new F()

// JS는 아래코드를 자동으로 실행합니다.
f.__proto__ = F.prototype
```

즉 아래 2가지로 해석가능합니다. 이때 2번째 경우는 `F === F.prototype.constructor`이기 때문에 가능합니다.

1. `instance.__proto__` === `F.prototype`
2. `instance.__proto__` === `F.prototype.constructor.prototype`

instance의 __proto__는 F.prototype을 가리키고, F.prototype.constructor는 F와 같기 때문에, instance.__proto__는 F.prototype과 F.prototype.constructor.prototype과 같습니다.

3. `instance.constructor === instance.__proto__.constructor`

js에서는 instance가 생성될 때, constructor라는 property를 추가시켜, F.prototype.constructor를 가리키고 있습니다.

```js
> function F() {}
undefined
> const f = new F()
undefined
> f.constructor
[Function: F]
> F.prototype.constructor === F
true
> F.prototype.constructor === f.__proto__.constructor
true
> F.prototype.constructor === f.constructor
true
> f.constructor = undefined
undefined
> f.constructor
undefined
> F.prototype.constructor
[Function: F]
```

즉 instance의 constructor property는 reference타입이라는 것을 알 수 있습니다.



### The prototype chain

![By stackoverflow Above link](/images/prototype_in_js2.jpeg)


js는 Built-in Object가 존재하며, User Defined Object들은 결국 빌트인 Object까지 chain을 타고 올라가며, 최종적으로 Object.[[Prototype]] === null에서 chainning이 마무리됩니다.

> The chain of objects connected by the `__proto__` property is called the prototype chain.

즉 js는 상속을 `__proto__` 필드를 사용해서 구현했으며, 실제로는 reference로 저장되기 때문에, singleton object들을 공유해서 상속하는 방식으로 되어있습니다.

아래는 위의 다이어그램을 js코드로 간단하게 표현해봤습니다.

```js
// 1. Built-in Objects in Javascript
function Object() {
    prototype: {
        constructor: function Object(),
        __proto__: null // Object.prototype.__proto__ === null
    }

    // Object Internal Property
    // ecma: https://262.ecma-international.org/5.1/#sec-8.6.2
    [[Prototype]]:{
        constructor: function Function(),
        __proto__: Object.prototype
    }

    get __proto__() {
        return this.[[Prototype]]
    }

    [[Call]]: ...

}

function Function() {
    prototype: Object.__proto__
    [[Prototype]]: Object.__proto__

    get __proto__() {
        return this.[[Prototype]]
    }

    [[Call]]: ...
}

// 2. User-defined Objects

// new를 사용하면 o.[[Prototype]] property에 포인터 of Constructor.prototype
// 즉 o.__proto__ === Object.prototype
const o = new Object() 

function Polygon() {
    // 함수 생성시, js엔진은 prototype property를 추가해준다. 이 property는 object로 constructor로 self를 가리키는 property를 지니고 있다.
    prototype: {
        constructor: function Polygon(), // self
        __proto__: Object.prototype // {}(literal object)로 생성할 때, __proto__에 Object.prototype의 reference가 담긴다.
    }
    __proto__: Function.prototype
}
```

## `__proto__` vs [[Prototype]]
> https://developer.mozilla.org/ko/docs/Web/JavaScript/Reference/Global_Objects/Object/proto

>> 주의: 객체의 [[Prototype]]을 변경하는 것은 최신 JavaScript 엔진이 속성 접근을 최적화하는 방식의 특성상 모든 브라우저 및 JavaScript 엔진에서 매우 느린 작업입니다. 상속 구조를 변경하는 것이 성능에 미치는 영향은 미묘하고 광범위하며, `obj.__proto__ = ...` 문에 소요되는 시간 뿐만 아니라 [[Prototype]]이 변경된 객체에 접근할 수 있는 모든 코드들에 대해서도 영향을 줄 수 있습니다. 성능에 관심이 있다면 객체의 [[Prototype]] 설정을 피해야 합니다. 대신 Object.create()를 사용하여 원하는 [[Prototype]]으로 새 객체를 만드세요.

>> 주의: `Object.prototype.__proto__`는 오늘날 대부분의 브라우저에서 지원되지만, 그 존재와 정확한 동작은 오직 웹 브라우저와의 호환성을 보장하기 위한 레거시 기능으로서 ECMAScript 2015 사양에서 비로소 표준화되었습니다. 더 나은 지원을 위해 대신 Object.getPrototypeOf()를 사용하세요.

그러니, 직접적인 `__proto__` 보다는 Object.getPrototypeOf()를 사용하는 것이 권장된다.


## Conclusion


- `__proto__`: getter of [[Prototype]].
- prototype: 함수에서는 constructor기능을 위해, 사용되며, object는 inheritance를 위해 사용되는 필드.
- [[Prototype]]: proto chain을 사용한 상속(코드 공유)를 위해 사용되는 internal property.


