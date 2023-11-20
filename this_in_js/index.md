# this, javascript


Q. What on earth, `this` is interpretated in js?
<!--more-->

## TL;DR
1. **`this`는 동적으로 해석된다.**
    - 일반적으로 호출하는 주체가 parameter로 전달된다.
    - 즉 `obj.method(...)`는 `method(this=obj, ...)`로 해석
2. **`Global context`**
    - `commonJS`: `this` === `globalThis` (`window` or `global`)
    - `type="module"`: **`this`는 언제나 `undefined`**
3. **`Function`**: method로 활용되지 않는 일반적인 함수에서 `this`는 2가지로 해석됩니다.
    - `use strict`: undefined
    - `non strict`: globalThis
4. **`Arrow functions`**: `outer scope`의 `this`를 `reference`하는 변수를 closure로 보존(`lexical scoping`)한다.
5. **`Callback`**: 일반적으로 this가 전달되지 않아 function과 동일하게 처리되지만, 일부 API (`JSON.parse(text, reviver)`)들은 내부적으로 this를 넣어준다.
6. **`Constructor` (`new`)**: `new`를 통해 호출되는 constructor는 내부적으로 `this`에 생성될 instance를 할당한다.
7. **`super`**: 부모의 context가 아닌, super.method()를 호출한 context의 this가 적용된다.
8. **`Class`**
    1. static vs instance
    2. `derived class constructor`에서 super()를 호출하지 않거나, return object하지 않는 이상 this는 생성되지 않는다.
9. **`EventHandler`**
    1. 대부분 브라우저는 `addEventListener`의 경우, handler에 현재 element를 this에 bind시켜서 줍니다.
    2. 인라인 이벤트 핸들러에서도 this에 현재 이벤트 리스닝되는 element가 바인드됩니다.
        1. `<button onclick="alert(this);">Show this</button>`


## 0. `this` intro
> [MDN: this](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Operators/this)

- Javascript에서 `this` 키워드는, 일반적으로 instance에 bind되는 대부분 언어와 달리, **동적으로 해석되며 호출한 방법에 의해 결정됩니다.** 

- 즉 동적으로 해석되는 this는 아래와 같이 호출되는 주체에 따라서 다르게 해석됩니다.

```js
function getThis() {
    return this;
}

const o1 = { name: "o1"};
const o2 = {
    __proto__: o1,
    name: "o2",
};

// 1. o1을 통해서 호출 되었을 때
o1.getThis = getThis;
console.log(o1.getThis()); // { name: 'o1', getThis: [Function: getThis] }

// 2. o2를 통해서 호출되었을 때
console.log(o2.getThis()); // { name : 'o2' }
```

- **전형적인 function 호출에서, this는 `function's prefix`(dot 앞에 있는 part)를 통해 implicitly하게 parameter로 전달됩니다.**

```js
o1.getThis();
// getThis(this=o1, ...나머지 args);
```

- 물론 explicitly하게 this를 지정하여 전달할 수도 있습니다.
    - Function.prototype.call()
    - Function.prototype.apply()
    - Reflect.apply()
- 또는 this를 bind시켜서 function을 새롭게 생성할 수도 있습니다.
    - Funciton.prototype.bind()


이제 차근차근 this가, context에 따라서 어떻게 다르게 해석되는지 알아보도록 하겠습니다.

## 1. Global context

`strict mode` 상관없이  `globalThis` property를 의미하며, 이는 실행환경에 따라서 2가지로 해석됩니다.

1. node: `global`
2. browser: `window`

```js
console.log(this === window) // true

a = 1
console.log(this.a) // 1
console.log(window.a) // 1

this.b = 2
console.log(window.b) // 2
console.log(b) // 2
```


## 2. Function declaration의 context

`use strict`를 사용하는지에 따라서 2가지 경우가 발생합니다.

```js
function f1() {
    return this;
}

function f2() {
    "use strict";
    return this
}

// 1. non-strict
// 브라우저
f1() === window; // true
// Node.js
f1() === global; // true

// 2. strict
f2() === undefined; // true
```

---

## 3. Callbacks
> `iterative array methods`, `Promise` constructor case

- **일반적으로 callback으로 함수를 넘겨준다면 this는 bind되지 않았기 때문에**
    1. "strict": **undefined**
    2. "non-strict": `globalThis`

```js
function print() {
    "use strict";
    console.log(this);
}

[1,2,3].forEach(print); // undefined, undefined, undefined

// 몇몇의 API들을 thisArg를 통해 this를 전달하도록 해줍니다.
const thisObj = {
    name: "john"
}

[1,2,3].forEach(print, thisObj) // {name: 'john'}, {name: 'john'}, {name: 'john'}
```
- **가끔 어떤 API들은 this를 넣어주는 경우도 있습니다.**
    - JSON.parse(text, reviver?)
    - JSON.stringify(value, replacer?)

