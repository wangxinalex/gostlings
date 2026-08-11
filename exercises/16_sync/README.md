# Sync：共享内存的生命周期

按 1–14 顺序练习。这里的主角是共享状态与标准库同步原语；消息流、worker
pipeline 和 request/reply 组合留给 channels 与 `24_concurrency_patterns`。

| 题目 | 重点 |
| ---: | --- |
| 1–2 | Mutex 保护计数器与 map 的 read-modify-write |
| 3–5 | `sync.Once`、WaitGroup 的 Add/Done/Wait 顺序与动态任务 |
| 6–7 | RWMutex cache、锁外慢计算与受保护提交 |
| 8 | 并发配置初始化只发生一次 |
| 9–10 | Cond predicate loop、Signal 与 Close/Broadcast |
| 11–12 | `sync.Map` 的 canonical LoadOrStore、`sync.Pool` 的借还边界 |
| 13 | 有界共享状态、等待容量和 shutdown 唤醒 |
| 14 | Once 初始化、活动计数、Finish 和服务 Wait |

所有 sync 题都在 `race.list` 中做选择性 race 检查；语义测试仍然先于 race
检查，因为“这次调度没有撞到 race”不等于实现正确。Cond 必须在锁内检查
predicate，并用 `for` 循环等待。

```sh
sh check.sh exercises/16_sync --run-all
sh check.sh solutions/16_sync --run-all --race
```
