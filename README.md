# gostlings

Go exercises inspired by [rustlings](https://github.com/rust-lang/rustlings): one small
program per concept. Fix it, run it, and you have learned the topic.

## Prerequisites

Go 1.23+ (check with `go version`)

## How to solve

1. Work through the topics in the table below, in order.
2. Open `exercises/<topic>/<exercise>/main.go` and read the header comment
   (Concept / Task / Expected output / Hint).
3. Edit the line marked `// TODO:`.
4. Verify it:

   ```sh
   go run ./exercises/01_variables/variables1
   ```

   (For the 14_testing topic, use `go test ./exercises/14_testing/testing1`.)
5. Output matches the expected output → done. Stuck? Peek at the matching path
   under `solutions/`.

## Topic order and Go Tour mapping

| Topic | # | Go Tour section |
|---|---|---|
| 00_intro | 2 | Basics 1 |
| 01_variables | 4 | Basics 8-12 |
| 02_functions | 4 | Basics 4-7 |
| 03_control_flow | 4 | Flowcontrol 1-13 |
| 04_pointers | 3 | Moretypes 1; Methods 5 |
| 05_slices | 5 | Moretypes 7-15; [sort](https://pkg.go.dev/sort) |
| 06_maps | 3 | Moretypes 19-22 |
| 07_structs | 3 | Moretypes 2-5 |
| 08_methods | 3 | Methods 1-6 |
| 09_interfaces | 4 | Methods 9-17 |
| 10_errors | 4 | Methods 19-20 |
| 11_generics | 3 | Generics 1-2 |
| 12_goroutines | 3 | Concurrency 1 |
| 13_channels | 6 | Concurrency 2-6 |
| 14_testing | 3 | [Testing tutorial](https://go.dev/doc/tutorial/add-a-test) |
| 15_context | 3 | [context](https://pkg.go.dev/context) |
| 16_sync | 3 | [sync](https://pkg.go.dev/sync) |
| 17_panic_recover | 2 | [builtin](https://pkg.go.dev/builtin) |
| 18_embedding | 2 | [Effective Go: Embedding](https://go.dev/doc/effective_go#embedding) |
| 19_json | 3 | [encoding/json](https://pkg.go.dev/encoding/json) |
| 20_io | 2 | [io](https://pkg.go.dev/io) |
| 21_time | 2 | [time](https://pkg.go.dev/time) |
| 22_strconv | 2 | [strconv](https://pkg.go.dev/strconv) |

## Checking your work

```sh
sh check.sh                      # run exercises/ in order, stop at the first failure and show its output
sh check.sh --run-all            # run every exercise and report all PASS/FAIL
sh check.sh solutions --run-all  # verify all 73 reference solutions pass
```

Exercises that ship in a runnable-but-wrong state (so you can experiment before
fixing them) come with a `main_test.go` that asserts stdout against the header's
Expected output. `check.sh` uses `go test` for those, `go run` for the rest.