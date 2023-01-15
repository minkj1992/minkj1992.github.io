# The rust programming language


`Rust`에서 공식적으로 지원하는 ["The rust programming language"](https://doc.rust-lang.org)을 통해서 `Rust` 핵심 문법적 특징과 예시코드 그리고 이면에 숨은 디자인 원칙들을 정리 해보겠습니다.
<!--more-->


# 3. Common Programming Concepts

## 3-1. Variables and Mutability

> **Rust에서 기본 변수는 불변성입니다.**

이를 통해 컴파일 타임에 실수로 immutable 변수를 변경하는 버그를 잡아내도록 강제합니다. `mut` 키워드를 사용하면 Mutability를 제공할 수 있습니다. (가변변수)

- immutable vs mutable

만약 `매우 큰 구조체`를 다루는 경우 mutable 인스턴스를 사용하는 것이 새로 인스턴스를 할당하고 반환하는 것보다 빠를 수 있습니다. 데이터 크기가 작을수록 새 인스턴스를 생성하고 FP(함수적) 프로그래밍 스타일로 작성하는 것이 더 합리적이고, 그렇기에 약간의 성능 하락을 통해 가독성을 확보할 수 있다면 더 가치있는 선택입니다.

### Constants
> const vs Variables

- 상수에 대해서는 `mut`을 사용하는 것이 허용되지 않습니다: 상수는 항상 불변합니다.
- 상수는 `let`키워드 대신 `const`키워드를 사용해야 하고, 값의 type을 선언해야 합니다.
- 상수는 can be declared in any scope(including the global scope)
- 상수는 may be set only to a `constant expression`(상수 표현식), **not the result of a value that could only be computed at runtime.**, 즉 컴파일 타임에 하드코드 되어야합니다.

### Shadowing
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

## 3-2. Data Types
- **Rust의 타입은 크게 2가지: `scalar`와 `compound` 두가지로 나뉩니다.**
- Rust는 `statically typed language`(타입이 고정된 언어)입니다. 

즉 Rust는 컴파일타임에 모든 변수의 타입이 정해집니다. 그러므로 명시적으로 타입을 지정 또는 컴파일러가 타입을 추측할 수 있도록 선택의 폭을 줄여주어야 합니다.


### scalar types
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

### Compound types
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

## 3-3. Functions

- Rust code uses `snake case` as the conventional style for function and variable names
- Rust는 당신의 함수의 위치를 신경쓰지 않습니다, 어디든 정의만 되어 있으면 됩니다.

### Statements and Expressions

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

### Functions with Return Values

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

반면

```rs
fn five() -> () {
    5;
}
```
와 같이 return 타입을 tuple로 변경하게 될 경우 컴파일 에러가 사라지는 것을 확인할 수 있습니다.
## 3-4. Comments

```rs
fn main() {
    // I’m feeling lucky today.
    let lucky_number = 7;
}
```

## 3-5. Control Flow

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

### Using `if` in a let Statement

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

### 반복문 (3)
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

# 4. `Ownership`
> `소유권(Ownership)`은 러스트의 가장 유니크한 특성이며, 러스트가 가비지 콜렉터 없이도 메모리 안정성 보장을 하게 해줍니다.


<center>

![](/images/rust_ownership.png)

</center>


- 4-1. 소유권은 무엇인가?
- 4-2. `References` and `Borrowing`
- 4-3. The `Slice`

## 4.1. 소유권은 무엇인가?

모든 프로그램은 실행하는 동안 컴퓨터의 메모리(Heap)를 사용하는 방법을 관리해야 합니다.
크게 3가지 방식이 있는데요.

1. java, python 같은 언어는 실행될 때 사용하지 않는 메모리들을 정리하는 GC를 사용하고
2. c / c++ 같은 언어에서는 프로그래머가 직접 explicit(명시적)으로 사용한 메모리를 할당하고 해제해야 합니다.

`Golang`의 경우에 초기 개발자가 java의 GC 개발자라는 점을 활용해 `GC`를 가지고 있습니다.

마지막으로 `Rust`의 경우에는 **3. 컴파일 타임에 정한 규칙들을 활용해 소유권을 시스템으로 메모리가 관리됩니다.**

### Ownership 규칙 (3)

1. 러스트의 각각의 값(each value)은 해당값의 오너(owner)라고 불리우는 변수를 갖고 있다.
2. 한번에 딱 하나의 오너만 존재할 수 있다.
3. 오너가 스코프 밖으로 벗어나는 때, 값은 버려진다(dropped).

{{< admonition note "Heap vs Stack" >}}
- Stack
  - LIFO로 동작하는 자료구조.
  - stack의 element들은 동일한 size를 가져야 한다. 예를 들면 instance는 heap에, 이를 가리키는 로컬 변수 pointer는 정해져있는 element이므로 stack에 저장
  - Scope 벗어나면 pop됩니다.
- Heap
  - pointer를 타고타고 찾는 구조. (Linked list) 이기 때문에 stack은 top만 보는 특성에 비해서, 접근하는 데 느립니다. Heap의 경우에는 프로세서가 메모리 내부를 레퍼런스들을 타고타면서 jump해야하지만, stack의 경우에는 데이터가 붙어있으니 덜 jump해도 됩니다. 그러므로 더 빠릅니다.
  - 이 때문에 stack은 modify 비용이 크지만, heap의 경우에는 modify 되는 다양한 size 데이터들을 저장할 수 있습니다.
  - 컴파일 타임에 크기가 결정되어 있지 않거나 크기가 변경될 수 있는 데이터를 위해서는, 힙에 데이터를 저장할 수 있습니다.
  - `allocating`(allocating on the heap): 데이터를 힙에 넣을때, 먼저 저장할 공간이 있는지 물어봅니다. 그러면 운영체제가 충분히 커다란 힙 안의 빈 어떤 지점을 찾아서 이 곳을 사용중이라고 표시하고, 해당 지점의 포인터를 우리에게 돌려주죠.
{{< /admonition  >}}

### Strings 타입

Rust에서 string literal은 compile시 binary로 저장됩니다. 즉 하드코딩 되어있습니다. 그러니 stack, heap 둘 중 어디에도 저장되지 않습니다. 

Binary로 하드코딩 되어있기 때문에, 당연히 immutable합니다.

{{< admonition question "Java's String Literal" >}}
비교를 위해서, JAVA7 이후 JVM의 경우 string literal은 `constant pool`에 저장됩니다. Constant Pool은 `Heap`의 일부이기 때문에 GC됩니다. (7이전에는 PermGen 영역)
{{< /admonition  >}}

Immutable한 `String Literal`을 런타임에 mutable하게 만들기 위해서는 

- `from()`: type convert, 내부적으로는 os에 메모리 요청
- `.push_str()`: append

```rs
{
  let mut s = String::from("hello");
  s.push_str(", minwook :)");
  // "hello, minwook :)"
}
```

### Memory and Allocation

String 타입은 변경 가능하고 커질 수 있는 텍스트를 지원하기 위해 만들어졌고, 우리는 힙에서 컴파일 타임에는 알 수 없는 어느 정도 크기의 메모리 공간을 할당받아 내용물을 저장할 필요가 있습니다. 이는 즉 다음을 의미합니다:

1. 런타임에 운영체제로부터 메모리가 요청되어야 한다.
2. String의 사용이 끝났을 때 운영체제에게 메모리를 반납(free)할 방법이 필요하다.

**즉 allocate와 free가 쌍으로 구현되어야 합니다.**

첫번째는 `String::from`을 호출하면 가능하며, 2번의 경우 대부분 언어는 GC를 통해 관리합니다. 

Rust는 2번을 위해서 "메모리는 변수가 소속되어 있는 scope 밖으로 벗어나는 순간 자동으로 free(반납)됩니다." 운영체제에게 메모리를 반납(free)시키기 위해 Rust는 `drop()` 함수가 존재하며, 중괄호가 `}` 닫힐 때 자동으로 `drop()`를 호출합니다.

```rs
{
    let s = String::from("hello"); // s는 여기서부터 유효합니다

    // s를 가지고 뭔가 합니다
}                                  // 이 스코프는 끝났고, s는 더 이상 
                                   // 유효하지 않습니다
```

C++에서는 이렇게 `아이템의 수명주기의 끝나는 시점에 자원을 해제하는 패턴`을 Resource Acquisition Is Initialization, `RAII`) 라고 부릅니다.

