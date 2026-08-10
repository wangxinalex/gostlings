# Channels：按真实开发频率递进练习

这一章不是把 channel 当作“带箭头的队列”来背 API，而是逐步建立一个生命周期模型：

> 谁发送？谁接收？谁知道不会再发送？谁负责关闭？调用者提前放弃时，谁负责取消？

前 5 题建立阻塞、缓冲、关闭和所有权；6–12 题练习真实服务代码里最常见的 `select`、超时、完成通知和取消；13–18 题拆解并组合 fan-out、fan-in 和 worker pool；19–21 题把同一套关闭/取消规则迁移到 pipeline；22–24 题覆盖并发上限、速率保护和一个低频高级技巧。

## 先记住四条规则

1. **发送方通常负责关闭。** 只有知道“不会再发送”的一方才应 `close(ch)`；接收方通常只负责 `range ch` 或 comma-ok 接收。
2. **关闭不是发送特殊值。** `close(ch)` 表示生命周期结束；缓冲区里的值仍会先被读完，之后接收才得到 `ok == false`。
3. **每个可能阻塞的操作都要有退出路径。** 生产者的 `out <- value`、消费者的 `<-input`、pipeline stage 的两端，都可能需要 `select` 监听取消信号。
4. **只关闭一次，并等待退出。** 多个 goroutine 可以接收同一个关闭信号，但输出 channel 应由一个明确的协调者在所有发送者退出后关闭。

常用的职责划分是：数据 channel 传值，`done` channel 表示完成，`stop`/`cancel` channel 广播取消，`WaitGroup` 等待一组 goroutine 退出。

## 24 题学习地图

按目录编号顺序完成，不需要跳题。频率标记表示真实开发中的常见程度，不代表题目难度。

| 题目 | 模式 | 频率 | 难度 | 前题提供的铺垫 |
| --- | --- | --- | --- | --- |
| 1 | 无缓冲交接 | 基础 | ★ | 认识发送和接收的同步阻塞 |
| 2 | 有缓冲与背压 | 高频 | ★ | 在 1 的阻塞模型上加入容量 |
| 3 | 关闭 + `range` | 高频 | ★ | 让消费者知道数据何时结束 |
| 4 | comma-ok 接收 | 高频 | ★★ | 显式区分关闭和真实零值 |
| 5 | 单向 channel + generator | 高频 | ★★ | 把关闭责任封装进生产者 API |
| 6 | 基础 `select` | 高频 | ★★ | 用多个可能的 receive 组织等待 |
| 7 | `select` 超时 | 高频 | ★★ | 给一次等待加截止时间 |
| 8 | `select` + `default` | 常见 | ★★ | 把阻塞尝试改成立即返回 |
| 9 | `done` 完成通知 | 高频 | ★★ | 区分“完成事件”和“结果数据” |
| 10 | 可取消的阻塞发送 | 高频 | ★★★ | 让 producer 响应调用者取消 |
| 11 | 超时 + 取消 + join | 高频 | ★★★ | 组合 7、9、10，返回前完成清理 |
| 12 | 广播取消多个 worker | 高频 | ★★★ | 把一个取消信号传播给任务组 |
| 13 | fan-out | 高频 | ★★★ | 多个 worker 共同消费一个 jobs channel |
| 14 | fan-in | 高频 | ★★★ | 等待多个输入后统一关闭输出 |
| 15 | 可取消 fan-in | 高频 | ★★★★ | 防止下游提前退出造成 forwarder 泄漏 |
| 16 | 基础 worker pool | 高频 | ★★★★ | 组合分发、处理、结果收集和生命周期 |
| 17 | worker pool 保序 | 常见 | ★★★★ | 处理顺序与完成顺序不同，携带 index 恢复 |
| 18 | 首错取消 worker pool | 高频 | ★★★★ | 错误路径也要取消并等待所有 worker |
| 19 | 单阶段 pipeline | 高频 | ★★★ | 把“输入关闭→输出关闭”封装为 stage |
| 20 | 多阶段 pipeline | 高频 | ★★★★ | 让多个 stage 串联且各自管理输出 |
| 21 | 可取消 pipeline | 高频 | ★★★★★ | 取消必须覆盖每个 stage 的接收和发送 |
| 22 | buffered channel semaphore | 常见 | ★★★★ | 用容量限制活跃任务数 |
| 23 | ticker/channel rate limiter | 常见 | ★★★★ | 用 tick 控制操作启动速率 |
| 24 | nil channel 动态开关 | 低频拓展 | ★★★★★ | 在 `select` 中动态禁用已结束输入 |

### 为什么是这个顺序？

- **1–5：语义和所有权。** 如果还不清楚无缓冲发送为什么阻塞、谁关闭 channel，后面的 fan-in 和 worker pool 会变成记模板。
- **6–12：等待和退出。** 先学 `select` 的等待方式，再学超时；然后将完成、取消、广播取消组合起来。第 11 题是第一个完整的“调用者放弃后仍然干净返回”的模式。
- **13–18：常用组合模式。** 先单独练 fan-out 和 fan-in，再组合成 worker pool；第 17 题补上业务经常需要的顺序约束，第 18 题补上真实服务经常出现的错误取消。
- **19–21：迁移到 pipeline。** pipeline 并没有新生命周期规则，只是把前面“每个 stage 关闭自己的输出”和“发送可取消”重复应用到更多 goroutine。
- **22–24：边界和拓展。** 并发上限和速率限制有实际用途；nil channel 很有用但不应成为初学者的主要写法，因此只放一题收尾。

## 关键时序

### 正常生产者：先发送，最后关闭

```text
producer                         consumer
   |                                |
   |------ value -----------------> |
   |------ value -----------------> |
   |------ close(out) ------------> |
   |                                | range 结束
```

