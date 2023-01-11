# The rust programming language


`Rust`에서 공식적으로 지원하는 ["The rust programming language"](https://doc.rust-lang.org)을 통해서 `Rust` 핵심 문법적 특징과 예시코드 그리고 이면에 숨은 디자인 원칙들을 정리 해보겠습니다.
<!--more-->


## Common Programming Concepts

### Variables and Mutability

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

