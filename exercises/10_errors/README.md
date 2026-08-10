# Errors 递进路线

本章按真实开发中错误处理的使用频率排列。先学会返回和检查错误，再学习如何在跨层调用时保留错误原因，最后练习验证、重试和应用边界转换。

| 题目 | 重点 | 练习目标 |
| --- | --- | --- |
| errors1 | 返回并检查 error | 调用者必须处理失败分支 |
| errors2 | 输入校验 | 用非 nil error 表示非法输入 |
| errors3 | sentinel + `%w` + `errors.Is` | 包装错误但保留分类 |
| errors4 | 自定义错误 + `errors.As` | 提取结构化错误信息 |
| errors5 | 标准库错误类型 + `errors.As` | 保留底层 `strconv.NumError` |
| errors6 | `errors.Join` | 同时报告多个独立校验错误 |
| errors7 | 自定义 wrapper + `Unwrap` | 让业务错误参与 `errors.Is` |
| errors8 | 应用边界转换 | 内部错误映射成稳定的用户语义 |
| errors9 | 可重试错误 | 只重试明确允许重试的失败 |
| errors10 | 错误组合综合题 | 跨层包装后仍可识别多个原因 |

解题时始终区分两件事：错误字符串用于日志和诊断，`errors.Is`/`errors.As` 用于程序逻辑。除非题目明确要求展示文本，不要通过比较 `err.Error()` 来判断错误类别。

参考： [Go Error Handling](https://go.dev/doc/effective_go#errors)、[`errors` package](https://pkg.go.dev/errors)。
