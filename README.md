# nocjk

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/nocjk.svg)](https://pkg.go.dev/github.com/aethiopicuschan/nocjk)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/nocjk)](https://goreportcard.com/report/github.com/aethiopicuschan/nocjk)
[![CI](https://github.com/aethiopicuschan/nocjk/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/nocjk/actions/workflows/ci.yaml)

`nocjk` is a simple CLI tool and library to detect CJK (Chinese, Japanese, and Korean) text.

## Installation

### As a CLI tool:

```sh
go install github.com/aethiopicuschan/nocjk/cmd/nocjk@latest
nocjk .
```

When CJK text is detected, the CLI exits with error code 1.

#### Ignore

If you want to ignore the detection of a specific language, you can use the following option.

- `--ignore-chinese` to ignore Chinese text
- `--ignore-japanese` to ignore Japanese text
- `--ignore-korean` to ignore Korean text

You can also define ignore rules using a `.nocjkignore` file in your project root. The format of `.nocjkignore` is fully compatible with `.gitignore`. Files and directories matching the patterns in this file will be skipped during CJK text detection. This is especially useful when you want to exclude certain generated files or third-party code from being analyzed.

### As a library:

```sh
go get -u github.com/aethiopicuschan/nocjk/pkg/nocjk
```

`nocjk` provides `FindChineseLines`, `FindJapaneseLines`, `FindKoreanLines` and `FindCJKLines` functions to detect CJK text in a string. More information can be found in the [documentation](https://pkg.go.dev/github.com/aethiopicuschan/nocjk).

## Usage with GitHub Actions

You can easily integrate `nocjk` into your GitHub Actions workflows to automatically detect CJK text during code changes. Here is a basic example:

```yaml
jobs:
  check-cjk:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - name: Detect CJK text
        uses: aethiopicuschan/nocjk/actions/nocjk-action@v1.0.0
```

See [nocjk-action on GitHub](https://github.com/aethiopicuschan/nocjk/tree/main/actions/nocjk-action) for more details.
