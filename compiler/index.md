# Typescript Compiler


How typescript api compiles ts to js?
<!--more-->


- I've already wrote a simple typescript compiler [blog post](https://minkj1992.github.io/typescript/) before, this is the more deeper version of typescript compiler api.

## Reference

- [How the TypeScript Compiler Compiles - understanding the compiler internal](https://www.youtube.com/watch?v=X8k_4tZ16qU)

- [Typescript Compiler API](https://github.com/microsoft/TypeScript/wiki/Using-the-Compiler-API)


## 1. program

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



