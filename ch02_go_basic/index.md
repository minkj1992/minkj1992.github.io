# [일주일 만에 배우는 GO] CH.2 Go Basic


[Effective go](https://go.dev/doc/effective_go)를 통해 go 핵심 문법을 빠르게 배워보겠습니다. 또한 나중에 레퍼런스 개념으로 개발할 때 찾아보기 위해서 필요해 보이는 정보들을 모으는 개념으로 글을 작성합니다.

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

- `new`: 새로 제로값으로 할당된 타입 T를 가리키는 포인터를 반환
  - `new(File)` == `&File{}`
- `make`: 내부 데이터 구조를 초기화하고 사용될 값을 준비한다.

먼저 new부터 살펴보면, 내장 함수로 메모리를 할당하지만 다른 언어에 존재하는 같은 이름의 기능과는 다르게 메모리를 초기화하지 않고, 단지 값을 제로화(zero) 한다. **다시 말하면, new(T)는 타입 T의 새로운 객체에 제로값이 저장된 공간(zeroed storage)을 할당하고 그 객체의 주소인, `*T`값을 반환한다.**

제로값의 유용함은 전이적인(transitive) 특성이 있다.

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}

p := new(SyncedBuffer)  // type *SyncedBuffer
var v SyncedBuffer      // type  SyncedBuffer
```

### Constructors and composite literals(합성 리터럴)

때로 제로값만으로는 충분치 않고 생성자(constructor)로 초기화해야 할 필요가 생긴다.

먼저 불필요한 boiler plate 코드 부터 확인해보자.

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

아래는 constructor를 활용한 방식이다.

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

C와는 달리, 로컬 변수의 주소를 반환해도 아무 문제가 없음을 주목하라; 변수에 연결된 저장공간은 함수가 반환해도 살아 남는다. 실제로, 합성 리터럴의 주소를 취하는 표현은 매번 실행될 때마다 새로운 인스턴스에 연결된다. 그러므로 마지막 두 줄을 묶어 버릴 수 있다.

```go
    return &File{fd, name, nil, 0}
```

합성 리터럴의 필드들은 순서대로 배열되고 반드시 입력해야 한다. 하지만, 요소들에 레이블을 붙여 필드:값 식으로 명시적으로 짝을 만들면, 초기화는 순서에 관계 없이 나타날 수 있다. 입력되지 않은 요소들은 각자에 맞는 제로값을 갖는다. 그러므로 아래와 같이 쓸 수 있다.

```go
    return &File{fd: fd, name: name}
```

{{< admonition tip "Composite literals(합성 리터럴)이란?" >}}
_Composite literals are used to construct the values for arrays, structs, slices, and maps_

```go
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"} // array
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"} // slice
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"} // map
```

{{< /admonition  >}}

### Allocation with make

`new`와 달리 `make`는 slices, maps, 그리고 channels에만 사용하고 (`*T`가 아닌) 타입 T의 (제로값이 아닌) 초기화된 값을 반환한다. 아래는 new와 make의 차이점을 보여준다.

```go
var p *[]int = new([]int)       // slice 구조체를 할당한다; *p == nil; 거의 유용하지 않다
var v  []int = make([]int, 100) // slice v는 이제 100개의 int를 갖는 배열을 참조한다

// 불필요하게 복잡한 경우:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// Go 언어다운 경우:
v := make([]int, 100)
```

**make는 maps, slices 그리고 channels에만 적용되며 포인터를 반환하지 않음을 기억해야 합니다. 포인터를 얻고 싶으면 new를 사용해서 메모리를 할당하거나 변수의 주소를 명시적으로 취해야 합니다.**

### Arrays

Go와 C에서는 배열의 작동원리에 큰 차이가 있다. Go에서는,

- 배열은 값이다.
- 한 배열을 다른 배열에 assign할 때 모든 값이 복사된다.
- 함수의 argument로 배열을 패스하면, 포인터가 아닌 copy된 array를 받는다.
- 배열의 크기는 타입의 한 부분이다. 타입 [10]int과 [20]int는 서로 다르다.

> 개인적으로 `배열의 크기는 타입의 한 부분이다. 타입 [10]int과 [20]int는 서로 다르다.`가 무슨 말인지 잘 모르겠다.

배열을 값(value)으로 사용하는 것이 유용할 수도 있지만 또한 비용이 큰 연산이 될 수도 있다; 만약 C와 같은 실행이나 효율성을 원한다면, 아래와 같이 배열 포인터를 보낼 수도 있다.

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // 명시적인 주소 연산자(&)를 주목하라.
```

**하지만 이런 스타일조차 Go언어 답지는 않다. 대신 slice를 사용하라.**

### Slices

Go에서는 변환 메스릭스와 같이 뚜렷한 차원(dimension)을 갖고 있는 항목들을 제외하고는, 거의 모든 배열 프로그래밍은 단순한 배열보다는 `slice`를 사용한다.

