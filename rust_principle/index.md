# Rust 디자인 원칙 🦀


[Rust Q&A](https://prev.rust-lang.org/ko-KR/faq.html)를 통해서 알아보는 Rust 프로그래밍 원리를 살펴보고, 유용했던 문구들을 스크랩 해둡니다.
<!--more-->

<br />
또한 앞으로 Rust 공부는 다음과 같은 순서로 진행할 예정입니다.

1. [Rust Q&A](https://prev.rust-lang.org/en-US/faq.html#syntax)
2. [Book: The Rust Programming language](https://doc.rust-lang.org/book/ch00-00-introduction.html)
3. [Effective Rust(35)](https://lurklurk.org/effective-rust/cover.html)
4. [Rust by Example](https://doc.rust-lang.org/rust-by-example/)

---

## 1. Rust 프로젝트

{{< admonition question "Rust의 목표는 무엇일까요?" >}}
**_To design and implement a safe, concurrent, practical systems language._**

Rust는 이 수준의 추상화와 효율을 추구하는 다른 언어들이 만족스럽지 못 하기에 존재합니다. 특히:

1. 안전성이 너무 덜 주목되어 있습니다.
2. 동시성 지원이 부족합니다.
3. 실용적으로 쓰기가 힘듭니다.
4. 자원에 대한 제어가 제한적입니다.

Rust는 효율적인 코드와 편안한 수준의 추상화를 제공하며, 동시에 위 4가지를 모두 개선하는 대안으로 만들어졌습니다.
{{< /admonition  >}}


{{< admonition question "Rust의 목표가 아닌 것은 무엇이 있나요?" >}}
1. 우리는 특별히 최신의 기술을 도입하지 않습니다. **오래되고 자리 잡힌 기술이 더 좋습니다.**
2. 우리는 **표현력, 최소주의 또는 우아함을 다른 목표에 우선하지 않습니다**. 이들은 바람직하긴 하지만 부수적인 목표입니다.
3. 우리는 C++나 기타 다른 언어의 모든 기능 집합을 커버하려 하지 않습니다. **Rust는 자주 쓰이는 기능들을 제공**할 것입니다.
4. 우리는 100% 정적이거나, 100% 안전하거나, 100% 반영적(reflective)이거나, 기타 어떤 의미에서도 너무 교조적이려 하지 않습니다. **트레이드 오프는 존재합니다.**
5. 우리는 Rust가 “가능한 모든 플랫폼”에서 동작할 걸 요구하지 않습니다. **언젠가 Rust는 널리 쓰이는 하드웨어와 소프트웨어 플랫폼에서 불필요한 타협 없이 동작할 것입니다.**
{{< /admonition  >}}

---

## 2. 성능

{{< admonition question "Rust는 얼마나 빠른가요?" >}}
_C++와 동일하게 Rust는 비용 없는 추상화를 주요 원칙으로 삼습니다. Rust에는 전역으로 성능을 떨어뜨리는 추상화가 존재하지 않으며, 런타임 시스템에서 부하가 발생하지도 않습니다._

_Rust가 LLVM에 기반해 있고 LLVM이 보기에 Clang과 비슷하게 보이려 한다는 걸 생각해 보면, LLVM에서 성능 개선이 일어난다면 Rust도 도움을 받게 됩니다. 장기적으로는 Rust 타입 시스템의 더 풍부한 정보로 C/C++ 코드에서는 어렵거나 불가능한 최적화도 가능해질 것입니다._
{{< /admonition  >}}

{{< admonition question "Rust는 쓰레기 수거(garbage collection, GC)를 하나요?" >}}
_아니요. Rust의 중요 혁신 중 하나는 쓰레기 수거 없이 메모리 안전성을 보장한다는 것입니다(즉, 세그폴트가 나지 않습니다)._

_Rust는 GC를 피한 덕에 여러 장점을 제공할 수 있었습니다. 자원들을 예측 가능하게 해제할 수 있고, 메모리 관리 오버헤드가 낮으며, 사실상 런타임 시스템이 없습니다. 이 모든 특징들 때문에 Rust는 아무 맥락에나 깔끔하게 포함(embed)하기 쉬우며, 이미 GC를 가지고 있는 언어에 Rust 코드를 통합하기에도 훨씬 쉽습니다._
{{< /admonition  >}}

{{< admonition question "제 프로그램이 왜 느린 거죠?" >}}
_Rust 컴파일러는 요청이 없다면 최적화 없이 컴파일을 하는데, 이는 최적화를 하면 컴파일이 느려지고 개발 과정에서는 보통 바람직하지 않기 때문입니다._

_cargo로 컴파일을 한다면 --release 플래그를 쓰세요. rustc를 직접 써서 컴파일을 한다면 -O 플래그를 쓰세요. 어느 쪽이나 최적화를 켜는 역할을 합니다._
{{< /admonition  >}}


{{< admonition question "Rust 컴파일이 느린 것 같습니다. 왜 그런 건가요?" >}}
코드를 기계어로 번역하고 최적화를 하기 때문입니다. Rust는 효율적인 기계어로 컴파일되는 고수준 추상화를 제공하고, 이 번역 과정은 특히 최적화를 할 경우 시간이 걸리게 마련입니다.


**Rust 컴파일이 느리다고 느끼는 주된 원인은 C++와 Rust가 컴파일 모델이 다르다는 점, 즉 C++의 컴파일 단위는 한 파일이지만 Rust는 여러 파일로 이루어진 크레이트라는 것 때문입니다.** 따라서 개발 도중에 C++ 파일 하나를 고치면 Rust에 비해 컴파일 시간이 훨씬 줄어들 수 있습니다. 현재 Rust 컴파일러를 리팩토링해서 증분 컴파일(Incremental Compile)을 가능하게 하려는 대형 작업이 진행 중이며, 완료되면 Rust에서도 C++ 모델과 같이 컴파일 시간이 개선될 것입니다.


(4) 컴파일 모델과는 별개로, Rust의 언어 설계에는 컴파일 시간에 영향을 미치는 요소가 여럿 있습니다.


1. Rust는 비교적 **복잡한 타입 시스템을** 가지고 있고, 실행 시간에 Rust를 안전하게 만들기 위한 제약 사항을 강제하는 데 무시할 수 없는 컴파일 시간을 사용해야 합니다.


2. Rust 컴파일러에는 오래된 기술 부채가 있으며, 특히 생성되는 LLVM IR의 품질이 좋지 못하기 때문에 LLVM이 시간을 들여 이를 “고쳐야” 합니다. **미래에는 MIR 기반 최적화 및 번역 단계가 Rust 컴파일러가 LLVM에 가하는 부하를 줄여 줄지도** 모릅니다.


3. Rust가 코드 생성에 LLVM을 쓰는 것은 양날의 검이라는 점입니다. **LLVM 덕분에 Rust는 세계구급 런타임 성능을 보여 주지만, LLVM은 컴파일 시간에 촛점을 맞추지 않은 거대한 프레임워크이며 특히 품질이 낮은 입력에 취약**합니다.


4. 마지막으로 **Rust가 일반화(제너릭) 타입을 C++와 비슷하게 단형화(monomorphise)하는 전략은 빠른 코드를 생성하지만, 다른 번역 전략에 비해 상당히 많은 코드를 생성해야 한다는 문제**가 있습니다. 이 코드 팽창은 트레이트 객체를 써서 동적 디스패치와 장단을 교환할 수 있습니다.

{{< /admonition  >}}

{{< admonition question "Rust의 HashMap은 왜 느린가요?" >}}
_Rust의 HashMap은 기본적으로 `SipHash` 해시 알고리즘을 사용합니다. 이 알고리즘은 해시 테이블 충돌 공격을 막으면서 여러 종류의 입력에 대해 적절한 성능을 내도록 설계되었습니다._

_SipHash가 많은 경우 경쟁력 있는 성능을 보여 주긴 하지만, SipHash는 정수 같이 키가 짧을 경우 다른 해시 알고리즘에 비해 현저히 느립니다. 이 때문에 종종 HashMap의 성능이 낮은 걸 볼 수 있습니다. 이런 경우에는 보통 `FNV 해시`를 추천하지만, 이 알고리즘이 충돌 공격에서 SipHash와 다른 특성을 보인다는 점은 염두에 두어야 합니다._
{{< /admonition  >}}


{{< admonition question "Rust는 꼬리 재귀(tail-call) 최적화를 하나요?" >}}
_일반적으로는 아닙니다. 제한적으로 꼬리 재귀 최적화를 하긴 하지만 보장되지는 않습니다. 이 기능은 언제나 요청되어 왔기 때문에 Rust에는 이를 위해 예약어(become)가 예약되어 있습니다_
{{< /admonition  >}}

{{< admonition question "Rust에는 런타임이 있나요?" >}}
_Java 같은 언어들에서 말하는 그런 통상의 런타임은 없습니다만, Rust 표준 라이브러리의 일부분은 힙(heap), 스택 추적(backtrace), 되감기(unwinding) 및 보호(guard)를 제공하는 “런타임”이라고 볼 수 있습니다. 사용자의 main 함수가 실행되기 전에는 소량의 초기화 코드가 실행됩니다. 또한 Rust 표준 라이브러리는 C 표준 라이브러리를 링크하는데 여기에서도 비슷한 런타임 초기화가 일어납니다. Rust 코드는 표준 라이브러리 없이 컴파일될 수 있으며 이 경우 런타임은 대략 C와 비슷해집니다._
{{< /admonition  >}}


## 3. 문법 (Syntax)

{{< admonition question "왜 중괄호인가요?" >}}
_또한 중괄호는 프로그래머 입장에서는 더 유연한 문법을 제공하고 컴파일러 입장에서는 더 간단한 파서를 가능하게 합니다._
{{< /admonition  >}}

{{< admonition question "if 조건에서 소괄호를 생략할 수 있는데, 그럼 한 줄짜리 블럭에는 왜 중괄호를 넣어야 하나요?" >}}
_C에서는 if 조건문에서 괄호가 필수이고 중괄호가 선택이지만, Rust에서는 반대로 합니다. 이렇게 해서 조건문 몸체와 조건을 명확하게 구분할 수 있고, 중괄호가 선택이라서 벌어지는 위험도 막을 수 있는데, 이는 Apple의 `goto fail` 버그와 같이 리팩토링 과정에서 흔히 생기고 잡기 어려운 오류들을 유발할 수 있습니다._
{{< /admonition  >}}


{{< admonition question "Why is there no literal syntax for dictionaries?" >}}
_**Rust의 전반적인 설계는 언어의 크기를 제한하되 강력한 라이브러리를 만들 수 있게 하는 쪽을 선호합니다.** Rust는 배열과 문자열 리터럴을 초기화하는 문법을 가지고 있지만 언어에 내장된 컬렉션 타입은 이걸로 전부입니다. 매우 널리 쓰이는 Vec 컬렉션 타입 같이, 라이브러리에서 정의하는 다른 타입들은 vec! 같은 매크로를 사용하여 초기화를 합니다._

_나중에는 Rust가 매크로를 써서 컬렉션을 초기화하는 설계가 다른 타입에도 일반적으로 사용할 수 있도록 확장될 수 있고, 그렇게 되면 HashMap이나 Vec 같은 것 뿐만이 아니라 BTreeMap 같은 다른 타입들도 간단하게 초기화할 수 있게 될 것입니다._
{{< /admonition  >}}


{{< admonition question "When should I use an implicit return?(암묵적인 반환)" >}}

Rust는 매우 수식 지향적인 언어이며 “암묵적인 반환”은 이 설계의 한 부분입니다. **if, match나 일반 블록들은 Rust에서는 다 수식입니다.** 예를 들어 다음 코드는 i64가 홀수인지 확인하고 결과를 단순히 값으로 내서 결과를 반환합니다:

```rs
fn is_odd(x: i64) -> bool {
    if x % 2 != 0 { true } else { false }
}

fn is_odd(x: i64) -> bool {
    x % 2 != 0
}
```

**두 예제에서 함수의 마지막 줄은 그 함수의 반환값입니다. 중요한 것은 함수가 세미콜론으로 끝난다면 그 반환값은 ()이고, 이는 반환값이 없다는 뜻이라는 점입니다**. 암묵적으로 반환하려면 세미콜론이 없어야 합니다.

명시적인 반환은 함수 몸체의 맨 끄트머리보다 이전에 반환을 해야 해서 암묵적인 반환이 불가능할 때만 쓰입니다. 물론 위 함수들도 `return ` 예약어와 세미콜론을 쓸 수는 있지만, 이는 Rust 코드 규약에 어긋날 뿐 아니라, 불필요하게 번잡해 보입니다.
{{< /admonition  >}}

{{< admonition question "왜 함수의 타입 서명(signature)들은 추론되지 않는 거죠?" >}}
Rust에서 선언은 타입을 명시적으로 쓰는 편이며 실제 코드는 타입을 추론하는 편입니다. 이 설계에는 몇 가지 이유가 있습니다:

- 선언의 서명을 명시적으로 쓰면 모듈 및 크레이트 수준에서 인터페이스 안정성을 강제하는 데 도움이 됩니다.
- 서명은 프로그래머가 코드를 더 잘 이해할 수 있게 하므로, IDE가 함수의 인자 타입들을 추측하려고 전체 크레이트에 추론 알고리즘을 돌릴 필요가 사라집니다. 
- 언제나 명시적이고 바로 옆에 있기 때문이죠.
기계적으로는 추론 과정에서 한 번에 한 함수만 보면 되므로 추론 알고리즘이 간단해집니다.
{{< /admonition  >}}

{{< admonition question "왜 `match`에는 모든 조건들이 들어 있어야 하나요?" >}}
**리팩토링을 돕고 코드를 명료하게 하기 위함입니다.**

먼저, match가 모든 가능성을 커버하고 있다면 enum에 새 변종(variant)을 넣을 때 실행 시간 **(execution time)에 오류가 나는 게 아니라 컴파일(compile time)이 실패하게 됩니다.** Rust에서 이런 종류의 컴파일러 도움은 두려움 없이 리팩토링을 가능하게 합니다.

두 번째로, 이러한 체크는 기본 선택지를 명시적으로 만듭니다. **일반적으로 모든 가능성을 커버하지 않는 match를 안전하게 만드는 방법은 아무 선택지도 선택되지 않았을 때 스레드를 패닉하게 만드는 것 뿐입니다.** Rust의 옛 버전에서는 match가 모든 가능성을 커버하지 않아도 되게 했는데 수많은 버그의 온상이 되었습니다.

기술되지 않은 선택지는 `_ 와일드 카드`로 간단하게 무시할 수 있습니다:

```rs
match val.do_something() {
    Cat(a) => { /* ... */ }
    _      => { /* ... */ }
}
```
{{< /admonition  >}}

---

## 4. 디자인 패턴

{{< admonition question "Rust는 객체 지향적(object-oriented)인가요?" >}}
_It is multi-paradigm. 객체 지향 언어에서 할 수 있는 많은 것들은 Rust에서도 할 수 있지만, 전부 가능한 건 아니고, 여러분에게 익숙한 추상화를 사용하지 않을 수도 있습니다._
{{< /admonition  >}}

{{< admonition question "How do I map object-oriented concepts to Rust?" >}}
_`다중 상속`과 같은 객체 지향 개념을 Rust로 옮기는 방법은 여럿 있습니다만, Rust는 객체 지향이 아니기에 객체 지향 언어들과는 상당히 다르게 보일 수 있습니다._
{{< /admonition  >}}

{{< admonition question "How do I handle configuration of a struct with optional parameters?" >}}
_가장 쉬운 방법은 `구조체 인스턴스`를 생성하는 어떤 함수에든 (보통 `new()`에) `Option 타입`을 쓰는 겁니다. 또 다른 방법은 `builder패턴`을 써서, 타입을 생성하기 전에 **멤버 변수를 인스턴스화하는 특정 함수들을 호출**해야 하도록 하는 것입니다._
{{< /admonition  >}}

{{< admonition question "Rust에서 전역 변수를 쓰려면 어떻게 하죠?" >}}
_Rust에서 `전역 변수`는 컴파일 시간에 계산된 전역 상수라면 `const` 선언을 쓸 수 있고, 변경 가능한 전역 변수는 `static`을 쓸 수 있습니다._

다만 `static mut` 변수를 변경하려면 unsafe가 필요한데, 이는 안전한 Rust에서는 발생하지 않는다고 보장하는 데이터 레이스(data race)가 일어날 수 있기 때문입니다. **const와 static 값의 중요한 차이는 static에서는 참조를 얻을 수 있지만 const는 지정된 메모리 위치를 가지지 않기 때문에 불가능하다는 점입니다.**
{{< /admonition  >}}

{{< admonition question "How can I set `compile-time constants` that are defined procedurally?" >}}
_You can define primitives using const declarations as well as define const functions and inherent methods._

(원시 값을 const 선언으로 정의할 수 있고, const 함수나 선천적인 메소드도 정의할 수 있습니다.)

_To define procedural constants that can’t be defined via these mechanisms, use the lazy-static crate, which emulates compile-time evaluation by automatically evaluating the constant at first use._

(이 방법으로 선언할 수 없는 `procedural 상수`를 선언하려면 `lazy-static crate`를 사용하세요. 이 크레이트는 컴파일 시간 evaluation를 상수가 처음 사용될 때, 자동으로 평가하여 procedural constants를 흉내냅니다.)
{{< /admonition  >}}

{{< admonition question "`main` 이전에 실행되는 초기화 코드를 만들 수 있나요?" >}}
_Rust에는 “main 이전의 life”라는 개념이 없습니다. `lazy-static 크레이트`가 가장 가까운 것일텐데, 이 크레이트는 “main보다 이전”이라는 시간을 `정적 변수`를 처음 사용할 때 지연하여 초기화하는 걸로 흉내냅니다._
{{< /admonition  >}}

{{< admonition question "Rust에서 상수 수식이 아닌 값을 전역에 넣을 수 있나요?" >}}
_아니요._

**Globals cannot have a non-constant-expression constructor and cannot have a destructor at all.** Static constructors are undesirable because portably ensuring a static initialization order is difficult. **`Life before main` is often considered a misfeature, so Rust does not allow it.**
{{< /admonition  >}}

--- 
## 5. Other Languages


{{< admonition question "How can I implement something like C's `struct X { static int X; };` in Rust?" >}}
_Rust는 static 필드가 없습니다. 대신 주어진 모듈에서만 접근할 수 있는 static 변수를 선언할 수 있습니다._
{{< /admonition  >}}

{{< admonition question "왜 Rust는 C 같이 안정화된 ABI가 없는 건가요? 그리고 왜 `extern`을 어노테이트 해야하는 거죠?" >}}
_Rust가 2015년 5월에야 1.0이 되었다는 걸 볼 때 안정된 ABI 같은 큰 투자를 하기에는 아직 너무 이릅니다. 하지만 미래에도 일어나지 않을 거라는 얘기는 아닙니다.

**`extern`를 쓰면 Rust가 잘 정의된 C ABI 같이 특정한 ABI를 써서 다른 언어와 상호작용하도록 할 수 있습니다.**
{{< /admonition  >}}

{{< admonition question "Rust 코드가 C 코드를 호출할 수 있나요?" >}}
_네. C 코드를 Rust에서 부르는 것은 C++에서 C 코드를 부르는 것만큼 효율적이도록 설계되었습니다._
{{< /admonition  >}}

{{< admonition question "C 코드가 Rust 코드를 호출할 수 있나요?" >}}
_네. **Rust 코드가 extern 선언으로 노출되어 C의 ABI와 호환되도록 만들어야 합니다.** 

이러한 함수는 C 코드에 함수 포인터로 전달되거나, `#[no_mangle]` 속성으로 `symbol mangling`을 껐을 경우, C 코드에서 바로 호출될 수 있습니다._
{{< /admonition  >}}

{{< admonition question "C++의 템플릿 특수화 같은 걸 Rust에서는 어떻게 할 수 있을까요?" >}}
Rust는 현재 템플릿 특수화(template specialization)와 완전히 같은 기능을 가지고 있지 않지만, [현재 작업이 진행 중](https://github.com/rust-lang/rfcs/pull/1210)이며 아마 곧 추가될 것입니다.

다만 [Associated types](https://doc.rust-lang.org/beta/rust-by-example/generics/assoc_items/types.html)으로 비슷한 결과를 얻을 수도 있습니다.
{{< /admonition  >}}

{{< admonition question "Rust에는 C++ 같은 생성자가 있나요?" >}}
**아니요.** 대신 함수가 생성자와 같은 역할을 수행합니다. 

Rust에서 생성자에 대응되는 함수의 일반적인 이름은 new()로, 이는 언어 규칙이 아니라 단순한 규약일 따름입니다. new() 함수는 다른 함수랑 다를 바가 없고, 이런 식으로 씁니다.

```rs
struct Foo {
    a: i32,
    b: f64,
    c: bool,
}

impl Foo {
    fn new() -> Foo {
        Foo {
            a: 0,
            b: 0.0,
            c: false,
        }
    }
}
```
{{< /admonition  >}}

{{< admonition question "Go와 Rust가 비슷한 점은 무엇이고 다른 점은 무엇인가요?" >}}

- Rust는 Go보다 저수준입니다. 예를 들어 Rust는 쓰레기 수거기(garbage collector)를 필요하지 않지만 Go는 필요로 합니다. 일반적으로 Rust는 C나 C++와 비견할 만한 제어 수준을 제공합니다.
- Rust의 촛점은 고수준의 편안함을 제공하면서도 안전함과 효율성을 보장하는 것이며, Go의 촛점은 빠르게 컴파일되고 수많은 도구와 함께 멋지게 동작할 수 있는 작고 간단한 언어가 되고자 하는 것입니다.
- Rust has strong support for generics, which Go does not.
- Rust는 함수형 프로그래밍에서 많은 영향을 받았으며, 여기에는 하스켈의 타입 클래스에서 유래한 타입 시스템이 포함됩니다. Go has a simpler type system, using `interfaces` for basic generic programming.
{{< /admonition  >}}


## 6. 모듈 및 크레이트

{{< admonition question "`모듈`과 `크레이트` 사이에 어떤 관계가 있나요?" >}}

- 크레이트는 컴파일 단위로, Rust 컴파일러가 다룰 수 있는 가장 작은 규모의 코드입니다.
- 모듈은 크레이트 안에 있는 코드 구조의 (중첩될 수도 있는) 단위입니다.
- 크레이트에는 암묵적이고 이름이 없는 최상위 모듈이 포함됩니다.
- 재귀 정의는 여러 모듈에 걸쳐 있을 수 있지만 여러 크레이트에는 걸칠 수 없습니다.
{{< /admonition  >}}


{{< admonition question "왜 모듈 파일을 정의하기 위해 크레이트 최상위에 mod를 넣어야 하나요? 그냥 use로 지정하면 안 되나요?" >}}

Rust에서 모듈은 제자리에 선언하거나 다른 파일에서 선언할 수 있습니다. 각각의 예제는 다음과 같습니다.


```rs
// main.rs에서
mod hello {
    pub fn f() {
        println!("hello!");
    }
}

fn main() {
    hello::f();
}
```

```rs
// main.rs에서
mod hello;

fn main() {
    hello::f();
}

// hello.rs에서
pub fn f() {
    println!("hello!");
}
```

첫 예제에서 모듈은 모듈이 사용되는 곳과 같은 파일에 정의되어 있습니다. 둘째 예제에서 메인 파일의 모듈 선언은 컴파일러에게 hello.rs나 hello/mod.rs를 찾아 보고 그 파일을 읽으라고 말해 줍니다.

**mod와 use의 차이를 주목하세요. mod는 모듈이 존재한다고 선언하지만, use는 다른 곳에 선언된 모듈을 참조하여 그 내용물을 현재 모듈의 범위 안에 가져 옵니다.**
{{< /admonition  >}}

---

## 7. 다중 플랫폼

{{< admonition question "Rust를 안드로이드 및 iOS 프로그래밍에 쓸 수 있나요?" >}}
_네 할 수 있습니다! 이미 Rust를 안드로이드와 iOS에서 사용하는 예제가 있습니다._

- android: https://github.com/rust-mobile/ndk
- ios
{{< /admonition  >}}


{{< admonition question "개인적인 Rust 프로그램을 웹 브라우저에서 실행할 수 있나요?" >}}
_아마도요. Rust는 `asm.js`와 `WebAssembly` 모두를 실험적으로 지원합니다._

- [WebAssembly의 동향에 대한 유용한 글, Naver D2](https://d2.naver.com/helloworld/8257914)
{{< /admonition  >}}

--- 

## 8. 저수준

{{< admonition question "Rust가 메모리 상에 값이 어떻게 배치될 지가 고정되어 있나요?" >}}
**기본적으로는 아닙니다.**
일반적으로 `enum`과 `struct`의 배치는 정의되지 않습니다. 따라서 컴파일러가 패딩을 구분값(discriminant)을 넣는데 재사용하거나, 중첩된 enum들의 변종(variant)들을 압축하거나, 패딩을 없애기 위해 필드를 재배치하는 등의 잠재적인 최적화를 할 수 있게 됩니다. 데이터를 들고 있지 않은 (“C와 비슷한”) enum은 정의된 표현을 가지도록 할 수 있습니다. 이러한 enum은 데이터를 들고 있지 않은 이름들만의 단순 목록이므로 쉽게 구분할 수 있습니다:

```rs
enum CLike {
    A,
    B = 32,
    C = 34,
    D
}
```
이러한 enum에 `#[repr(C)]` 속성을 적용하면 대응되는 C 코드가 가질 표현과 같은 표현이 되도록 할 수 있습니다. 

따라서 `FFI`(foreign function interface) 코드에서 C enum이 쓰일 대부분의 상황에서 Rust enum을 쓸 수 있습니다. **마찬가지로 struct에도 이 속성을 적용하면 C struct가 가질 배치와 같은 배치가 되도록 할 수 있습니다.**
{{< /admonition  >}}


## 9. 디버깅 / Tool


{{< admonition question "Rust 프로그램은 어떻게 디버깅하나요?" >}}
Rust 프로그램은 C나 C++와 같이 `gdb`나 `lldb`로 디버깅할 수 있습니다. 

사실은 모든 Rust 설치과정에는 (플랫폼 지원에 따라) `rust-gdb`나 `rust-lldb` 둘 중 하나가 함께 들어 있습니다. 이들은 gdb와 lldb에 Rust 값을 보기 좋게 출력해 주도록 감싼 것입니다.
{{< /admonition  >}}

{{< admonition question "rustc가 표준 라이브러리 코드에서 패닉(panic)이 일어났다고 하는데, 제 코드의 실수를 어떻게 찾을 수 있을까요?" >}}
이 오류는 보통 사용자 코드에서 None이나 Err을 `unwrap()`해서 일어납니다. `RUST_BACKTRACE=1` 환경 변수를 설정해서 스택 추적(backtrace)을 켜는 게 더 많은 정보를 얻는데 도움이 됩니다. 

디버그 모드로 컴파일하거나(`cargo build의 기본값`), 함께 들어 있는 `rust-gdb나` `rust-lldb` 같은 디버거를 쓰는 것도 도움이 됩니다.
{{< /admonition  >}}


- `rust`의 `gofmt` 같은 포맷팅 툴: [rustfmt](https://github.com/rust-lang/rustfmt)


## 10. 오류

{{< admonition question "Rust에는 왜 예외(exception)가 없나요?" >}}
예외는 제어 흐름을 이해하기 복잡하게 만들고, 타입 시스템을 넘어서는 유효성/무효성을 표현하며, (Rust의 주요 촛점인) 멀티스레딩된 코드와 잘 상호작용하지 않습니다.

**Rust는 오류 처리에 타입 기반의 접근을 선호하며, 이는 Rust의 제어 흐름, 동시성 및 여타 다른 것들에 더 잘 맞아 들어 갑니다.**
{{< /admonition  >}}


{{< admonition question "여기 저기 보이는 `unwrap()`를 어떻게 할 수 없나요?" >}}
`unwrap()`은 `Option`이나 `Result` 안에 있는 값을 뽑아 내고 아무 값도 없으면 패닉을 일으키는 함수입니다.

`unwrap()`이 잘못된 사용자 입력 따위의 **"예상할 수 있는 오류들을 다루는 기본 방법"** 이 되어서는 안 됩니다. 
현업 코드에서 이는 값이 비어 있지 않으며 만에 하나 비어 있다면 프로그램이 깨지는 **단언(assertion)처럼 취급되어야 합니다.**

또한 **unwrap()은 아직 오류를 처리하고 싶지 않은 빠른 프로토타입이나, 오류 처리가 주요 논점을 흐릴 수 있는 곳**에서도 유용합니다.
{{< /admonition  >}}


{{< admonition question "모든 곳에 `Result를` 쓰는 것 말고 더 쉽게 오류를 처리할 방법이 없나요?" >}}
다른 사람의 코드에 있는 `Result`를 처리하지 않는 방법을 원한다면 항상 unwrap()를 쓸 수 있지만, 아마도 원하는 게 아닐 겁니다. 
**Result는 어떤 계산이 성공적으로 끝나거나 끝나지 않을 수 있다는 표시입니다.** 

이러한 실패를 처리하도록 요구하는 건 Rust가 튼튼한 코드를 권장하는 방법 중 하나입니다. Rust는 실패를 더 편리하게 처리할 수 있도록 try! 매크로 같은 도구를 제공합니다.
정말로 오류를 처리하고 싶지 않다면 unwrap()를 쓰세요. **하지만 이렇게 하면 실패시 코드가 패닉을 일으키고, 보통 이는 프로세스를 종료시킨다는 점을 유의하시길 바랍니다.**

{{< /admonition  >}}


## 10. 숫자

{{< admonition question "부동 소숫점 계산을 할 때 f32와 f64 중 어느 쪽을 선호해야 하나요?" >}}
_잘 모르겠으면 정밀도를 우선시해서 f64를 선택하세요._
{{< /admonition  >}}

{{< admonition question "실수들을 비교하거나 HashMap 및 BTreeMap의 키로 쓸 수 없는 이유는 뭔가요?" >}}

(us 문서 기준 개인 해석입니다.)

Floats(실수)타입은 `total ordering`과 `totla equality`가 없습니다. 그렇기 때문에 `Eq` trait를 구현할 수 없습니다.

`total ordering`과 `total equality`가 Float형에 존재하지 않는 이유는 `floating-point`값인 `NaN`이 다른 floating point 값들 그리고 그 자신과도 **비교불가**하기 때문입니다.

(그래서) `float`타입이 `Eq`(total equality 구현체)와 `Ord`(total order 구현체)를 구현하지 못하기 때문에, 이 두 trait를 사용하는 타입들에서 사용될 수 없습니다. 해당 타입에는 `BTreeMap`, `HashMap` 또한 포함됩니다.

두 Map 타입에서는 `key`값에 total ordering 또는 total equality가 필요합니다.

(3rd 인) 이 [ordered-float](https://crates.io/crates/ordered-float) 크레이트에서는 `f32`, `f64` 타입들을 랩핑하여 Ord와 Eq 구현체들을 제공합니다. 그러므로 특정한 상황에서 유용하게 사용할 수 있습니다.
{{< /admonition  >}}

{{< admonition question "How can I convert between numeric types?" >}}
두 가지 방법이 있는데, 하나는 `as 예약어`로 원시 타입 사이에서 간단한 변환을 하는 것이고, 다른 하나는 `Into`와 `From` 트레이트를 써서 타입 변환을 하는 것입니다(직접 트레이트를 구현해서 변환을 추가할 수도 있습니다). 

`Into`와 `From` 트레이트는 변환에서 손실이 일어나지 않을 때만 구현되어 있습니다. 이를테면 `f64::from(0f32)`는 컴파일이 되지만 `f32::from(0f64)`는 아닙니다. 한편 `as`는 원시 타입들 사이에서는 모두 변환이 가능하며 필요하다면 값을 잘라냅니다(truncating).
{{< /admonition  >}}

## 11. 문자열

{{< admonition question "`&str`과 `String`에는 어떤 차이가 있나요?">}}
`String`은 힙에 할당된 UTF-8 바이트를 소유하는 버퍼입니다. 변경 가능한 `String`은 수정할 수 있고 필요에 따라 그 용량(capacity)을 늘릴 수 있습니다. `&str`은 다른 데 (보통 힙에) 할당되어 있는 `String`으로부터 참조된 슬라이스나, 문자열 리터럴의 경우 정적 메모리를 가리키는, 용량이 고정된 "view"입니다.
`&str`은 Rust 언어가 구현하는 원시 타입이지만 `String`은 표준 라이브러리에 구현되어 있습니다.
{{< /admonition  >}}

{{< admonition question "String의 각 문자를 O(1), 즉 상수 시간에 접근하려면 어떻게 해야 하나요?" >}}
**불가능합니다.**

Rust string들은 `UTF-8`로 인코딩되어 있습니다. `UTF-8`에서 `A single visual character`(보여지는 하나의 문자)는 반드시 `single byte`가 아닐 수도 있습니다(ASCII-encoded string일 경우).

각 바이트는 `“code unit”` (in UTF-16, code units are 2 bytes; in UTF-32 they are 4 bytes)라고 불립니다. `“Code points”`는 하나 또는 2이상의 `code unit`들로 구성되어 있으며, 
문자를 가장 가까이 근사한다고 할 수 있는 `“자소(grapheme) 클러스터”`는 여러 개의 `코드 포인트`로 구성되어 있습니다.

따라서 UTF-8 문자열에서 바이트를 인덱싱할 수 있다 하더라도 상수 시간에 i번째 코드포인트나 자소 클러스터를 얻어낼 수는 없습니다. **하지만 원하는 `코드포인트`나 `grapheme 클러스터`가 어느 바이트에서 시작하는지 안다면 그건 상수 시간에 접근할 수 있습니다. `str::find()`나 `정규식 검색 결과`는 `바이트 인덱스`를 반환하므로 이 방법으로 접근하는 게 가능합니다.**
{{< /admonition  >}}

{{< admonition question "Why are strings `UTF-8` by default?" >}}
`str` 타입이 `UTF-8`인 것은 현실에서, 특히 엔디안이 정해져 있지 않은 네트워크 전송에서 이 인코딩이 널리 쓰이기 때문이고, I/O를 할 때 어느 방향에서도 코드포인트를 다시 변환할 필요가 없는 것이 최선이라고 생각하기 때문입니다.

**물론 이는 문자열 안의 특정 유니코드 코드포인트의 위치를 찾는데 O(n) 연산이 필요하다는 뜻이긴 합니다.** 이미 시작하는 바이트 인덱스를 알고 있을 경우에는 예상대로 O(1) 시간이 걸리겠지만요. 어떻게 보면 바람직하지 않을 수도 있지만, 어떻게 보면 이 문제 자체가 트레이드오프로 가득 차 있기에 다음 중요한 점들을 지적할 필요가 있겠습니다:

**`str`에서 `ASCII 영역`의 코드포인트를 훑는 건 바이트 단위로 안전하게 할 수 있습니다.** 예를 들어 `.as_bytes()`를 쓸 경우 `u8`타입(8-bit unsigned int type)을 얻는 건 `O(1)` 연산이며 이 값은 `ASCII` 범위의 char로 변환하거나 비교할 수 있습니다. 그러니 이를테면 `'\n'`로 줄 바꿈을 찾는다면 바이트 단위로 검색해도 됩니다. UTF-8은 원래부터 이렇게 설계되었거든요.

**대부분의 “문자 기반” 텍스트 연산들은 “ASCII 범위의 코드포인트 한정” 같이 매우 제약된 언어 가정이 있어야만 동작합니다.** ASCII 범위를 벗어나면 언어학적인 단위들(glyphs[글리프], 낱말, 문단)의 경계를 찾기 위해 (상수 시간이 아닌) 복잡한 알고리즘을 써야 하기 마련입니다. 저희는 언어학적으로 올바르며 "honest"하다고 **유니코드에서 인증한 알고리즘을 권장**합니다.

`char` 타입은 `UTF-32`입니다. 한 번에 한 코드포인트를 들여다 보는 알고리즘이 정말로 필요하다고 생각한다면 `type wstr = [char]`을 정의하여 `str`로부터 한번에 읽어들인 뒤 `wstr`에서 연산을 하면 됩니다. **다르게 말하면, 언어가 “기본적으로 UTF-32로 디코딩하지 않는다”고 해서 UTF-32로 디코딩하거나 다시 인코딩하는 것 자체가 불가능한 건 아니라는 말입니다.**

왜 UTF-8이 UTF-16이나 UTF-32보다 보통 더 선호되는지 자세한 설명을 원한다면 [UTF-8 Everywhere manifesto](http://utf8everywhere.org/)를 읽어 보시길 바랍니다.
{{< /admonition  >}}

{{< admonition question "어떤 문자열 타입을 써야 하죠?" >}}
Rust는 네 쌍의 문자열 타입이 있고 각각 다른 역할을 합니다. 각 쌍마다 “소유된” 문자열 타입과 “슬라이스” 문자열 타입이 따로 있고, 다음과 같이 구성되어 있습니다

<center>

![](/images/rust_string_types.png)

</center>

Rust의 서로 다른 문자열 타입은 각자 다른 목적을 가집니다. 

- String과 str은 UTF-8로 인코딩된 일반 목적의 문자열입니다. - OsString과 OsStr은 현재 플랫폼에 맞춰 인코딩되어 있고 운영체제와 상호작용할 때 쓰입니다. 
- CString과 CStr은 C 문자열의 Rust 버전으로 FFI 코드에 사용됩니다.
- PathBuf와 Path는 OsString과 OsStr에 편의를 위해 경로 조작을 위한 메소드들을 추가한 것입니다.
{{< /admonition  >}}

{{< admonition question "&str와 String을 동시에 받는 함수를 어떻게 짤 수 있나요?" >}}

- 함수가 소유된 문자열을 필요로 하지만 아무 문자열 타입이나 받고 싶다면, `Into<String>` 제약을 쓰세요.
- 함수가 문자열 슬라이스를 필요로 하지만 아무 문자열 타입이나 받고 싶다면, `AsRef<str>` 제약을 쓰세요.
- 함수가 문자열 타입에 대해 신경쓰지 않고 두 가능성을 일관되게 처리하고 싶다면, 입력 타입으로 `Cow<str>`을 쓰세요.

{{< /admonition  >}}

#### `Into<String>`

이 예제에서 함수는 소유된 문자열과 문자열 슬라이스를 둘 다 받으며, 어느 쪽인지에 따라 함수 몸체 안에서 아무 일도 하지 않거나 입력을 소유된 문자열로 변환합니다. 참고로 변환은 명시적으로 해야 하며 안 그러면 변환되지 않을 것입니다.

```rs
fn accepts_both<S: Into<String>>(s: S) {
    let s = s.into();   // s를 `String`으로 변환합니다.
    // ... 함수의 나머지 내용
}
```

#### `AsRef<str>`

이 예제에서 함수는 소유된 문자열과 문자열 슬라이스를 둘 다 받으며, 어느 쪽인지에 따라 아무 일도 하지 않거나 입력을 문자열 슬라이스로 변환합니다. 이는 입력을 참조로 받아서 다음과 같이 자동으로 일어나게 할 수 있습니다:

```rs
fn accepts_both<S: AsRef<str>>(s: &S) {
    // ... 함수의 몸체
}
```


#### `Cow<str>`

이 예제에서 함수는 Cow<str>을 받는데, 이는 일반화된 타입이 아니라 컨테이너로서 필요에 따라 소유된 문자열이나 문자열 슬라이스를 담을 수 있습니다.

```rs
fn accepts_cow(s: Cow<str>) {
    // ... 함수의 몸체
}
```

## 12. 컬렉션

{{< admonition question "How can I iterate over a collection without moving/consuming it?" >}}

가장 쉬운 방법은 컬렉션의 `IntoIterator` 구현체를 사용하는 겁니다. 
`IntoIterator`는 `&Vec`과 `&mut Vec`에 구현되어 있으며, 아래는 `&Vec`를 활용해 이를 설명합니다.

```rs
let v = vec![1,2,3,4,5];
for item in &v {
    print!("{} ", item);
}
println!("\nLength: {}", v.len());
```
Rust의 for 반복문은 반복하고자 하는 대상에 대해 (`IntoIterator` 트레이트에 정의된) `into_iter()`를 호출합니다. `IntoIterator` 트레이트를 구현하는 모든 value들은 for 반복문에서 사용될 수 있습니다. 

**`into_iter()`는 컬렉션을 옮기거나 소모하는 것이 아니라, 그 내용물을 빌리도록 합니다. 다른 표준 컬렉션에 대해서도 똑같은 관계가 성립합니다.**

만약 옮기거나 소모하는 반복자가 필요하다면 for 반복문에서 반복할 때 **`&`나 `&mut` 없이 쓰세요.**

빌리는 반복자를 직접 접근하고 싶다면 보통 **iter() 메소드**를 써서 얻을 수 있습니다.

{{< /admonition  >}}

{{< admonition question "Why do I need to type the array size in the array declaration?" >}}

**꼭 그럴 필요 없습니다.** 

배열을 직접 선언한다면 원소의 갯수로부터 크기가 추론됩니다. 하지만 고정된 크기의 배열을 받는 함수를 선언한다면 컴파일러가 배열이 얼마나 클 지를 알아야 합니다.

하나 짚고 넘어가야 하는 게 있는데, **Rust는 현재 서로 다른 크기의 배열에 대해 `generics`를 지원하지 않습니다. 만약 갯수가 바뀔 수 있는 값들의 연속된 컨테이너를 받고자 한다면 (소유권이 필요하냐 마냐에 따라) `Vec`이나 `slice`를 사용하세요.**
{{< /admonition  >}}

## 13. 소유권
> 추가로 다시 봐야할 듯 싶어 [link](https://prev.rust-lang.org/en-US/faq.html#ownership)를 적어둡니다.

{{< admonition question "What is the difference between `passing by value`, `consuming`, `moving`, and `transferring ownership`?" >}}

**다 같은 뜻입니다.**

네 가지 경우에서 모두, 값이 새 소유자에게 옮겨가고, 원 소유자가 소유를 잃어버려 더 이상 쓸 수 없게 됩니다. 단 만약 타입이 Copy 트레이트를 구현한다면 원 소유자의 값은 무효화되지 않아 계속 쓸 수 있습니다.

{{< /admonition  >}}

{{< admonition question "왜 어떤 타입은 함수에 넘긴 뒤에도 재사용할 수 있지만 다른 타입은 그렇지 않나요?" >}}

타입이 `Copy 트레이트`를 구현하면 함수에 전달될 때 복사됩니다. **Rust의 모든 numeric type들은 Copy를 구현하지만**, struct type들은 기본적으로 Copy를 구현하지 않기 때문에 대신 `이동`이 일어납니다. 


**This means that the struct can no longer be used elsewhere, unless it is moved back out of the function via the return.**

(즉 구조체는 함수에서 다시 반환되거나 하지 않는 한 더 이상 다른 데서 사용할 수 없게 됩니다.)


{{< /admonition  >}}

## 14. Lifetimes(수명)

{{< admonition question "Why lifetimes?" >}}
`lifetime`은 메모리 안전성에 대한 Rust의 해답입니다. 

Rust는 수명을 사용해 쓰레기 수거(garbage collection)의 성능 비용 없이 메모리 안전성을 보장합니다. 
{{< /admonition  >}}

{{< admonition question "Why is the lifetime syntax the way it is?" >}}
`'a` 문법은 `ML(meta language) 계열`의 프로그래밍 언어에서 따 왔는데, 여기서 `'a` 문법은 `generic type parameter`를 나타내는 데 사용됩니다. Rust의 경우 수명 문법은 모호하지 않고, 눈에 띄어야 했으며 타입 선언에서 트레이트와 참조와 함께 쓰기 좋아야 했습니다. 다른 문법도 의논되었으나 이보다 확실히 더 좋은 문법이 제시되진 않았습니다.
{{< /admonition  >}}

{{< admonition question "Why do some references have lifetimes, like `&'a T`, and some do not, like `&T`?" >}}

사실 모든 참조 타입에는 수명이 있지만, 대부분의 경우 직접 쓸 필요가 없습니다. 규칙은 다음과 같습니다:

1. `function body`에서는 수명을 명시적으로 쓸 필요가 전혀 없습니다.(항상 올바른 값이 추론될 것입니다).
2. `function signature` (예를 들어 인자 타입이나 반환 타입) 에서는 수명을 명시적으로 써야 할 수도 있습니다. 여기에서는 `수명 탈락(elision)`라는 간단한 기본값이 적용되는데 이는 다음 세 규칙으로 구성되어 있습니다
    1. function argument들에서 탈락된(elided) 각각 수명들은 서로 다른 인자가 됩니다.
    2. 입력된 수명이 하나 뿐이면, 그게 탈락되었든 아니든 그 함수의 `return value들`의 all elided lifetimes(탈락된 수명들)에 할당됩니다.
    3. 입력 수명이 여럿 있지만 그 중 하나가 `&self`거나 `&mut self`라면, `self`의 lifetime이  all elided output lifetimes(모든 탈락된 출력 수명들)에 할당됩니다.
3. 마지막으로 struct와 enum 정의에서는 모든 수명이 명시적으로 선언되어야 합니다.


만약 이 규칙이 컴파일 에러(compilation errors)를 일으킨다면, Rust 컴파일러는 일어난 에러를 가리키는 메시지를 제공하며 그 에러가 일어난 inference process(추론 단계)에 따라 필요한 solution을 제시할 것입니다.

{{< /admonition  >}}


{{< admonition question "How can Rust guarantee no `null pointers` and no `dangling pointers`(유령 포인터)?" >}}
`&Foo`와 `&mut Foo` 타입의 값을 만드는 유일한 방법은 `이미 존재하는 Foo 타입의 값을 reference가 point하는 값으로 specify하는 것 뿐입니다. 

reference는 region of code(즉, the lifetime of the reference) 안에서 원래 값을 “빌리며”, reference가 값을 "빌리는" 동안에는 original value를 옮기거나 소멸시킬 수 없습니다.

{{< /admonition  >}}

{{< admonition question "How do I express the absence of a value without null?" >}}

You can do that with the `Option type`, which can either be `Some(T)` or `None`. `Some(T)` indicates that a value of type T is contained within, while `None` indicates the absence of a value.
{{< /admonition  >}}

## 15. 제너릭

`Generic`에 대해서는 복잡한 내용들이 많아 링크로 대체합니다.

- [Generic: en](https://prev.rust-lang.org/en-US/faq.html#generics)
- [GenericL kr](https://prev.rust-lang.org/ko-KR/faq.html#generics)

