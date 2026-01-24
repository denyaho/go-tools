### `echo-plus`（CLI基礎）

- 目的：標準入力を読み、行番号付きで出力
- 要件：`n`（開始行番号）、`s`（区切り文字）フラグ対応、空行スキップ `-skip-empty`
- ヒント：`flag`、`bufio.Scanner`、`fmt`
- 伸ばす：`-json` で `{line, text}` をJSON出力