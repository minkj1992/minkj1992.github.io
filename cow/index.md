# COW 🐄 in Rust 🦀


dfinity/ic 코드를 보다 이해되지 않는 코드들이 있어, Cow를 왜 사용하는지 정리하려 합니다.
<!--more-->


## TL;DR

```rs
pub enum Cow<'a, B>
where
    B: 'a + ToOwned + ?Sized,
{
    Borrowed(&'a B),
    Owned(<B as ToOwned>::Owned),
}
```

- Cow의 약어 뜻은 clone-on-write로 read가 아닌 write시 clone시킬 수 있는 기능을 가지고 있습니다.
- Cow는 Borrowed, Owned를 구분하는 enum 타입입니다. 
- 제너릭 `B`는 'a와, ToOwned, ?Sized로 바운드 되어있습니다.

- `ToOwned` trait
    - borrowed 데이터의 `Clone`에 대한 일반화
    - clone시켜서, Owned 타입을 만들어 낼 수 있다.
- `?Sized` trait
    - 컴파일 타임에 constant size를 알 수 있는 타입
    - 모든 타입 파라미터는 implicit bound로 `Sized`를 보유하고 있다.
    - `?`를 사용하면 이 bound를 remove시켜줄 수 있다.

즉 Cow는 `Borrowed` 또는 `Owned` 둘 모두를 사용하고 싶을 때 사용한다. 예를 들면 &str, String타입 모두 사용하길 원하는 경우.


```rs
#[derive(CandidType, Deserialize, Clone)]
struct LogoResult {
    logo_type: Cow<'static, str>,
    data: Cow<'static, str>,
}

// ...
const DEFAULT_LOGO: LogoResult = LogoResult {
    data: Cow::Borrowed(include_base64!("logo.png")), //&str
    logo_type: Cow::Borrowed("image/png"), //&str
};

```


## Cow를 사용하는 이유
> [6 thing you can do with the 🐄 in 🦀](https://dev.to/kgrech/6-things-you-can-do-with-the-cow-in-rust-4l55)

### 1. A function rarely modifying the data

불필요하게 clone을 하게되는 경우를 막기 위해서 cow를 사용할 수 있습니다.
FYI `to_string()`은 복사본을 전달한다.

#### before

```rs
fn remove_whitespaces(s: &str) -> String {
    s.to_string().replace(' ', "")
}

fn main() {
    let value = remove_whitespaces("Hello world!");
    println!("{}", value);
}
```

#### after

```rs
use std::borrow::Cow;

fn remove_whitespaces(s: &str) -> Cow<str> {
    if s.contains(' ') {
        Cow::Owned(s.to_string().replace(' ', ""))
    } else {
        Cow::Borrowed(s)
    }
}

fn main() {
    let value = remove_whitespaces("Hello world!");
    println!("{}", value);
}
```

### 2. A struct optionally owning the data


### 3. A clone on write struct
### 4. Keep your own type inside it
### 5. Borrow the type as dyn Trait
### 6. Implement safe wrapper over FFI type