> The reviver is called with the object containing the property being processed as this.

- JSON.parse는 `reviver` 함수에 parse처리된 상태를 this로 넣어서 실행시켜준다.

```js
function print(a,b) { 
    "use strict"; 
    console.log(this); 
}

print(1,2,) // undefined

JSON.parse('{"result":true, "count":42}', print)
// { result: true, count: 42 }
// { count: 42 }
// { '': {} }
```

## 4. Arrow functions

- `Arrow function`은 lexical context의 this를 유지합니다. 다른 말로 일반 function과 달리 호출 방식에 따라서 dynamic하게 this가 변경되지 않습니다. 

> `lexical context`, a.k.a `static context`, 동적으로 this가 처리되는 것이 아닌, 코드 작성 위치에서 this가 지정됨.

- **Arrow function는 closure처럼 구현되어 `this` value를 감싸고 있는 scope에서 들고 있다고 생각하면 됩니다. (`auto-bound`).**

- 또한 arrow function은 call(), bind(), apply()를 통해서 this를 동적으로 묶어주더라도 무시합니다.

```js
const a = {
    hi: () => {
        return this
    }
}

a.hi() // window
```

object literal은 아래와 같이 this값이 그 자체로 없기 때문에, outer scope의 this인 globalThis를 받습니다.

```js
const t = {
    who: this
}

t.who // window
```

- 경험상에 따르면 **마치 arrow function이 static하게 this를 지정하고 있는 것으로 이해되고 있지만, 사실 arrow function은 outer scope의 this를 가리키는 reference를 closure로 들고 있는 것입니다.**

outer의 this를 가리키는 reference를 closure로 들고 있다는 것이, 무슨 뜻인지는 아래 예시를 보면 더 확실해집니다.

```js
// function declaration은 호출되는 .에 따라서 this가 parameter로 전달됨.
const a1 = {
    hi: function() {
        return this
    }
}

a1.hi() // {hi: ƒ}
```

- 만약 arrow function이 function declaration안에 존재한다면? arrow function의 this는 function declaration의 this를 가리키게 됩니다.

```js

const a2 = {
    hi: function() {
        const f = () => {
            return this
        }
        return f()
    }
}

a2.hi() // 1. {hi: ƒ}

// 2. 🧐 what the fuck? why chnaged?

const hi = a2.hi
hi() // undefined (use strict) 
```

**2번 케이스를 보면 arrow function의 this값이 변경되는 것처럼 보입니다.**  `hi` property는 function declaration을 가지고 있으며, function declaration은 `dot`앞의 주체를 this로 하여 param에 전달되는 것처럼 동작합니다. 


하지만 const hi는 dot앞의 주체가 없기 때문에 strict mode에서는 this가 undefined로 할당되게 됩니다. arrow function은 외부 scope의 function의 this를 reference하고 있기 때문에, 해당 function의 this가 undefined이기 때문에, 마치 변경된 것 처럼 동적으로 변화되어 undefined을 return합니다. **그러므로 arrow function의 this 또한 동적으로 호출 방법에 따라서 변경되는 것처럼 동작가능합니다.**

이런 현상 때문에 위에서 arrow function의 this는 lexical scope의 outer scope의 this를 closure의 this로 reference하고 있다고 표현한 것입니다. 즉 arrow function의 this인 reference는 여전히 변경되지 않았기 때문입니다.



## 5. Constructors (`new operator`)

- **function이 `new`를 통해 constructor로 사용되면, js는 내부적으로 constructor 함수 안의 this를 생성되는 instance로 할당합니다**

```js
function Person(name) {
    this.name = name;
}

const p = new Person("John");
p.name // John
```

즉 원래 function declaration이 global context에서 사용되면, this가 window 또는 undefined로 해석되지만, `new` operator를 사용하게 될 경우 내부적으로 this를 instance로 할당해서 처리하게 됩니다.

## 6. super

- 자녀에서 super의 method를 호출했을 때, super의 method안에 this는, super의 값과 상관없이 **`super.method()`를 감싸고 있는 context의 this로 처리됩니다.**

```js
class Parent {
    constructor() {
        this.name = "parent"
    }

    getName() {
        // 여기서 this는 child
        return this.name
    }
}

class Child extends Parent {
    constructor() {
        super();
        this.name = "child";
    }
}

(new Child()).getName() // child
```

**즉 위와 같은 경우 Parent의 메서드안에 this는 Parent가 아니라, child의 this를 따릅니다.** 왜냐하면 super.getName을 감싸고 있는 child.getName의 this는 child를 가리키고 있기 때문입니다.


## 7. Class

1. static context: `this` = `Class`
    1. `static method`
    2. `static field (initializer / block)`
