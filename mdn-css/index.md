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


## REFS
- [css 디버깅](https://developer.mozilla.org/ko/docs/Learn/CSS/Building_blocks/Debugging_CSS)
