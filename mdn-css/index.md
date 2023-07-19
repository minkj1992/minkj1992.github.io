# CSS (MDN)


[CSS: Cascadubg Style Sheets](https://developer.mozilla.org/ko/docs/Web/CSS)
<!--more-->

## 1. CSS first steps

- list-style-type
    - @counter-style
    - display: list-item

```css
li {
  list-style-type: none;
}
```

```css
@counter-style thumbs {
  system: cyclic;
  symbols: "\1F44D";
  suffix: " ";
}

ul {
  list-style: thumbs;
}
```

- 여러개 동시에 적용시킬 때, `,`

```css
li.special,
span.special {
  color: orange;
  font-weight: bold;
}
```

- element 하위요소안에 적용할 때 (space)

```css
li em {
  color: rebeccapurple;
}
```

- sibling
    - h1, p 둘다 font 200%
    - h1 바로 뒤에 나오는 p에 대해서 적용된다.

```svelte

<style>
h1 + p {
  font-size: 200%;
}
</style>

<h1>I am a<p>minwook</p> level one heading</h1>

<p>This is a paragraph of text. In the text is a <span>span element</span> 
and also a <a href="http://example.com">link</a>.</p>
```

```css
a:hover {
  text-decoration: none;
}
```

### CSS의 구조

#### 계단식 및 상속
> https://developer.mozilla.org/ko/docs/Learn/CSS/Building_blocks/Cascade_and_inheritance


1. cascade
    1. Importance
    2. 우선 순위
    3. 소스 순서
2. ID > .(class) > tag

**일부 속성은 상속되지 않습니다.** 너비 (위에서 언급 한 것처럼), 마진, 패딩 및 테두리와 같은 것은 상속되지 않습니다. 


- 상속 제어하기
    - inherit: 부모 영향 받기
    - initial: 기본 브라우저 스타일
    - unset: 상속값 무시


#### Properties and Value

```css
{
    properties: value
}
```

#### @rules
> "at-rules"


- @media
- @import



### CSS 작동 방식

![](https://developer.mozilla.org/ko/docs/Learn/CSS/First_steps/How_CSS_works/rendering.svg)

1. 브라우저는 HTML (예: 네트워크에서 HTML 을 수신) 을 로드합니다.
2. HTML 을 DOM (Document Object Model) 로 변환합니다. DOM 은 컴퓨터 메모리의 문서를 나타냅니다. DOM 은 다음 섹션에서 좀 더 자세히 설명됩니다.
3. 그런 다음 브라우저는 포함된 이미지 및 비디오와 같은 HTML 문서에 연결된 대부분의 리소스와 연결된 CSS 를 가져옵니다! JavaScript 는 작업에서 나중에 처리되므로 더 간단하게 하기위해 여기에서는 다루지 않습니다.
4. 브라우저는 가져온 CSS 를 구문 분석하고 선택자 유형별로 다른 규칙을 다른 "buckets" 으로 정렬합니다. 예: 요소, class, ID 등 찾은 선택자를 기반으로 DOM 의 어느 노드에 어떤 규칙을 적용해야 하는지 결정하고, 필요에 따라 스타일을 첨부합니다 (이 중간 단계를 render tree 라고 합니다).
5. render tree 는 규칙이 적용된 후에 표시되어야 하는 구조로 배치됩니다.
6. 페이지의 시각적 표시가 화면에 표시됩니다 (이 단계를 painting 이라고 함).

## 2. CSS 구성 블록
> https://developer.mozilla.org/ko/docs/Learn/CSS/Building_blocks

### CSS 선택자

- 유형
  1. type
  2. class
  3. id

- 속성 선택자

```css
a[title] { }
```

- Pseudo-classes 및 pseudo-elements

```css
/* <article> 요소의 모든 자식 요소 중에서 첫 번째 자식 요소에 스타일을 적용  */
article :first-child {
  font-weight: bold;
}

/* article :first-child와 같음. */
article *:first-child {}

/*  <article> 요소가 부모 요소의 첫 번째 자식으로 있을 때에만 스타일이 적용됩니다. 즉, <article> 요소가 문서 내에서 최상위 요소가 되거나, 다른 요소들 사이에서 첫 번째로 등장하는 경우에만 스타일이 적용 */
article:first-child {}

p::first-line { } // pseudo-elements
```

### ' ' vs '>'

하위 결합자 (Descendant combinator)와 자식 결합자 (Child combinator)는 CSS에서 요소를 선택하는 데 사용되는 두 가지 다른 방식입니다.

- 하위 결합자 (` `):
이 결합자는 한 요소의 모든 하위 요소를 선택하는 데 사용됩니다. 예를 들어, article p는 article 요소 안에 있는 모든 p 요소를 선택합니다. 이는 중첩된 p 요소가 있더라도 해당 p가 어디에 위치하든 간에 모든 p 요소를 선택합니다.

- 자식 결합자 (`>`):
이 결합자는 바로 아래에 있는 자식 요소만을 선택합니다. article > p라는 선택자는 article 요소의 직접적인 자식인 p 요소만을 선택합니다. 만약 p 요소가 더 깊게 중첩되어 있으면, 그 요소는 선택되지 않습니다.

즉, 이 둘의 주요한 차이는 article p가 article 내의 모든 수준에서 p 요소를 선택하는 반면, article > p는 오직 article의 직접적인 자식인 p 요소만을 선택한다는 것입니다.


### 2이상 클래스 적용

```svelte
<style>
.notebox {
  border: 4px solid #666;
  padding: .5em;
}

.notebox.warning {
  border-color: orange;
  font-weight: bold;
}

.notebox.danger {
  border-color: red;
  font-weight: bold;
}
</style>

<div class="notebox">
    This is an informational note.
</div>

<div class="notebox warning">
    This note shows a warning.
</div>

<div class="notebox danger">
    This note shows danger!
</div>

<div class="danger">
    This won't get styled — it also needs to have the notebox class
</div>    
```
### Attribute Selector
> https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Selectors/Attribute_selectors#presence_and_value_selectors

| 선택자         | 예                            | 설명                                                                                                           |
| -------------- | ----------------------------- | -------------------------------------------------------------------------------------------------------------- |
| [attr]         | a[title]                      | attr 속성(**이름**은 대괄호 안의 값임)이 있는 요소와 일치합니다 .                                              |
| [attr=value]   | a[href="https://example.com"] | 값이 **정확히 value** (따옴표 안의 문자열) 인 attr 속성이 있는 요소와 일치합니다 .                             |
| [attr~=value]  | p[class~="special"]           | 값이 정확히 value 인 attr 속성이 있는 요소와 일치하거나 (공백으로 구분된) 값 목록에 있는 값을 **포함**합니다 . |
| [attr\|=value] | div[lang\|="zh"]              | 값이 정확히 value 이거나 **바로 뒤에 하이픈**이 오는 value 로 시작하는 attr 속성 이 있는 요소와 일치합니다 .   |


```css
/* li에 class가 있는 요소만 */
li[class] {
    font-size: 200%;
}

/* li의 class 이름이 정확히 a */
/* 적용되지 않음:  <li class="a b">Item 2</li> */
li[class="a"] {
    background-color: yellow;
}

/* a를 포함 */
/* 적용됨 <li class="a b">Item 2</li> */
/* 적용되지 않음 <li class="ab">Item 4</li> */
li[class~="a"] {
    color: red;
}
```

- 대소문자 구분하지 않음 (`i`)

```css

li[class^="a" i] {
    color: red;
}
```

### Pseudo Class (`:`)
> Pseudo Class는 특정 상태에 있는 요소를 선택하는 선택기입니다

- `:hover`
- `:first-child`


```css
article p:first-child {
    font-size: 120%;
    font-weight: bold;
}   
```

### Pseudo Element (`::`)
> 기존 요소에 클래스를 적용하는 것이 아니라 완전히 새로운 HTML 요소를 마크업에 추가한 것처럼 작동합니다.

### Cascade Layer
> https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/Cascade_layers


## REFS
- [css 디버깅](https://developer.mozilla.org/ko/docs/Learn/CSS/Building_blocks/Debugging_CSS)
