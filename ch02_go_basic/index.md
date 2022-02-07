# [일주일 만에 배우는 GO] CH.2 Go Basic


[Effective go](https://go.dev/doc/effective_go)를 통해 go 핵심 문법을 빠르게 배워보겠습니다.

<!--more-->
<br />

## Formatting

> `gofmt` 패키지

아래 코드는 소스 파일이 아닌 패키지 레벨에서 실행됩니다.

```shell
$ go fmt
```

{{< admonition note "들여쓰기" >}}
_들여쓰기를 위해 탭(tabs)을 사용하며, gofmt는 기본값으로 탭을 사용한다. 만약 꼭 써야하는 경우에만 스페이스(spaces)를 사용하라._
{{< /admonition  >}}

{{< admonition note "한 줄 길이" >}}
_Go는 한 줄 길이에 제한이 없다. 길이가 길어지는것에 대해 걱정하지 마라. 만약 라인 길이가 너무 길게 느껴진다면, 별도의 탭을 가지고 들여쓰기를하여 감싸라_
{{< /admonition  >}}

{{< admonition note "괄호" >}}
_Go는 C와 Java에 비해 적은 수의 괄호가 필요하다. 제어 구조들(`if`, `for`, `switch`)의 문법엔 괄호가 없다._
{{< /admonition  >}}

## Comment

> godoc

{{< admonition note "package 주석" >}}
_패키지에서 최상위 선언의 바로 앞에있는 주석이 그 선언의 문서주석으로 처리된다. 패키지 내부에서 최상위 선언 바로 이전의 주석은 그 선언을 위한 문서주석이다. 프로그램에서 모든 외부로 노출되는 (대문자로 시작되는) 이름은 문서주석이 필요하다._
{{< /admonition  >}}

첫 문장은 선언된 이름으로 시작하는 한 줄짜리 문장으로 요약되어야 한다.

```go
// Compile parses a regular expression and returns, if successful,
// a Regexp that can be used to match against text.
func Compile(str string) (*Regexp, error) {
```

패키지는 각각의 문서 주석을 패키지명과 함께 시작하기 때문에 만약 아래와 같은 명령을 터미널에 활용하면 효율적이다.

```shell
$ godoc regexp | grep parse
    Compile parses a regular expression and returns, if successful, a Regexp
    parsed. It simplifies safe initialization of global variables holding
    cannot be parsed. It simplifies safe initialization of global variables
```

## Names

### Package Name

`Go`에서는 이름의 첫 문자가 대문자인지 아닌지에 따라서 이름의 패키지 밖에서의 노출여부가 결정된다.

{{< admonition note "package name" >}}
_패키지가 임포트되면, 패키지명은 패키지 내용들에 대한 접근자가 된다._

```go
import "bytes"

bytes.Buffer // usage
```

{{< /admonition  >}}

- 관례적으로, 패키지명은 소문자, 한 단어로만 부여하며 언더바(`_`)나 대소문자 혼용에 대한 필요가 없어야한다.

- 또 다른 규칙은 패키지명은 소스 디렉토리 이름 기반이라는 것이다. **src/encoding/base64에 있는 패키지는 encoding/base64로 임포트가 된다. base64라는 이름을 가지고 있지만, encoding_base64나 encodingBase64를 쓰지 않는다.**

- `import .`표현을 사용하지 말라.

{{< admonition tip "Package Naming Convention" >}}
_bufio 패키지에 있는 버퍼 리더는 BufReader가 아닌 Reader로 불린다. 왜냐하면 사용자는 이를 bufio.Reader로 보게되며, 이것이 더 명확하고 간결하기 때문이다. 게다가 임포트된 객체들은 항상 패키지명과 함께 불려지기 때문에 bufio.Reader는 io.Reader와 충돌하지 않는다._

... 중략 ...
_Go에 존재하는 ring.Ring이라는 구조체의 인스턴스를 만드는 함수는 보통은 NewRing으로 불릴테지만, Ring은 패키지 밖으로 노출된 유일한 타입이며, 패키지가 ring으로 불리기 때문에, 이 함수는 그냥 New라고 부르고 ring.New와 같이 사용한다._
{{< /admonition  >}}

{{< admonition note "Comment is better than long naming" >}}
_또 다른 간단한 예시는 once.Do이다. once.Do(setup)는 읽기가 쉬우며 once.DoOrWaitUntilDone(setup)으로 개선될게 없다. 긴 이름은 좀 더 쉽게 읽는것을 방해한다. 문서에 주석을 다는것이 긴 이름을 사용하는 것보다 더 좋을 것이다._

클린코드에서는 네이밍을 길게 가져가고 주석을 없애는, 즉 코드로 설명이 가능하도록 하자는 적략을 취하는데 Effective go를 쓴 저자는 무조건 짧은게 최고다는 느낌을 준다.

이러다 보니 history 차원이 아닌 설명을 위한 주석을 달때는 죄의식을 느꼈는데, 개인적으로 미니멀리즘을 좋아하니 go 방식이 더 끌리는 것 같다.
{{< /admonition  >}}

### Getter and Setter

Go는 getters와 setters를 자체적으로 제공하지 않는다.

{{< admonition note "getter and setter naming" >}}
_getter의 이름에 Get을 넣는건 Go언어 답지도, 필수적이지도 않다. 만약 owner(첫 문자가 소문자이며 패키지 밖으로 노출되지 않는다.)라는 필드를 가지고 있다면 getter 메서드는 GetOwner가 아닌 Owner(첫 문자가 대문자이며, 패키지 밖으로 노출됨)라고 불러야한다._

... 중략 ...

_만약 필요하다면, setter 함수는 SetOwner라고 불릴 것이다_

{{< /admonition  >}}

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### Interface

관례적으로, 하나의 메서드를 갖는 인터페이스는 메서드 이름에 -er 접미사를 붙이거나 에이전트 명사를 구성하는 방식을 사용한다.

- Reader
- Writer
- Formatter
- CloseNotifier
- ...

### MixedCaps

go는 camelCase를 사용한다.

## Semicolons

C언어 처럼, Go의 정식문법은 구문을 종료하기 위하여 세미콜론을 사용한다. 하지만 C언어와는 달리 세미콜론은 소스상에 나타나지 않는다. 대신 구문분석기(lexer)는 간단한 규칙을 써서 스캔을 하는 과정에 자동으로 세미콜론을 삽입한다. **그래서 소스작성시 대부분 세미콜론을 사용하지 않는다.**

{{< admonition warning "세미콜론과 중괄호" >}}
_세미콜론 입력규칙의 중요한 한가지는 제어문(if, for, switch, 혹은 select)의 여는 중괄호(`{`)를 다음 라인에 사용하지 말아야 한다._

```go
// This is good
if i < f() {
    g()
}
```

```go
// This sucks
if i < f()  // wrong!
{           // wrong!
    g()
}
```

{{< /admonition  >}}

## Control structures

Go언어에서는 do 나 while 반복문이 존재하지 않으며, `for`, `switch` `select`가 존재한다.

### if

**중괄호를 의무적으로 사용**해야 하기 때문에, 다중 라인에서 if 구문들이 간단하게 작성된다.

```go
if x > 0 {
    return y
}
```

**if와 switch가 초기화 구문을 허용**하므로 지역변수를 설정하기 위해 사용된 초기화 구문을 흔히 볼 수 있다.

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

### Redeclaration and reassignment

먼저 아래의 코드를 보자

```go
f, err := os.Open(name)
d, err := f.Stat()
```

이런 경우처럼 `err`가 위/아래 곳 모두에서 사용되는데, **이런 선언 중복은 허용된다.**

{{< admonition tip "Function variable" >}}
_Go언어에서 함수 파라미터와 리턴 값들은, 함수를 감싸고 있는 브래이스들(braces)밖에 위치해 있음에도, 그 스코프는 함수 body의 스코프와 동일하다는 점을 주목할 가치가 있다._
{{< /admonition  >}}

### for

```go
// C언어와 같은 경우
for init; condition; post { }

// C언어의 while 처럼 사용
for condition { }

// C언어의 for(;;) 처럼 사용
for { }
```

아래는 go-style for문들입니다.

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

만약 배열, slice, string, map, 채널로 부터 읽어 들이는 반복문을 작성한다면, range 구문이 이 반복문을 관리가능합니다.

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

- 이렇게 index를 날릴수도 있습니다.

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

- UTF-8 파싱이 덜된 string의 경우 rune으로 변환된다.

```go
for pos, char := range "日本\x80語" { // \x80 은 합법적인 UTF-8 인코딩이다
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
/*
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
*/
```

다음은 `parallel assignment`를 사용한 for문이다.

```go
for i, j := 0, len(a) -1; i<j; i,j = i+1, j-1 {
    a[i], a[j] = a[j], a[i]
}
```

### Switch

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}

