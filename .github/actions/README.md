# nocjk

Detect CJK (Chinese, Japanese, and Korean) text in your repository using GitHub Actions.

This action runs the `nocjk` CLI built from the same Git reference as the action itself.  
This guarantees that the action version and the CLI version always match.

## Features

- Detect Chinese, Japanese, and Korean text in files
- Fail the workflow when CJK text is found
- Ignore specific languages via inputs
- Support `.nocjkignore` (compatible with `.gitignore`)
- No version mismatch between the action and the CLI

## Usage

```yaml
jobs:
  check-cjk:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4
      - name: Detect CJK text
        uses: aethiopicuschan/nocjk@v2
```

## Inputs

| Name            | Description                        | Default |
| --------------- | ---------------------------------- | ------- |
| args            | Paths or arguments passed to nocjk | .       |
| ignore_chinese  | Ignore Chinese text                | false   |
| ignore_japanese | Ignore Japanese text               | false   |
| ignore_korean   | Ignore Korean text                 | false   |

## Exit Codes

- 0: No CJK text found
- 1: CJK text detected

## Ignore Rules

You can define ignore rules using a `.nocjkignore` file in your project root.  
The format of `.nocjkignore` is fully compatible with `.gitignore`.

## License

MIT
