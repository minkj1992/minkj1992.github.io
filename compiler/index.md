# Typescript Compiler


How typescript api compiles ts to js?
<!--more-->


- I've already wrote a simple typescript compiler [blog post](https://minkj1992.github.io/typescript/) before, this is the more deeper version of typescript compiler api.

## Reference

- [How the TypeScript Compiler Compiles - understanding the compiler internal](https://www.youtube.com/watch?v=X8k_4tZ16qU)

- [Typescript Compiler API](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API)


## Program

- [typescript/src/compiler/program.ts](https://github.com/microsoft/TypeScript/blob/c790dc1dc7ff67e619a5a60fc109b7548f171322/src/compiler/program.ts#L1513)

```TypeScript
import * as ts from "typescript";

const program = ts.createProgram(fileNames, options); // code -> ast
const checker = program.getTypeChecker(); // binding, type check
program.emit() // ast -> code
```

typescript는 program을 생성하고, type check를 한 뒤, emit하는 방식으로 컴파일이 진행됩니다. ms doc에 따르면 아래 helper function을 사용하면 더 쉽게 컴파일이 가능합니다.

```ts
import * as ts from "typescript";

const source = "let x: string  = 'string'";

let result = ts.transpileModule(source, { compilerOptions: { module: ts.ModuleKind.CommonJS }});

console.log(JSON.stringify(result));
```

이를 real world로 예를 든다면 아래 와 같이 isolated 환경에서 dynamic하게 custom code를 받아서 transpile 후 실행 시킬 수 있습니다.

```ts
import { Script, createContext } from "node:vm";
import ts from "typescript";

import AuthorDbWrapper from './dbwrapper.js'

// "import dayjs from \"dayjs\";
// import timezone from \"dayjs/plugin/timezone.js\";
// import utc from \"dayjs/plugin/utc.js\";

// export default async (input: Input): Promise<Record<string, any>> => {
//   dayjs.extend(timezone);
//   dayjs.extend(utc);

//   interface People {
//     name: string;
//     email: string;
//     age: number;
//     birth: Date | null;
//   }
//   // imported from context
//   const authors: People[] = AuthorDbWrapper.select(
//     {
//       deleted_at: undefined
//     }
//   )
//   return { authors }
// }"
const customCode = "import dayjs from \"dayjs\";\nimport timezone from \"dayjs/plugin/timezone.js\";\nimport utc from \"dayjs/plugin/utc.js\";\n\nexport default async (input: Input): Promise<Record<string, any>> => {\n  dayjs.extend(timezone);\n  dayjs.extend(utc);\n\n  interface People {\n    name: string;\n    email: string;\n    age: number;\n    birth: Date | null;\n  }\n  // imported from context\n  const authors: People[] = AuthorDbWrapper.select(\n    {\n      deleted_at: undefined, name:input.name\n    }\n  )\n  return { authors }\n}"

const exportTargets = {
    AuthorDbWrapper,
    // ...   
}


const logContainer = new LogContainer();
const output = ts.transpileModule(customCode, {
    compilerOptions: {
      module: ts.ModuleKind.CommonJS,
      esModuleInterop: true,
      sourceMap: true,
    },
  });

const transpiled = output.outputText;
// Instances of the `vm.Script` class contain precompiled scripts that can be executed in specific contexts.
const script = new Script(transpiled);
  const context = createContext({
    require,
    exports: {},
    module: {
      exports: {},
    },
    console: logContainer,
    setTimeout,
    ...exportTargets, // already compiled codes
  });

const func = script.runInNewContext(context);

// result
const authors = await func("minwook");
```

자 이제 그럼 다시 compiler로 들어가보도록 하겠습니다.

앞서 3줄의 코드에서 보였듯, ts program은 3가지 단계에 거쳐 진행됩니다.

1. Source Code to Data (syntax tree)
2. Type Checking (bind / check)
3. Creating Files (emit)


## 1. Source Code to Data

Source code에서 의미있는 data로 변경하기 위해서는 `syntax tree` 개념이 존재합니다.

> Compiler는 frontend와 backend로 나눠지고, frontend에서 input code를 받아, 의미 있는 내부 형태로 구성하며 이를 기반으로 backend를 통해서 target output으로 변경을 합니다. 이런 구조를 취하는 이유는 adptor pattern 처럼 원하는 target language에 따라서 편하게 새로운 backend 또는 frontend를 갈아 끼우기 위해서 입니다. 또한 이런 유연한 구조를 위해서 내부적으로 AST (abstract syntax tree)를 도입하였습니다.


![](https://upload.wikimedia.org/wikipedia/commons/thumb/c/c7/Abstract_syntax_tree_for_Euclidean_algorithm.svg/800px-Abstract_syntax_tree_for_Euclidean_algorithm.svg.png)


Syntax tree는 크게 2가지로 나뉩니다.

1. Scanner (src/compiler/scanner.ts), aka `Lexer`
2. Parser (src/compiler/parser.ts)


Scanner은 text를 syntax tokens으로 변경합니다.

![](/images/tsc_scanner.png)

Parsesr은 syntax tokens를 tree로 변경합니다.

![](/images/tsc_parser.png)

실제 코드를 보면 

1. parser가 source code를 받음
2. parser가 scanner instance 생성
3. parser가 scanner의 [`textToToken`](https://github.com/microsoft/TypeScript/blob/c790dc1dc7ff67e619a5a60fc109b7548f171322/src/compiler/scanner.ts#L212C7-L212C18) 를 통해 source code의 text 청크들 syntax token으로 변환
4. parser가 syntax token의 stream들을 tree 형식으로 구성 [`parseJsonText`](https://github.com/microsoft/TypeScript/blob/c790dc1dc7ff67e619a5a60fc109b7548f171322/src/compiler/parser.ts#L1624)

정리하면 

- Scanner: 소스코드를 컴파일러가 이해할 수 있는 토큰 시퀀스로 변환합니다.
- Parser: 토큰 스트림을 분석하여 AST를 생성합니다.
    - 함수 선언, 변수 선언, 조건문, 반복문 등의 구조를 식별합니다.

## 2. Type Checking
> Checking the syntax tree

Parsing process를 통해 stream of Syntax tree가 생성되면 


1. **Binder**, Syntax -> Symbols
    - Converts identifiers in syntax tree to symbols
2. **Type Check**, Use binder and syntax tree to look for issues in code.

### 2.1. Type Checking: Binder

Binder, post-parser grab bag는 아래와 같은 역할을 수행합니다.

1. identifiers가 정의되어있는 Symbol Tables 생성
2. Sets up 'parent' on all syntax tree nodes
3. Make flow node for narrowing
4. Validate script vs module conformance

1번에서 symbol이란 scope에 따른 identifier를 말합니다. 

![](/images/tsc_symbols.png)


2번이 무슨말인지 찾아보니, AST는 비록 트리구조이지만, parser에 의해 생성된 AST는 아직 각 노드에 부모 노드에 대한 정보가 없어 상위 context를 파악하기 어렵다고 합니다. Binder는 AST를 방문(매우 무거운 작업)하면서 각 노드에 대한 부모 노드 정보를 설정할 수 있습니다.

3번의 flow nodes는 위에서 설명한 대로 scope를 찾아가면서 if conditional, functional scope 처럼 flow를 파악하는 것을 뜻합니다. 

![](/images/tsc_binder1.png)

이렇게 flow graph가 만들어지게 되면 나중에 type check시에 typescript는 scope를 narrow하면서 check할 수 있습니다. 이 덕분에 typeof 같은 Type Guard가 typescript 타입 체크에 영향을 줄 수 있습니다.

![](/images/tsc_binder2.png)

Binder가 Symbols table를 생성하면 `Program.emit()`을 호출하여 Emit worker가 AST가 javascript source code로 변환될 수 있도록 합니다.

Emitter가 running할때 `getDiagnostics()`를 호출하며, 이 함수를 통해 `Type Checker` (src/compiler/checker.ts)가 실행됩니다.

Emitter는 AST를 traverse(walk)하며 each node에 상응하는 check function를 실행합니다.


### 2.2. Type Checking: Type Checker


모든 syntax tree node에는 그에 상응하는 check function이 정의되어 있습니다. 이를 통해서 syntax tree에 정의되어있는 node들을 타고 가면서 type check가 가능합니다.

![](/images/tsc_checker.png)

타입스크립트는 source, target로 타입을 구분하는데, const greet: string = "Hello World'의 경우에 string은 source, 'Hello World'는 target이 되어 각각에 대해 check됩니다.

방대한 양의 type check이다 보니, 이번 post에서는 여기까지만 파악하고 다음으로 넘어가도록 하겠습니다.

## 3. Creating Files
> Syntax tree -> JS

TS에는 여러 transformers가 존재하며, 지정된 버전에 따라서 실행되는 transformers가 달라집니다. 

![](/images/tsc_transformers.png)

이 과정을 거쳐서 typescript는 *.js, *.map, *.d.t를 만들어냅니다.


