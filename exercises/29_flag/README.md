# Flag: command-line flag parsing

Work through 1–2 in order. Practice parsing flags with `flag.NewFlagSet`
instead of relying on the global `flag.CommandLine`, which keeps the code
testable and reusable. Verify with `go test` or `check.sh`:

```sh
sh check.sh exercises/29_flag/flag1
go test ./exercises/29_flag/flag2
```

| Exercise | Focus | Goal |
| --- | --- | --- |
| flag1 | `flag.NewFlagSet` + `String/Int/Bool` | Parse flags of several types |
| flag2 | `fs.Args()` | Mix flags and positional arguments |

Reference: [flag package](https://pkg.go.dev/flag)