이제 이 `free`를 힙에 할당시킨 데이터를 사용하는 여러 개의 변수를 사용하고자 할 경우, 즉 좀더 복잡한 상황들을 살펴보겠습니다.

1. 변수와 데이터가 상호작용하는 방법: `Move`
2. 변수와 데이터가 상호작용하는 방법: `Clone`
3. 스택에만 있는 데이터: `Copy`


#### 1. 변수와 데이터가 상호작용하는 방법: `Move`
```rs
let s1 = String::from("hello");
let s2 = s1;
```

`String`은 아래 그림의 왼쪽과 같이 세 개의 부분으로 이루어져 있습니다.
문자열의 내용물을 담고 있는 메모리의 포인터, 길이, 그리고 용량입니다. 이 데이터의 그룹은 스택에 저장됩니다. 내용물을 담은 오른쪽의 것은 힙 메모리에 있습니다.


<center>

![](/images/rust_mem1.svg)

</center>

s2에 s1을 assign 하게될경우 대부분 프로그래밍 언어에서는 총 2가지 현상이 일어날 수 있습니다.

1. 얕은 복사(Shallow Copy)
2. 깊은 복사(Deep Copy)


- 얕은 복사(Shallow Copy): heap 메모리 상의 데이터는 복사되지 않는 것, 즉 stack의 레퍼런스(포인터)만 복사 되는 것.
  - 참조해서 데이터를 수정하게 될 경우, 원하지 않는 현상이 일어날 수 있습니다.
<center>

![](/images/rust_mem2.svg)

</center>

