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

