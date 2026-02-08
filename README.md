# nocjk

[![License: MIT](https://img.shields.io/badge/License-MIT-brightgreen?style=flat-square)](/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/aethiopicuschan/nocjk.svg)](https://pkg.go.dev/github.com/aethiopicuschan/nocjk)
[![Go Report Card](https://goreportcard.com/badge/github.com/aethiopicuschan/nocjk)](https://goreportcard.com/report/github.com/aethiopicuschan/nocjk)
[![CI](https://github.com/aethiopicuschan/nocjk/actions/workflows/ci.yaml/badge.svg)](https://github.com/aethiopicuschan/nocjk/actions/workflows/ci.yaml)
[![nocjk Action](https://img.shields.io/badge/GitHub%20Action-nocjk-blue?logo=github-actions)](https://github.com/marketplace/actions/nocjk)

Detect CJK (Chinese, Japanese, and Korean) text in your repository.

- GitHub Action: fail the workflow when CJK text is detected
- CLI: scan files and exit with code 1 when CJK text is detected
- Library: functions to find CJK lines in strings

## GitHub Action

This action runs the `nocjk` CLI built from the same Git reference as the action itself, ensuring the action version and the CLI version always match.

### Usage

```yaml
jobs:
  check-cjk:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Detect CJK text
        uses: aethiopicuschan/nocjk@v1
        with:
          args: "."
```

### Inputs

| Name            | Description                          | Default |
| --------------- | ------------------------------------ | ------- |
| args            | Paths or arguments passed to `nocjk` | .       |
| ignore_chinese  | Ignore Chinese text                  | "false" |
| ignore_japanese | Ignore Japanese text                 | "false" |
| ignore_korean   | Ignore Korean text                   | "false" |

### Exit Codes

- 0: No CJK text found (workflow succeeds)
- 1: CJK text detected (workflow fails)

### Ignore Rules

You can define ignore rules using a `.nocjkignore` file in your project root. The format is fully compatible with `.gitignore`. Files and directories matching the patterns in this file will be skipped during CJK text detection.

## CLI

### Installation

```sh
go install github.com/aethiopicuschan/nocjk/cmd/nocjk@v1
```

### Usage

```sh
nocjk .
```

When CJK text is detected, the CLI exits with error code 1.

### Options

- `--ignore-chinese` to ignore Chinese text
- `--ignore-japanese` to ignore Japanese text
- `--ignore-korean` to ignore Korean text

`.nocjkignore` is supported for skipping specific files and directories.

## Library

### Installation

```sh
go get -u github.com/aethiopicuschan/nocjk/pkg/nocjk
```

See [GoDoc](https://pkg.go.dev/github.com/aethiopicuschan/nocjk/pkg/nocjk) for usage instructions.