`range` 不会因为“暂时没有值”结束，只会因为 channel 已关闭且值已取完结束。producer 必须拥有关闭权；调用者不应该替 producer 关闭它的输出。

### 第 11 题：超时不等于取消

```text
caller                         slow producer
  |                                |
  |--- wait result or timeout -----|
  |--- timeout ------------------->|
  |--- close(stop) --------------->|
  |                                | select sees stop
  |<---------- close(done) --------|
  | return "timed out"             |
```

`time.After` 只让 caller 停止等待；它不会杀掉 producer。正确顺序是：关闭取消信号，等待 producer 退出，再向 caller 返回。第 11 题通过 `done` 参数让测试可以确认 `run` 返回时 producer 已经完成。

### 第 14/15 题：谁能关闭 fan-in 的输出？

多个 forwarder 都会向 `out` 发送，所以任何一个 forwarder 都不能单独关闭它。正确时序是：

```text
input A -> forwarder A --+
input B -> forwarder B ---+--> out -> consumer
input C -> forwarder C --+       ^
                                  |
                         WaitGroup coordinator
                         waits all, then close(out)
```

第 15 题再加入 `stop`：forwarder 既可能卡在读取 input，也可能卡在向 out 发送，因此两个位置都必须能接收取消信号。

### 第 16/18 题：worker pool 的关闭顺序

```text
jobs producer --close(jobs)--> workers --send--> results --close(results)--> caller
                                      |
                                WaitGroup
                                      |
                           coordinator waits workers
```

只能由 jobs producer 关闭 jobs；只能由协调者在所有 worker 退出后关闭 results。第 18 题的错误路径还要关闭 stop、停止继续投递，并保证 stop 只关闭一次。

## 高频模式的最小模板

### 可取消发送

```go
select {
case out <- value:
	// 发送成功
case <-stop:
	return
}
```

凡是下游可能提前退出的 producer、fan-in forwarder 和 pipeline stage，都要考虑这个模板。

### 等待一组 worker 后关闭输出

```go
var wg sync.WaitGroup
for i := 0; i < workerCount; i++ {
	wg.Go(func() {
		for job := range jobs {
			results <- process(job)
		}
	})
}

go func() {
	wg.Wait()
	close(results)
}()
```

Go 1.26 可以用 `WaitGroup.Go` 把 `Add(1)`、启动 goroutine 和 `Done()` 封装起来；本质仍然是“所有发送者结束后，单独的协调者关闭输出”。如果使用旧版本 Go，则用 `wg.Add(1)`、`defer wg.Done()` 写出同一生命周期。

### nil channel 作为 `select` 开关

```go
for first != nil || second != nil {
	select {
	case value, ok := <-first:
		if !ok {
			first = nil // 永久禁用 first case
			continue
		}
		use(value)
	case value, ok := <-second:
		if !ok {
			second = nil
			continue
		}
		use(value)
	}
}
```

nil channel 的发送和接收永远不会就绪，所以它在 `select` 中等价于暂时移除该 case。它是第 24 题的高级技巧，不应替代普通的清晰控制流。

## 常见错误和排查方式

- `range ch` 不结束：检查发送方是否关闭了 channel，以及是否真的已经发送完所有值。
- 同一个 channel 被多个 goroutine 关闭：把关闭责任交给一个 producer 或 coordinator。
- worker pool 卡死：检查是否在启动结果消费者之前就同步投递所有 jobs；无缓冲 results 可能让 worker 卡住，进而让 producer 也卡住。
- 超时后仍有 goroutine：超时分支只返回是不够的；给阻塞发送/接收加入 stop，并在返回前等待 done。
- 用 `time.Sleep` 等待 goroutine：Sleep 只能改变概率，不能证明完成；使用 close(done)、range 结束或 WaitGroup。
- `select` + `default` 消耗 CPU：没有实际工作或退避的无限 default 循环会 busy loop。
- 关闭 channel 后出现很多零值：comma-ok 的 `ok` 为 false 后应停止读取；如果是多路 `select`，把关闭的输入设为 nil。

### 如何 detect goroutine 泄漏？

练习和单元测试优先使用可观察的生命周期信号：让函数返回一个 done channel，或等待输出 channel 关闭，然后用短 timeout 作为死锁保护。不要只用 `runtime.NumGoroutine` 做断言，它会受到测试框架和运行时 goroutine 影响。

实际服务中可以组合使用：

1. `go test -race -timeout 5s ./...`：race detector 能发现共享状态问题，timeout 能暴露永久阻塞。
2. 用 `runtime/pprof` 或 HTTP `/debug/pprof/goroutine` 查看 goroutine profile，重点找大量卡在 `chan send`、`chan receive`、`select` 或 `sync.WaitGroup.Wait` 的栈。
3. 为每个长期运行的 goroutine 定义明确的退出条件，并在测试中关闭取消信号后等待退出，而不是只检查主函数返回。

生产请求链一般优先使用 `context.Context` 传递截止时间和取消信号；本章用 channel 直接展示底层时序，相关 API 见 [`15_context`](../15_context)。

## 做完后的自测清单

- 我能指出每个 channel 的发送者、接收者和关闭者。
- 我能解释为什么 `close` 后仍可能读到缓冲区里的值。
- 我能区分 result channel、done channel 和 stop channel。
- 所有 `range ch` 都有明确的关闭路径。
- 所有可能阻塞的发送，在调用者提前退出时都有取消分支。
- fan-in/worker pool 的输出只被一个协调者关闭。
- 我没有把结果到达顺序误当成输入顺序。
- 我能用 `-race` 和 goroutine profile 辅助定位泄漏，而不是依赖 Sleep 猜时序。
