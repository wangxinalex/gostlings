# Flag：命令行参数解析

按 1–2 顺序完成。练习用 `flag.NewFlagSet` 解析 flag，避免依赖全局
`flag.CommandLine`，方便测试和复用。验证用 `go test` 或 `check.sh`：

```sh
sh check.sh exercises/29_flag/flag1
go test ./exercises/29_flag/flag2
```

| 题目 | 重点 | 练习目标 |
| --- | --- | --- |
| flag1 | `flag.NewFlagSet` + `String/Int/Bool` | 解析多个类型的 flag |
| flag2 | `fs.Args()` | flag 与位置参数混用 |

参考：[flag package](https://pkg.go.dev/flag)
