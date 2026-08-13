# gostlings

Go exercises inspired by [rustlings](https://github.com/rust-lang/rustlings): one small
program per concept. Fix it, run it, and you have learned the topic.

## Prerequisites

Go 1.26.5+ (check with `go version`)

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
   (For the 26_files topic, the exercises read files next to `main.go`, so run
   them from their own directory: `cd exercises/26_files/files1 && go run .`.)
5. Output matches the expected output → done. Stuck? Peek at the matching path
   under `solutions/`.

## Topic order, levels, and Go Tour mapping

Topics 00–22 form the core and applied tracks. Topics 23–24 are intermediate
scenarios: each combines multiple standard-library contracts and uses focused
tests instead of relying only on stdout. Topics 25–26 are applied additions
(closures and file I/O) that can be done any time after `02_functions` and
`05_slices`. Topics 27–33 are standard-library additions (strings, formatting,
flags, sorting, regexp, subprocesses, and templating) that can be done any
time after the Core topics; each chapter README notes its prerequisites.

| Level | What it verifies |
|---|---|
| Core | Compilation, values, functions, and one concept at a time |
| Applied | Coordination, cancellation, synchronization, and standard-library APIs |
| Intermediate | HTTP behavior, channel ownership, cancellation, error paths, and race-safe state |

Recommended progression: Core (`00_intro`–`11_generics`) → Applied
(`12_goroutines`–`22_strconv`, plus the stdlib additions `25_closures`–`26_files`
and `27_strings`–`33_template`) → Intermediate
(`23_http`–`24_concurrency_patterns`).

| Topic | # | Go Tour section |
|---|---|---|
| 00_intro | 2 | Basics 1 |
| 01_variables | 4 | Basics 8-12 |
| 02_functions | 6 | Basics 4-7 |
| 03_control_flow | 5 | Flowcontrol 1-13 |
| 04_pointers | 3 | Moretypes 1; Methods 5 |
| 05_slices | 5 | Moretypes 7-15; [sort](https://pkg.go.dev/sort) |
| 06_maps | 3 | Moretypes 19-22 |
| 07_structs | 3 | Moretypes 2-5 |
| 08_methods | 3 | Methods 1-6 |
| 09_interfaces | 5 | Methods 9-17 |
| 10_errors | 10 | [Errors progression](exercises/10_errors/README.md) |
| 11_generics | 3 | Generics 1-2 |
| 12_goroutines | 10 | [Goroutines progression](exercises/12_goroutines/README.md) |
| 13_channels | 50 | [Channel patterns guide](exercises/13_channels/README.md) |
| 14_testing | 11 | [Testing progression](exercises/14_testing/README.md) |
| 15_context | 14 | [Context progression](exercises/15_context/README.md) |
| 16_sync | 14 | [sync progression](exercises/16_sync/README.md) |
| 17_panic_recover | 2 | [builtin](https://pkg.go.dev/builtin) |
| 18_embedding | 2 | [Effective Go: Embedding](https://go.dev/doc/effective_go#embedding) |
| 19_json | 3 | [encoding/json](https://pkg.go.dev/encoding/json) |
| 20_io | 2 | [io](https://pkg.go.dev/io) |
| 21_time | 8 | [time progression](exercises/21_time/README.md) |
| 22_strconv | 2 | [strconv](https://pkg.go.dev/strconv) |
| 23_http | 3 | [net/http](https://pkg.go.dev/net/http), [httptest](https://pkg.go.dev/net/http/httptest) |
| 24_concurrency_patterns | 18 | [Concurrency patterns progression](exercises/24_concurrency_patterns/README.md) |
| 25_closures | 2 | builds on [Go Tour: Basics 4-7](https://go.dev/tour/basics/4); closures are not in the Tour |
| 26_files | 3 | [os](https://pkg.go.dev/os), [bufio](https://pkg.go.dev/bufio) |
| 27_strings | 4 | [strings progression](exercises/27_strings/README.md) |
| 28_fmt | 3 | [fmt progression](exercises/28_fmt/README.md) |
| 29_flag | 2 | [flag progression](exercises/29_flag/README.md) |
| 30_sort | 3 | [sort progression](exercises/30_sort/README.md) |
| 31_regexp | 3 | [regexp progression](exercises/31_regexp/README.md) |
| 32_os_exec | 2 | [os/exec progression](exercises/32_os_exec/README.md) |
| 33_template | 2 | [text/template progression](exercises/33_template/README.md) |

## Checking your work

```sh
sh check.sh                      # run exercises/ in order, stop at the first failure and show its output
sh check.sh --run-all            # run every exercise and report all PASS/FAIL
sh check.sh exercises/13_channels/channels6 # check one exercise directly
sh scripts/verify_solutions.sh   # overlay exercise tests and verify every reference solution passes
sh check.sh solutions --run-all  # compile/run every reference solution as a stdout smoke
sh check.sh solutions --run-all --race # verify selectively listed reference solutions with the race detector
sh check.sh solutions --run-all --race-all # audit every selected reference solution with the race detector
```

Every exercise ships with a `main_test.go` that asserts the program's stdout
against the header's Expected output (order-insensitively where the header says
"any order"). Exercises that do not compile or that deadlock when unmodified
simply fail their test until fixed, so a solution that merely exits 0 cannot
pass. `check.sh` uses `go test` whenever a `_test.go` is present, runs packages
with `Benchmark` functions using `-bench=.`, and falls back to `go run` for any
directory without tests. It also accepts an individual exercise directory as
the target.

`scripts/verify_solutions.sh` temporarily copies each exercise's `*_test.go`
into the matching solution directory, runs `go test ./solutions/...`, then
removes the copies. The exercise tests therefore stay the single source of
truth, and the reference answers are checked against the same behavioral
contract without committing duplicate test files under `solutions/`.

For intermediate exercises, stdout is only part of the contract. Run the focused
package tests to check response status and headers, error paths, channel closure,
and cancellation. Before considering a solution complete, also run:

~~~sh
find exercises solutions -name '*.go' -print0 | xargs -0 gofmt -d
GOCACHE=/tmp/gostlings-go-cache sh scripts/verify_solutions.sh
GOCACHE=/tmp/gostlings-go-cache sh check.sh solutions --run-all --race
GOCACHE=/tmp/gostlings-go-cache sh check.sh solutions --run-all --race-all
~~~

The exercise tree is intentionally incomplete; `verify_solutions.sh` checks the
reference answers against the exercise tests, and the learner validates each
exercise after fixing it.

## Release workflow

semantic-release runs after the existing checks succeed on a push to `main`.
It determines the next version from existing `vX.Y.Z` tags, then creates the
matching tag and GitHub Release. Releases do not write a version commit,
package manifest, or changelog back to `main`.

Use Conventional Commits to describe release impact:

- `fix:` or `perf:` → patch release
- `feat:` → minor release
- `!` after the type/scope or a `BREAKING CHANGE:` footer → major release
- `docs:`, `test:`, `ci:`, `chore:`, and `refactor:` do not release by themselves

Before merging the release configuration, mark the current `main` commit as
the `0.1.0` baseline and push the tag. This prevents semantic-release from
interpreting the first untagged release as `1.0.0`:

```sh
git tag -a v0.1.0 "$(git rev-parse main)" -m "chore: mark 0.1.0 baseline"
git push origin v0.1.0
```