```

go에서는 switch에 Label(예시에서는 `Label`)을 넣어서 `escape`하는 방식도 가끔이지만 쓰인다.

```go
Loop:
    for n := 0; n < len(src); n += size {
        switch {
        case src[n] < sizeOne:
            if validateOnly {
                break
            }
            size = 1
            update(src[n])

        case src[n] < sizeTwo:
            if n+1 >= len(src) {
                err = errShortInput
                break Loop
            }
            if validateOnly {
                break
            }
            size = 2
            update(src[n] + src[n+1]<<shift)
        }
    }
```

### Type switch

스위치 구문은 인터페이스 변수의 동적 타입을 확인하는데 사용될 수도 있다.

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

## Functions

### Multiple return values

- 값과 에러를 같이 내리는 signature

```go
func (file \*File) Write(b []byte) (n int, err error)
```

- index와 value를 return

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```

### Named result parameters

Go 함수에서는 반환 "인자"나 결과 "인자"에 이름을 부여하고 인자로 들어온 매개변수처럼 일반 변수로 사용할 수 있다. 이름을 부여하면, **해당 변수는 함수가 시작될 때 해당 타입의 제로 값으로 초기화 된다.**

```go
func nextInt(b []byte, pos int) (value, nextPos int) {
}
```

return parameter는 선언되기 때문에 다음과 같이 사용될 수도 있다.

