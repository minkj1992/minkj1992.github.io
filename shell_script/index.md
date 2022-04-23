# Depp dive into shell script line by line



Useful shell script commands :)
<!--more-->
<br />

## Shebang
- `#!`: `Shebang`
- `#!/usr/bin/<bash|python|perl|php...>`
  - 스크립트를 실행시켜줄 인터프리터의 절대경로를 지정하는 역할
- `#!/usr/bin/env <languate>`
  - env는 환경 변수에서 지정한 언어의 위치를 찾아서 실행됩니다. 다양한 환경에서 실행되는 스크립트라면 "env"를 사용하는 것이 좋습니다.
  - e.g. `#!/usr/bin/env bash`


## set commands
> [refs](https://kvz.io/bash-best-practices.html)

```sh
#!/usr/bin/env bash

set -o nounset -o errexit -o errtrace -o pipefail
```


만일, 미리 선언되지 않은 변수를 사용했을 때 스크립트를 종료시키고 싶은 경우
```sh
$ set -o nounset
or
$ set -u
```

exit if a command yields a non-zero exit status. (exit when the command fails.)
```sh
$ set -o errexit
or
$ set -e
```

prints commands arguments during execution. Useful for debugging

```sh
$ set -o xtrace
or
$ set -x
```

파이프 사용시, 이전 단계의 오류(non-zero exit code)를 승계하도록 하는 설정
파이프 사용시 오류 코드(non-zero exit code)를 이어받는다.
- e.g. `$ mysqldump | gzip` commands fails, than gzip commands returns non-zero
- `-e` option과 함께 사용된다.
```sh
$ set -o pipefail
```

When errtrace is enabled, the ERR trap is also triggered when the error (a command returning a nonzero code) occurs inside a function or a subshell.
```sh
$ set -o errtrace
# or
$ set -E
```

- *e.g.*

```sh
#!/bin/bash

set -o errtrace

function x {
    echo "X begins."
    false
    echo "X ends."
}

function y {
    echo "Y begins."
    false
    echo "Y ends."
}

trap 'echo "ERR trap called in ${FUNCNAME-main context}."' ERR
x
y
false
true
```

- if `set -o errtrace` enabled
```
X begins.
ERR trap called in x.
X ends.
Y begins.
ERR trap called in y.
Y ends.
ERR trap called in main context.
```
- disabled

```
X begins.
X ends.
Y begins.
Y ends.
ERR trap called in main context.
```

## trap
> `trap` defines and activates handlers to run when the shell receives signals or other special conditions.

```sh
# trap [action] [signal]

trap 'echo "Fatal: Exits abnormally at line #$(caller 0 || echo ${LINENO})" >&2' ERR
```

- ERR

A SIGNAL_SPEC of ERR means to execute [action] each time a command's failure would cause the shell to exit when the -e option is enabled.


- `#$(caller 0 || echo ${LINENO})`

## redirection (`>`)

```sh
# >1은 커맨드의 표준 출력을 다음에 나오는 파일 디스크립터에 전달한다.
$ > 1
# 쉘 스크립트의 표준 에러를 다음에 나오는 파일 디스크립터에 전달한다.
$ >2
# >&2는 모든 출력을 강제로 쉘 스크립트의 표준 에러로 출력한다. 
$ >&2
```




