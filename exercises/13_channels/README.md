# Channels：从原始协议到并发组合

这一章用 50 个小练习建立 channel 的生命周期模型：谁发送、谁接收、谁
知道不会再发送、谁负责关闭，以及调用者提前放弃时谁负责取消。这里坚持
使用 raw channels；`context.Context`、`sync`、计时器和综合并发模式属于后续
章节。完成本章后，把更完整的 pipeline 工作交给
[`24_concurrency_patterns`](../24_concurrency_patterns)。

本章不构造可复用的 timer/ticker 抽象；少数超时或取消测试/示例仍可能使用
`time.After`。`21_time` 专门负责 `time.Timer`/`time.Ticker` 的构造、重置、
生命周期管理和周期性设计。

## 所有权模型

- 发送方通常负责关闭；只有知道“不会再发送”的一方才能 `close(ch)`。
- 接收方通常只 `range ch` 或使用 comma-ok 接收，不替生产者关闭输出。
- 关闭表示生命周期结束，不是一个特殊数据值；缓冲区中的值仍会先被读完。
- `done` channel 表示完成，数据 channel 传值，`stop`/`cancel` channel 广播取消。
- 一个 shared output 只能由明确的 coordinator 在所有发送者退出后关闭。
- 每个可能阻塞的 receive、send、token wait 都要在取消协议中有退出路径。

## 技能矩阵与难度

| 技能 | 练习 |
| --- | --- |
| 无缓冲/缓冲、关闭、comma-ok、drain | 1–5 |
| generator、方向约束、`select`、`default`、超时 | 6–11 |
| 可取消 receive/send、done、result、stop、join | 12–20 |
| fan-out、fan-in、worker pool、错误与方向 | 21–30 |
| 首结果、request/reply、共享输出协调、控制状态 | 31–40 |
| nil channel、取消生产者、信号量、限速、服务收尾 | 41–50 |

| 难度带 | 练习 | 目标 |
| --- | --- | --- |
| 基础 | 1–10 | 理解阻塞、关闭和一次性的 `select` |
| 过渡 | 11–20 | 把超时、取消和完成组成可等待的生命周期 |
| 组合 | 21–30 | 组织 worker、共享输出和 backpressure |
| 高级 | 31–40 | 处理首结果、reply、状态机和协调关闭 |
| Capstone | 41–50 | 在 bounded、ordered、cancelable 服务中保持所有权清晰 |

## 1–50 学习地图