```go
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

### Defer

Go 의 defer 문은 defer 를 실행하는 함수가 반환되기 전에 즉각 함수 호출(연기된 함수)을 실행하도록 예약한다.

이를 통해 기존의 언어가 자원 해제를 context를 사용했던 것과 달리 defer라는 키워드를 통해서 자원 해지가 가능해진다. (python에서는 with을 사용해서 context가 끝날때 **exit** 호출을 처리해 주었다.)

다음은 mutex에 lock을 풀거나 잠그는 코드입니다.

```go
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(resut, buf[0:n]...)
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err
        }
    }
    return string(result), nil
}
```

close를 delay 시킴으로써 크게 2가지 장점을 얻게된다.

1. **파일을 닫는 것을 잊어버리는 실수를 하지 않도록 보장해 준다.**
2. **open 근처에 close 가 위치하면 함수 맨 끝에 위치하는 것 보다 훨씬 명확한 코드가 되는것을 의미한다.**

defer는 함수가 종료될 때 실행되기 때문에 하나의 defer 호출 위치에서 여러개의 함수 호출을 delay 시킬 수 있다.

```go
for i:=0; i < 5; i++ {
    defer fmt.Println(i)
}
```

**지연된 함수는 `LIFO` 순서로 실행된다. (4 3 2 1 0)** (와우 궁금했던 부분인데)

추가로 defer안에 인자들에 대한 평가는 기존 함수 실행 순서에 따라 진행된다.
다음은 조금 더 복잡한 defer 예시이다.

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

위 함수들에 대한 결과는 아래와 같다.

```
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

나중에 나올 `panic` 과 `recover`를 사용하면 더욱 멋진 것들을 만들 수 있다고 한다.

## Data

### Allocation with new

Go에는 메모리를 할당하는 두가지 기본 방식이 있는데, 내장(built-in) 함수인 `new`와 `make`이다.

```go

```

```go

```

```go

```

```go

```

## conclustion

<center>- 끝 -</center>

