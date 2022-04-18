# Google like python on vscode


Describes python google style editor settings on `vscode`

<!--more-->
<br />

This is base vscode settings related format, python 

*settings.json*
```json
{
...
    "editor.fontSize": 13,
    "editor.tabCompletion": "on",
    "editor.suggestSelection": "first",
    "vsintellicode.modify.editor.suggestSelection": "automaticallyOverrodeDefaultValue",

    "[python]": {
        "editor.defaultFormatter": "ms-python.python",
        "editor.formatOnSave": true,
        "editor.codeActionsOnSave": {
            "source.organizeImports": true,
        },
    },
}
```

## Yapf

```json
{
    ...
    "python.formatting.provider": "yapf",
    // or "python.linting.pylintPath": "${workspaceFolder}/.venv/bin/yapf",
    "python.formatting.yapfPath": "/Library/Frameworks/Python.framework/Versions/3.7/bin/yapf",
    ...
    "python.formatting.yapfArgs": [
        "--style",
        "google",
    ],
}
```

## Isort
> It helps to sort automatically whenever file is saved

- [google import formatting guide](https://google.github.io/styleguide/pyguide.html#313-imports-formatting)


*vscode's settings.json*
```json
{
...

    "python.sortImports.path": "/Library/Frameworks/Python.framework/Versions/3.7/bin/isort",
    // or
    // "python.sortImports.path": "${workspaceFolder}/.venv/bin/isort",
    "python.sortImports.args": [
        "--settings-file=${workspaceFolder}/.isort.cfg",
    ],
...
}
```

*${workspaceFolder}/.isort.cfg*
```apacheconf
[settings]
py_version=37
profile=google
src_paths=api,common,core,infra,logs,tests
multi_line_output=3
use_parentheses=True
force_single_line=False

# profile google default
# force_single_line: True
# force_sort_within_sections: True
# lexicographical: True
# single_line_exclusions: ('typing',)
# order_by_type: False
# group_by_package: True
```

readable references below
- [isort options](https://pycqa.github.io/isort/docs/configuration/options.html)
- [isort configs](https://pycqa.github.io/isort/docs/configuration/config_files.html)



## Pylint
> TODO: force to avoid string double quotes

```json
{
    ...

    "python.linting.enabled": true,
    "python.linting.lintOnSave": true,
    "python.linting.pylintEnabled": true,
    // or "python.linting.pylintPath": "${workspaceFolder}/.venv/bin/pylint",
    "python.linting.pylintPath": "/Library/Frameworks/Python.framework/Versions/3.7/bin/pylint",
    "python.linting.pylintArgs": [
        "--load-plugins",
        "pylint_django",
        "pylint_quotes", // related string quotes
    ],
    ...
}

```

## IntelliSense

> *general term for various code editing features including: code completion, parameter info, quick info, and member lists*


