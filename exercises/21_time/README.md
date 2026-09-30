# Time: timers and periodic tasks

Work through 1–10 in order. This chapter is dedicated to creating, stopping,
draining, resetting, and setting deadlines for `time.Timer`/`time.Ticker`;
channel token streams and integrated worker protocols remain in later
chapters.

| Exercise | Focus |
| ---: | --- |
| 1–2 | Go reference layout, read a ticker twice and Stop |
| 3 | Duration units, strings, and boundary arithmetic |
| 4–5 | Timer normal/early cancellation, result/closed/timeout three-way select |
| 6 | Reusable timer ordering: Stop, drain, Reset |
| 7 | Every exit path and done in a ticker-driven loop |
| 8 | Combine context, ticker, deadline, and final resource release |
| 9 | Measure elapsed work against a budget through an injected clock |
| 10 | Retry with a timer between attempts and stop at the first success |

Tests use injected gates or watchdogs, never a fixed `Sleep`, to prove that a
goroutine has finished. Every timer/ticker created in an implementation must
be stopped on all return paths.

```sh
sh check.sh exercises/21_time --run-all
sh scripts/verify_exercise_starters.sh exercises/21_time
```
