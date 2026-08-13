# Strings：常用字符串操作

按 1–4 顺序完成。这里覆盖日常编码最常用的字符串 API；格式化动词在
`28_fmt`，正则匹配在 `31_regexp`。

| 题目 | 重点 | 练习目标 |
| --- | --- | --- |
| strings1 | `strings.Contains` | 判断子串是否存在 |
| strings2 | `strings.Join` / `strings.Split` | 拼接与切分 |
| strings3 | `strings.TrimSpace` / `strings.ReplaceAll` | 清洗输入 |
| strings4 | `strings.Builder` | 高效拼接多个片段 |

注意：不要在循环里用 `+` 反复拼接长字符串；需要反复追加时优先
`strings.Builder`。用 `strings.ReplaceAll` 代替手写循环替换。

参考：[strings package](https://pkg.go.dev/strings)
