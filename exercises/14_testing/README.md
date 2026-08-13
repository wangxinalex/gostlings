# Testing: progression path

This chapter starts with the smallest readable test and progressively
practices the fixtures, test doubles, HTTP testing, error classification,
fuzzing, and benchmarks that appear in real projects. Every exercise should
fail first, then pass after you fix the test or the test helper code.

| Exercise | Focus | Goal |
| --- | --- | --- |
| testing1 | Unit test | Meet `go test` and basic assertions |
| testing2 | Table-driven test | Cover boundaries with multiple inputs |
| testing3 | Subtests | Organize output with `t.Run` |
| testing4 | Test helpers | Keep failure locations readable with `t.Helper` |
| testing5 | File fixtures | Isolate temp files with `t.TempDir` |
| testing6 | Environment fixtures | Auto-restore environment variables with `t.Setenv` |
| testing7 | Fake + dependency injection | Test business logic without a database |
| testing8 | `httptest` | Test handlers without listening on a real port |
| testing9 | Fuzz | Discover boundary inputs with invariants |
| testing10 | Error-classification assertions | Test stable semantics with `errors.Is` |
| testing11 | Benchmark | Measure hot paths with `b.ResetTimer` and `b.ReportAllocs` |

Common commands:

```sh
go test ./exercises/14_testing/testing4
go test -v ./exercises/14_testing/testing8
go test -run=^$ -fuzz=FuzzReverse -fuzztime=2s ./exercises/14_testing/testing9
go test -bench=. -benchmem ./exercises/14_testing/testing11
```

Reference: [Add a test](https://go.dev/doc/tutorial/add-a-test),
[Fuzzing](https://go.dev/doc/tutorial/fuzz),
[`testing` package](https://pkg.go.dev/testing).
