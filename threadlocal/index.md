# std::thread_local!


A brief summary about `std::thread_local!` 
<!--more-->

## `thread_local!`

`thread_local!`은 rust의 std 매크로이며, thread_local 내부에 선언된 변수를 wrap한뒤  `std::thread::LocalKey`를 제공합니다. pub과 `#[]`(attribute)가 허용됩니다.


## `std::thread::LocalKey`

`thread_local` storage에 들어있는 값에 접근할 수 있는 `key`입니다.

`.with()`는 thread_local에 들어있는 변수에 대한 reference를 yield합니다. 이 값은 sent across threads 할 수 없고, given closure를 escape할 수 없습니다.

`Drop`을 통해서 escape될 때 destruct됩니다.


## 예시1
> rust docs의 예시1


```rust
use std::cell::RefCell;
use std::thread;

thread_local!(static FOO: RefCell<u32> = RefCell::new(1));

FOO.with(|f| {
    assert_eq!(*f.borrow(), 1);
    *f.borrow_mut() = 2;
});

// each thread starts out with the initial value of 1
let t = thread::spawn(move|| {
    FOO.with(|f| {
        assert_eq!(*f.borrow(), 1);
        *f.borrow_mut() = 3;
    });
});

// wait for the thread to complete and bail out on panic
t.join().unwrap();

// we retain our original value of 2 despite the child thread
FOO.with(|f| {
    assert_eq!(*f.borrow(), 2);
});
```

- thread_local! 매크로를 거친 `FOO`는 `LocalKey<RefCell<u32>>` 타입입니다.
- `with()` 매서드에 `|f| {}`라는 anonymous function을 넣어줍니다. 
- `thread::spawn()`이 없는 경우는 main thread이다.
- main thread에서 spawn 한 쓰레드에서 `borrow_mut()`를 3으로 변경하더라도, main_thread의 FOO 변수값은 2 그대로이다.

## 예시2: `icp nft`

```rs
use std::mem;

thread_local! {
    static STATE: RefCell<State> = RefCell::default();
}


#[pre_upgrade]
fn pre_upgrade() {
    let state = STATE.with(|state| mem::take(&mut *state.borrow_mut()));
    ...
}

// init or update 
#[init]
fn init(args: InitArgs) {
    STATE.with(|state| {
        let mut state = state.borrow_mut();
        state.custodians = args
            .custodians
            .unwrap_or_else(|| HashSet::from_iter([api::caller()]));
        state.name = args.name;
        state.symbol = args.symbol;
        state.logo = args.logo;
    });
}


// retrieve
#[query(name = "balanceOfDip721")]
fn balance_of(user: Principal) -> u64 {
    STATE.with(|state| {
        state
            .borrow()
            .nfts
            .iter()
            .filter(|n| n.owner == user)
            .count() as u64
    })
}
```

- thread_local에 wrap된 변수를 변경하기 위해서는 RefCell에 접근하는 `borrow_mut()`메서드가 필요하다.
- Read만 할 경우에는 `.borrow()`를 사용하면 된다.
