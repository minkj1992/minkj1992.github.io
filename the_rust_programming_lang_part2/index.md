# [Part2] The rust programming language


[Part1. Rust 프로그래밍](https://minkj1992.github.io/the_rust_programming_lang/)에 이어서 Rust의 고급 문법들을 익혀보겠습니다.
<!--more-->

# 11. Testing

## 11-1. Test 작성하기

가장 단순하게 말하면, 러스트 내의 테스트란 test 속성(`attribute`)이 주석으로 달려진 (annotated) 함수입니다.

`attribute`란 러스트 코드 조각에 대한 메타데이터입니다.

```rs
#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}
```


- `#[test]`를 fn위에 올리게된다면, 이 함수는 test 함수라는 것을 뜻합니다.
- `assert_eq!`, `assert_ne!` 매크로는 `PartialEq`와 `Debug` 트레잇을 구현해야합니다. 즉 `#[derive(PartialEq, Debug)]` 어노테이션이 필요합니다.


test module이 외부 fn에 대해 테스트를 작성하기 위해서 일반적으로 glob을 활용한 `use super::*`을 내부에 적습니다.

```rs
pub fn add_numbers(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_works() {
        let expect = 4;
        assert_eq!(expect, add_numbers(2,2));
    }
}


// $ cargo test
```


- `should_panic`을 활용하면 에러 처리하는 부분 또한 테스트할 수 있습니다.

```rs
pub struct Guess {
    value: u32,
}

impl Guess {
    pub fn new(value: u32) -> Guess {
        if value < 1 {
            panic!("Guess value must be greater than or equal to 1, got {}.",
                   value);
        } else if value > 100 {
            panic!("Guess value must be less than or equal to 100, got {}.",
                   value);
        }

        Guess {
            value
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    #[should_panic(expected = "Guess value must be less than or equal to 100")]
    fn greater_than_100() {
        Guess::new(200);
    }
}
```

## 11-2. Test 실행하기

- `cargo test`는 기본적으로 스레드를 이용해 병렬적으로 수행됩니다.
만일 병렬 테스트를 실행하고 싶지 않을 경우, 스레드 갯수를 그저 1개로 줄이면 됩니다.

```bash
$ cargo test -- --test-threads=1
```

cargo test에는 총 2가지 종류의 argument가 존재합니다.

```bash
> cargo test --help
Execute all unit and integration tests and build examples of a local package

Usage: cargo test [OPTIONS] [TESTNAME] [-- [args]...]
```

1. (컴파일 옵션) cargo test 커맨드라인에 테스트 옵션/파일이름을 위해 전달되는  argument
    1. OPTIONS
    2. TESTNAME
2. (컴파일후 생성된 바이너리 실행 옵션) cargo test 바이너리를 실행시 전달할 옵션입니다.


`cargo test`는 위의 2가지 argument를 구분하기 위해 `--`를 사용합니다.


### `--nocapture`
> `--nocapture`를 사용하면 stdout까지 print되는 것을 막을 수 있습니다.

```bash
$ cargo test -- --nocapture
```

### filtering
> cargo는 test 뒤에 나오는 네이밍을 regex로 필터링합니다.

만약 아래와 같은 코드가 있고, `$ cargo test add`를 실행하게 된다면

```rs
pub fn add_two(a: i32) -> i32 {
    a + 2
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn add_two_and_two() {
        assert_eq!(4, add_two(2));
    }

    #[test]
    fn add_three_and_two() {
        assert_eq!(5, add_two(3));
    }

    #[test]
    fn one_hundred() {
        assert_eq!(102, add_two(100));
    }
}
```

```bash
$ cargo test add
    Finished dev [unoptimized + debuginfo] target(s) in 0.0 secs
     Running target/debug/deps/adder-06a75b4a1f2515e9

running 2 tests
test tests::add_two_and_two ... ok
test tests::add_three_and_two ... ok

test result: ok. 2 passed; 0 failed; 0 ignored; 0 measured; 1 filtered out
```

다음과 같은 결과가 일어납니다.

### `ignore`
> 만약 특정 테스트들을 무시하고 싶으면 `#[ignore]`를 사용하면됩니다.

```rs
#[test]
fn it_works() {
    assert_eq!(2 + 2, 4);
}

#[test]
#[ignore]
fn expensive_test() {
    // code that takes an hour to run
}
```

역으로 ignore된 테스트들만 실행시키고 싶다면, --ignored를 추가하면 됩니다.

```
$ cargo test -- --ignored
```

## 11-3. 테스트 조직화

1. Unit test
2. Integration test

### `unit test`

관례는 각 파일마다 테스트 함수를 담고 있는 tests라는 이름의 모듈을 만들고, 이 모듈에 `cfg(test)`라고 어노테이션 하는 것입니다.

`cfg(test)`

- 이 어노테이션은 러스트에게 우리가 **cargo build를 실행시킬 때가 아니라 cargo test를 실행시킬 때에만 컴파일하고 실행**시키라고 말해줍니다.
- 통합 테스트는 다른 디렉토리에 위치하기 때문에, 이 어노테이션이 필요없습니다.

또한 rust의 `private`은 test에서는 접근가능 합니다.

```rs
pub fn add_two(a: i32) -> i32 {
    internal_adder(a, 2)
}

fn internal_adder(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn internal() {
        assert_eq!(4, internal_adder(2, 2));
    }
}
```

마지막으로 rust의 모듈 특성을 이용하면, test 파일을 소스 코드와 분리해서 관리할 수도 있습니다.

[Should unit tests really be put in the same file as the source?](https://users.rust-lang.org/t/should-unit-tests-really-be-put-in-the-same-file-as-the-source/62153)

```
src/
  ...
  ops.rs
  ops/
    test.rs
```



### `Integration test`

러스트의 통합테스트는 `tests/` 디렉토리로 완전히 `src/`와 분리되서 관리됩니다.

```bash
src/ 
  lib.rs
  ops.rs
  ops/
    test.rs

tests/
```


## 12. I/O 커맨드 라인 프로그램

[CLI repo](https://github.com/minkj1992/rust_playground/tree/main/cli)

## 13. Funtional Programming

- `Closure`
- `Iterator`


### 13.1 `Closure`
- simple define closure

```rs
fn  add_one_v1   (x: u32) -> u32 { x + 1 }
let add_one_v2 = |x: u32| -> u32 { x + 1 };
let add_one_v3 = |x|             { x + 1 };
let add_one_v4 = |x|               x + 1  ;

{
    let expensive_closure = |num| {
    println!("calculating slowly...");
    thread::sleep(Duration::from_secs(2));
    num
    };
}
```

클로저는 만약 좁은 범위에서 사용되기 때문에, 문맥상 타입이 명확하다면 오히려 불필요하게 타입을 쓰지 않아도 좋은 접근인 것 같다.

- 한번 호출 이후 타입은 고정됩니다.

```rs
let example_closure = |x| x;

let s = example_closure(String::from("hello"));
let n = example_closure(5);
```

```bash
error[E0308]: mismatched types
 --> src/main.rs
  |
  | let n = example_closure(5);
  |                         ^ expected struct `std::string::String`, found
  integral variable
  |
  = note: expected type `std::string::String`
             found type `{integer}`
```


- full example
```rs
#[derive(Debug, PartialEq, Copy, Clone)]
enum ShirtColor {
    Red,
    Blue,
}

struct Inventory {
    shirts: Vec<ShirtColor>,
}

impl Inventory {
    fn giveaway(&self, preference: Option<ShirtColor>) -> ShirtColor {
        // |no parameter| => self.most_stocked call closure
        preference.unwrap_or_else(|| self.most_stocked())
    }
    
    fn most_stocked(&self) -> ShirtColor {
        let (mut n_red, mut n_blue) = (0,0);
        
        for c in &self.shirts {
            match c {
                ShirtColor::Red => n_red +=1,
                ShirtColor::Blue => n_blue +=1,
            }
        }
        
        if n_red > n_blue {ShirtColor::Red} else {ShirtColor::Blue}
    }
}

fn main() {
    let store = Inventory {
        shirts: vec![ShirtColor::Blue, ShirtColor::Red, ShirtColor::Blue],
    };

    let user_pref1 = Some(ShirtColor::Red);
    let giveaway1 = store.giveaway(user_pref1);
    println!(
        "The user with preference {:?} gets {:?}",
        user_pref1, giveaway1
    );

    let user_pref2 = Some(ShirtColor::Red);
    let giveaway2 = store.giveaway(user_pref2);
    println!(
        "The user with preference {:?} gets {:?}",
        user_pref2, giveaway2
    );
}
```
#### `Memoization` (`lazy evaluation`)
> with `Fn trait`

각 클로저 인스턴스는 자신의 유일한 익명 타입을 갖습니다: 즉, 두 클로저가 동일한 타입 서명을 갖더라도 그들의 타입은 여전히 다른 것으로 간주 됩니다.


- 구조체 필드에 클로저를 구현한 경우. (`Fn trait`)
- impl 
```rs
pub struct Cacher<T>
where
    T: Fn(u32) -> u32,
{
    calculation: T,
    value: Option<u32>,
}

impl<T> Cacher<T>
where
    T: Fn(u32) -> u32,
{
    pub fn new(calculation: T) -> Cacher<T> {
        Cacher {
            calculation,
            value: None,
        }
    }
    pub fn value(&mut self, arg: u32) -> u32 {
        match self.value {
            Some(v) => v,
            None => {
                let v = (self.calculation)(arg);
                self.value = Some(v);
                v
            }
        }
    }
}

```

## `Iterator`

- 모든 iterator는 lazy하게 evaluation합니다. (python의 `range`와 같다.)


```rs
let v1 = vec![1, 2, 3];

let v1_iter = v1.iter();

for val in v1_iter {
    println!("Got: {}", val);
}
```

iterator는 내부적으로 next()를 사용하여 item들을 참조합니다. 또한 next는 `&mut self`로 참조하는데, 이를 통해 next가 호출 될 때마다, item들이 소비됩니다.


- 모든 반복자는 표준 라이브러리에 정의된 Iterator 라는 이름의 트레잇을 구현 합니 다. 트레잇의 정의는 아래와 같습니다.

```rs
trait Iterator {
    type Item;

    fn next(&mut self) -> Option<Self::Item>;

    // methods with default implementations elided
}

...

#[test]
fn iterator_demonstration() {
    let v1 = vec![1, 2, 3];

    let mut v1_iter = v1.iter();

    assert_eq!(v1_iter.next(), Some(&1));
    assert_eq!(v1_iter.next(), Some(&2));
    assert_eq!(v1_iter.next(), Some(&3));
    assert_eq!(v1_iter.next(), None);
}
```

- next 호출로 얻어온 값들은 벡터 안에 있는 값들에 대한 불변 참조라는 점 역시 유의 하세요. 
- `iter()` 불변 참조에 대한 반복자를 만듭니다. 
- 만약 v1 의 소유권을 갖고 소유된 값들을 반환하도록 하고 싶다면, `iter` 대신 `into_iter` 를 호출해야 합니다. 비슷하게, 가변 참조에 대한 반복자를 원한다면, iter 대신 `iter_mut` 을 호출할 수 있습니다.


### 반복자를 소비하는 메서드들

- sum 또한 item을 소비합니다.

```rs
#[test]
fn iterator_sum() {
    let v1 = vec![1, 2, 3];

    let v1_iter = v1.iter();

    let total: i32 = v1_iter.sum();

    assert_eq!(total, 6);
}
```
sum 은 호출한 반복자의 소유권을 갖기 때문에, sum 을 호출한 후 v1_iter 은 사용할 수 없습니다.

### 다른 반복자를 생성하는 메서드들

- `map()`

```rs
let v1: Vec<i32> = vec![1, 2, 3];

v1.iter().map(|x| x + 1); // 새로운 iterator를 생성
```

iterator는 lazy하기 때문에, consume 되기전까지는 evaluate되지 않습니다.
그렇기 때문에 위에 코드는 아래와 같은 경고를 만들게 되는데요.

```
warning: unused `std::iter::Map` which must be used: iterator adaptors are lazy
and do nothing unless consumed
 --> src/main.rs:4:5
  |
4 |     v1.iter().map(|x| x + 1);
  |     ^^^^^^^^^^^^^^^^^^^^^^^^^
  |
  = note: #[warn(unused_must_use)] on by default
```

이를 해결하기 위해서는 iterator를 소비해주면 됩니다. 


- `collect()`

```rs
let v1: Vec<i32> = vec![1, 2, 3];

let v2: Vec<_> = v1.iter().map(|x| x + 1).collect();

assert_eq!(v2, vec![2, 3, 4]);
```

- `filter`

```rs
#[derive(PartialEq, Debug)]
struct Shoe {
    size: u32,
    style: String,
}

fn shoes_in_my_size(shoes: Vec<Shoe>, shoe_size: u32) -> Vec<Shoe> {
    shoes.into_iter()
        .filter(|s| s.size == shoe_size)
        .collect()
}

#[test]
fn filters_by_size() {
    let shoes = vec![
        Shoe { size: 10, style: String::from("sneaker") },
        Shoe { size: 13, style: String::from("sandal") },
        Shoe { size: 10, style: String::from("boot") },
    ];

    let in_my_size = shoes_in_my_size(shoes, 10);

    assert_eq!(
        in_my_size,
        vec![
            Shoe { size: 10, style: String::from("sneaker") },
            Shoe { size: 10, style: String::from("boot") },
        ]
    );
}
```

### 성능 비교하기: 루프 vs. 반복자

[performance compare](https://rinthel.github.io/rust-lang-book-ko/ch13-04-performance.html)

오히려 iterator가 loop보다 빠르게 측정된다. 빠르다는 것이 중요한 것은 아니고
비록 iterator가 고수준의 abstract임에도, 컴파일이 진행되면 low level 코드와 같은 수준까지 내려갑니다. 이를 `zero cost abstraction`라고 러스트에서는 부릅니다.

즉 iterator와 closure 코드는 고수준이지만, 컴파일러의 zero cost abstraction 덕분에 런타임 성능 걱정없이 사용할 수 있습니다.

# 15. 스마트 포인터

- 러스트에서 스마트 포인터는 보통 구조체를 이용해서 구현된 기능이 추가된 포인터입니다.
- `String`, `Vec<T>` 또한 스마트 포인터의 일종입니다. 이유는 이들이 얼마간의 메모리를 소유하고, 개발자들이 다루도록 허용하기 때문입니다. 또한 메타데이터와, 추가 능력(확장) 기능을 가지고 있습니다.

{{< admonition note "Smart pointer" >}}
_In computer science, a smart pointer is an abstract data type that simulates a pointer while providing added features, such as automatic memory management or bounds checking. Such features are intended to reduce bugs caused by the misuse of pointers, while retaining efficiency._

_Smart pointers typically keep track of the memory they point to, and may also be used to manage other resources, such as network connections and file handles. Smart pointers were first popularized in the programming language C++ during the first half of the 1990s as rebuttal to criticisms of C++'s lack of automatic garbage collection._
{{< /admonition  >}}


스마트 포인터가 일반적인 구조체와 구분되는 특성은 바로 `Deref`, `Drop` 트레잇을 구현한다는 것입니다.

- `Deref`: 스마트 포인터 구조체의 인스턴스가 참조자처럼 동작하도록 하여 참조자나 스마트 포인터 둘 중 하나와 함께 작동하는 코드를 작성하게 해줍니다.
- `Drop`: 스마트 포인터의 인스턴스가 스코프 밖으로 벗어났을 때, 실행되는 코드

표준 라이브러리에는 가장 대표적으로 아래의 스마트 포인터들이 있습니다.

- `Box<T>`: 값을 `힙`에 할당
- `Rc<T>`: Reference Counting 타입, 복수개의 소유권 가능하도록 함.
- 빌림 규칙을 컴파일 타임 대신 런타임에 강제하는 타입인, `RefCell<T>`를 통해 접근 가능한 `Ref<T>`와 `RefMut<T>`


## `Box<T>`
> 데이터를 스택이 아닌 힙에 저장하도록 합니다.
>
> 스택 대신 힙에 저장한다는 점 외에는, 성능적인 오버헤드는 없습니다. (stack vs heap 자료구조에 상에서 성능 오버헤드를 뜻하는 듯)

아래 3가지 경우에 자주 사용합니다.

1. 컴파일 타임에 크기를 알 수 없는 타입을 갖고 있고, 정확한 사이즈를 알 필요가 있는 맥락 안에서 해당 타입의 값을 이용하고 싶을 때
2. 커다란 데이터를 가지고 있고 소유권을 옮기고 싶지만 그렇게 했을 때 데이터가 복사되지 않을 것이라고 보장하기를 원할 때
3. 어떤 값을 소유하고 이 값의 구체화된 타입을 알고 있기보다는 특정 트레잇을 구현한 타입이라는 점만 신경 쓰고 싶을 때 (`trait object` 17장)

이번 장에서는 1번 상황을 설명합니다. 2번의 경우는 그저 stack에 여러개 올리기 부담스러울 정도로 큰 데이터 또는 copy가 일어날지 불명확할 떄 heap에 저장시킨다는 뜻입니다. 

```rs
{
    let bx = Box::new(5);
}
```

## Recursive type

컴파일 타임에서, 러스트는 어떤 타입이 얼마나 많은 공간을 차지하는지를 알 필요가 있습니다. 컴파일 타임에는 크기를 알 수 없는 한 가지 타입이 바로 재귀적 타입 (recursive type)입니다.

하지만 이때 재귀적 타입 정의 안에 Box를 사용하면 가능합니다.

- without `Box`, 컴파일 에러

```rs
enum List {
    Cons(i32, List),
    Nil,
}

use List::{Cons, Nil};

fn main() {
    let list = Cons(1, Cons(2, Cons(3, Nil)));
}
```

- with `Box`, 성공.

```rs
enum List {
    Cons(i32, Box<List>),
    Nil,
}

use List::{Cons, Nil};

fn main() {
    let list = Cons(1,
        Box::new(Cons(2,
            Box::new(Cons(3,
                Box::new(Nil))))));
}
```

<center>

![](/images/rust_box.svg)

</center>


## `Deref trait`

`Deref` 트레잇을 구현한다는 것은, `dereference` operator(역참조 연산자) 즉 `*`의 동작을 커스터마이징 하는 것을 허용합니다.

우선 다음은 일반적인 역참조 연산자 입니다.

```rs
fn main() {
    let mut x = 5;
    let y = &x;
    assert_eq!(5, x);
    assert_eq!(5, *y); // 역참조, deref
}
```

### `Box<T>`를 참조자처럼 사용하기

Box는 `Deref trait`를 구현하고 있기 때문에, 다음과 같이 동작할 수 있습니다.

```rs
fn main() {
    let x = 5;
    let y = Box::new(x);
    assert_eq!(5, x);
    assert_eq!(5, *y); // 역참조
}
```

### 커스텀 Box 타입
> deref를 지원하는 커스텀 box 타입

```rs
use std::ops::Deref;

struct MyBox<T>(T);

impl<T> MyBox<T> {
    fn new(x: T) -> MyBox<T> {
        MyBox(x)
    }
}

impl<T> Deref for MyBox<T> {
    type Target = T;

    fn deref(&self) -> &T {
        dbg!("deref is called");
        &self.0 // 도대체 self.0이 뭘가르키는 거지....
    }
}


fn main() {
    let x = 5;
    let y = MyBox::new(x);
    
    dbg!(*y); 
    // "deref is called"
    // *y = 5
}
```
러스트는 `*y`를 뒤에서 다음과 같이 호출합니다.

```rs
*(y.deref());
```

### deref coercion (역참조 강제, 암묵적 역참조)


> 역참조 강제란(암묵적 역참조) 우리가 특정 타입의 값에 대한 참조자를 함수 혹은 메소드의 인자로 넘기는 중 정의된 파라미터 타입에는 맞지 않을 때 자동적으로 발생합니다.

Deref 트레잇을 구현한 타입은 컴파일러가 `암묵적 역참조`를 자동으로 처리해줍니다.

- 암묵적 역참조
```rs
fn hello(name: &str) {
    println!("Hello, {}!", name);
}

fn main() {
    let m = MyBox::new(String::from("Rust"));
    hello(&m);
}
```

- 만약 암묵적 역참조가 없다면

```rs
fn main() {
    let m = MyBox::new(String::from("Rust"));
    hello(&(*m)[..]);
}
```

Deref 트레잇이 관련된 타입에 대해 정의될 때, 러스트는 해당 타입을 분석하여 파라미터의 타입에 맞는 참조자를 얻기 위해 필요한 수만큼의 Deref::deref를 사용할 것입니다. Deref::deref가 삽입될 필요가 있는 횟수는 컴파일 타임에 분석되므로, 역참조 강제의 이점을 얻는 데에 관해서 어떠한 런타임 페널티도 없습니다!

### `Mutable reference`의 암묵적 역참조(deref coercion)

불변 참조자에 대한 `*`를 오버 라이딩하기 위해 `Deref 트레잇`을 이용하는 방법과 비슷하게, 러스트는 가변 참조자에 대한 `*`를 오버 라이딩하기 위한 `DerefMut 트레잇`을 제공합니다.

러스트 컴파일러는 다음 3가지 경우에 해당 하는 타입을 만나면 역참조 강제를 수행합니다.

1. `T: Deref<Target=U`>일때 `&T`에서 `&U`로
2. `T: DerefMut<Target=U>`일때 `&mut T`에서 `&mut U`로
3. `T: Deref<Target=U>`일때 `&mut T`에서 `&U`로


1번은 앞서 보았던 역참조 강제이며, 2번은 mut reference에서도 역참조 강제가 일어난다는 것을 뜻합니다.

마지막 세 번째 경우는 좀 더 교묘합니다: 

러스트는 가변 참조자를 불변 참조자로 강제할 수도 있습니다. 하지만 그 역은 불가능합니다: 불변 참조자는 가변 참조자로 결코 강제되지 않을 것입니다. 빌림 규칙 때문에, 만일 여러분이 가변 참조자를 가지고 있다면, 그 가변 참조자는 해당 데이터에 대한 유일한 참조자임에 틀림없습니다 (만일 그렇지 않다면, 그 프로그램은 컴파일되지 않을 것입니다). 가변 참조자를 불변 참조자로 변경하는 것은 결코 빌림 규칙을 깨트리지 않을 것입니다. 불변 참조자를 가변 참조자로 변경하는 것은 해당 데이터에 대한 단 하나의 불변 참조자가 있어야 한다는 요구를 하게 되고, 이는 빌림 규칙이 보장해줄 수 없습니다. 따라서, 러스트는 불변 참조자를 가변 참조자로 변경하는 것이 가능하다는 가정을 할 수 없습니다.

## `Drop trait`
## `Rc<T>`와 `레퍼런스 카운팅` 스마트 포인터
## `RefCell<T`와 `내부 가변성` 패턴
## reference cycle (순환참조)
