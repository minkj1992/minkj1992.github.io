# Elderjs



`Elder.js`: an SEO first Svelte Framework & Static Site Generator
<!--more-->



# Elder.js
> https://elderguide.com/tech/elderjs

## cli 

```bash
pnpm start

npm run build
# preview
npx sirv-cli public
# To Run in SSR Mode for Production:
npm run serve
```

## `Elder.js의 Markdown-plugin`

- [Getting all Markdown For a Route](https://github.com/Elderjs/plugins/tree/master/packages/markdown#getting-all-markdown-for-a-route)
```
data.markdown
data.markdown.blog
```

- By default, if there is a **date field in your frontmatter** it will sort all of the markdown for that route by it.
- If there is a **slug field in your frontmatter** it will use that for the slug, if not it falls back to the **filename**.
- If `draft: true` is in a file's **frontmatter** or a **slug is prefixed with draft-** these markdown files will be hidden when `process.env.NODE_ENV === 'production'`. They will also be prefixed with DRAFT: [Post Title Here] to make this functionality more obvious.

markdown 플러그인이 data hook때, 어떤 데이터들을 추가하는지 나타내는 코드

- [markdown plugin data hook](https://github.com/Elderjs/plugins/blob/f7e8fd8f2503881342220baf9df6b2d535427f24/packages/markdown/src/index.ts#L222)
- [elder js `data` hook](https://elderguide.com/tech/elderjs/#data)

```js
// https://github.com/Elderjs/plugins/blob/f7e8fd8f2503881342220baf9df6b2d535427f24/packages/markdown/src/index.ts#L222
{
      hook: 'data',
      name: 'addFrontmatterAndHtmlToDataForRequest',
      description: 'Adds parsed frontmatter and html to the data object for the specific request.',
      priority: 50,
      run: async ({ request, data }) => {
        if (data.markdown && data.markdown[request.route]) {
          const markdown = data.markdown[request.route].find((m) => m.slug === request.slug);
          if (markdown) {
            await markdown.compileHtml();
            const { html, frontmatter, data: addToData } = markdown;

            return {
              data: {
                ...data,
                ...addToData,
                html,
                frontmatter,
              },
            };
          }
        }
      },
},
```

### `Markdown plugin's` Ecosystem

- [Markdown plugin](https://github.com/Elderjs/plugins/tree/master/packages/markdown)

- markdown
    - remarkjs (Markdown Abstract Syntax Tree format -> unified)
        - `remark-frontmatter`: 마크다운 frontmatter 처리
        - `remark-gfm`: github flavored markdown 형식 변환
        - [`remark-slug`](https://github.com/remarkjs/remark-slug/tree/8e6394c): (Deprecated) html tag id 채워주기
        - [`remark-html`](https://github.com/remarkjs/remark-html)
            - serializing HTML
            - unified와 remark 그리고 rehype를 이어준다.
            - [복수의 (unified, rehype) markdown -> html 컴파일러가 필요한 이유](https://github.com/remarkjs/remark-html#when-should-i-use-this)
        - `remark` unified, remark-parse, and remark-stringify, useful when input and output are markdown
    - unified
        - transform markdown -> HTML
    - `rehype`: Markdown -> HTML (uses `hast`)


즉 unified와 rehype는 같은 기능(md -> html)을 하는 녀석들이지만, 서로 보충해주는 역활을 한다. 이유는 markdown의 문법이 여러 경우의 수가 있어서 html 전환을 할때 경우의 수가 많기 때문이다. 
보통 `.use(remarkRehype).use(rehypeStringify)`를 사용한다.



## `Route.js`
- `all` -> pass requests -> `permalink` transforms links
- **Skinny request objects. Fat data functions**
    - Fetching, preparing, and processing data should be done in your `data function`.
    - it is recommended that you only include the bare minimum required to query your database, api, file system, or data store on the `request object`.


```js
all: async ({ settings, query, data, helpers }): Array<Object> => {
  // settings: this describes the Elder.js settings at initialization.
  // query: an empty object that is usually populated on the 'bootstrap' hook with a database connection or api connection. This is sharable throughout all hooks, functions, and shortcodes.
  // data: any data set on the 'bootstrap' hook.
  return Array<Object>;
}

permalink: ({ request, settings, helpers }): String => {
  // NOTE: permalink must be sync. Async is not supported.

  // request: this is the object received from the all() function. Generally, we recommend passing a 'slug' parameter but you can use any naming you want.
  // settings: this describes the Elder.js bootstrap settings.
  // helpers: Elder.js helpers and user helpers from the ./src/helpers/index.js` file.
  // NOTE: You should avoid using helpers here as helpers.permalinks default helper (see below) doesn't support it.
  return String;
};

data: async ({
  data, // any data set by plugins or hooks on the 'bootstrap' hook
  helpers, // Elder.js helpers and user helpers from the ./src/helpers/index.js` file.
  allRequests, // all of the `request` objects returned by a route's all() function.
  settings, // settings of Elder.js
  request, // the requested page's `request` object.
  errors, // any errors
  perf, // the performance helper.
  query, // search for 'query' in these docs for more details on it's use.
}): Object => {
  // data is any data set from plugins or hooks.
  return Object;
};
```


## Pages.ts

[Elder js에서 page에서 사용할 수 있는 props](https://github.com/Elderjs/elderjs/blob/6dae253bc12b438d71b12bb78365991284f536fa/src/utils/Page.ts#L152)
  

## Hooks.ts

[About hooks](https://elderguide.com/tech/elderjs/#hooks)

## `Elder.js` x `Svelte`

- `Svelte Components`
    - ./src/components/ folder
    - `Svelte Templates` and `Svelte Layouts` 에서 호출
- `Svelte Templates`
    - for each route (eg: ./src/routes/blog/Blog.svelte)
    - **only rendered on the server** because they receive **props of data, helpers, request, and settings.**
- `Svelte Layouts`
    - ./src/layouts/ folder
    - **only rendered on the server** because they receive **props of data, helpers, request, settings, and templateHtml.**

대부분의 **Svelte 파일은 ./src/components/ 폴더에 있으며** 아래에 정의된 대로 Svelte 템플릿 또는 Svelte 레이아웃 내에서 파일을 hydrate(ssr or csr) 해서 사용할 수 있다.

<details>
<summary>Sveltekit에서 hydrate이 뭐야?</summary>

A: SvelteKit에서 "hydrate"는 서버에서 렌더링된 HTML을 클라이언트에서 동적으로 활성화하는 과정을 의미합니다. 이 과정은 서버 사이드 렌더링(SSR)과 클라이언트 사이드 렌더링(CSR) 사이의 간극을 메우는 중요한 단계입니다.

서버에서 페이지가 렌더링되면, 그 결과는 정적 HTML로 클라이언트에 전송됩니다. 이 HTML은 브라우저에서 빠르게 렌더링되어 사용자에게 페이지를 즉시 보여줄 수 있습니다. 그러나 이 HTML은 초기에는 정적이므로, 사용자의 상호작용에 반응하기 위해서는 JavaScript가 필요합니다.

이때 "hydration" 단계가 시작됩니다. SvelteKit는 이 단계에서 서버에서 렌더링된 HTML을 가져와서 Svelte 컴포넌트로 "hydrate"하거나 활성화합니다. 이렇게 하면 페이지가 동적으로 반응하고 사용자의 상호작용에 응답할 수 있게 됩니다.

이 과정은 사용자에게 빠른 초기 로딩 시간을 제공하면서도, 페이지가 완전히 인터랙티브하게 동작할 수 있도록 하는 데 중요합니다. 이것이 SvelteKit와 같은 프레임워크에서 SSR과 CSR을 결합하는 이유입니다.

</details>


### Partial Hydration

어떤 것이 hydrate되어야 하고 어떤 것이 hydrate되지 않아야 하는지 확실하지 않은 경우, 일반적인 규칙은 **if a component needs to be interactive on the client, you need to hydrate it.**


## `Shorcodes`

현재 굳이 필요없어 패스

[link](https://elderguide.com/tech/elderjs/#overview)

## Elder.js Data flow
> [link](https://elderguide.com/tech/elderjs/#data-flow)

<center>

![Server hooks](/images/server_hooks.png)

</center>

1. route.js module export 실행
2. Elder의 `bootstrap hook` 실행
    1. `all` function in routes are called
3. `allRequests hook`  run
    1. allows users to modify the allRequests array.
4. Full `request` Objects are Built
    1. 이는 추후 data func의 파라미터로 들어간다.
    2. 또한 Svelte templates, Svelte layouts 파일로 전달된다.

- i.g
```js
request = {
  slug: `why-kittens-rock`,
  // ... any other keys from the `request` object returned from the `all` function.

  // below is then added by Elder.js
  permalink: "/blog/why-kittens-rock",
  route: "blog",
  type: "build", // server or build.
};
```



5. 남은 `Hooks`들이 실행된다. (Until the `Route's data` function)
    1. 즉 위의 그림에서`middleware`, `request` 실행됨
6. The 'data' hook is Executed
7.  The `data Object` is passed to the `Svelte template`

```svelte
// i.g Blog.svelte
<script>
  export let data; // here is the 'data' object we've been following.
  export let settings; // Elder.js settings
  export let helpers; // Elder.js helpers and user helpers.
  export let request; // 'request' object from above. ....
</script>
```

8.  The HTML returned by Blog.svelte is passed into Layout.svelte
    - svelte layout도 svelte template과 같은 props를 받는다. 추가로 `templateHtml` props도 받는다.
9.  Page Generation Completes
    1.  All further hooks are run until the 'request' has been completed.
    2.  This includes user hooks, system hooks, and plugin hooks.


## Elder.js vs Sveltekit

Elder와 SvelteKit을 비교하는 방법은 매우 일반적인 질문입니다. 주요 차이점은 Elder.js가 SEO를 염두에 두고 설계되었으며 대규모 정적 사이트를 보다 쉽게 ​​구축할 수 있는 도구를 제공한다는 것.

## 더 읽어볼거리

- [자바스크립트 검색엔진 최적화의 기본사항 이해하기](https://developers.google.com/search/docs/crawling-indexing/javascript/javascript-seo-basics?hl=ko)
- [Google hydration guide](https://web.dev/rendering-on-the-web/)

