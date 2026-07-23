# gostlings 设计文档

日期：2026-07-23
状态：已获用户批准

## 目标

仿照本地 rustlings 项目，为 Go 学习者提供一套可执行的填空/修错练习题。
核心版：15 个主题、51 题，英文注释，附带完整参考答案。

## 决策（用户已确认）

- 规模：核心版 ~45 题（实际 51）
- 运行方式：纯 `go run` + README，不写 CLI
- 语言：英文注释（代码内字符串与预期输出同样用英文；README 用中文）
- 附带 `solutions/` 参考答案

## 仓库结构

位置：`/Users/wangxinalex/SelfStudy/Rust/gostlings/`（rustlings 的同级目录）

```
gostlings/
├── go.mod                  # module gostlings, go 1.23
├── README.md               # 做题顺序表 + 运行方法 + 与 Go Tour 章节映射
├── check.sh                # 跑通全部 solutions 的验证脚本
├── exercises/
│   └── <主题>/<题名>/main.go    # 14_testing 主题为 _test.go
└── solutions/              # 与 exercises 完全同构
```

- 单 module、每题独立目录（package main），避免同目录多 main 冲突，gopls 友好
- 只用标准库，零外部依赖
- exercises/ 中的题故意编译不过或逻辑缺失，属预期行为
- solutions/ 每题必须 `go run`（testing 主题用 `go test`）通过

## 题目格式约定

每题一个 `main.go`，题头注释统一四段式（英文）：

```go
// Concept: <the single concept this exercise teaches>
// Task: <what to do>
// Expected output: <what the program should print>
// Hint: <one-line hint + Go Tour section>
```

- 用 `// TODO:` 标出需要修改/补充的位置（英文）
- 题型一：修错——代码有编译错误或逻辑错误，修到能跑
- 题型二：填空——删掉关键语句（留 TODO），补上
- 每题只考一个知识点，难度循序渐进

## 题目清单（51 题）

| 主题 | 题目 | 考点 |
|---|---|---|
| 00_intro | intro1 | 修复缺失的 fmt 导入和 Println 调用 |
| | intro2 | 修复 package/func main 结构 |
| 01_variables | variables1 | var 声明 |
| | variables2 | 短变量声明 := |
| | variables3 | 零值 |
| | variables4 | 常量与 iota |
| 02_functions | functions1 | 定义被调用的函数 |
| | functions2 | 参数类型匹配 |
| | functions3 | 多返回值 |
| | functions4 | 可变参数 ...int |
| 03_control_flow | if1 | if 条件（含短语句） |
| | for1 | for 循环求和 |
| | switch1 | switch 分支 |
| | defer1 | defer 执行顺序 |
| 04_pointers | pointers1 | & 取地址、* 解引用 |
| | pointers2 | 通过指针修改值 |
| | pointers3 | nil 指针与 new |
| 05_slices | slices1 | 切片字面量、len/cap |
| | slices2 | append |
| | slices3 | 切片表达式 s[lo:hi] |
| | slices4 | 共享底层数组 / copy |
| 06_maps | maps1 | 创建、增删改查 |
| | maps2 | comma-ok 判断键存在 |
| | maps3 | range 遍历与 delete |
| 07_structs | structs1 | 定义并初始化结构体 |
| | structs2 | 字段名初始化 vs 位置初始化 |
| | structs3 | 嵌套结构体 |
| 08_methods | methods1 | 值接收者方法 |
| | methods2 | 指针接收者修改字段 |
| | methods3 | 选择正确的接收者类型 |
| 09_interfaces | interfaces1 | 实现接口 |
| | interfaces2 | 实现 fmt.Stringer |
| | interfaces3 | 空接口 any + 类型断言 |
| | interfaces4 | comma-ok 类型断言 / type switch |
| 10_errors | errors1 | 返回并检查 error |
| | errors2 | errors.New / fmt.Errorf |
| | errors3 | %w 包装 + errors.Is |
| | errors4 | errors.As + 自定义错误类型 |
| 11_generics | generics1 | 泛型函数（cmp.Ordered） |
| | generics2 | 自定义约束 interface |
| | generics3 | 泛型类型 Stack[T] |
| 12_goroutines | goroutines1 | go 关键字启动协程 |
| | goroutines2 | sync.WaitGroup 等待 |
| | goroutines3 | 循环变量捕获（传参解决） |
| 13_channels | channels1 | 无缓冲收发 |
| | channels2 | 缓冲通道 |
| | channels3 | range + close |
| | channels4 | select 多路复用 |
| 14_testing | testing1 | 编写 TestXxx（`go test` 验证） |
| | testing2 | 表驱动测试 |
| | testing3 | t.Run 子测试 |

## README 内容

- 项目简介（一段）
- 前置要求：Go 1.23+
- 使用方法：按主题顺序做题，`go run ./exercises/<主题>/<题名>` 验证；卡住看 `solutions/`
- 做题顺序表：主题 → 题数 → Go Tour 章节映射（仿 rustlings 的 Book Chapter 表）

## check.sh

约 15 行 bash：遍历 solutions/ 下每个题目录，14_testing 下的题目录用 `go test`，其余用 `go run`；任一失败即退出非零。

## 验证计划

1. `sh check.sh` 全绿（51 个 solutions 全部通过）
2. 抽查若干 exercises 题目确认按预期编译失败
3. `gofmt -l .` 无输出（全部格式化）

## 明确不做（YAGNI）

- CLI 工具（list/run/hint/watch）
- quiz 综合题
- context、反射、mutex、time 主题
- CI 配置、lint 配置
