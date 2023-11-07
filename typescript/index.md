# Deep dive into Typescript


Deep dive into typescript
<!--more-->


# tsc and ts-node

## `tsc`(Typescript Compiler)

tsc는 타입스크립트 공식 컴파일러로 Typescript 코드를 Javascript 코드로 변환합니다. 이 과정에서 타입 검사를 수행하고 설정된 대상 버전의 JS 코드를 생성합니다.

## `ts-node`

`ts-node`는 TypeScript 코드를 직접 실행할 수 있는 Node.js의 런타임입니다. 기본적으로 Node.js는 .js 파일만 실행 할 수 있지만, ts-node를 사용하면 Node.js의 런타임을 확장하여 .ts 파일을 직접 실행할 수 있는 기능을 추가합니다.

일반적으로 ts-node는 개발환경에서 TS코드를 신속하게 실행하고 테스트하기 위해 사용되기 때문에 production 의존성에서는 제외시키는 것이 좋습니다.

## tsc vs ts-node

개발자는 ts-node를 사용해 빠른 피드백을 얻고 프로덕션 빌들르 생성할 때는 tsc를 사용해 최적화된 Javascript 코드를 생성할 수 있습니다.

# eslint, tslint, prettier

2019년 `tslint`는 deprecated되었고, eslint는 코딩 컨벤션을 위배하거나 안티 패턴을 자동 검출하는 Linting(소스코드를 분석하여 프로그램 오류, 버그, 스타일 오류등을 찾아내는 도구다.) tool입니다. 공식문서에서는 타입스크립를 위한 eslint로 tslint대신 [`typescript-eslint`](https://typescript-eslint.io/)를 추천합니다. 
`prettier`는 코딩 스타일 교정을 위한 포맷터입니다.

광범위한 scope이다 보니, eslint와 prettier이 충돌되는 경우가 있습니다. (예를 들면 tab옵션)
충돌되지 않도록 하기 위해서는 일반적으로 Prettier를 먼저 실행하고 eslint를 실행하는 것이 좋습니다.

> `Code` -> `prettier` -> `eslint --fix` -> `Formatted Code`

`Prettier`로 코드를 Formatting한 뒤(`prettier`) ESLint로 수정(`eslint --fix`)한 것과 같은 결과물을 출력한다.

- [LogRocket: Linting in Typescript using eslint and prettier](https://blog.logrocket.com/linting-typescript-eslint-prettier/)
- [TypeScript ESLint + Prettier 함께 사용하기 on vscode](https://pravusid.kr/typescript/2020/07/19/typescript-eslint-prettier.html)


# CommonJS vs ES6

## CommonJS

CommonJS는 자바스크립트를 위한 모듈 표준 중 하나입니다. Node.js는 이 CommonJS 모듈 시스템을 사용해 파일과 모듈을 관리하며, 대표적으로 `require`, `module.exports` 구문을 사용해 모듈을 가져오고 내보내는 것이 표준의 특징입니다.

## ES6 (ECMAScript 2015)

ES6는 `import`, `export` 구문을 사용하는 모듈 시스템을 도입한 버전입니다. 

# Bundler

번들러는 웹 개발에서 사용되는 여러 자바스크립트 파일과 리소스를 하나 또는 여러개의 최적화된 파일로 결합하는 도구입니다. 이를 통해 웹 애플리케이션의 로딩 시간을 줄이고 성능을 향상시키며, 브라우저 간 호환성 문제를 해결하는데 도움을 줍니다.

1. (FE) 성능 최적화: 사용자가 웹사이트에 접속할떄, 모든 js파일을 개별적으로 로드하는 것은 많은 시간이 소요되기 때문에, 번들러는 하나의 파일로 결합하여 네트워크 요청의 수를 줄이고 결과적으로 페이지 로딩시간을 단축시킵니다.
2. 코드 최적화: 번들러는 타입스크립트를 순수 js로 컴파일하고 사용되지 않는 코드(트리 쉐이킹)를 제거하여 최종 파일 크기를 줄입니다.
3. 보안 / 유지보수:  번들러는 소스코드를 압축하고 난독화하여 보안을 강화합니다.
4. 환경 일관성: 다양한 브라우저와 장치에서 애플리케이션을 동일하게 실행하기 위해, 번들러는 필요한 폴리필과 트랜스파일링을 적용하여 호환성을 보장합니다.
    - 트랜스파일링: 호환성을 위해 최신 js 문법 -> 구 버전 JS문법 변환
    - 폴리필: 특정 브라우저가 지원하지 않는 기능을 구현하는 코드 조각. 예를 들어 구형 브라우저에서 Promise나 fetch API를 사용할 수 있게함.
5. env: ENV 주입, phase관리(dev, sandbox, prod)

대표적으로 아래와 같은 번들러들이 존재합니다.

- `Webpack`: 매우 강력하고 유연한 플러그인 시스템을 가지고 있으며, 대규모 프로젝트와 복잡한 구성에 적합합니다. 많은 커뮤니티 지원과 풍부한 플러그인 생태계를 가지고 있습니다.

- `Vite`: 최신 프론트엔드 도구로, `esbuild`를 사용하여 매우 빠른 cold start와 HMR을 제공합니다. 간단한 설정과 빠른 빌드 속도로 인기를 얻고 있습니다.

- `Rollup`: 특히 라이브러리 개발에 적합하며, 효율적인 트리 쉐이킹과 코드 분할 기능을 제공합니다. 또한, 결과물이 깔끔하고 작은 번들을 생성합니다.

- `Parcel`: 설정이 필요 없는 번들러로, 빠른 설정과 개발 시작이 가능합니다. 작은 프로젝트나 간단한 웹 애플리케이션에 적합합니다.

- `esbuild`: Go로 작성되어 매우 빠른 빌드 속도를 자랑합니다. 대규모 프로젝트의 빌드 시간을 단축시키는 데 유용합니다.

# TS Compile process

> [TypeScript / How the compiler compiles](https://www.huy.rocks/everyday/04-01-2022-typescript-how-the-compiler-compiles)

> https://www.nextree.io/typescript-compile-process/

![](/images/typescripts/tsc1.png)

높은 차원에서 타입스크립트 컴파일러는 .ts파일을 *.js, *.d.ts, *.map로 생성합니다. 이제 좀 더 자세히 알아보겠습니다.

내부적으로 컴파일 프로세스는 아래와 같은 과정들을 포함합니다.

![](/images/typescripts/tsc2.png)

1. `$ tsc` command
2. Read `tsconfig.json`

tsc는 `src/compiler/program.ts`에 정의되어있는 Program 객체로 컴파일 컨텍스트를 생성한 뒤 tsconfig에 정의된 모든 입력파일과 import를 로드합니다. 이후 각각의 파일을 AST(Abstract Syntax Tree)로 변환시키기 위해 `Parser`(src/compiler/parser.ts)를 호출합니다.

![](/images/typescripts/tsc3.png)

내부적으로 `Parser`는 `Scanner`(src/compiler/scanner.ts) 인스턴스를 생성합니다. `Scanner` 인스턴스는 소스코드를 스캔하여 스트림형태의 토큰(`SyntaxKind`라고 불림)을 생성합니다.

![](/images/typescripts/tsc6.jpeg)

정리하면 Scanner가 타입스크립트로 입력된 코드 문자열은 각각 예약어, 콜론, 부호등의 토큰으로 분리시키고, Parser는 Scanner가 분리해준 토큰을 구문의 구조에 따라 트리 구조로 만들어냅니다.
그러면서 추가로 코드가 올바른 문법인지 분석하여 구문 오류를 잡아냅니다.

3. Scanner: tokenize
4. Parser: Build AST

![](/images/typescripts/tsc4.png)

이후 `Binder`(src/compiler/binder.ts)는 AST 즉 Tree를 traverse하면서, Symbols table인 HashMap을 생성합니다.
바인더는 AST의 각 노드를 방문하여 식별자(변수, 함수, 클래스, 인터페이스 등)를 수집하며, 이 식별자를 key로 하는 value인 symbol을 생성합니다. (아래 그림의 키값이 잘못됨 실제로는 string타입인 identifier로 저장됨)
심볼에는 해당 식별자의 **타입**, 스코프, 선언 위치 등의 메타데이터를 포함합니다. Binder의 Symbol테이블을 통해 식별자 네이밍 컬리션을 방지할 수 있으며, HashMap 특성상 상수시간에 식별자에 대한 symbol에 접근하여 효율적으로 컴파일이 진행될 수 있습니다.

앞서 AST가 tree라고 하였는데, Binder는 symbol table이외에도 AST 트리를 타고 가면서 flow nodes를 생성해 추적합니다.
flow nodes(flow container)는 flow conditional(조건식)을 기준으로 분기가 나눠집니다. (FYI, flow nodes는 이후 타입 체킹과정에서 Type Guards, Type Inference와 같은곳에서 사용됩니다.)

![](/images/typescripts/tsc7.jpeg)

즉 binder는 AST를 통해서 `symbol table`과 `flow nodes`를 생성합니다.

5. Binder: Create Symbol table and Flow Nodes

![](/images/typescripts/tsc7.jpeg)

이후 타입스크립트의 핵심이 Type Check단계가 진행됩니다. (`Type Checker`, src/compiler/types.ts) 
여기에서 타입 추론과 타입체크 등이 발생합니다.

6. TypeChecker: type checking

![](/images/typescripts/tsc5.png)

이후 AST와 type Checker를 통해 자바스크립트 코드로 변환하는 emitter가 실행되며 emitter는 크게 2가지가 존재합니다.

1. emitter.ts (.js): 자바스크립트 소스코드와 source map을 생성합니다.
2. declarationEmitter.ts (.d.ts): 타입 정의 파일을 생성합니다.

이를 emitter과정이 끝나면 최종적으로 3가지 파일형태가 생성됩니다.

- *.js: transpile된 js
- *.d.ts: 타입스크립트 타입 정보
- *.map: 소스맵 파일, 주로 js파일의 minify 또는 transpile과정에서 생성됩니다. 소스 맵 파일로, 트랜스파일된 코드의 각 부분이 원본 소스 코드의 어느 부분에 해당하는지를 나타내는 정보를 담고 있습니다. 개발자 도구에서 원본 소스 코드처럼 보이게 하여 디버깅을 용이하게 합니다.




# npx, nvm

# `d.ts`








## Node JS Architecture

1. namespace
2. build process
3. nodejs internal process
4. v8 핵심만 정리
5. nodejs deployd
