# gostlings

仿 [rustlings](https://github.com/rust-lang/rustlings) 的 Go 入门练习题：每题一个小程序，
修好它、跑通它，你就学会了这个知识点。

## 前置要求

Go 1.23+（`go version` 确认）

## 怎么做题

1. 按下表顺序做题
2. 打开 `exercises/<主题>/<题名>/main.go`，读题头注释（Concept / Task / Expected output / Hint）
3. 修改 `// TODO:` 标记的地方
4. 运行验证：

   ```sh
   go run ./exercises/01_variables/variables1
   ```

   （14_testing 主题用 `go test ./exercises/14_testing/testing1`）
5. 输出符合预期 = 过关。卡住了看 `solutions/` 里的同路径答案

## 做题顺序与 Tour 章节映射

| 主题 | 题数 | Go Tour 章节 |
|---|---|---|
| 00_intro | 2 | Basics 1 |
| 01_variables | 4 | Basics 8-12 |
| 02_functions | 4 | Basics 4-7 |
| 03_control_flow | 4 | Flowcontrol 1-13 |
| 04_pointers | 3 | Moretypes 1; Methods 5 |
| 05_slices | 5 | Moretypes 7-15; [sort](https://pkg.go.dev/sort) |
| 06_maps | 3 | Moretypes 19-22 |
| 07_structs | 3 | Moretypes 2-5 |
| 08_methods | 3 | Methods 1-6 |
| 09_interfaces | 4 | Methods 9-17 |
| 10_errors | 4 | Methods 19-20 |
| 11_generics | 3 | Generics 1-2 |
| 12_goroutines | 3 | Concurrency 1 |
| 13_channels | 6 | Concurrency 2-6 |
| 14_testing | 3 | [Testing tutorial](https://go.dev/doc/tutorial/add-a-test) |
| 15_context | 3 | [context](https://pkg.go.dev/context) |
| 16_sync | 3 | [sync](https://pkg.go.dev/sync) |
| 17_panic_recover | 2 | [builtin](https://pkg.go.dev/builtin) |
| 18_embedding | 2 | [Effective Go: Embedding](https://go.dev/doc/effective_go#embedding) |
| 19_json | 3 | [encoding/json](https://pkg.go.dev/encoding/json) |
| 20_io | 2 | [io](https://pkg.go.dev/io) |
| 21_time | 2 | [time](https://pkg.go.dev/time) |
| 22_strconv | 2 | [strconv](https://pkg.go.dev/strconv) |

## 校验全部答案

```sh
sh check.sh   # 跑通 solutions/ 下全部 73 题
```