Slice는 내부의 배열을 가리키는 레퍼런스를 쥐고 있어, 만약에 다른 slice에 할당(assign)되어도, 둘 다 같은 배열을 가리킨다. 함수가 slice를 받아 그 요소에 변화를 주면 호출자도 볼 수 있는데, 이것은 내부의 배열를 가리키는 포인터를 함수에 보내는 것과 유사하다.

slice의 용량은, 내장함수 cap을 통해 얻을 수 있는데, slice가 가질 수 있는 최대 크기를 보고한다. 아래를 보면 slice에 데이터를 부착(append)할 수 있는 함수가 있다. 만약 데이터가 용량을 초과하면, slice의 메모리는 재할당된다. 결과물인 slice는 반환된다.

```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) { // 재할당의 경우
        doubleLength := (l+len(data))*2
        newSlice := make([]byte, doubleLength)

        // copy 함수는 사전에 선언되어 있고 어떤 slice 타입에도 사용될 수 있다.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0: l + len(data)]
    copy(slice[1:], data)
    return slice
}
```

**slice는 꼭 처리후 반환되어야 한다.** Append가 slice의 요소들을 변경할 수 있지만, slice 자체(포인터, 크기, 용량을 갖고 있는 런타임 데이터 구조)는 값으로 패스되었기 때문이다. 참고로 slice에는 append가 구현되어있다.

### Two-dimensional slices

다음은 `go`에서 이차원 slices 또는 배열을 정의하고 init하는 방법입니다.

```go
type Transform [3][3]float64
type LinesOfText [][]byte

text := LinesOfText{
    []byte("Leoo is awesome"),
    []byte("Life is fun"),
    []byte("Life is full of love"),
    []byte("Let's give our love and fire to the world")
}
```

예를 들어 사진을 스캔하는 상황이 온다면 2가지 방식으로 이를 해결할 수 있다.

- 일반적으로 2차원 배열을 만들어 사용하는 방식

```go
height := 300
width := 300

picture := make([][]uint8, height)
for i:= range picture {
    picture[i] = make([]uint8, width)
}
```

- 하나의 긴 slice에 width만큼 자르면서 이차원 배열에 포인터를 전달하는 방식

```go
picture := make([][]uint8, height)
pixels := make([]uint8, width * height)

for i := range picture {
    picture[i], pixels = pixels[:width], pixels[width:]
}
```

### Maps

> {key: value}

`key`는 `equality`연산이 정의되어 있는 어떤 타입이라도 가능하다.

- int
- float
- string
- pointer
- interface(equality 구현된)
- structs
- array

`slice`의 경우에는 map의 key로 사용이 될 수 없는데, 이유는 equality가 정의되어 있지 않기 때문이다.

{{< admonition question "왜 slice에는 equality가 없을까?" >}}
[go에 제시되었던 issue](https://github.com/golang/go/issues/21829)에 레퍼런스된 [slice equality에 대한 golang discussion 링크](https://groups.google.com/g/golang-nuts/c/ajXzEM6lqJI)를 보면서 일부분을 정리하면, **slice가 value로 비교해야할지, pointer타입으로 비교해야할지 혼돈을 줄 수 있기 때문이라고 합니다.**

_This would probably introduce unnecessary confusion. People are used to the equality operator comparing values in go, as opposed to references. It's much better if the slices finally support the equality operator, even though the comparison speed will depend on the number of items in the slices._
{{< /admonition  >}}

Slice와 마찬가지로 map 역시 내부 데이터 구조를 가진다. 함수에 map을 입력하고 map의 내용물을 변경하면, 그 변화는 호출자에게도 보인다.

Map 또한 콜론으로 분리된 key-value 짝을 이용한 합성 리터럴로 생성될 수 있으며, 초기화중에 쉽게 만들 수 있다.

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

```go
offset := timeZone["EST"]
```

go는 keyError를 내지 않고 타입별로 0을 의미하는 값을 리턴한다. 그러므로 아래와 같은 경우가 가능하다.

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // 만약 person이 맵에 없다면 false일 것이다.
    fmt.Println(person, "was at the meeting")
}
```

만약 value가 bool인 경우같이 keyError와 value(false)를 구분하고 싶다면 아래와 같이한다. 이것을 "comma ok" 관용구라고 부른다. 이 예제에서, 만약 tz가 있다면, seconds는 적절히 세팅될 것이고 ok는 true가 된다

```go
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

다음은 에러헨들링 하는 방식이다.

