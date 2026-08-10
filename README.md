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
`05_slices`.

| Level | What it verifies |
|---|---|
| Core | Compilation, values, functions, and one concept at a time |
| Applied | Coordination, cancellation, synchronization, and standard-library APIs |
| Intermediate | HTTP behavior, channel ownership, cancellation, error paths, and race-safe state |

Recommended progression: Core (`00_intro`–`11_generics`) → Applied
(`12_goroutines`–`22_strconv`, plus `25_closures`–`26_files`) → Intermediate
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
| 12_goroutines | 3 | Concurrency 1 |
| 13_channels | 24 | [Channel patterns guide](exercises/13_channels/README.md) |
| 14_testing | 11 | [Testing progression](exercises/14_testing/README.md) |
| 15_context | 3 | [context](https://pkg.go.dev/context) |
| 16_sync | 3 | [sync](https://pkg.go.dev/sync) |
| 17_panic_recover | 2 | [builtin](https://pkg.go.dev/builtin) |
| 18_embedding | 2 | [Effective Go: Embedding](https://go.dev/doc/effective_go#embedding) |
| 19_json | 3 | [encoding/json](https://pkg.go.dev/encoding/json) |
| 20_io | 2 | [io](https://pkg.go.dev/io) |
| 21_time | 2 | [time](https://pkg.go.dev/time) |
| 22_strconv | 2 | [strconv](https://pkg.go.dev/strconv) |
| 23_http | 3 | [net/http](https://pkg.go.dev/net/http), [httptest](https://pkg.go.dev/net/http/httptest) |
| 24_concurrency_patterns | 3 | [context](https://pkg.go.dev/context), [sync/atomic](https://pkg.go.dev/sync/atomic) |
| 25_closures | 2 | builds on [Go Tour: Basics 4-7](https://go.dev/tour/basics/4); closures are not in the Tour |
| 26_files | 3 | [os](https://pkg.go.dev/os), [bufio](https://pkg.go.dev/bufio) |

## Checking your work

```sh
sh check.sh                      # run exercises/ in order, stop at the first failure and show its output
sh check.sh --run-all            # run every exercise and report all PASS/FAIL
sh check.sh solutions --run-all  # verify all 120 reference solutions pass
sh check.sh solutions --run-all --race # verify reference solutions with the race detector
```

Every exercise ships with a `main_test.go` that asserts the program's stdout
against the header's Expected output (order-insensitively where the header says
"any order"). Exercises that do not compile or that deadlock when unmodified
simply fail their test until fixed, so a solution that merely exits 0 cannot
pass. `check.sh` uses `go test` whenever a `_test.go` is present and falls back
to `go run` for any directory without tests.

For intermediate exercises, stdout is only part of the contract. Run the focused
package tests to check response status and headers, error paths, channel closure,
and cancellation. Before considering a solution complete, also run:

~~~sh
find exercises solutions -name '*.go' -print0 | xargs -0 gofmt -d
GOCACHE=/tmp/gostlings-go-cache go test ./solutions/...
GOCACHE=/tmp/gostlings-go-cache sh check.sh solutions --run-all --race
~~~

The exercise tree is intentionally incomplete, so package-wide tests target
`solutions/...`; the learner validates each exercise after fixing it.

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
