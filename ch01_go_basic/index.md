# [일주일 만에 배우는 GO] CH.1 Go Language Design


[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)를 기반으로`Go`에 녹아있는 배경/철학을 분석합니다. 이후 사용되는 필드들을 정리하고 학습에 도움되는 best practice 레퍼런스들을 정리합니다.

<!--more-->
<br />

## tl;dr

> Go is a compiled, concurrent, garbage-collected, statically typed language developed at Google

## Introduction

> 웹 서비스가 많아지고 컴퓨터 환경이 발전하면서 관리해야 할 코드와 서비스가 많아졌고 이를 효율적으로 관리하기 위해 go는 만들어졌습니다.

공식문서의 go 디자인 레퍼런스를 발견해, 이를 토대로 분석글을 작성합니다.[Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article) 추가로 [FAQ](https://go.dev/doc/faq)를 참조하였습니다.

{{< admonition note "go를 개발한 이유">}}
_The Go programming language was conceived in late 2007_

... 중략 ...

_mostly C++, Java, and Python, had been created_

... 중략 ...

_We were not alone in our concerns. After many years with a pretty quiet landscape for programming languages, Go was among the first of several new languages—Rust, Elixir, Swift, and more—that have made programming language development an active, almost mainstream field again._

... 중략 ...

_The problems introduced by multicore processors, networked systems, massive computation clusters, and the web programming model were being worked around rather than addressed head-on. Moreover, the scale has changed: today's server programs comprise tens of millions of lines of code, are worked on by hundreds or even thousands of programmers, and are updated literally every day. To make matters worse, build times, even on large compilation clusters, have stretched to many minutes, even hours._

{{< /admonition >}}

go는 2007년에 만들어졌습니다. 이 당시의 언어로써는 다음과 같은 문제들을 해결하기 어려웠기 때문에 go를 개발하기 시작했다고 보입니다.

1. multicore processors 고려
2. 개발자들의 수요 증가(코드의 scale 증가 -> 깔끔한 코드 니즈 증가)
3. 서버 scale increase
4. build time too long

이때는 참고로 Rust, Elixir, Swift 같은 대안 언어들이 없었다고 합니다.

{{< admonition note "go 디자인 목표">}}
_Go was designed and developed to make working in this environment more productive. Besides its better-known aspects such as built-in concurrency and garbage collection, Go's design considerations include rigorous dependency management, the adaptability of software architecture as systems grow, and robustness across the boundaries between components._
{{< /admonition >}}

- built-in concurrency
- garbage collection
- dependency management
- adaptability of software architecture

{{< admonition note "google at go">}}
_The goals of the Go project were to eliminate the slowness and clumsiness of software development at Google, and thereby to make the process more productive and scalable. The language was designed by and for people who write—and read and debug and maintain—large software systems._
{{< /admonition >}}

구글의 현존하는 수많은 legacy 코드들과 앞으로 늘어날 코드들을 효율적으로 관리하기 위해서 go를 개발하기 시작했다고 합니다.

## Pain points

아래는 go를 만들 때 고려되었던 `pain point`들입니다.

- slow builds
- uncontrolled dependencies
- each programmer using a different subset of the language
- poor program understanding (code hard to read, poorly documented, and so on)
- duplication of effort
- cost of updates
- version skew
- difficulty of writing automatic tools
- cross-language builds

## c.f go는 왜 { }를 사용하나요?

<center>

**안정성과 신뢰성을 위해 python style indentation 편의성을 포기하고 `{}`를 그래도 도입함.**

</center>

{{< admonition note "go는 왜 중괄호를 사용하나요? ">}}
_Our position is therefore that, although spaces for indentation is nice for small programs, it doesn't scale well, and the bigger and more heterogeneous the code base, the more trouble it can cause. It is better to forgo convenience for safety and dependability, so Go has brace-bounded blocks._
{{< /admonition >}}

구글이 가진 노하우에 따르면 작은 프로그램에서는 `python`의 indentation이 좋지만, 코드 기반이 클수록 & cross-language build 환경에서는 python 스타일이 좋지 않기 때문에 `{ }`를 그대로 쓰기로 결정했다고 합니다.

## Dependencies in Go

> go가 의존성을 대하는 전략에 대해서 정리합니다.

아래와 같은 c의 guard기능으로 인해 불필요하게 의존성을 읽어들이는 현상이 발생했습니다.

```c
/* Large copyright and licensing notice */

#ifndef _SYS_STAT_H_
#define _SYS_STAT_H_

/* Types and other definitions */
#endif
```

{{< admonition note >}}
_The first step to making Go scale, dependency-wise, is that the language defines that unused dependencies are a compile-time error (not a warning, an error)._

... 중략 ...

_This guarantees by construction that the dependency tree for any Go program is precise, that it has no extraneous edges. That, in turn, guarantees that no extra code will be compiled when building the program, which minimizes compilation time._
{{< /admonition >}}

구글 같이 코드 베이스가 많은 경우, 빌드타임이 오래걸리는 것은 업무의 효율성을 떨어뜨리기 때문에 **`go`에서는 사용하지 않는 dependency들에 대해서 compile-time error를 뱉어내도록 하고 있습니다.**

더 나아가서 `go`는 `include of include file`와 같은 상황에도 효율적인 전략을 취하고 있습니다.

```
package A imports package B;
package B imports package C;
package A does not import package C
```

다음과 같이 A에서 B를 직접 참조하고, C는 간접적으로 참조하는 상황이 있을 때, go 컴파일러는 2가지 전략을 취합니다.

1. **C -> B -> A 순으로 compile한다.**
2. **A에서 B를 참조할 때, 컴파일 시 생성된 B `Object 파일`에 B public interface가 영향받을 dependency들의 `type information`을 넣어둔다.**

{{< admonition  >}}
_In other words, when B is compiled, the generated object file includes type information for all dependencies of B that affect the public interface of B._
{{< /admonition >}}

예를 들어서 A에서 B안에 들어있는 `I/O package`를 import하고, B안에 들어있는 I/O 패키지는 C에 정의되어있는 `buffered I/O`를 사용해서 구현(implementation)되어있을 때,
C -> B -> A 순으로 컴파일되면서 link됩니다. 이때 A가 컴파일될 때 컴파일러는 B의 `Object file`을 읽어(주의! 소스코드를 읽는게 아님)들이는데, 이때 해당 파일에는 A에서 `import B`를 실행할 때 컴파일러에 필요한 필요한 모든 타입 정보가 들어있습니다.

> _This is, of course, reminiscent of the Plan 9 C (as opposed to ANSI C) approach to dependency management, except that, in effect, the compiler writes the header file when the Go source file is compiled_

이런 종속성 관리 접근 방식은 `Plan 9 C(ANSI C와 반대)` 접근 방식과 유사하며, 단 실제로 Go 소스 파일이 컴파일될 때 컴파일러가 헤더 파일을 작성해 준다는 점만 다릅니다.

{{< admonition  >}}
_To make compilation even more efficient, the object file is arranged so the export data is the first thing in the file, so the compiler can stop reading as soon as it reaches the end of that section._
{{< /admonition  >}}

참고로 컴파일 할 때 고언어 컴파일러는 export할 data를 Object file 맨 앞단에 위치시켜서 export 해줄 정보만 빠르게 찾아볼 수 있도록 구현되어 있다고 합니다.

{{< admonition  >}}
_Go places the export data in the object file; some languages require the author to write or the compiler to generate a second file with that information. That's twice as many files to open. In Go there is only one file to open to import a package. Also, the single file approach means that the export data (or header file, in C/C++) can never go out of date relative to the object file._
{{< /admonition  >}}

또한 export data를 object파일에 위치시키기 때문에 import package시킬 때, 컴파일러는 파일 당 하나의 object 파일만 읽어들이면 되며 부가적으로 관리해야 할 포인트가 하나로 줄어든다는 장점이 있습니다.

{{< admonition  >}}
_Another feature of the Go dependency graph is that it has no cycles._

...중략...

_The lack of circular imports causes occasional annoyance but keeps the tree clean, forcing a clear demarcation between packages. As with many of the design decisions in Go, it forces the programmer to think earlier about a larger-scale issue (in this case, package boundaries) that if left until later may never be addressed satisfactorily._
{{< /admonition  >}}

**마지막으로 go언어는 `circular import`(= `cyclic import`)를 컴파일 타임에 에러를 내 줌으로써 효율적인 `package boundaries`에 대해서 개발자가 고민할 수 있도록 해줍니다.**

이런 방식들을 통해 go는 기존 언어보다 획기적으로 build time을 줄일 수 있었습니다.

## Packages

{{< admonition  >}}
_It's important to recognize that package paths are unique, but there is no such requirement for package names. The path must uniquely identify the package to be imported, while the name is just a convention for how clients of the package can refer to its contents. The package name need not be unique and can be overridden in each importing source file by providing a local identifier in the import clause. These two imports both reference packages that call themselves package log, but to import them in a single source file one must be (locally) renamed:_
{{< /admonition >}}

- package의 path는 unique해야 합니다.
- package name은 unique할 필요없으며, 중복 될 경우 이를 사용하는(import) 영역에서 naming을 override해서 사용합니다.

## remote package

개인적으로 신기했던 부분인데, go에서는 import 해주는 package path가 `url`이 될 수도 있습니다.

예를 들어서 github에서 doozer 패키지를 가져오고 싶으면 아래와 같이 해주면 됩니다.

```shell
$ go get github.com/4ad/doozer // Shell command to fetch package
```

```go
import "github.com/4ad/doozer" // Doozer client's import statement

var client doozer.Conn         // Client's use of package
```

이렇게 한번 fetch(package를 install)해주면 그 다음부터는 일반적인 package와 마찬가지로 import 해주면 됩니다.

이런 방식을 도입함으로써 explicit하게 dependencies들을 보여줄 수 있게 되었습니다.

{{< admonition  >}}
_Also, the allocation of the space of import paths is delegated to URLs, which makes the naming of packages decentralized and therefore scalable, in contrast to centralized registries used by other languages._
{{< /admonition >}}

**또한 URL로 `allocation of the space of import paths`가 위임됨으로써, package의 naming이 decentralized & scalable하게 만들었습니다.** (저는 allocation of the space of import paths라는 뜻을 url의 unique한 장점을 `package path`로 그대로 가져와서 global하게 unique한 decentralized system을 만들었다고 이해했습니다.)

## Syntax

> _Go was therefore designed with clarity and tooling in mind, and has a clean syntax._

go declaration(선언) 문법중에서 C 스타일 프로그래머들을 놀라게 하는 부분이 있습니다.

- go style

```
var fn func([]int) int
type T struct { a, b int }
```

- c style

```c
int (*fn)(int[]);
struct T { int a, b; }
```

> The declared name appears before the type and there are more keywords.

이런 `Declarations introduced by keyword` 문법은 `Pascal`에 더 가깝다고 합니다. 이를 통해 개발자는 더 수월하게 코드 분석이 가능하며, `type syntax`를 가지는 것이 C언어 같이 `expression syntax` 보다 컴퓨터가 parsing 성능에 상당히 더 유리하다고 합니다.

**go는 `type syntax`라는 문법을 추가해서 코드가 늘어나지만, 이를 통해 모호성을 제거하였습니다.** 단 편의를 위해 go에서는 var 키워드를 삭제하고 `:=`라는 키워드를 관용적으로 사용합니다.

```go
var buf *bytes.Buffer = bytes.NewBuffer(x) // explicit
buf := bytes.NewBuffer(x) // derived
```

마지막으로 Go에서는 `default function arguments`를 의도적으로 누락시켰습니다. 기본 인자가 가지는 모호성으로 코드를 사용하는 부분이 명시적이지 못할 경우가 많다는 단점이 있기 때문입니다. 물론 같은 네이밍을 가질 함수가 가질 수 있는 모든 interface(function signature)를 구현해줘야 한다는 단점을 가지긴합니다.

아래는 go에서 함수 / 메서드를 표현하는 방법입니다. 코트린 처럼 `fun`을 키워드로 가졌으면 더 좋았을 것 같네요.

```go
func Abs(x T) float64      // function declaration
func (x T) Abs() float64   // method declaration
```

go는 `first-class` function / closures를 지원합니다.

```go
negativeAbs := func(x T) float64 { return -Abs(x)} // lambda
```

go는 multiple value return이 가능합니다.

```go
func ReadByte() (c byte, err error)

c, err := ReadByte()
if err != nil { ... }
```

{{< admonition  >}}
_Those functions all need separate names, too, which makes it clear which combinations exist, as well as encouraging more thought about naming, a critical aspect of clarity and readability._
{{< /admonition >}}

엇? java 처럼 오버로딩을 하는게 아니라 signature 바뀔때마다 모두 새로운 함수명을 만들어줘야한다는 의미일까요? 아무튼 go는 clarity를 가장 최우선으로 두는 것 같습니다.

{{< admonition tip >}}
개인적으로 default argument가 사라진 건 너무 좋은 것 같은게, python기준으로 함수가 여기저기 많이 사용될수록 default argument 때문에 불필요하게 사용되는 모든 코드들을 뒤져야 할 때가 많았습니다.

예를 들어

```python
def create_person(name, age=30):
   if age < 30:
   ...
   else:
   ...
```

이런 함수가 있고 이 함수 관련해서 에러가 나면 `create_person`를 호출하는 모든 코드들을 다 뒤져봐야 할 경우가 있었습니다. 특히 mongodb처럼 데이터의 스키마가 고정되지 않은 데이터를 함수의 dto로 받는 경우에는 정말 짜증납니다. (kotlin migration 과정에서 엄청 스트레스를 줬던 콘 `mcard`...)
{{< /admonition >}}

## Naming

`Go`에서는 신기하게도 `visibility of an identifier`를 대소문자로 구별합니다. (타 언어에서는 private / public따위의 키워드를 씀)

{{< admonition note "name as visibility identifier" >}}
_Go the name itself carries the information: the case of the initial letter of the identifier determines the visibility. If the initial character is an upper case letter, the identifier is exported (public); otherwise it is not:_
{{< /admonition  >}}

- upper case initial letter: Name is visible to clients of package
- otherwise: name (or `_Name`) is not visible to clients of package

이런 rule은 모든 곳에 적용됩니다.

- variables
- types
- functions
- methods
- constants
- fields

이런 특이한 전략을 취한데에는 naming으로 visibility를 관리하는 것이 identifier에 비해 더욱 깔끔한 Public api관리에 도움되는 전략이라고 판단한 google의 노하우가 반영되었다고 합니다.

개인적으로는 코드 검색을 할 때, identifier를 가지고 검색하면 관련 identifier 리스트들을 한번에 확인이 가능한 반면 네이밍으로 visibility를 관리하면 정확한 네이밍을 알아야 가능하다는 점에서 솔직히 별로 인 것 같습니다.

{{< admonition note "scope hierarchy" >}}
_Another simplification is that Go has a very compact scope hierarchy:_
{{< /admonition  >}}

- `universe` (predeclared identifiers such as int and string)
- `package` (all the source files of a package live at the same scope)
- `file` (for package import renames only)
- `function`
- `block`

{{< admonition note "naming scope" >}}
_There is no scope for name space or class or other wrapping construct. Names come from very few places in Go, and all names follow the same scope hierarchy: at any given location in the source, an identifier denotes exactly one language object, independent of how it is used. (The only exception is statement labels, the targets of break statements and the like; they always have function scope.)_

...중략...

_top-level predefined names such as int, (the first component of) every name is always declared in the current package._

{{< /admonition  >}}

go에서는 일부 예외(`statement labels`, `argets of break statements`는 function scope)를 제외하면 naming은 `package scope`로 관리됩니다.

{{< admonition note "package scope로 naming이 관리되는 이유" >}}
_exported name to a package can never break a client of that package. The naming rules decouple packages, providing scaling, clarity, and robustness._
{{< /admonition  >}}

{{< admonition note "function / method 오버로딩 불가능" >}}
_method lookup is always by name only, not by signature (type) of the method. In other words, a single type can never have two methods with the same name_
{{< /admonition  >}}

`java`와 달리 `go`에서는 function signature만 다르고 네이밍이 같은 메서드 / 함수를 만드는 것은 불가능합니다. 개발자가 네이밍을 좀 더 신경써야 된다는 불편함은 있겠지만, 개인적으로 더 깔끔하게 함수들을 관리할 수 있는 방법이라고 생각되네요 (왠지 function naming관련된 convention tip들이 있을 것 같음)

- [effective go about Functions](https://go.dev/doc/effective_go#functions)

## Semantics

go는 기본적으로 `c`와 많이 닮아 있지만, modern언어에 익숙한 개발자들을 위해 몇가지 차이점을 두었습니다.

- there is no pointer arithmetic
- there are no implicit numeric conversions
- array bounds are always checked
- there are no type aliases (after type X int, X and int are distinct types not aliases)
- ++ and -- are statements not expressions
- assignment is not an expression
- it is legal (encouraged even) to take the address of a stack variable

아래는 C, C++, and even Java와 비교했을 때, 크게 변화한 부분입니다. (참고로 go 초기 개발자인 Robert Griesemer. Java hotspot compiler(JVM)을 개발했었습니다. )

- concurrency
- gc
- interface type
- reflection
- type switches

## Concurrency

> **Go is not purely memory safe in the presence of concurrency. Sharing is legal and passing a pointer over a channel is idiomatic (and efficient).**

`Go`는 [Communicating sequential processes, CSP](https://en.wikipedia.org/wiki/Communicating_sequential_processes)의 `first-class channel`개념을 가져왔습니다. (which is about message passing via channels in concurrent system)

{{< admonition note "Don't communicate by sharing memory, share memory by communicating" >}}
_Some concurrency and functional programming experts are disappointed that Go does not take a write-once approach to value semantics in the context of concurrent computation, that Go is not more like Erlang for example. Again, the reason is largely about familiarity and suitability for the problem domain. Go's concurrent features work well in a context familiar to most programmers. Go enables simple, safe concurrent programming but does not forbid bad programming. We compensate by convention, training programmers to think about message passing as a version of ownership control._
{{< /admonition  >}}

Go는 concurrency context에 `write-once` 접근을 하지 않습니다. 이는 의도적으로 bad programming을 막지 않은 것인데요, convention으로 이런 bad practice를 방지하고, 프로그래머들이 message passing에 대해서 더 생각하도록 유도하기 위함이라고 합니다.

이런 철학은 go의 motto에서도 드러나 있습니다.

<center>

**"Don't communicate by sharing memory, share memory by communicating."**

</center>

## Garbage collection

go는 jvm 개발자가 있어서 그런지 c / c++ / rust와 달리 gc를 가져왔습니다.

{{< admonition note "interior pointer" >}}
_The X.buf field in the example above lives within the struct but it is legal to capture the address of this inner field, for instance to pass it to an I/O routine. In Java, as in many garbage-collected languages, it is not possible to construct an interior pointer like this, but in Go it is idiomatic._
{{< /admonition  >}}

java와 비교해 `interior pointers`(interior pointers to objects allocated in the heap) 기능을 추가해 커스텀하게 gc 기능이 동작하도록 적용했다고 합니다.

{{< admonition note "interior 디자인이 끼칠 영향" >}}
_This design point affects which collection algorithms can be used, and may make them more difficult, but after careful thought we decided that it was necessary to allow interior pointers because of the benefits to the programmer and the ability to reduce pressure on the (perhaps harder to implement) collector._

... 중략 ...

_The garbage collector remains an active area of development. The current design is a parallel mark-and-sweep collector and there remain opportunities to improve its performance or perhaps even its design. (The language specification does not mandate any particular implementation of the collector.) Still, if the programmer takes care to use memory wisely, the current implementation works well for production use._

{{< /admonition  >}}

듣기로는 gc기능이 퍼포먼스 이슈가 있어 Discord에서는 기존에 go로 짜여있는 코드들을 rust로 옮겼다고 하네요.

## Composition not inheritance

> _there is no type hierarchy_

go에는 기존 oop 언어들과 달리 `type hierarchy`가 없습니다.

{{< admonition note "interface" >}}
_In Go an interface is just a set of methods_

...중략...

_All data types that implement these methods satisfy this interface implicitly; there is no implements declaration. That said, **interface satisfaction is statically checked at compile time so despite this decoupling interfaces are type-safe.**_
{{< /admonition  >}}

자바와 달리 고의 interface는 behavior만 정의하며 subclassing이 없으므로 **상속이란 개념이 존재하지 않습니다.** 대신 `composition`(embedding)을 활용한다고 합니다.

{{< admonition tip "subclassing vsv subtyping" >}}
`go`에는 `subclassing`이 없다고 하는데요, subclassing이 무엇인지 그리고 subtyping 또한 무엇인지 비교해보겠습니다.

- 서브클래싱은 구현되어 있는 클래스를 상속하는 것
- 서브타이핑은 정의되어 있는 인터페이스를 구현하는 것

먼저 좀 더 Subclassing이란 ? Superr Class에 구현된 코드와 내부 표현 구조를 Sub Class(하위 클래스)가 이어받는 기능을 뜻합니다. 클래스 inheritance라고도 불리며, 이를 통해 하위클래스에서 슈퍼 클래스에 구현된 코드의 재사용이 가능합니다. 그렇기 때문에 sub class는 overriding을 통해 같은 이름의 비슷하지만 커스텀한 행동들을 정의할 수 있습니다.

이와 달리 Subtying이란, Super Type의 객체가 수행할 행동(behavior only)의 약속(프로토콜, api)를 Sub Type이 이어 받습니다.
행동들을 공통된 타입으로 묶어 runtime에 super type의 객체의 타입으로 sub type을 대체가능하도록 합니다. 이를 통해 프로그램 변경에 대한 영향을 최소화 할 수 있습니다. 즉 core한 behavior들을 공통적으로 관리 가능합니다.
{{< /admonition  >}}

{{< admonition note "go가 inheritance를 버린 이유" >}}
_**that the behavior of data can be generalized independently of the representation of that data. The model works best when the behavior (method set) is fixed, but once you subclass a type and add a method, the behaviors are no longer identical.** If instead the set of behaviors is fixed, such as in Go's statically defined interfaces, the uniformity of behavior enables data and programs to be composed uniformly, orthogonally, and safely._
{{< /admonition  >}}

이런 전략을 취했던 이유는 behavior가 fix되어 코드가 작성되면 data representation을 담당하는 model이 works best한다는 철학이 녹아들어있다고 합니다.

TODO: 글을 읽다보니 composition에 대해서는 어느정도 이해가 되는데, 이런 철학이 왜 고려되어야 하는지가 정확하게 와닫지는 않는 것 같아서 나머지 부분은 실제 코드를 만져보고 다시 읽어보려고 합니다.

[Composition not inheritance](https://go.dev/talks/2012/splash.article#TOC_15.)

## Errors

Go에는 일반적인 의미의 예외 기능이 없습니다. 즉, 오류 처리와 관련된 제어 구조가 없습니다.

{{< admonition note "go가 error를 보는 방식" >}}
_Errors are just values and programs compute with them as they would compute with values of any other type_
{{< /admonition  >}}

{{< admonition note "go가 error를 value로 취급하는 첫번째 이유" >}}
_First, there is nothing truly exceptional about errors in computer programs. For instance, the inability to open a file is a common issue that does not deserve special linguistic constructs; if and return are fine._

```go
f, err := os.Open(fileName)
if err != nil {
    return err
}
```

{{< /admonition  >}}

go 철학에서는 error를 특별한 예외라고 생각할 필요가 전혀 없다고 생각하기 떄문입니다. 그냥 value가 return되고 if 분기로 이를 핸들링해주면 그만이라고 주장합니다.

{{< admonition note "go가 error를 value로 취급하는 두번째 이유" >}}
_There is no question the resulting code can be longer, but the clarity and simplicity of such code offsets its verbosity. Explicit error checking forces the programmer to think about errors—and deal with them—when they arise._
{{< /admonition  >}}

결과 코드가 더 길어질 수 있다는 점에는 의심의 여지가 없지만 그러한 코드의 명확성과 단순성은 장황함을 상쇄합니다. 명시적 오류 검사는 프로그래머가 오류에 대해 생각하고 오류가 발생할 때 처리하도록 합니다.

## Useful references

- [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)
- [golang cheat sheet](https://github.com/a8m/golang-cheat-sheet)
- [gin](https://github.com/gin-gonic/gin)

## conclustion

이상으로 go의 디자인 철학에 대해서 분석해보았습니다. 개인적으로 하루정도를 투자하려 했지만, 실제로는 와닿지 않는 내용들 때문에 시간이 조금 더 지체되었던 것 같네요 (소요시간: 대략 2일)

<center>- 끝 -</center>