```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

값이 필요없다면 이렇게 한다.

```go
_, present := timeZone[tz]
```

Map의 엔트리를 제거하기 위해서는, 내장 함수 delete을 쓰는데, map과 제거할 key를 인수로 쓴다. map에 key가 이미 부재하는 경우에도 안전하게 사용할 수 있다.

```go
delete(timeZone, "PDT")  // Now on Standard Time
```

### Printing

정수(integer)를 소수로 바꾸는 예와 같은 기본적인 변환을 원할 경우는, 다목적 용도 포맷인 %v(value라는 의미로)를 사용할 수 있다

```go
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)

// map[CST:-21600 PST:-28800 EST:-18000 UTC:0 MST:-25200]
```

물론, map의 경우 key들은 무작위로 출력될 수 있다. struct를 출력할 때는, 수정된 포맷인 `%+v`를 통해 구조체의 필드에 주석으로 이름을 달며, 대안 포맷인 `%#v`를 사용하면 어떤 값이든 완전한 Go 문법을 출력한다.

```go
type T struct {
    a int
    b float64
    c string
}

t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```

```
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string] int{"CST":-21600, "PST":-28800, "EST":-18000, "UTC":0, "MST":-25200}
```

또 다른 유용한 포맷은 %T로, 값의 타입을 출력한다.

```go
fmt.Printf("%T\n", timeZone)

// map[string] int
```

{{< admonition note "커스텀 타입 print 포맷 지정하는 방법" >}}

커스텀 타입의 기본 포맷을 조종하기 위해 해야할 것은 단지 String() string의 시그너처를 갖는 메서드를 정의해 주는 것이다. (위에 정의된) 단순한 타입 T는 아래와 같은 포맷을 가질 수 있다.

```go
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)

// 7/-2.35/"abc\tdef"
```

{{< /admonition  >}}

위에 예제에서 struct 타입에 포인터를 사용한 이유는 더 효율적이고 Go 언어다운 선택이기 때문이다.

**String 메서드가 Sprintf를 호출할 수 있는 이유는 print 루틴들의 재진입(reentrant)이 충분히 가능하고 예제와 같이 감싸도 되기 때문이다**. 하지만 이 방식에 대해 한가지 이해하고 넘어가야 하는 매우 중요한 디테일이 있는데: **String 매서드를 만들면서 Sprintf를 호출할 때 다시 String 매서드로 영구히 재귀하는 방식은 안 된다는 것이다.** Sprintf가 리시버를 string처럼 직접 출력하는 경우에 이런 일이 발생할 수 있는데, 그렇게 되면 다시 같은 메서드를 호출하게 되고 말 것이다. 흔하고 쉽게 하는 실수로, 다음의 예제에서 살펴보자.

```go
package main

import "fmt"

type MyString string

func (m MyString) String() string {
	return fmt.Sprintf("MyString=%s", m) // 에러: 영원히 재귀할 것임.
}

func main() {
	var s MyString = "test"
	fmt.Printf("%v\n", s)
}
```

해결책은 `string()` 시켜주면 된다. 인수를 기본적인 문자열 타입으로 변환하면, 같은 메서드가 없기 때문이다.

```go
package main

import "fmt"

type MyString string

func (m MyString) String() string {
	return fmt.Sprintf("MyString=%s", string(m))
}

func main() {
	var s MyString = "test"
	fmt.Printf("%v\n", s)
}
```

또 다른 출력 기법으로는 출력 루틴의 인수들을 직접 또 다른 유사한 루틴으로 대입하는 것이다. Printf의 시그너처는 마지막 인수로 임의적인 숫자의 파라미터가 포맷 다음에 나타날 수 있음을 명시하기 위해 타입 `...interface{}`를 사용한다.

```go
// Println 함수는 fmt.Println처럼 표준 로거에 출력한다.
func Println(v ...interface{}) {
    std.Output(2, fmt.Sprintln(v...))  // Output 함수는 (int, string) 파라미터를 받게된다.
}
```

```go
log.Println("Hello", "世界", 1, 2, 3, 4, 5, 6, 7, 8)

// 2009/11/10 23:00:00 Hello 世界 1 2 3 4 5 6 7 8
```

`Sprintln`을 부르는 중첩된 호출안에 v 다음에 오는 `...`는 컴파일러에게 v를 인수 리스트로 취급하라고 말하는 것이고; 그렇지 않은 경우는 v를 하나의 slice 인수로 대입한다.

`...` 파라미터는 특정한 타입을 가질 수도 있는데, 예로 integer 리스트에서 최소값을 선택하는 함수인 min에 대한 `...int`를 살펴보자

```go
func Min(a ...int) int {
    min := int(^uint(0) &gt; &gt; 1) // wtf????
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

### Append

go에 내장되어있는 append 함수의 signature는 다음과 같다.

```go
// slice는
func append(slice []T, elements ...T) []T
```

기본적인 사용법

```go
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

slice 끼리 append

```go
x := []int{1,2,3}
y := []int{4,5,6}
x = append(x, y...)
fmt.Println(x)
```

## conclustion

<center>- 끝 -</center>

