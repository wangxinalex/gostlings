# Strings: common string operations

Work through 1–4 in order. This covers the string APIs used most often in
day-to-day code; format verbs live in `28_fmt` and regular expressions in
`31_regexp`.

| Exercise | Focus | Goal |
| --- | --- | --- |
| strings1 | `strings.Contains` | Check whether a substring exists |
| strings2 | `strings.Join` / `strings.Split` | Join and split |
| strings3 | `strings.TrimSpace` / `strings.ReplaceAll` | Clean input |
| strings4 | `strings.Builder` | Build long strings efficiently |

Note: do not repeatedly concatenate long strings with `+` in a loop; prefer
`strings.Builder` when appending repeatedly. Use `strings.ReplaceAll` instead
of a hand-written replace loop.

Reference: [strings package](https://pkg.go.dev/strings)
