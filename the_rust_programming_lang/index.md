# The rust programming language


`Rust`에서 공식적으로 지원하는 ["The rust programming language"](https://doc.rust-lang.org)을 통해서 `Rust` 핵심 문법적 특징과 예시코드 그리고 이면에 숨은 디자인 원칙들을 정리 해보겠습니다.
<!--more-->


## 3. Common Programming Concepts

### 3-1. Variables and Mutability

> **Rust에서 기본 변수는 불변성입니다.**

이를 통해 컴파일 타임에 실수로 immutable 변수를 변경하는 버그를 잡아내도록 강제합니다. `mut` 키워드를 사용하면 Mutability를 제공할 수 있습니다. (가변변수)

- immutable vs mutable

만약 `매우 큰 구조체`를 다루는 경우 mutable 인스턴스를 사용하는 것이 새로 인스턴스를 할당하고 반환하는 것보다 빠를 수 있습니다. 데이터 크기가 작을수록 새 인스턴스를 생성하고 FP(함수적) 프로그래밍 스타일로 작성하는 것이 더 합리적이고, 그렇기에 약간의 성능 하락을 통해 가독성을 확보할 수 있다면 더 가치있는 선택입니다.

#### Constants
> const vs Variables

- 상수에 대해서는 `mut`을 사용하는 것이 허용되지 않습니다: 상수는 항상 불변합니다.
- 상수는 `let`키워드 대신 `const`키워드를 사용해야 하고, 값의 type을 선언해야 합니다.
- 상수는 can be declared in any scope(including the global scope)
- 상수는 may be set only to a `constant expression`(상수 표현식), **not the result of a value that could only be computed at runtime.**, 즉 컴파일 타임에 하드코드 되어야합니다.

#### Shadowing
> "선언한 변수와 같은 이름의 새로운 변수를 선언할 수 있고, 새 변수는 이전 변수를 shadows하는 것"

```rs
fn main() {
    let x = 5;

    let x = x + 1;

    {
        let x = x * 2;
        println!("The value of x in the inner scope is: {x}"); // x is 12
    }

    println!("The value of x is: {x}"); // x is 6
}
```

`shadowing`과 `mut`은 크게 2가지 차이가 있습니다.

1. 문법 차이

```rs
// shadowing
{
    let x = 5;
    let x = x + 1;
}


// mut
{
    let mut x = 5;
    x = x + 1;
}
```

2. `shadowing`은 같은 이름을 유지하면서, 다른 타입을 사용할 수 있습니다. 즉 네이밍을 깔끔하게 사용할 수 있습니다.

```rs
// shadowing: 깔끔한 네이밍 유지 가능하다.
{
  let spaces = "   ";
  let spaces = spaces.len();
}

// mut: 컴파일 에러
{
  let mut spaces = "   ";
  spaces = spaces.len() // error[E0308]: mismatched types
}

// immutable: 더럽다.
{
  let spaces_str = "    ";
  let spaces_len = spaces_str.len();
}
```

### 3-2. Data Types
- **Rust의 타입은 크게 2가지: `scalar`와 `compound` 두가지로 나뉩니다.**
- Rust는 `statically typed language`(타입이 고정된 언어)입니다. 

즉 Rust는 컴파일타임에 모든 변수의 타입이 정해집니다. 그러므로 명시적으로 타입을 지정 또는 컴파일러가 타입을 추측할 수 있도록 선택의 폭을 줄여주어야 합니다.


#### scalar types
> A scalar type represents a single value. Rust는 4가지의 primary 스칼라 타입을 가지고 있습니다.

- `Integer Types`

<span class="caption">Table 3-1: Integer Types in Rust</span>

| Length  | Signed  | Unsigned |
| ------- | ------- | -------- |
| 8-bit   | `i8`    | `u8`     |
| 16-bit  | `i16`   | `u16`    |
| 32-bit  | `i32`   | `u32`    |
| 64-bit  | `i64`   | `u64`    |
| 128-bit | `i128`  | `u128`   |
| arch    | `isize` | `usize`  |

`arch`는 32-bits, 64-bits 같은 컴퓨터 아키텍처를 뜻합니다.

<br />

- `Integer Literals`

<span class="caption">Table 3-2: Integer Literals in Rust</span>

| Number literals  | Example       |
| ---------------- | ------------- |
| Decimal          | `98_222`      |
| Hex              | `0xff`        |
| Octal            | `0o77`        |
| Binary           | `0b1111_0000` |
| Byte (`u8` only) | `b'A'`        |

**확실하게 정해진 경우가 아니면 Rust의 기본 값인 i32가 일반적으로는 좋은 선택입니다.**

{{< admonition note "Integer Overflow" >}}
Integer overflow란 type의 값 scope를 벗어나는 경우를 뜻합니다.

```rs
{
  let n: u8 = 256;
}
```


이 경우 `rust`에서는 2가지 모드 `--debug`, `--release`에 따라서 다르게 동작합니다.

1. debug 모드로 컴파일
  - integer overflow를 런타임에 체크하여, `“Unrecoverable Errors with panic!”`을 일으킵니다.
2. release 모드로 컴파일
  - panic 대신 `two’s complement wrapping`을 실시합니다.
  - u8의 경우 256이면 최소값인 0으로 값이 변환됩니다.

```rs
{
    // $ cargo build --release
    let mut a: u8 = 0;
    let mut b: u8 = 255;
    println!("{a}, {b}"); // 0, 255

    a = a - 1;
    b = b + 1;
    println!("{a}, {b}"); // 255, 0
}

```

{{< /admonition  >}}

- `Floating-Point Types`

```rs
fn main() {
    let x = 2.0; // f64
    let y: f32 = 3.0; // f32
}
```

- `The Boolean Type`

```rs
fn main() {
    let t = true;
    let f: bool = false; // with explicit type annotation
}
```

- `The Character Type`

```rs
fn main() {
    let c = 'z';
    let z: char = 'ℤ'; // with explicit type annotation
    let heart_eyed_cat = '😻';
}
```

`char` literal은 `single quotes`를 사용해야 합니다.

또한 Rust의 char타입은 four bytes`Unicode` Scalar를 표현하는 값입니다. (ASCII 보다 많은 표현 가능)
즉 한국어/중국어/일본어 표의 문자, 이모티콘, 넓이가 0인 공백문자를 `char`타입 변수로 받을 수 있습니다.

#### Compound types
> Compound types can **group multiple values** into one type. Rust는 2가지의 primative 컴파운드 타입(`tuples` and `arrays`)을 가지고 있습니다.

- `The Tuple Type`

튜플에 포함되는 각 값의 타입이 동일할 필요없이 서로 달라도 됩니다.

```rs
// 다른 타입들을 사용할 경우
{
  let tup: (i32, f64, u8) = (500, 6.4, 1);
}

// 단일 타입을 사용할 경우
{
    let tup = (500, 6.4, 1); 
    let (x, y, z) = tup; // 패턴 매칭 destructuring
}
```

`마침표(.)`를 통해서 튜플의 index 접근이 가능합니다.

```rs
{
    let x: (i32, f64, u8) = (500, 6.4, 1);
    let five_hundred = x.0;
    let six_point_four = x.1;
}
```

- `The Array Type`
  1. 튜플과는 다르게, 배열의 모든 요소는 모두 같은 타입이어야 합니다.
  2. Rust에서는 배열은 고정된 길이를 갖습니다. (선언되면 크기가 커지거나 작아지지 않는다.)

```rs
{
   let a = [1, 2, 3, 4, 5];

}
```

배열이 유용할 때는 당신의 데이터를 heap보다 stack에 할당하는 것을 원하거나, 항상 고정된 숫자의 요소(element)를 갖는다고 확신하고 싶을 때입니다. (vector 타입은 가변적)

```rs
{
  let months = ["January", "February", "March", "April", "May", "June", "July",
                "August", "September", "October", "November", "December"];
  
  let first = months[0];
  let second = months[1];   
}
```


index를 사용해 요소에 접근하려고 하면 Rust는 지정한 색인이 배열 길이보다 작은지 확인합니다. index가 array 길이보다 크면 *패닉(panic)*을 발생시킵니다.

또한 index 에러는 컴파일 시에는 아무런 에러도 발생시키지 않습니다만, 프로그램의 결과는 실행 중에 에러가 발생했고 성공적으로 종료되지 못했다고 나옵니다.

```js
$ cargo run
   Compiling arrays v0.1.0 (file:///projects/arrays)
    Finished dev [unoptimized + debuginfo] target(s) in 0.31 secs
     Running `target/debug/arrays`
thread '<main>' panicked at 'index out of bounds: the len is 5 but the index is
 10', src/main.rs:6
note: Run with `RUST_BACKTRACE=1` for a backtrace.
```

### 3-3. Functions

- Rust code uses `snake case` as the conventional style for function and variable names
- Rust는 당신의 함수의 위치를 신경쓰지 않습니다, 어디든 정의만 되어 있으면 됩니다.

#### Statements and Expressions

- `Statements`(구문) are instructions that perform some action and **do not return a value**.
- `Expressions`(표현식) **evaluate to a resultant value**.

```rs
// statement
{
  let y = 6;
  let x = (let y = 6); // compile error, return value가 없기 때문
}
// 
```

- `{ }` 또한 표현식입니다.

```rs
fn main() {
  let x = 5;

  let y = {
    let x = 3;
    x + 1 // expression, evaluated return value
  }; // let y = 4;
}
```

- `Expression`은 경우 종결을 나타내는 세미콜론(;)을 사용하지 않습니다.

만약 세미콜론을 표현식 마지막에 추가하면, 이는 구문으로 변경되고 반환 값이 아니게 됩니다. 이후부터 함수의 반환 값과 표현식을 살펴보실 때 이 점을 유의하세요.

#### Functions with Return Values

- `return` 키워드와 값을 써서 함수로부터 일찍 반환할 수 있지만, 대부분의 함수들은 암묵적으로 마지막 표현식을 반환합니다. 

```rs
fn five() -> i32 {
    5
}
```

위의 코드의 경우 `return 5`가 동작하게 됩니다.
이와 반대로

```rs
fn five() -> i32 {
    5;
}
```

와 같이 세미콜론을 붙이게 된다면 `()`(비어있는 튜플)을 반환하게 되어, `mismatched typed` 에러가 발생합니다.

### 3-4. Comments

```rs
fn main() {
    // I’m feeling lucky today.
    let lucky_number = 7;
}
```

### 3-5. Control Flow

- if의 조건문은 반드시 명시적으로 `bool` 타입이어야 합니다.

```rs
fn main() {
    let number = 3;

    // mismatched types
    if number {
        println!("number was three");
    }
}
```

#### Using `if` in a let Statement

```rs
{
  let number = if condition {
      5
  } else {
      6
  };
}
```

- 변수가 가질 수 있는 타입이 오직 하나여야 합니다. 그러므로 아래와 같은 코드는 에러입니다.


```rs
{
    let number = if condition {
        5
    } else {
        "six"
    };  
}
```

Rust는 컴파일 타임에 number 변수의 타입이 뭔지 확실히 정의해야 합니다. 그래야 `number`가 사용되는 모든 곳에서 유효한지 검증할 수 있으니까요. 

Rust는 number의 타입을 런타임에 정의되도록 할 수 없습니다. **컴파일러가 모든 변수의 다양한 타입을 추적해서 알아내야 한다면 컴파일러는 보다 복잡해지고 보증할 수 있는 것은 적어지게 됩니다.**

#### 반복문과 반복 (3)
> `loop`, `while`, `for`

- `loop`

```rs
{
  loop {
    do_something();
  }
}


```

어라? 이 방식은 `while true { }`와 큰 차이점이 없어 보입니다.
그래서 리서치해보니 loop은 expression으로 값을 return할 수 있습니다. 반면에 while과 for는 statement로 값을 return할 수 없습니다.

```rs
{
  let mut cnt = 0;
  let result = loop {
    cnt += 1;
    if cnt == 10 {
      break cnt * 2;
    }
  };
  
  assert_eq!(result, 20);
}
```

FYI 위 코드에서 **loop의 마지막 부분에 `};`이 사용되었다는 것을 알 수 있습니다.**

- `while`

```rs
{
  let mut i = 0;

  while i < 5 {
    do_something();
    i = i + 1
  }
}
```

그러나 이런 방식은 에러가 발생하기 쉽습니다.

- 개발자가 정확한 index를 사용하지 못하면 프로그램은 패닉을 발생합니다. 
- 또한 느립니다.
  - 이유는 컴파일러가 실행 간에 반복문을 통해 반복될 때마다 요소에 대한 조건 검사를 수행하는 런타임 코드를 추가하기 때문입니다.

이에 대한 대안으로 `for`을 사용합니다.

- `for`

```rs
{
  let arr = [1,2,3,4,5];

  for e in arr.iter() {
    do_somethin(e);
  }
}
```

이를 통해 `index`에 대한 실수를 줄일 수 있습니다.

만약 배열의 길이 만큼이 아닌 **특정한 횟수**만큼 반복하고 싶다면 `Range`를 사용합니다.

```rs
{
  for n in (1..4).rev() {
    do_something();
  }
}
```


