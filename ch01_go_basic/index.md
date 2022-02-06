# [일주일 만에 배우는 GO] CH.1 Go Language Design


`Go` 공식문서에 녹아있는 배경/철학을 분석합니다. 이후 사용되는 필드들을 정리하고 학습에 도움되는 best practice 레퍼런스들을 정리합니다.

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

이런 함수가 있고 이 함수 관련해서 에러가 나면 `create_person`를 호출하는 모든 코드들을 다 뒤져봐야 할 경우가 있었습니다. 특히 mongodb처럼 데이터의 스키마가 고정되지 않은 데이터를 함수의 dto로 받는 경우에는 정말 짜증납니다. (카카오콘 mcard...)
{{< /admonition >}}

## Useful references

- [Go at Google: Language Design in the Service of Software Engineering](https://go.dev/talks/2012/splash.article)
- [golang cheat sheet](https://github.com/a8m/golang-cheat-sheet)
- [gin](https://github.com/gin-gonic/gin)

## conclustion

<center>- 끝 -</center>

