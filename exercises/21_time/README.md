# Time：计时器和周期任务

按 1–8 顺序完成。这里专门负责 `time.Timer`/`time.Ticker` 的创建、停止、
drain、Reset 和 deadline；channel token 流与综合 worker 协议仍在后续章节。

| 题目 | 重点 |
| ---: | --- |
| 1–2 | Go reference layout、ticker 读取两次并 Stop |
| 3 | Duration 单位、字符串和边界算术 |
| 4–5 | timer 正常/提前取消、result/closed/timeout 三路 select |
| 6 | Stop、drain、Reset 的可复用 timer 顺序 |
| 7 | ticker 驱动循环的每条退出路径和 done |
| 8 | context、ticker、deadline 和最终资源释放的组合 |

测试使用注入的 gate 或 watchdog，不用固定 Sleep 证明 goroutine 已结束。
实现中创建的每一个 timer/ticker 都要在所有返回路径停止。

```sh
sh check.sh exercises/21_time --run-all
sh scripts/verify_exercise_starters.sh exercises/21_time
```