2. instance context: `this` = `instance`
    1. `constructor`
    2. `method`
    3. `instance field`

```js
class C {
    static staticField = this;
    instanceField = this;
}

const c = new C();
console.log(C.staticField === C); // true
console.log(c.instanceField === c); // true
```

### Derived class(child) constructor(`extends`)


- **`파생 클래스 생성자`**: **파생 클래스(자식 클래스) 생성자는 기본 클래스(부모 클래스) 생성자와 달리 초기에 this 바인딩이 없습니다.** 

```js
class Base {}
class Child extends Base {
    name ="child"
    
    constructor() {
        console.log(this.name)
    }
}
const c = new Child()
// Uncaught ReferenceError: Must call super constructor in derived class before accessing 'this' or returning from derived constructor
```

- super()를 호출하면 생성자 내에 this 바인딩이 생성되고, 이것은 사실상 `this = new Base();`라는 코드를 실행하는 것과 같은 효과를 가집니다. 여기서 Base는 기본 클래스를 의미합니다.

- **주의 사항**: `super()`를 호출하기 전에 `this`를 참조하려고 하면 오류가 발생합니다(당연히 this가 없으니), 그러므로 생성자안에서 this를 사용한다면, 그 보다 더 위에 super()가 존재해야 합니다.

```js
class Base {
    name = "Base"
}

class Child extends Base {
    constructor() {
        super();
        console.log(this.name);
    }
}

const c = new Child(); // Base
```


- **`super()` 호출 규칙**: 파생 클래스의 constructor는 `super()`를 호출하지 않고 반환해서는 안 됩니다. **단, 생성자가 객체를 반환하여 this 값을 덮어쓰는 경우나 클래스에 생성자가 전혀 없는 경우는 예외입니다.**

```js
class Good extends Base {
  constructor() {
    return { a: 5 };
  }
}

class Bad extends Base {
  constructor() {}
}
```

- `JS`의 child class에서 constructor를 명시적으로 작성하지 않으면, 내부적으로 constructor를 생성하고, 이 생성자에서는 super()를 자동으로 호출합니다.

```js
class AlsoGood extends Base {}
```


## 8. DOM Event Handler

### 8.1. 함수 이벤트 핸들러
- 대부분의 브라우저에서, 이벤트 핸들러로 사용되는 함수의 `this`는 **리스너가 부착된 DOM 요소에 바인딩 시킵니다.**

```js
function bluify(e) {
    "use strict"
    // 원래라면 undefined이지만, addEventListener는 target을 this로 bind시킨다.
    this.style.backgroundColor = "#A5D9F3";

    console.log(this === e.currentTarget) // true
    console.log(this === e.target) // currentTarget과 target이 같은 객체일 때 true
}

const elements = document.getElementsByTagName("*");
for (const ele of elements) {
    ele.addEventListener("click", bluify, false);
}
```

### 8.2. 인라인 이벤트 핸들러

- 인라인에서 사용되는 this는 이벤트 리스너가 부착된 element 입니다.


```html
<!-- [object HTMLButtonElement] -->
<button onclick="alert(this);">Show this</button>
```

- **하지만, 내부 scope를 추가로 가지게 된다면, global context로 해석됩니다.**


```html
<!-- undefined -->
<button onclick="alert((function () { 'use strict'; return this; })());">
```

즉 이는 다른말로, function을 정의해서 인라인에 집어넣더라도 동일하게 global context로 해석된다는 뜻입니다.

```html

<button onclick="print()">
<script>
    "use strict";
    function print() {
        // undefined
        alert(this); 
    }
</script>
```


지금까지 내용들을 정리하면 아래와 같은 테스트 코드를 작성해볼 수 있습니다.

```html
<!DOCTYPE html>
<html>
  <head>
    <title>Event Handler Reference Test</title>
  </head>
  <body>
    <!-- 1. use outside function -->
    <!-- 1.1. just function: undefined -->
    <!-- 1.2. addEventListener: [object HTMLButtonElement] -->
    <button id="btn1" onclick="print()">
      Show inner this (print Function)
    </button>
    
    <hr />


    <!-- 2. [object HTMLButtonElement] -->
    <button onclick="alert(this);">Show this</button>
    
    <hr />
    
    <!-- 3. undefined -->
    <button onclick="alert((function () { 'use strict'; return this; })());">
      Show inner this (Anonymous Function, Nothing Happen!)
    </button>

    <script>
      "use strict";
      // 1.1. just function
      function print() {
        alert(this); // undefined
      }
      // 1.2. addEventListener
      document
        .getElementById("btn1")
        .addEventListener("click", printHandler, false);
      function printHandler(e) {
        console.log(this); // [object HTMLButtonElement]
        console.log(this === e.currentTarget); // true
      }
    </script>
  </body>
</html>
```



<center>- 끝 -</center>

