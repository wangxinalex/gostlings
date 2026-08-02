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

## Topic order, levels, and Go Tour mapping

The first 23 topics form the core and applied tracks. The final two topics are
intermediate scenarios: each combines multiple standard-library contracts and
uses focused tests instead of relying only on stdout.

| Level | What it verifies |
|---|---|
| Core | Compilation, values, functions, and one concept at a time |
| Applied | Coordination, cancellation, synchronization, and standard-library APIs |
| Intermediate | HTTP behavior, channel ownership, cancellation, error paths, and race-safe state |

Recommended progression: Core (`00_intro`–`11_generics`) → Applied
(`12_goroutines`–`22_strconv`) → Intermediate (`23_http`–
`24_concurrency_patterns`).

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
| 23_http | 3 | [net/http](https://pkg.go.dev/net/http), [httptest](https://pkg.go.dev/net/http/httptest) |
| 24_concurrency_patterns | 3 | [context](https://pkg.go.dev/context), [sync/atomic](https://pkg.go.dev/sync/atomic) |

## Checking your work

```sh
sh check.sh                      # run exercises/ in order, stop at the first failure and show its output
sh check.sh --run-all            # run every exercise and report all PASS/FAIL
sh check.sh solutions --run-all  # verify all 79 reference solutions pass
sh check.sh solutions --run-all --race # verify reference solutions with the race detector
```

Exercises that ship in a runnable-but-wrong state (so you can experiment before
fixing them) come with a `main_test.go` that asserts stdout against the header's
Expected output. `check.sh` uses `go test` for those, `go run` for the rest.

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
`VERSION` is the canonical project version; successful releases update
`CHANGELOG.md`, create a GitHub Release, and create the matching `vX.Y.Z` tag.

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
