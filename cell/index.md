# Cell and RefCell


You need to write intro in here
<!--more-->

## TL;DR
참고로 Cell과 RefCell는 thread-safe하지 않다. 즉 single-threaded way에서만 safe를 보장한다.

- `Cell`: immutable한 타입을 명시적으로 런타임에 mutable하게 사용할 수 있고 value를 return한다.
- `RefCell`: Cell의 Reference버전


- 차이점
    - Cell provides you values, RefCell with references
    - Cell never panics, RefCell can panic
    - panic이 없는 Cell을 선호하게 짜는게 유리하며, 몇몇 operation에서는 Cell을 사용할 수 없으니 그때 RefCell을 사용한다.


## 1. Cell
> Interior mutability

- [Cell](https://fongyoong.github.io/easy_rust/Chapter_41.html?highlight=cell#cell)

<center>

![](/images/cell.png)

</center>

`mut` 키워드 없이, mutable하게 변수를 변경할 수 있는 방법이 뭐가 있을까?
러스트에서는 몇가지 방법을 제공하는데 그 중 가장 simple한 방법은 `Cell`이다.

> 언제 사용할까? 예를 들어 immutable struct에서 특정 필드만 mutable하게 관리하고 싶을 때.

- before

```rs
struct PhoneModel {
    company_name: String,
    model_name: String,
    screen_size: f32,
    memory: usize,
    date_issued: u32,
    on_sale: bool,
}

fn main() {
    let super_phone_3000 = PhoneModel {
        company_name: "YY Electronics".to_string(),
        model_name: "Super Phone 3000".to_string(),
        screen_size: 7.5,
        memory: 4_000_000,
        date_issued: 2020,
        on_sale: true,
    };

}
```

- after

```rs
use std::cell::Cell;

struct PhoneModel {
    company_name: String,
    model_name: String,
    screen_size: f32,
    memory: usize,
    date_issued: u32,
    on_sale: Cell<bool>,
}

fn main() {
    let super_phone_3000 = PhoneModel {
        company_name: "YY Electronics".to_string(),
        model_name: "Super Phone 3000".to_string(),
        screen_size: 7.5,
        memory: 4_000_000,
        date_issued: 2020,
        on_sale: Cell::new(true),
    };

    // 10 years later, super_phone_3000 is not on sale anymore
    super_phone_3000.on_sale.set(false);
}
```

- **`Cell`은 value를 return하기** 때문에 `Copy types`에 가장 적홥하다.


## 2. `RefCell`
> [RefCell](https://fongyoong.github.io/easy_rust/Chapter_42.html)

<center>

![](/images/refcell.png)

</center>


`Cell`과 마찬가지로 `mut`없이 value를 change할 수 있지만, `RefCell`는 `Copy` 대신 reference를 사용한다.

```rs
use std::cell::RefCell;

#[derive(Debug)]
struct User {
    id: u32,
    year_registered: u32,
    username: String,
    active: RefCell<bool>,
    // Many other fields
}

fn main() {
    let user_1 = User {
        id: 1,
        year_registered: 2020,
        username: "User 1".to_string(),
        active: RefCell::new(true),
    };

    println!("{:?}", user_1.active);
}
```

RefCell은 총 2가지 메서드를 통해 RO(read only), RW(Read/Write) 를 처리한다. 이는 `&`와 `&mut` 각각과 대응된다.

- `.borrow()`
- `.borrow_mut()`

당연히 `borrow_mut()`는 최대 1개만 사용되어야 한다. 또한 readonly인 `borrow()`는 여러개 사용되어도 상관없다.

**마지막으로 `borrow()`와 `borrow_mut()`는 동시에 사용할 수 없다.**


- `replace()`
```rs
user_1.active.replace(false);
```

- `replace_with()`: 조건을 통한 replace
```rs
// 🚧
let date = 2020;

user_1
    .active
    .replace_with(|_| if date < 2000 { true } else { false });
println!("{:?}", user_1.active);
```


**또한 Cell과 마찬가지로 Runtime에 동작하기 때문에 panic이 일어날 수 있다. 예를들면 borrow_mut()를 2번사용하게 되면 컴파일 타임에러가 아닌 런타임에 에러가 난다.**


```rs
...

fn main() {
    let user_1 = User {
        id: 1,
        year_registered: 2020,
        username: "User 1".to_string(),
        active: RefCell::new(true),
    };

    let borrow_one = user_1.active.borrow_mut(); // first mutable borrow - okay
    let borrow_two = user_1.active.borrow_mut(); // second mutable borrow - not okay
}

// thread 'main' panicked at 'already borrowed: BorrowMutError', C:\Users\mithr\.rustup\toolchains\stable-x86_64-pc-windows-msvc\lib/rustlib/src/rust\src\libcore\cell.rs:877:9
// note: run with `RUST_BACKTRACE=1` environment variable to display a backtrace
// error: process didn't exit successfully: `target\debug\rust_book.exe` (exit code: 101)
```



