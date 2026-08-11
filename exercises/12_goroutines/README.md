# Goroutines: launch, finish, and join

This chapter is a deliberately small foundation. Each exercise launches work in
goroutines and makes completion observable before the function returns. The
completion primitive for all ten exercises is `sync.WaitGroup`; channels begin
in the next chapter, where they carry values, completion, and cancellation.

## Progression

Work through the directories in order:

| Exercise | Focus |
| --- | --- |
| 1 | Launch one goroutine and wait for it before `main` returns |
| 2 | Keep a loop value correct inside a goroutine closure |
| 3 | Pass values explicitly as goroutine arguments |
| 4 | Build the basic `WaitGroup` lifecycle for a worker count |
| 5 | Wait for a dynamically sized batch of jobs |
| 6 | Mark each started worker complete exactly once |
| 7 | Join one batch before launching the next batch |
| 8 | Pass immutable job input into each worker |
| 9 | Use deferred `Done` safely across an early return |
| 10 | Review the complete parameterized-worker lifecycle, including empty input |

The recurring rule is: call `Add` before `go`, call `Done` in the worker (usually
with `defer`), and call `Wait` before returning. Do not use `Sleep` as a join.
These exercises intentionally avoid channels, mutexes, timers, contexts, and
integrated concurrency patterns. Learn those protocols in `13_channels`, then
the later `15_context`, `16_sync`, `21_time`, and `24_concurrency_patterns`
chapters.

## Checking

Run one exercise while you work:

```sh
sh check.sh exercises/12_goroutines/goroutines1
```

Run the full chapter, or verify every untouched starter against its focused
tests:

```sh
sh check.sh exercises/12_goroutines --run-all
sh scripts/verify_exercise_starters.sh exercises/12_goroutines
```

The normal checker rejects an exercise target whose `main.go` still contains a
`TODO:`. The starter verifier is different: it temporarily checks that every
starter fails its behavioral test, so a vacuous implementation cannot pass.
For race coverage, `--race` adds `-race` only to exact paths listed in
`race.list`; `--race-all` applies `-race` to every selected target and is the
expensive full audit:

```sh
sh check.sh exercises/12_goroutines/goroutines1 --race
sh check.sh solutions --run-all --race
sh check.sh solutions --run-all --race-all
```
