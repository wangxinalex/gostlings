# Context：从取消信号到请求收尾

按顺序完成 1–14，不要跳题。每题都保留一个明确的 `TODO:` seam；focused
test 会检查真实的阻塞点、context 透传和 cleanup，而不是只检查一个字面结果。

| 题目 | 重点 |
| ---: | --- |
| 1–2 | `WithCancel`/`WithTimeout`、`Done`、capacity-one result、defer cancel |
| 3–5 | typed value、`errors.Is`、绝对 deadline 和 cancel |
| 6–7 | 父子传播、child 隔离、所有 child 退出后再 done |
| 8–9 | 可取消的 blocked receive/send |
| 10–12 | worker join、取消前置检查、helper 链透传同一个 context |
| 13 | `WithCancelCause` 与 `context.Cause` |
| 14 | deadline、worker、result、取消和最终 join 的请求级组合 |

Hints 给出每题的最小协议：阻塞操作必须同时监听 `ctx.Done()`，创建派生
context 的函数必须释放，完成信号必须等所有 worker 返回。原始 channel 的
fan-in/fan-out 属于 `13_channels`，共享内存原语属于 `16_sync`。

```sh
sh check.sh exercises/15_context --run-all
sh scripts/verify_exercise_starters.sh exercises/15_context
```
