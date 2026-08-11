# Concurrency patterns：综合并发题单

完成前面的 goroutine、channel、context、sync 和 time 章节后，按 1–18 顺序
练习。本章每题至少组合两个已学原语，并新增一个生命周期或失败策略；不要
把基础 channel 关闭题跳回本章重复做。

| 题目 | 综合能力 |
| ---: | --- |
| 1–3 | generator/reducer、可取消 stage、atomic counter + join |
| 4–6 | 多 stage cancel、bounded pool、typed result/error pipeline |
| 7–9 | size/time batch、token rate limit、graceful service shutdown |
| 10–12 | fan-in coordinator、ordered results、cancellable retry/backoff |
| 13–15 | context-aware Once、atomic metrics、bounded load shedding |
| 16–18 | buffered cancellation、deadline shutdown、request/reply capstone |

共同检查顺序是：停止继续接收或投递 → 让阻塞操作退出 → join workers → 由
唯一 coordinator 关闭输出。题目中的错误用 `errors.Is` 语义判断；结果要有
顺序/数量/关闭断言，不能用空 slice、忽略错误或提前返回制造通过。

```sh
sh check.sh exercises/24_concurrency_patterns --run-all
sh scripts/verify_exercise_starters.sh exercises/24_concurrency_patterns
```
