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


# 12. I/O 커맨드 라인 프로그램

[CLI repo](https://github.com/minkj1992/rust_playground/tree/main/cli)

# 13. Funtional Programming

- `Closure`
- `Iterator`


## 13.1 `Closure`

Closures are functions that can capture the enclosing environment. 

- `|val| val + x`

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
### `Memoization` (`lazy evaluation`)
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

### `Iterator`

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


## 15.1 `Box<T>`
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

### Recursive type

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


## 15.2 `Deref trait`

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

## 15.3 `Drop trait`

```rs
struct CustomSmartPointer {
    heap_data: String,
}

impl Drop for CustomSmartPointer {
    fn drop(&mut self) {
        println!("Drop CustomSmartPointer `{}`", self.heap_data);
    }
}

fn main() {
    let a = CustomSmartPointer {
        heap_data: String::from("leoo"),
    };
    let b = CustomSmartPointer {
        heap_data: String::from("wants to learn more loves."),
    };

    println!("CustomSmartPointers are created.");
}

// CustomSmartPointers are created.
// Drop CustomSmartPointer `wants to learn more loves.`
// Drop CustomSmartPointer `leoo`
```

- Drop 트레잇은 `prelude`에 포함되어 있으므로, 이를 가져오지 않아도 됩니다.
- `drop 함수`의 본체는 여러분이 만든 타입의 인스턴스가 스코프 밖으로 벗어났을 때 실행시키고자 하는 어떠한 로직이라도 위치시킬 수 있는 곳입니다.
- drop은 stack에 스택에 따라 처리되기 때문에, 최근에 선언된 스마트포인터일 수록 더 먼저 처리됩니다.


일반적이지는 않지만 아주 가끔, 여러분은 **값을 일찍 정리하기를 원할 지도 모릅니다**. 한 가지 예는 락을 관리하는 스마트 포인터를 이용할 때입니다.

단 러스트는 default로 `.drop()` 메서드(소멸자, `destructor`)가 호출되는 것을 허용하지 않습니다. 만약 명시적으로 drop을 일찍 시켜주고 싶다면 `std::mem::drop` 함수를 이용할 수 있습니다.

std::mem::drop 함수는 Drop 트레잇 내에 있는 drop 메소드와 다릅니다. 우리가 일찍 버리도록 강제하길 원하는 값을 인자로 넘김으로써 이를 호출할 수 있습니다. 이 함수는 프렐루드에 포함되어 있습니다. 

```rs
fn main() {
    let c = CustomSmartPointer { data: String::from("some data") };
    println!("CustomSmartPointer created.");

    // method가 아닌, 함수로 처리
    drop(c);
    println!("CustomSmartPointer dropped before the end of main.");
}

// CustomSmartPointer created.
// Dropping CustomSmartPointer with data `some data`!
// CustomSmartPointer dropped before the end of main.
```

## 15.4 `Rc<T>`와 `레퍼런스 카운팅` 스마트 포인터

대부분의 경우에서, 소유권은 명확합니다: 여러분은 어떤 변수가 주어진 값을 소유하는지 정확히 압니다. 그러나, 하나의 값이 여러 개의 소유자를 가질 수도 있는 경우가 있습니다. 

예를 들면, 그래프 데이터 구조에서, 여러 에지가 동일한 노드를 가리킬 수도 있고, 그 노드는 개념적으로 해당 노드를 가리키는 모든 에지들에 의해 소유됩니다. 노드는 어떠한 에지도 이를 가리키지 않을 때까지는 메모리 정리가 되어서는 안됩니다.

- 복수 소유권을 가능하게 하기 위해서, 러스트는 `Rc<T>`라 불리우는 타입을 가지고 있습니다. 
- 이 이름은 참조 카운팅 (`reference counting`) 의 약자입니다.
- 이는 어떤 값이 계속 사용되는지 혹은 그렇지 않은지를 알기 위해 해당 값에 대한 참조자의 갯수를 계속 추적하는 것입니다.

- before: `RC<T>`없이 소유권을 나눠가질 때: 컴파일 에러

```rs
enum List {
    Cons(i32, Box<List>),
    Nil,
}

use List::{Cons, Nil};

fn main() {
    let a = Cons(5,
        Box::new(Cons(10,
            Box::new(Nil))));
    let b = Cons(3, Box::new(a));
    let c = Cons(4, Box::new(a));
}
```
```
error[E0382]: use of moved value: `a`
  --> src/main.rs:13:30
   |
12 |     let b = Cons(3, Box::new(a));
   |                              - value moved here
13 |     let c = Cons(4, Box::new(a));
   |                              ^ value used here after move
   |
   = note: move occurs because `a` has type `List`, which does not implement
   the `Copy` trait
```

Cons variant는 이것이 가지고 있는 데이터를 소유하므로, 우리가 b리스트를 만들때, a는 b 안으로 이동되고 b는 a를 소유합니다. 그 뒤, c를 생성할 때 a를 다시 이용하는 시도를 할 경우, 이는 a가 이동되었으므로 허용되지 않습니다.

우리는 Cons가 대신 참조자를 갖도록 정의를 변경할 수도 있지만, 그러면 라이프타임 파라미터를 명시해야 할 것입니다. 라이프타임 파라미터를 명시함으로써, 리스트 내의 모든 요소들이 최소한 전체 리스트만큼 오래 살아있도록 명시될 것입니다. 빌림 검사기는 예를 들면 `let a = Cons(10, &Nil);`을 컴파일되도록 하지 않게 할텐데, 이는 일시적인 Nil 값은 a가 그에 대한 참조자를 가질 수도 있는 시점 이전에 버려질 것이기 때문입니다.

- after: `RC<T>`를 사용.

```rs
enum List {
    Cons(i32, Rc<List>),
    Nil,
}

use List::{Cons, Nil};
use std::rc::Rc;

fn main() {
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    let b = Cons(3, Rc::clone(&a));
    let c = Cons(4, Rc::clone(&a));
}
```

- Rc<T>는 프렐루드에 포함되어 있지 않으므로 우리는 이를 가져오기 위해 `use std::rc::Rc`가 필요합니다.

`Rc::clone(&a)` 보다는 `a.clone()`을 호출할 수도 있지만, 위의 경우 러스트의 관례는 `Rc::clone`를 이용하는 것입니다.

`Rc::clone`의 구현체는 대부분의 타입들의 `clone` 구현체들이 하는 것처럼 모든 데이터의 깊은 복사 (deep copy) 를 만들지 않습니다. **`Rc::clone`의 호출은 오직 참조 카운트만 증가 시키는데, 이는 큰 시간이 들지 않습니다.** 

이를 통해서 코드 내에서 성능 문제가 있어 문제가 될 부분들을 찾고 있다면, 깊은 복사 클론만 고려할 필요가 있고 Rc::clone 호출은 무시할 수 있습니다.

- 레퍼런스 카운트 출력하기
```rs
fn main() {
    let a = Rc::new(Cons(5, Rc::new(Cons(10, Rc::new(Nil)))));
    println!("count after creating a = {}", Rc::strong_count(&a));
    let b = Cons(3, Rc::clone(&a));
    println!("count after creating b = {}", Rc::strong_count(&a));
    {
        let c = Cons(4, Rc::clone(&a));
        println!("count after creating c = {}", Rc::strong_count(&a));
    }
    println!("count after c goes out of scope = {}", Rc::strong_count(&a));
}
```

```
count after creating a = 1
count after creating b = 2
count after creating c = 3
count after c goes out of scope = 2
```

## 15.5 `RefCell<T>`와 `내부 가변성` 패턴

`interior mutability`(내부 가변성)이란 어떤 데이터와 관련된 immutable reference가 있더라도, 여러분이 데이터를 변형할 수 있게 해주는 러스트의 디자인 패턴입니다. 보통 borrow rule에 의해서 이는 허용되지 않지만, `unsafe`코드를 사용하여 이를 우회할 수 있습니다.

만약 우리가 런타임에 borrow rule을 따릇 것이라는 것을 보장할 수 있다면, 컴파일러가 이를 보장하지 못하더라도 내부 가변성 패턴을 이용하는 타입을 사용할 수 있습니다.

- `unsafe 코드`는 안전한 API로 감싸져 있고, 외부에서는 여전히 불변하게 동작합니다.

`RefCell<T>`는 대표적으로 내부 가변성(`interior mutability`)를 따르는 타입입니다.


`Box<T>`, `Rc<T>`, 혹은 `RefCell<T>`을 선택하는 이유의 요점은 다음과 같습니다:

- `Rc<T>`는 동일한 데이터에 대해 복수개의 소유자를 가능하게 합니다; `Box<T>`와 `RefCell<T>`은 단일 소유자만 갖습니다.
- `Box<T>`는 컴파일 타임에 검사된 불변 혹은 가변 빌림을 허용합니다; `Rc<T>`는 오직 컴파일 타임에 검사된 불변 빌림만 허용합니다; `RefCell<T>`는 런타임에 검사된 불변 혹은 가변 빌림을 허용합니다.
- `RefCell<T>`이 런타임에 검사된 가변 빌림을 허용하기 때문에, `RefCell<T>`이 불변일 때라도 `RefCell<T>` 내부의 값을 변경할 수 있습니다.


- `RefCell<T>` 예시

```rs
pub trait Messenger {
    fn send(&self, msg: &str);
}

pub struct LimitTracker<'a, T: 'a + Messenger> {
    messenger: &'a T,
    value: usize,
    max: usize,
}

impl<'a, T> LimitTracker<'a, T>
    where T: Messenger {
    pub fn new(messenger: &T, max: usize) -> LimitTracker<T> {
        LimitTracker {
            messenger,
            value: 0,
            max,
        }
    }

    pub fn set_value(&mut self, value: usize) {
        self.value = value;

        let percentage_of_max = self.value as f64 / self.max as f64;

        if percentage_of_max >= 0.75 && percentage_of_max < 0.9 {
            self.messenger.send("Warning: You've used up over 75% of your quota!");
        } else if percentage_of_max >= 0.9 && percentage_of_max < 1.0 {
            self.messenger.send("Urgent warning: You've used up over 90% of your quota!");
        } else if percentage_of_max >= 1.0 {
            self.messenger.send("Error: You are over your quota!");
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use std::cell::RefCell;

    struct MockMessenger {
        sent_messages: RefCell<Vec<String>>, // 여기에서 사용되었다.
    }

    impl MockMessenger {
        fn new() -> MockMessenger {
            MockMessenger { sent_messages: RefCell::new(vec![]) }
        }
    }

    impl Messenger for MockMessenger {
        fn send(&self, message: &str) {
            self.sent_messages.borrow_mut().push(String::from(message));
        }
    }

    #[test]
    fn it_sends_an_over_75_percent_warning_message() {
        // --snip--

        assert_eq!(mock_messenger.sent_messages.borrow().len(), 1);
    }
}
```

### Rc<T>와 RefCell<T>를 조합하여 가변 데이터의 복수 소유자 만들기

**RefCell<T>를 사용하는 일반적인 방법은 Rc<T>와 함께 조합하는 것입니다.** 

- Rc<T>이 어떤 데이터에 대해 복수의 소유자를 허용하지만, 그 데이터에 대한 불변 접근만 제공하는 것을 상기하세요. 
- 만일 우리가 RefCell<T>을 들고 있는 Rc<T>를 갖는다면, 우리가 변경 가능하면서 복수의 소유자를 갖는 값을 가질 수 있습니다.

우리가 어떤 리스트의 소유권을 공유하는 여러 개의 리스트를 가질 수 있도록 하기 위해 Rc<T>를 사용했던 cons 리스트 예제를 상기해보면, `Rc<T>`가 오직 불변의 값만을 가질 수 있기 때문에, 우리가 이들을 일단 만들면 리스트 안의 값들을 변경하는 것은 불가능했습니다. 

이 리스트 안의 값을 변경하는 능력을 얻기 위해서 `RefCell<T>`을 추가해 봅시다.

```rs
#[derive(Debug)]
enum List {
    Cons(Rc<RefCell<i32>>, Rc<List>),
    Nil,
}

use List::{Cons, Nil};
use std::rc::Rc;
use std::cell::RefCell;

fn main() {
    let value = Rc::new(RefCell::new(5));

    let a = Rc::new(Cons(Rc::clone(&value), Rc::new(Nil)));

    let b = Cons(Rc::new(RefCell::new(6)), Rc::clone(&a));
    let c = Cons(Rc::new(RefCell::new(10)), Rc::clone(&a));

    *value.borrow_mut() += 10;

    println!("a after = {:?}", a);
    println!("b after = {:?}", b);
    println!("c after = {:?}", c);
}
```
Cons 정의 내에 RefCell<T>를 사용함으로써 우리가 모든 리스트 내에 저장된 값을 변경할 수 있음을 보여줍니다.

```
a after = Cons(RefCell { value: 15 }, Nil)
b after = Cons(RefCell { value: 6 }, Cons(RefCell { value: 15 }, Nil))
c after = Cons(RefCell { value: 10 }, Cons(RefCell { value: 15 }, Nil))
```

이 기술은 매우 깔끔합니다! RefCell<T>을 이용함으로써, 우리는 표면상으로는 불변인 List를 갖고 있습니다. 하지만 우리는 내부 가변성 접근을 제공하여 우리가 원할때 데이터를 변경시킬 수 있는 RefCell<T> 내의 메소드를 사용할 수 있습니다. 빌림 규칙의 런타임 검사는 데이터 레이스로부터 우리를 지켜주고, 우리 데이터 구조의 이러한 유연성을 위해서 약간의 속도를 트레이드 오프 하는 것이 때때로 가치있습니다.

표준 라이브러리는 내부 가변성을 제공하는 다른 타입을 가지고 있는데, 이를 테면 Cell<T>는 내부 값의 참조자를 주는 대신 값이 복사되어 Cell<T> 밖으로 나오는 점만 제외하면 비슷합니다. 또한 Mutex<T>가 있는데, 이는 스레드들을 건너가며 사용해도 안전한 내부 가변성을 제공합니다.

## 15.5 reference cycle (순환참조)

러스트의 memory safety(memory leak 안정보장)은 뜻하지 않게 해제 되지 않는 메모리 생성을 힘들게 하지만, `Rc<T>`, `RefCell<T>` 처럼 메모리 릭을 허용하는 것이 있다는 것을 알 수 있습니다. 

즉 아이템들 끼리 서로를 순환 참조하는 참조자를 만드는 것이 가능합니다. 이로 인해 메모리릭이 발생되는데, 서로 참조하는 cycle에서 reference count는 결코 0이 되지 않을 것이고, 그렇게 되면 해당 값들은 버려지지 않게 됩니다.

- 순환 참조가 발생하는 코드

```rs
use std::cell::RefCell;
use std::rc::Rc;
use List::{Cons, Nil};

#[derive(Debug)]
enum List {
    Cons(i32, RefCell<Rc<List>>),
    Nil,
}

impl List {
    fn tail(&self) -> Option<&RefCell<Rc<List>>> {
        match *self {
            Cons(_, ref item) => Some(item),
            Nil => None,
        }
    }
}

fn custom_print(name: &str, l: &Rc<List>) {
    println!("### About {name} ###");
    println!("rc count = {}", Rc::strong_count(&l));
    println!("next block = {:?}", l.tail());
}

fn main() {
    let a = Rc::new(Cons(5, RefCell::new(Rc::new(Nil))));
    custom_print("a", &a);

    let b = Rc::new(Cons(10, RefCell::new(Rc::clone(&a))));
    custom_print("a", &a);
    custom_print("b", &b);

    // reference cycle point (memory leak);
    // a -> b;
    if let Some(link) = a.tail() {
        *link.borrow_mut() = Rc::clone(&b);
    }
    // thread 'main' has overflowed its stack
    // fatal runtime error: stack overflow
    custom_print("a", &a);
    custom_print("b", &b);
}
```

[5, Nil]를 가진 리스트 a를 만든 뒤, [10, a]를 가진 리스트 b를 생성하였습니다.
이후 `borrow_mut()`를 사용해 a의 tail인 Nil이 b로 변경되도록 하였습니다.

- `a`: [5, b]
- `b`: [10, a]

<center>

![](/images/rust_reference_cycle.svg)

</center>

순환참조에 의한 메모리릭 문제는 러스트 컴파일러에 기대어서는 안되며 테스트, 코드리뷰 등으로 파악해야 하는 논리적인 에러입니다.

순환참조를 피하는 다른 해결책으로는 각 인스턴스들의 `그래프`에 따라서 소유권이 필요한 노드와 그렇지 않은 노드를 파악해서 끊어주면 됩니다.

### 순환 참조 방지하기
> `Rc<T>`를 `Weak<T>`로 변경하기

- `Rc::downgrade()`: 참조자(reference)를 weak reference(`Weak<T>`로 변경시킵니다.

`Weak<T>` 타입의 스마트 포인터는 `clone()`이 호출될 시, Rc<T>인스턴스의 strong_count를 +=1 시키는 것이 아니라, `weak_count`라는 필드를 1 증가시킵니다.

weak_count와 strong_count의 차이점 즉, Weak<T>과 Rc<T>의 차이점은

- `Rc<T>`가 제거되기(`free`) 위해서는 `strong_count == 0`이어야 하지만
- `Weak<T>`는 제거되기 위해서 `weak_count`가 0일 필요가 없습니다.

`강한 참조`는 여러분이 `Rc<T>` 인스턴스의 소유권을 공유할 수 있는 방법입니다. **약한 참조는 소유권 관계를 표현하지 않습니다.** (`Weak<T>`)


그렇기 때문에 `Weak<T>`가 참조하고 있는 값은 이미 버려져 있을지도 모릅니다. `Weak<T>`가 가리키고 있는 값을 가지고 어떤 일을 하기 위해서는, 반드시 그 전에 참조하고 있는 값의 존재여부를 확인해야 합니다.

이를 위해서 `Weak<T>`의 `.upgrade()` 메소드를 호출합니다.
이 메소드는 `Option<Rc<T>>`를 반환할 것이고, `Some`이 return된다면 값이 있다는 것이며, None의 경우네는 값이 free된 경우입니다.

아래는 Weak<T>를 활용해 `트리 데이터: 자식 노드를 가진 Node`를 만들어 보겠습니다.

```rs
use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
struct Node {
    value: i32,
    children: RefCell<Vec<Rc<Node>>>,
}

fn main() {
    let leaf = Rc::new(Node {
        value: 3,
        children: RefCell::new(vec![]),
    });

    let branch = Rc::new(Node {
        value: 5,
        children: RefCell::new(vec![Rc::clone(&leaf)]),
    });
}
```
- `dbg!(branch);` 결과값
```

[src/main.rs:21] branch = Node {
    value: 5,
    children: RefCell {
        value: [
            Node {
                value: 3,
                children: RefCell {
                    value: [],
                },
            },
        ],
    },
}
```

이를 통해 branch는 leaf에 접근가능하게 되었으며, leaf는 2개의 strong reference count를 가지게 되었습니다. 다음으로 leaf가 branch에 접근가능하도록 코드를 수정 해보겠습니다.


```rs
use std::cell::RefCell;
use std::rc::{Rc, Weak};

#[derive(Debug)]
struct Node {
    value: i32,
    parent: RefCell<Weak<Node>>,
    children: RefCell<Vec<Rc<Node>>>,
}

fn generate_empty_node() -> RefCell<Weak<Node>> {
    RefCell::new(Weak::new())
}

fn main() {
    let leaf = Rc::new(Node {
        value: 3,
        parent: generate_empty_node(),
        children: RefCell::new(vec![]),
    });

    // leaf parent = None
    println!("leaf parent = {:?}", leaf.parent.borrow().upgrade());

    let branch = Rc::new(Node {
        value: 5,
        parent: generate_empty_node(),
        children: RefCell::new(vec![Rc::clone(&leaf)]),
    });

    // link to parent
    *leaf.parent.borrow_mut() = Rc::downgrade(&branch);
    
    // leaf parent = Some(Node { value: 5, parent: RefCell { value: (Weak) }, <- branch 자신
    // children: RefCell { value: [Node { value: 3, parent: RefCell { value: (Weak) }, <- leaf
    // children: RefCell { value: [] } }] } }) <- leaf의 children
    println!("leaf parent = {:?}", leaf.parent.borrow().upgrade());
}
```

leaf 노드가 branch를 parent로 가리키도록 하였습니다. 다음으로 아래 custom print함수를 사용해, weak count와 strong count 갯수를 확인해보겠습니다.


```rs
fn print_node(node: &Rc<Node>) {
    println!(
        "value = {} ,strong = {}, weak = {}",
        node.value,
        Rc::strong_count(node),
        Rc::weak_count(node),
    );
}
```

```rs
fn main() {
    let leaf = Rc::new(Node {
        value: 3,
        parent: generate_empty_node(),
        children: RefCell::new(vec![]),
    });
    print_node(&leaf); // value = 3 ,strong = 1, weak = 0

    {
        let branch = Rc::new(Node {
            value: 5,
            parent: generate_empty_node(),
            children: RefCell::new(vec![Rc::clone(&leaf)]),
        });

        // link to parent
        *leaf.parent.borrow_mut() = Rc::downgrade(&branch);
        print_node(&branch); // value = 5 ,strong = 1, weak = 1
        print_node(&leaf); // value = 3 ,strong = 2, weak = 0
    }
    
    // leaf parent = None
    println!("leaf parent = {:?}", leaf.parent.borrow().upgrade());
    print_node(&leaf); // value = 3 ,strong = 1, weak = 0
}
```

leaf 노드는 내부 scope에 정의된 `branch`를 parent로 하였고, 이에대한 참조를 weak reference로 해두었기 때문에 `branch`는 scope를 벗어날 때 free됩니다. (strong count가 0). 

이때 branch는 `weak count = 1`인 leaf의 weak reference를 가지고 있지만 이는 free에는 어떤 영향도 주지 않습니다. 이를 통해 어떠한 메모리 릭도 발생되지 않습니다.

`leaf parent = None`을 보시면, 스코프 끝 이후에 leaf의 부모에 접근을 시도하였기 때문에 `None`이 반환됩니다.

참조 카운트들과 버리는 값들을 관리하는 모든 로직은 Rc<T>와 Weak<T>, 그리고 이들의 Drop 트레잇에 대한 구현부에 만들어져 있습니다. 자식으로부터 부모로의 관계가 Node의 정의 내에서 Weak<T> 참조자로 되어야 함을 특정함으로서, 여러분은 순환 참조와 메모리 릭을 만들지 않고도 자식 노드를 가리키는 부모 노드 혹은 그 반대의 것을 가지게 될 수 있습니다.

# 16. 동시성

- 들어가기 앞서 이번장에서는 `동시성`과 `병렬성`을 구분하지 않고 모두 `동시성`이라고 칭합니다.
- 또한 `런타임`의 범위를 프로그래밍 언어의 모든 바이너리 내에 포함되는 코들르 의미합니다.

대표적으로 스레드는 코드 snippet에 대해 실행 순서를 보장하지 않기 때문에, 발생하는 문제점들은 다음과 같습니다. (3)

1. race condition
2. deadlock
3. 특정한 상황에서만 발생되어 재현하기와 안정적으로 수정하기가 힘든 버그들

1:1 스레드라는 것은 프로그래밍 언어에서 운영체제 API가 제공하는 스레드와 1:1로 상응하는 스레드를 의미합니다.

반면 `green thread`의 경우에는 운영체제 스레드와 M:N관계를 가집니다.
그린 스레드 M:N 구조는 자체 스레드들을 관리하기 위해 더 `큰 언어 런타임`이 필요하게 됩니다. 이런 트레이드 오프 때문에 러스트의 `std 라이브러리`는 오직 1:1 스레드 구현만 제공합니다. 이런 트레이드 오프(오버헤드)를 감수하더라도 `context switching`에 더 저렴한 cost를 원한다면 M:N 스레드를 구현한 `crate`들도 존재합니다.

## `thread::spawn()`
> 새로운 스레드 생성하기

```rs
{
    thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1));
        }
    });
}
```

## `.join().unwrap()`
> join 핸들을 사용하여 모든 스레드 끝날때까지 기다리기

개의 경우 메인 스레드가 종료되는 이유로 생성된 스레드가 조기에 멈출 뿐만 아니라, 생성된 스레드가 모든 코드를 실행할 것임을 보장해 줄수도 없습니다. 그 이유는 스레드들이 실행되는 순서에 대한 보장이 없기 때문입니다.

이를 해결하기 위해서는 .join()을 사용하면 됩니다. thread::spawn()은 `JoinHandle`을 리턴하며, 이를 변수에 담아 .join() 메서드를 호출시키면 스레드가 끝날때까지 기다릴 수 있습니다.

```rs
use std::thread;
use std::time::Duration;

fn main() {
    let handle = thread::spawn(|| {
        for i in 1..10 {
            println!("hi number {} from the spawned thread!", i);
            thread::sleep(Duration::from_millis(1));            
        }
    });

    for i in 1..5 {
        println!("hi number {} from the main thread!", i);
        thread::sleep(Duration::from_millis(1));
    }

    // join spwaned thread.
    handle.join().unwrap();
}
```

```
hi number 1 from the main thread!
hi number 2 from the main thread!
hi number 1 from the spawned thread!
hi number 3 from the main thread!
hi number 2 from the spawned thread!
hi number 4 from the main thread!
hi number 3 from the spawned thread!
hi number 4 from the spawned thread!
hi number 5 from the spawned thread!
hi number 6 from the spawned thread!
hi number 7 from the spawned thread!
hi number 8 from the spawned thread!
hi number 9 from the spawned thread!
```

## `move` 클로저
> 스레드 간 데이터 소유권 이동

move 클로저는 thread::spawn와 함께 자주 사용되는데 그 이유는 이것이 여러분으로 하여금 어떤 스레드의 데이터를 다른 스레드 내에서 사용하도록 해주기 때문입니다.

**클로저의 파라미터 목록 앞에 move 키워드를 이용하여 클로저가 그 환경에서 사용하는 값의 소유권을 강제로 갖게 할 수 있습니다.** 이 기술은 값의 소유권을 한 스레드에서 다른 스레드로 이전하기 위해 새로운 스레드를 생성할 때 특히 유용합니다.

move가 필요한 코드를 먼저 보여드리겠습니다.

```rs
use std::thread;

fn main() {
    let v = vec![1,2,3];

    let handle = thread::spawn(|| {
        println!("Here's a vector: {:?}", v); // 만약 v에 대해서 레퍼런스를 주었다면, v가 언제까지 살아있을지 확신을 할 수 없습니다.
    });

    drop(v); // v를 main thread에서 제거
    handle.join().unwrap();
}
```

이러한 이유로, 러스트 컴파일러는 다음과 같은 에러 메시지를 제공합니다.

```
help: to force the closure to take ownership of `v` (and any other referenced
variables), use the `move` keyword
  |
6 |     let handle = thread::spawn(move || {
  |                                ^^^^^^^
```

즉 클로저 안에서 v에 대한 소유권을 main으로부터 받아오라고 말해줍니다.

```rs
use std::thread;

fn main() {
    let v = vec![1,2,3];

    let handle = thread::spawn(move || {
        println!("Here's a vector: {:?}", v);
    });

    // drop(v);
    // error[E0382]: use of moved value: `v`

    handle.join().unwrap();
}
```

## Message Passing
스레드에 대한 Go의 슬로건 중 하나는 다음과 같습니다.

> "Do not communicate by sharing memory; instead, share memory by communicating."

이 처럼 안전한 동시성을 보장하는 인기있는 방법은 message passing 입니다. 러스트 또한 go처럼 channel을 활용합니다.

프로그래밍에서 채널은 둘로 나뉘어져 있습니다.

1. transmitter (송신자)
    - abbr. `tx`
2. receiver (수신자)
    - abbr. `rx`

`transmitter` 측은 여러분이 강에 고무 오리를 띄우는 상류 위치이고, `receiver` 측은 하류에 고무 오리가 도달하는 곳입니다. 

여러분 코드 중 한 곳에서 여러분이 보내고자 하는 데이터와 함께 송신자의 메소드를 호출하면, 다른 곳에서는 도달한 메세지에 대한 수신 종료를 검사합니다. 송신자 혹은 송신자가 드롭되면 채널이 닫혔다 (closed) 라고 말합니다.

```rs
use std::sync::mpsc;

fn main() {

}
```

채널을 사용하기 위해서는 std 라이브러리인 `mpsc`를 활용합니다. `mpsc`는 `multiple producer, single consumer`의 약자입니다.

다시 말해 표준 라이브러리 mpsc가 채널을 구현한 방법은 **한 채널이 값을 생성하는 복수개의 송신 단말을 가질 수 있지만 값을 소비하는 단 하나의 수신 단말을 가질 수 있음을 의미합니다.**

`mpsc::channel()`는 튜플을 반환합니다.

```rs
use std::sync::mpsc;
use std::thread;

fn main() {
    let (tx, rx) = mpsc::channel(); // 튜플 반환

    thread::spawn(move || {
        let v = String::from("Hi");
        tx.send(v).unwrap(); // Result<T, E>
    });

    let received = rx.recv().unwrap();
    println!("Got: {}", received);
}
```

`rx`는 2가지 유용한 메소드를 활용해 메시지를 받을 수 있습니다.

1. `.recv()`
2. `.try_recv()`

`recv()`는 **block**한 상태로 메시지가 보내질 때까지 기다립니다. 그리고 그 전달된 값은 Result<T,E>형태로 전달됩니다.


`try_recv()`는 블록하지 않는 대신, 해당 시점에 `Result<T,E>`형태로 값을 전달해줍니다. 만약 메시지가 전달되었다면 `Ok`, 없다면 `Err`입니다. 만약 메시지를 기다리면서 다른 작업을 해야한다면 유용하게 사용할 수 있습니다.

## 채널간 메시지 전달과 소유권 처리

다음으로 tx에서 rx로 값을 내려보낸 뒤에, 그 값을 사용한다면 소유권 체크가 어떻게 되는지 확인해보겠습니다.

```rs
use std::sync::mpsc;
use std::thread;

fn main() {
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let v = String::from("Hi");
        tx.send(v).unwrap();
        println!("Here!!!! is Problem {val}");
    });

    let received = rx.recv().unwrap();
    println!("Got: {}", received);
}
```

```
error[E0382]: use of moved value: `val`
  --> src/main.rs:10:31
   |
9  |         tx.send(val).unwrap();
   |                 --- value moved here
10 |         println!("val is {}", val);
   |                               ^^^ value used here after move
   |
   = note: move occurs because `val` has type `std::string::String`, which does
not implement the `Copy` trait
```

당연히 소유권체크에서 컴파일러가 똑똑하게 잡아줍니다.

## multiple producer

마지막으로 `mpsc`의 multiple producer를 사용하는 코들르 작성하겠습니다.

```rs
use std::sync::mpsc;
use std::thread;
use std::time::Duration;

fn main() {
    let (tx, rx) = mpsc::channel();
    let tx1 = tx.clone();
    thread::spawn(move || {
        let vals = vec![
            String::from("tx1: hi"),
            String::from("tx1: from"),
            String::from("tx1: the"),
            String::from("tx1: thread"),
        ];
        for v in vals {
            tx1.send(v).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    thread::spawn(move || {
        let vals = vec![
            String::from("tx: more"),
            String::from("tx: messages"),
            String::from("tx: for"),
            String::from("tx: you"),
        ];

        for v in vals {
            tx.send(v).unwrap();
            thread::sleep(Duration::from_secs(1));
        }
    });

    // receiver
    for r in rx {
        println!("Got: {}", r);
    }
}
```

```
Got: tx: more
Got: tx1: hi
Got: tx: messages
Got: tx1: from
Got: tx: for
Got: tx1: the
Got: tx: you
Got: tx1: thread
```

결과값 순서는 다르게 나올 수 있습니다.
