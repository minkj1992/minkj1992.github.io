# Go Value Parts


Real implementation of value parts type in golang
<!--more-->



## Value Parts
> https://go101.org/article/value-part.html



~~하나의 value가 하나의 메모리 block을 차지하는 `C`와 달리~~, golang은 몇몇 types들이 하나 이상의 memory block에 할당될 수도 있습니다. **이렇게 다른 메모리 블록들에서 part되어 분포되는 value의 구성요소들을 value parts라고 칭합니다.** 하나 이상의 메모리 블록에 hosting되는 value는 `direct value part`와 여러개의 `underlying indirect parts`로 구성됩니다.

> go101 문서가 비약이 좀 많은 것 같습니다. 예를 들면 "Each C value in memory occupies one memory block (one continuous memory segment)." 라는 주장에 대해서 아래와 같이 반박할 수 있습니다.
>> C 언어에서도 모든 값이 단일 메모리 블록에 저장된다는 주장은 정확하지 않습니다. 구조체, 배열, 포인터와 같은 데이터 타입들은 여러 메모리 블록에 걸쳐 저장될 수 있습니다. 예를 들어, 구조체는 각 멤버 변수가 서로 다른 메모리 위치에 저장될 수 있고, 큰 배열은 여러 블록에 분포될 수 있으며, 포인터가 가리키는 변수는 별도의 메모리 블록에 저장됩니다. 따라서 Go 언어와 마찬가지로 C 언어에서도 값의 복잡성에 따라 메모리 분포가 다양할 수 있다.
>> https://github.com/go101/go101/issues/270 에 관련된 doc fix issue를 넣었습니다.



- **Solo Direct Value Part**는 단일 메모리 블록에 저장되는 값을 의미합니다. 즉, 값 전체가 하나의 연속된 메모리 공간에 존재합니다.
- **Direct value part**:  포인터의 value처럼 reference 하는 address value
- **serveral underlying indirect parts**: 여러 메모리 블록에 분산되어 있는 값의 각 부분

아래는 golang에서 지원하는 type을 multiple value parts 여부 (메모리 블록 갯수)로 나눈 테이블입니다.

| Types whose values each is only hosted on one single memory block (solo direct value part) | Types whose values each may be hosted on multiple memory blocks (direct part -> underlying direct part) |
|--------------------------------------------------------------------|-----------------------------------------------------------------|
| ![single value part](https://go101.org/article/res/value-parts-single.png) | ![multiple value parts](https://go101.org/article/res/value-parts-multiple.png) |
| boolean types<br/>numeric types<br/>pointer types<br/>unsafe pointer types<br/>struct types<br/>array types | slice types<br/>map types<br/>channel types<br/>function types<br/>interface types<br/>string types |


> \* Note
>> - 인터페이스와 문자열 값에 기본 부분이 포함될 수 있는지 여부는 컴파일러에 따라 다릅니다. 
>> - 표준 Go 컴파일러 구현의 경우 인터페이스 및 문자열 값에 기본 부분이 포함될 수 있습니다.
>> - 함수 값에 기본 부분이 포함될 수 있는지 여부를 증명하는 것은 거의 불가능합니다.


> 어째서 101문서에서 pointer, unsafe pointer를 solo direct value part로 구분했는지 모르곘다. 관련된 문의 pr을 올렸다. 
> https://github.com/go101/go101/issues/269

## Internal definitions

그럼 이제 실제 2번째 type들의 내부 definitions들을 살펴보겠습니다.

### `map`, `channel` and `function types`


```go
// map types
// map types
type _map *hashtableImpl

// channel types
type _channel *channelImpl

// function types
type _function *functionImpl
```

3가지 유형은 내부적으로 그냥 포인터 유형입니다.

### slice

```go
type _slice struct {
	// referencing underlying elements
	elements unsafe.Pointer
	// number of elements and capacity
	len, cap int
}
```

슬라이스 유형은 pointer wrapper struct types입니다.


> Unsafe Pointer
> Go의 `unsafe` 패키지에 정의된 Unsafe Pointer는 언어의 타입 안전성 시스템을 우회하여 직접 메모리 접근을 수행할 수 있게 해줍니다. 이를 통해 성능 최적화, 시스템 레벨 프로그래밍, interfacing with non-Go 코드 등에 필요할 수 있습니다. 또한 일반 pointer 타입이 형변환이 불가한 것과 달리, 타입 시스템을 우회하는 unsafe pointer는 형변환이 가능합니다. (물론 일반 pointer도 reflect을 사용해서 runtime에 형변환을 시키는 방법도 있습니다.)

### string

```go
type _string struct {
	elements *byte // referencing underlying bytes
	len      int   // number of bytes
}
```

string 또한 pointer wrapper struct type입니다.

### interface

- blank interface type

```go
type _interface struct {
	dynamicType  *_type         // the dynamic type
	dynamicValue unsafe.Pointer // the dynamic value
}
```

**standard go compiler에서 위의 정의를 blank interface types에만 사용합니다.**

> Blnak interface types are the interface types which don't specify any methods.

- non-blacnk interface type

```go
type _interface struct {
	dynamicTypeInfo *struct {
		dynamicType *_type       // the dynamic type
		methods     []*_function // method table
	}
	dynamicValue unsafe.Pointer // the dynamic value
}
```


## copy
> Underlying Value Parts Are Not Copied in Value Assignments

Golang에서 parameter passing을 포함한 value Assignments는 shallow value copy입니다. (단, destination과 source value의 타입이 같을 경우에만) 만약 타입이 다른 value끼리의 value assignment는 implicitly converted 되어 assignment가 진행됩니다.


일반적인 값 할당 (shallow copy)

1. `direct part`만 복사됩니다.
2. `underlying value part`은 참조만 복사됩니다.
3. 결과적으로 원본 값과 복사된 값은 `underlying value part`을 공유하게 됩니다.


문자열과 인터페이스의 특수 케이스:
- 위의 shallow copy와 달리 [Go FAQ](https://go.dev/doc/faq#pass_by_value)에 따르면, 인터페이스 value가 copy될 때,`underlying value part` 또한 copy되어야 한다고 합니다.
- 이론적으로는 `underlying value part`도 함께 복사되어야 하지만, 내부적으로는 그렇지 않습니다.
- **실제 동작에서는 인터페이스의 dynamic value는 read only이기 때문에, Go 컴파일러/런타임은 최적화를 위해 `underlying value part`를 복사하지 않습니다.**
- 이는 string또한 똑같이 적용됩니다.


결론:
- 표준 Go 컴파일러/런타임에서 모든 타입의 값 할당은 `direct part`만 복사하고, `underlying value part`은 공유합니다.
- 문자열과 인터페이스는 컴파일러 최적화로 인해 이론과 약간 차이가 있습니다.
- `unsafe.Sizeof` 함수는 `direct part`의 크기만 반환하며, `underlying value part`은 포함하지 않습니다.


