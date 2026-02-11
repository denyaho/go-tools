*`log-status-count`（ログ解析）*

- 目的：Nginx/ALB風のアクセスログからステータスコード集計
- 要件：`stdin` から読み、`2xx/3xx/4xx/5xx` と上位コード（200,404…）をカウント
- ヒント：`strings.Fields`、`map[int]int`
- 伸ばす：時間帯（分単位）で集計したCSVも出す