앞서 우리는 변수가 스코프 밖으로 벗어날 때, 러스트는 자동적으로 drop함수를 호출하여 해당 변수가 사용하는 힙 메모리를 제거한다고 했습니다.
그러므로 이렇게 shallow copy가 일어나게 된다면, s2와 s1은 동시에 메모리를 해제(free)하려 합니다. 이는 두번-해제(`double free`)오류 라고 알려져있습니다. (메모리 안전성 관련 버그 중 하나)


메모리를 두번이상 해제하는 것은 memory corruption(손상)의 원인이 되며, 보안 취약성을 일으킬 수 있습니다.



- 깊은 복사(Shallow Copy): heap 메모리 상의 데이터까지 복사되는 것

<center>

![](/images/rust_mem3.svg)

</center>

깊은 복사의 경우에는, 힙 안의 데이터가 클 경우 s2 = s1 연산은 런타임 상에서 매우 느려질 가능성이 있습니다.

그래서 Rust는 `Move`라는 개념을 도입합니다.


```rs
{
    let s1 = String::from("hello");
    let s2 = s1;
    println!("{s1}, {s2}"); // compile error
}
```

실제 위의 코드는 아래와 같은 컴파일 에러를 발생시킵니다.
```javascript
error[E0382]: use of moved value: `s1`
 --> src/main.rs:4:27
  |
3 |     let s2 = s1;
  |         -- value moved here
4 |     println!("{}, world!", s1);
  |                            ^^ value used here after move
  |
  = note: move occurs because `s1` has type `std::string::String`,
which does not implement the `Copy` trait
```

- `Move`: `shallow copy`에서 `첫번째 변수의 무효화` 된 개념.

러스트는 s2에 s1을 대입하게 될 경우, 첫번째 변수인 `s1`을 무효화 시킵니다. 그래서 아래와 같은 현상이 내부적으로 발생합니다.


<center>

![](/images/rust_mem4.svg)

</center>

그래서 위의 코드에서 move된 s1에 대한 참조를 없애게 된다면 정상적으로 동작하게 됩니다.

```rs
{
    let s1 = String::from("hello");
    let s2 = s1;
    println!("{s2}"); // "hello"
}
```

이런 `Move` 개념을 통해서 Rust는 아래의 문제들을 해결하면서

1. shallow: `double free` 에러.
2. shallow: 원치 않게 오리지널 `heap` 데이터 변경.
3. deep: `heap`의 데이터 복사에 의한 퍼포먼스 저하.

동시에 allocate한 메모리에 대해서 **운영체제에게 메모리를 반납(free)**을 할 수 있게 되었습니다.

#### 2. 변수와 데이터가 상호작용하는 방법: `Clone`
> Rust의 Deep copy 방법

만일 String의 스택 데이터 만이 아니라, 힙 데이터를 깊이 복사하기를 정말 원한다면, clone이라 불리우는 공용 메소드를 사용할 수 있습니다. 

```rs
let s1 = String::from("hello");
let s2 = s1.clone();

println!("s1 = {}, s2 = {}", s1, s2); // s1 = hello, s2 = hello
```

러스트는 `.clone()`이라는 명시적인 제약을 두어, 성능상의 문제가 발생할 때 손쉽게 찾아볼 수 있게 문법적으로 강제하였습니다.

#### 3. 스택에만 있는 데이터: `Copy`

`String`에서는 불가능했던, 코드가 앞서 보았던 int예시에서는 문제없이 동작하는 것을 알 수 있습니다.

```rs
{
  let x = 5;
  let y = x;
  println!("x = {}, y = {}", x, y);  // x = 5, y = 5
}
```

**위의 코드의 `int`타입의 경우에는 `clone`을 호출하지 않았지만, x도 유효하며 y로 `Move`(이동)하지 않았습니다.**


그 이유는 정수형과 같이 컴파일 타임에 결정되어 있는 크기의 타입은 스택에 모두 저장되기 때문에, 실제 값의 복사본이 빠르게 만들어질 수 있습니다. 이는 변수 `y`가 생성된 후에 `x`가 더 이상 유효하지 않도록 해야할 이유가 없다는 뜻입니다. **바꿔 말하면, 여기서는 깊은 복사와 얕은 복사 간의 차이가 없다는 것으로, `clone`을 호출하는 것이 보통의 얕은 복사와 아무런 차이점이 없어 우리는 이를 무시할 수 있다는 것입니다.**

러스트는 정수형과 같이 **스택에 저장할 수 있는 타입**에 대해 `Copy trait`(카피 트레잇)이라고 불리우는 특별한 어노테이션(annotation)을 가지고 있습니다. 만일 어떤 타입이 Copy 트레잇을 갖고 있다면, 대입 과정 후에도 예전 변수를 계속 사용할 수 있습니다.

그리고 당연하게도 만약 `Copy trait`을 어노테이트한 타입이, `Drop trait`도 어노테이트 했다면 에러를 내도록합니다.

아래는 Rust가 지원하는 타입중에서, `Copy`가 가능한 타입 리스트입니다.

