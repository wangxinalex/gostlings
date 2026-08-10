# Testing 递进路线

本章从最小可读测试开始，逐步练习真实项目中常见的 fixture、依赖替身、HTTP 测试、错误分类、fuzz 和 benchmark。每道题都应该先让测试失败，再修改测试或测试辅助代码使它通过。

| 题目 | 重点 | 练习目标 |
| --- | --- | --- |
| testing1 | 单元测试 | 认识 `go test` 和基本断言 |
| testing2 | 表格测试 | 用多个输入覆盖边界 |
| testing3 | 子测试 | 用 `t.Run` 组织输出 |
| testing4 | 测试辅助函数 | 使用 `t.Helper` 保持失败位置可读 |
| testing5 | 文件 fixture | 用 `t.TempDir` 隔离临时文件 |
| testing6 | 环境 fixture | 用 `t.Setenv` 自动恢复环境变量 |
| testing7 | fake + 依赖注入 | 不连接数据库也能测试业务逻辑 |
| testing8 | `httptest` | 不监听真实端口测试 handler |
| testing9 | fuzz | 用不变量发现边界输入 |
| testing10 | 错误分类断言 | 用 `errors.Is` 测试稳定语义 |
| testing11 | benchmark | 用 `b.ResetTimer` 和 `b.ReportAllocs` 测量热路径 |

常用命令：

```sh
go test ./exercises/14_testing/testing4
go test -v ./exercises/14_testing/testing8
go test -run=^$ -fuzz=FuzzReverse -fuzztime=2s ./exercises/14_testing/testing9
go test -bench=. -benchmem ./exercises/14_testing/testing11
```

参考： [Add a test](https://go.dev/doc/tutorial/add-a-test)、[Fuzzing](https://go.dev/doc/tutorial/fuzz)、[`testing` package](https://pkg.go.dev/testing)。