| # | 模式 |
| ---: | --- |
| 1 | 无缓冲交接：发送和接收必须在不同 goroutine |
| 2 | 缓冲 channel 解耦发送和接收 |
| 3 | 关闭 channel 让 `range` 停止 |
| 4 | comma-ok 区分关闭与真实零值 |
| 5 | drain 关闭 channel 前已缓冲的值 |
| 6 | producer 拥有并关闭 generator 输出 |
| 7 | directional channel 表达 generator 所有权 |
| 8 | `select` multiplex 多个输入 |
| 9 | `select` + `default` 非阻塞接收 |
| 10 | `select` + `default` 非阻塞发送 |
| 11 | `select` 超时 |
| 12 | receive 和 send 两端都可取消的 relay |
| 13 | 关闭 done channel 广播完成 |
| 14 | capacity-one 的异步 result |
| 15 | 分离 result 与 done 信号 |
| 16 | 可取消 producer，释放阻塞发送 |
| 17 | 两端可阻塞的可取消 forwarder |
| 18 | 用 raw channels 启动、广播停止并 join workers |
| 19 | timeout → cancel → join 的完整清理 |
| 20 | stop、join workers、再报告 graceful shutdown |
| 21 | fan-out workers 共享 jobs，再 fan-in results |
| 22 | fan-in 多输入并协调关闭输出 |
| 23 | 可取消 fan-in 的两端退出路径 |
| 24 | worker pool 的 jobs receive 与 result send 都可取消 |
| 25 | fan-in 处理 nil、已关闭和带缓冲输入 |
| 26 | jobs producer、workers、results closer 的基本 pool |
| 27 | 携带 index，恢复 worker pool 的输入顺序 |
| 28 | 首个错误停止新工作并 join 所有 workers |
| 29 | 通过 channel direction 暴露最小接口 |
| 30 | 有界 jobs channel 提供 backpressure |
| 31 | 首个完成结果取消 peers 并等待退出 |
| 32 | request 携带私有 reply channel |
| 33 | 多个 forwarder 由一个 coordinator 关闭共享输出 |
| 34 | worker 数为零时定义 pool 行为 |
| 35 | 可取消 worker group 保护 receive 与 send |
| 36 | result envelope 同时传值和错误 |
| 37 | 下游放弃时通知上游停止 |
| 38 | 一个 done channel 广播给多个 observer |
| 39 | pause/resume/stop 控制命令与工作共存 |
| 40 | relay 从阻塞 receive 或 send 中取消 |
| 41 | 关闭输入设为 nil，避免 closed channel 永远 ready |
| 42 | stop 释放无人接收的 producer send |
| 43 | buffered token channel 作为 semaphore，恢复输入顺序 |
| 44 | semaphore 获取和结果发布都可取消 |
| 45 | caller 提供 tokens 的 raw-channel rate limiter |
| 46 | rate limiter 的 input/token/output 全部可取消 |
| 47 | raw-channel service 的 results/done 由 coordinator 关闭 |
| 48 | bounded、ordered、cancelable worker pool |
| 49 | 首个 job error 取消新工作并由单一 owner 关闭 cancel |
| 50 | bounded request/reply 服务：backpressure、取消、顺序、错误和关闭 |

## 取消时序

正常关闭：

```text
producer -- values --> out -- range --> consumer
   |                       |
   +------ close(out) -----+--> range ends after buffered values drain
```

下游提前退出时，发送端必须可取消；取消后还要 join：

```text
caller -- close(stop) --> producer/forwarders -- exit --> coordinator
caller <-- return only after done/result closure ----------------------+
```

fan-in、worker pool 和 capstone 的共同顺序是：停止继续投递，等待所有
发送者退出，然后由一个 coordinator 关闭 shared output（以及需要的 done）。
不能靠 `Sleep` 猜测 goroutine 已经结束。

## Hints 是严格边界

每个 `main.go` 的 Concept / Task / Expected behavior / Hint 是该题的
最小协议，不是可选建议。尤其是 hint 指定的关闭者、channel direction、
取消覆盖范围、buffer 容量、index、reply 次数和 join 顺序都应保持；不能用
空 slice、提前返回、`Sleep` 或忽略错误制造“看起来通过”的实现。遇到阻塞，
先画出发送者、接收者、关闭者和每条退出边，再改最小的一处。

## Checker 与练习顺序

单题、全章和 starter 验证：

```sh
sh check.sh exercises/13_channels/channels1
sh check.sh exercises/13_channels --run-all
sh scripts/verify_exercise_starters.sh exercises/13_channels
```

普通 checker 会拒绝仍含 `TODO:` 的 exercise target；starter verifier 则
要求每个未修改 starter 的 focused behavioral test 确实失败。`--race` 只对
`race.list` 中精确匹配的路径追加 `-race`，`--race-all` 对所选目标全部追加，
后者是昂贵的全量审计：

```sh
sh check.sh solutions/16_sync/sync1 --race
sh check.sh solutions --run-all --race
sh check.sh solutions --run-all --race-all
```

实用顺序是先完成 1–20 的语义和退出协议，再完成 21–40 的组合，最后做
41–50 的边界与服务收尾。`15_context`、`16_sync`、`21_time` 和
`24_concurrency_patterns` 会把这些 raw-channel 时序迁移到更高层 API；本章
不要提前用它们替代 channel 协议。