- int 타입들 (i.g `u32`)
- bool 타입들
- float 타입들
- `char` 타입들
- Copy가 가능한 타입만으로 구성된 튜플 (즉 숫자형 또는 bool 또는 char인 경우)
  - i.g. `(i32, i32)`는 Copy가 되지만, `(i32, String)`은 안됩니다.
  - i.g. `(i32, f64)`는 혼합되어 있지만 Copy가 됩니다. 왜냐하면 숫자형, bool, char안에 포함되는 타입들이기 때문에.

### 소유권과 함수(Ownership and Functions)
> 함수에 변수를 넘기는 것(args, param) 또한 대입(`=`)과 마찬가지로 이동(Move)하거나 복사(Move)가능합니다.

```rs
fn main() {
    let s = String::from("hello");  // s가 스코프 안으로 들어왔습니다.

    takes_ownership(s);             // s의 값이 함수 안으로 이동했습니다.
                                    // 그리고 여기 부터는 더이상 유효하지 않습니다.
    let x = 5;                      // x가 스코프 안으로 들어왔습니다.

    makes_copy(x);                  // x가 함수 안으로 이동했습니다만,
                                    // i32는 Copy가 되므로, x를 이후에 계속
                                    // 사용해도 됩니다.

} // 여기서 x는 스코프 밖으로 나가고, s도 그 후 나갑니다. 하지만 s는 이미 이동되었으므로,
  // 별다른 일이 발생하지 않습니다.

fn takes_ownership(some_string: String) { // some_string이 스코프 안으로 들어왔습니다.
    println!("{}", some_string);
} // 여기서 some_string이 스코프 밖으로 벗어났고 `drop`이 호출됩니다. 메모리는
  // 해제되었습니다.

fn makes_copy(some_integer: i32) { // some_integer이 스코프 안으로 들어왔습니다.
    println!("{}", some_integer);
} // 여기서 some_integer가 스코프 밖으로 벗어났습니다. 별다른 일은 발생하지 않습니다.
```

```js
hello
5
```

만일 우리가 `s`를 `takes_ownership` 함수를 호출한 이후에 사용하려 한다면, 러스트는 컴파일 타임 오류를 낼 것입니다.


> 저는 개인적으로 맞는지 아닌지는 모르겠지만, scalar 타입들에 대해서는 `String`이 아니면, 모두 `Copy` 기억하려고 합니다.

또한 문득 개인적으로 Copy vs Clone에 대해서는 명확하게 구분이 불가능한 것 같아 

{{< admonition question "Copy vs Clone" >}}
- [Youtube: copy vs clone](https://www.youtube.com/watch?v=jC8uR6aaQKI)을 정리할 예정입니다.
{{< /admonition  >}}

### Return Values and Scope
> `return` 또한 `소유권`을 이동시킵니다.

```rs
fn main() {
    let s1 = gives_ownership();         // gives_ownership은 반환값을 s1에게
                                        // 이동시킵니다.

    let s2 = String::from("hello");     // s2가 스코프 안에 들어왔습니다.

    let s3 = takes_and_gives_back(s2);  // s2는 takes_and_gives_back 안으로
                                        // 이동되었고, 이 함수가 반환값을 s3으로도
                                        // 이동시켰습니다.

} // 여기서 s3는 스코프 밖으로 벗어났으며 drop이 호출됩니다. s2는 스코프 밖으로
  // 벗어났지만 이동되었으므로 아무 일도 일어나지 않습니다. s1은 스코프 밖으로
  // 벗어나서 drop이 호출됩니다.

fn gives_ownership() -> String {             // gives_ownership 함수가 반환 값을
                                             // 호출한 쪽으로 이동시킵니다.

    let some_string = String::from("hello"); // some_string이 스코프 안에 들어왔습니다.

    some_string                              // some_string이 반환되고, 호출한 쪽의
                                             // 함수로 이동됩니다.
}

// takes_and_gives_back 함수는 String을 하나 받아서 다른 하나를 반환합니다.
fn takes_and_gives_back(a_string: String) -> String { // a_string이 스코프
                                                      // 안으로 들어왔습니다.

    a_string  // a_string은 반환되고, 호출한 쪽의 함수로 이동됩니다.
}
```

어떤 값을 다른 변수에 대입하면 값이 이동됩니다. 힙에 데이터를 갖고 있는 변수가 스코프 밖으로 벗어나면, 해당 값은 데이터가 다른 변수에 의해 소유되도록 이동하지 않는한 `drop`에 의해 제거될 것입니다.

만일 함수에게 값을 사용할 수 있도록 하되 소유권은 갖지 않도록 하고 싶다면요? 함수의 본체로부터 얻어진 결과와 더불어 우리가 넘겨주고자 하는 어떤 값을 다시 쓰고 싶어서 함께 반환받아야 할 경우가 대표적입니다.

이런 경우에는 아래와 같이 튜플을 이용하여 여러 값을 돌려받는 식으로 가능하긴 합니다.

```rs
fn main() {
    let s1 = String::from("hello");
    let (s2, len) = calculate_length(s1);

    println!("The length of '{}' is {}.", s2, len);
}

fn calculate_length(s: String) -> (String, usize) {
    let length = s.len(); // len()함수는 문자열의 길이를 반환합니다.

    (s, length)
}
```
하지만 이건 보편화 되기에 너무 많이 과한 작업이 됩니다. 다행히, 러스트는 이런 경우를 위해 참조자(references)라는 feature 문법을 도입했습니다.

> for using a value without transferring ownership, called `references`.


## 4.2. `References` and `Borrowing`


- References (참조자, 불변 참조자, immutable)
- Borrowing
- Mutable references (가변 참조자)
- Dangling References
- The Rules of References

### References

바로 위의 코드에서, s1의 소유권을 넘기는 대신 Rust는 함수 signature를 reference 타입으로 변경시켜서 처리할 수 있습니다.

```rs
fn main() {
  let s1 = String::from("hello");
  let len = calculate_length(&s1);
  
  println!("The length of '{}' is {}.", s1, len);
}

fn calculate_length(s: &String) -> usize {
    s.len()
} // 여기서 s는 스코프 밖으로 벗어났습니다. 하지만 가리키고 있는 값에 대한 소유권이 없기
  // 때문에, 아무런 일도 발생하지 않습니다.
```

1. s2 선언부가 사라지고, calculate_length() 튜플 return이 사라졌습니다.
2. calculate_length에 `&s1`을 사용합니다.

**`&`를 `references`(참조)라고 부르며, 어떤 값의 소유권을 넘기지 않고 참조하도록 할 수 있습니다.**

<center>

![](/images/rust_ref.svg)

</center>

위 도표와 같이, `calculate_length`의 param인 참조자 `s`라는 포인터 타입 local변수가 생성되고, 이 포인터가 s1을 참조하는 형식입니다.

`Reference`를 활용하게 되면, **참조자는 소유권을 갖고 있지는 않기 때문에, 이 참조자가 가리키는 값은 참조자가 스코프 밖으로 벗어났을 때도 메모리가 반납되지 않을 것입니다.**

**`&s`(참조자)를 파라미터로 가지게 되어, 함수는 소유권을 가지지 않게 될 수 있습니다. 그러므로 당연히 소유권을 되돌려주기 위해 값을 다시 반환할 필요도 없게 됩니다.**

### Borrowing

함수의 파라미터로 `참조자`(`&s`)를 만드는 것을 `Borrowing`(빌림)라고 부릅니다.

어떤 무언가를 빌렸다는 것은, 함부로 빌린 물건을 대해서는 안됩니다. 우리가 빌린 무언가를 고치려고 시도한다면 무슨 일이 생길까요? 

```rs
fn main() {
  let s1 = String::from("hello");
  modify(&s1);
}

fn modify(s: &String) {
  s.push_str(", minwook");
}
```

```js
error[E0596]: cannot borrow `*s` as mutable, as it is behind a `&` reference
 --> src/main.rs:7:3
  |
6 | fn modify(s: &String) {
  |              ------- help: consider changing this to be a mutable reference: `&mut String`
7 |   s.push_str(", minwook");
  |   ^^^^^^^^^^^^^^^^^^^^^^^ `s` is a `&` reference, so the data it refers to cannot be borrowed as mutable
```

**변수가 기본적으로 불변인 것처럼, 참조자도 마찬가지입니다.**

### Mutable references (가변 참조자)
> 변할 수 있는 참조자 (`&mut`)

빌린 물건을 의도적으로 수정하고 싶다면, `&mut`를 붙여주면 됩니다. 또한 원래 물건(변수) 또한 mut 처리해주어야 합니다.

```rs
fn main() {
  let mut s1 = String::from("hello");
  modify(&mut s1);
}

fn modify(s: &mut String) {
  s.push_str(", minwook");
}
```

하지만 가변 참조자는 딱 한가지 큰 제한이 있습니다.

**특정한 스코프 내에 특정한 데이터 조각에 대한 가변 참조자를 딱 하나만 만들 수 있다는 겁니다.**

아래 코드는 실패할 겁니다.

```rs
{
  let mut s = String::from("hello");
  let r1 = &mut s;
  let r2 = &mut s;
}

// error[E0499]: cannot borrow `s` as mutable more than once at a time.
```

이런 불편한 제한사항 덕분에, 여러분이 가질 수 있는 이점은 바로 **러스트가 컴파일 타임에 데이터 레이스(data race)를 방지할 수 있도록 해준다는 것입니다.**

아래는 Data race가 발생될 수 있는 race condition 조건입니다.

1. 두 개 이상의 포인터가 동시에 같은 데이터에 접근한다.
2. 그 중 적어도 하나의 포인터가 데이터를 쓴다.
3. 데이터에 접근하는데 동기화(sync)를 하는 어떠한 수단 없다.

Rust는 하나의 원본값(변수)에 대해 같은 scope 안에서 2개 이상의 &mut를 만들 수 없도록, 문법적 강제를 합니다. 이로써 1번의 조건을 사전에 차단하였습니다.

Rust의 scope를 사용하면, "동시"에 만드는 것을 우회하여, 여러개의 가변 참조자를 만들 수 있습니다.

```rs
let mut s = String::from("hello"):
{
  let r1 = &mut s;
} // 여기서 r1은 스코프 밖으로 벗어났으므로, 우리는 아무 문제 없이 새로운 참조자를 만들 수 있습니다.

let r2 = &mut s;
```

mutable reference(가변 참조자)와 immutable reference(불변 참조자, 디폴트)를 혼용에 대한 규칙도 존재합니다.

```rs
let mut s = String::from("hello");

let r1 = &s;
let r2 = &s; // ok
```

immutable reference는 중복해도 modify 할 수없으니(Read Only) 문제 없습니다.

반면에 혼용하게 될 경우에는 에러가 생깁니다.

```rs
let mut s = String::from("hello");

let r1 = &s;
let r2 = &mut s;
// error[E0502]: cannot borrow `s` as mutable because it is also borrowed as immutable.
```

위와 같이 불변 참조자를 가지고 있을 동안 가변 참조자를 만들 수 없습니다. 
불변 참조자의 사용자는 사용중인 동안에 값이 값자기 바뀌리라 예상하지 않기 때문입니다.

이는 순서를 반대로 해도 같습니다. (&mut 선언 후 & 재선언)

```rs
let mut s = String::from("hello");

let r2 = &mut s;
let r1 = &s;
// error[E0502]: cannot borrow `s` as immutable because it is also borrowed as mutable
```

### Dangling References

포인터가 있는 언어에서는 자칫 잘못하면 댕글링 포인터(dangling pointer, 허상 포인터)를 만들기 쉬운데, 댕글링 포인터란 어떤 메모리를 가리키는 포인터를 보존하는 동안, 포인터가 가리키고 있는 메모리를 해제함으로써 다른 개체에게 사용하도록 줘버렸을 지도, 또는 제거되었을지도 모를 메모리를 참조하고 있는 포인터를 말합니다.

<center>

![](/images/dangling_ptr.jpeg)

</center>

이와는 반대로, 러스트에서는 컴파일러가 모든 참조자들이 댕글링 참조자가 되지 않도록 보장해 줍니다. 

**만일 우리가 어떤 데이터의 참조자를 만들었다면, 컴파일러는 그 참조자가 스코프 밖으로 벗어나기 전에는 데이터가 스코프 밖으로 벗어나지 않을 것임을 확인해 줍니다.** 댕글링 참조자를 만들어보며 보도록 하겠습니다.

```rs
fn main() {
  let dangle_ref = dangle();
}

fn dangle() -> &String {
  let s = String::from("hello");

  &s
} // s는 }를 벗어나는 시점에 free되므로, &s는 dangling reference입니다.
```

```js
error[E0106]: missing lifetime specifier

...

help: this function's return type contains a borrowed value, but there is no value for it to be borrowed from. 
(해석: 이 함수의 반환 타입은 빌린 값을 포함하고 있는데, 빌려온 실제 값은 없습니다.)

```

여기서의 해법은 String을 직접 반환하는 것입니다. 소유권이 밖으로 이동되었고, 아무것도 할당 해제되지 않습니다.

```rs
fn no_dangle() -> String {
    let s = String::from("hello");

    s
}
```

### The Rules of References
> 지금 까지 Reference(참조자)에 대해 논의한 것들을 정리해봅니다.

1. 어떠한 경우이든, Rust는 아래 두가지 중에서, **오직 하나만 가질 수 있습니다.**
  - `one mutable reference`(하나의 가변 참조자, `&mut`)
  - `any number of immutable references`(임의 개수의 불변 참조자들, `&`)
2. 참조자는 항상 유효해야만 한다.


다음으로, 우리는 다른 종류의 참조자인 `슬라이스(slice)`를 알아보겠습니다.
## 4.3. `Slices`
> 슬라이스는 여러분이 컬렉션(collection) 전체가 아닌 컬렉션의 연속된 일련의 요소들을 참조할 수 있게 합니다.

- 소유권을 갖지 않는 또다른 데이터 타입은 슬라이스입니다.

1. slice의 필요성
2. string slice란?
3. 그 밖의 슬라이스들

### slice가 없다면?

여기 작은 프로그래밍 문제가 있습니다.

스트링을 입력 받아 그 스트링에서 찾은 첫번째 단어를 반환하는 함수를 작성해보세요. 

만일 함수가 공백문자를 찾지 못한다면, 이는 전체 스트링이 한 단어라는 의미이고, 이때는 전체 스트링이 반환되어야 합니다.

```rs
fn first_word(s: &String) -> ?
```

이 함수 first_word는 &String을 파라미터로 갖습니다. 우리는 소유권을 원하지 않으므로, 이렇게 해도 좋습니다. 하지만 뭘 반환해야할까요? 우리는 스트링의 일부에 대해 표현할 방법이 없습니다.


하지만 단어의 끝부분의 인덱스를 반환할 수는 있겠습니다.

```rs
// String 파라미터의 바이트 인덱스 값을 반환하는 first_word 함수
fn first_word(s: &String) -> usize {
  // 공백인지 확인할 필요가 있기 때문에, String은 as_bytes 메소드를 이용하여 바이트 배열로 변환.
  let bytes = s.as_bytes();

  for (i, &item) in bytes.iter().enumerate() {
    // 공백 문자를 찾았다면, 이 위치를 반환합니다.
      if item == b' ' {
          return i;
      }
  }

  s.len()
}
```

이제 우리에게 스트링의 첫번째 단어의 끝부분의 인덱스를 찾아낼 방법이 생겼습니다. usize를 그대로 반환하고 있지만, 이는 `&String`의 내용물 내에서만 의미가 있습니다. 

바꿔 말하면, 이것이 String로부터 분리되어 있는 숫자이기 때문에, 아래 코드 처럼 이것이 나중에도 여전히 유효한지를 보장할 길이 없습니다.

```rs
fn main() {
    let mut s = String::from("hello world");

    let word = first_word(&s); // word는 5를 갖게 될 것입니다.

    s.clear(); // 이 코드는 String을 비워서 ""로 만들게 됩니다.

    // word는 여기서 여전히 5를 갖고 있지만, 5라는 값을 의미있게 쓸 수 있는 스트링은 이제 없습니다.
    // word는 이제 완전 유효하지 않습니다!
}
```

이처럼 `word`의 인덱스가 `s` 데이터와 `싱크`가 안맞을 것은, 지겹고 쉽게 발생할 수 있는 오류입니다. 이러한 인덱스들을 관리하는 것은 우리가 `second_word` 함수를 작성했을 때 더더욱 다루기 어려워집니다. 이 함수의 시그니처는 아래와 같은 모양이 되어야 할 것입니다.

```rs
fn second_word(s: &String) -> (usize, usize) {

}
```

이로써, 모든 개발자들은 매번 동기화를 유지할 필요가 있는, 원본 데이터와 분리된 세 개의 변수들을 가지게 되었습니다. 이를 해결하기 위해서 도입한 문법이 바로

`String slice`입니다.

### 스트링 슬라이스 
> string slice는 String의 일부분에 대한 reference(참조자)입니다.

```rs
let s = String::from("hello world");

let hello = &s[0..5]; // 0, 1, 2, 3, 4
// let hello = &s[..5];와 동일

let world = &s[6..11]; // 6, 7, 8, 9, 10
// let len = s.len();
// let slice = &s[3..len]; 와 동일
// let slice = &s[3..]; 와 동일
```


<center>

![](/images/str_slice.svg)

</center>

전체 스트링의 슬라이스를 만들기 위해 양쪽 값을 모두 생략할 수 있습니다. 따라서 아래 두 줄의 표현은 동일합니다.

```rs
let s = String::from("hello");

let n = s.len();

let slice = &s[0..n];
let slice = &s[..];
```

이 모든 정보를 잘 기억하시고, first_word가 슬라이스를 반환하도록 다시 작성해봅시다. “스트링 슬라이스”를 나타내는 type은 `&str`로 씁니다.

```rs
fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}
```

만약 아까 와 같이, clear()한 경우 발생하는 문제는 그럼 어떻게 해결될 까요?

```rs
fn main() {
    let mut s = String::from("hello world");

    let word = first_word(&s);

    s.clear(); // Error!

    println!("the first word is: {}", word);
}
```

```js
17:6 error: cannot borrow `s` as mutable because it is also borrowed as
            immutable [E0502]
    s.clear(); // Error!
    ^
15:29 note: previous borrow of `s` occurs here; the immutable borrow prevents
            subsequent moves or mutable borrows of `s` until the borrow ends
    let word = first_word(&s);
                           ^
18:2 note: previous borrow ends here
fn main() {

}
```

`Borrowing`(빌림) 규칙에서 만일 무언가에 대한 불변 참조자를 만들었을 경우, 가변 참조자를 만들 수 없다는 점을 상기해보세요. 

`clear()`가 String을 잘라낼 필요가 있기 때문에, 이 함수는 가변 참조자를 갖기 위한 시도를 할 것이고, 이는 실패하게 됩니다.

- **스트링 리터럴은 슬라이스입니다**

스트링 리터럴이 바이너리 안에 저장된다고 하는 얘기를 상기해봅시다. 이제 슬라이스에 대해 알았으니, 우리는 스트링 리터럴을 적합하게 이해할 수 있습니다.

```rs
let s = "Hello, world!"; // s는 &str 타입
```

여기서 s의 타입은 &str입니다.

이것은 바이너리의 특정 지점을 가리키고 있는 슬라이스입니다. 이는 왜 스트링 리터럴이 불변인가도 설명해줍니다; &str은 불변 참조자이기 때문입니다.



- **파라미터로서의 스트링 슬라이스**

리터럴과 String의 슬라이스를 얻을 수 있다는 것을 알게 되었다면, `first_word`는 

```rs
fn first_word(s: &str) -> &str {
```

로 시그니처를 변경시킬 수 있습니다.

```rs
fn main() {
    let my_string = String::from("hello world");

    // first_word가 `String`의 슬라이스로 동작합니다.
    let word = first_word(&my_string[..]);

    let my_string_literal = "hello world";

    // first_word가 스트링 리터럴의 슬라이스로 동작합니다.
    let word = first_word(&my_string_literal[..]);

    // 스트링 리터럴은 또한 스트링 슬라이스이기 때문에,
    // 아래 코드도 슬라이스 문법 없이 동작합니다!
    let word = first_word(my_string_literal);
}
```

이런 Rust의 확장성은 python 처럼 slice를 편하게 만들어주네요.

### 그 밖의 슬라이스들

`slice`는 스트링 이외에도 array, vector에 모두 동작합니다.

# 5. `Structs`
> Using Structs to Structure Related Data.

OOP의 핵심, 데이터 속성과 메소드, 그리고 `associated functions`를 묶어주는 `struct`를 러스트 또한 제공합니다.

- 5.1. Defining and Instantiating Structs
- 5.2. An Example Program Using Structs
- 5.3. Method Syntax

## 5.1. Defining and Instantiating Structs

- 튜플과 유사하게, 구조체의 구성요소들은 각자 다른 타입을 지닐 수 있습니다.
- 구조체를 정의할 때는 struct 키워드를 먼저 입력하고 명명할 구조체명을 입력하면 됩니다.

```rs
struct User {
  name: String,
  email: String,
  sign_in_count: u64,
  is_active: bool,
}
```

- 구조체를 통해 인스턴스를 생성할때, 필드들의 순서가 정의한 필드의 순서와 같을 필요는 없습니다.
- User 구조체 정의에서, `&str` 문자 슬라이스 타입 대신 `String`타입을 사용했습니다.
- 이는 의도적인 선택으로, **구조체 전체가 유효한 동안 구조체가 그 데이터를 소유하게 하고자 함입니다.**

```rs
#[derive(Debug)]
struct User {
  name: String,
  email: String,
  sign_in_count: u64,
  is_active: bool,
}

fn main() {

    let mut user = User {
      email: String::from("leoo_is_cool@leoo.com"),
      name: String::from("leoo.j"),
      is_active: true,
      sign_in_count: 1,
    };
    user.is_active = false;    
    println!("{:#?}", user);
}
```


- 변수명이 필드명과 같을 때 간단하게 필드 초기화하기

```rs
fn build_user(email: String, name: String) -> User {
  User {
    email, // email: email과 동일
    name, // name: name과 동일
    is_actieve: true,
    sign_in_count: 1,
  }
}
```

### `struct update syntax`
> 구조체 갱신법을 이용하여 기존 구조체 인스턴스로 새 구조체 인스턴스 생성하기

`..` 연산자를 활용하면 쉽게 인스턴스를 생성할 수 있습니다.

- before

```rs
let user2 = User {
    email: String::from("another@example.com"),
    name: String::from("anotherusername567"),
    is_active: user1.is_active,
    sign_in_count: user1.sign_in_count,
};
```

- after

```rs
let user2 = User {
    email: String::from("another@example.com"),
    name: String::from("anotherusername567"),
    ..user1
};
```

### `tuple structs`
> 이름이 없고 필드마다 타입은 다르게 정의 가능한 튜플 구조체.

러스트의 `tuple struct`는 파이썬의 `Namedtuple`과 비슷해 보입니다.

```python
Point = namedtuple('Point', ['x', 'y'])
p = Point(11, y=22)
```


```rs
extern crate assert_type_eq;

struct Color(i32, i32, i32);
struct Point(i32, i32, i32);

fn main() {
  let black = Color(0,0,0);
  let origin = Point(0,0,0);
}
```

`black`과 `origin`은 다른 튜플 구조체이기 때문에, 다른타입 입니다.

### `unit-like structs`
> 필드가 없는 유사 유닛 구조체

러스트에서 `필드가 없는 튜플`인 `()`을 `unit` 또는 `unit type`라고 부릅니다.

- 러스트에서 어떤 필드도 없는 `구조체` 역시 정의 가능합니다. 이를 유닛처럼 동작하는 구조체, `unit-like structs`라고 부릅니다.

유사 유닛 구조체는 특정한 타입의 트레잇(`trait`)을 구현해야하지만 타입 자체에 데이터를 저장하지 않는 경우에 유용합니다.


### 구조체 안의 데이터 소유권 (`Ownership`)

- 구조체가 소유권이 없는 데이터의 참조를 저장할수는 있지만, `라이프타임(lifetimes)`의 사용을 전제로 합니다. 
- 라이프타임은 구조체가 존재하는동안 참조하는 데이터를 계속 존재할 수 있도록 합니다. 라이프타임을 사용하지 않고 참조를 저장하고자 하면 에러가 발생합니다.

```rs
struct User {
    username: &str, // 참조(reference) 데이터
    email: &str, // // 참조(reference) 데이터
    sign_in_count: u64,
    active: bool,
}
fn main() {
    let user1 = User {
        email: "someone@example.com",
        username: "someusername123",
        active: true,
        sign_in_count: 1,
    };
}

// error[E0106]: missing lifetime specifier
```

