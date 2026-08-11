# Task 9 Implementation Report — Checker package discovery

Verified in `/Users/wangxinalex/SelfStudy/Rust/gostlings/.worktrees/concurrency-practice` on 2026-08-11 with:

```sh
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH
GOCACHE=/private/tmp/gostlings-go-cache
```

## Recorded red phase

Before changing `check.sh`, the requested reproduction returned success without
executing a package:

```sh
sh check.sh solutions/12_goroutines --run-all
# ----
# 0/0 passed
# All solutions/12_goroutines pass ✓
# exit 0
```

An assertion requiring the output line `10/10 passed` then exited 1. The root
cause was the directory-only discovery command:

```sh
find "$target" -mindepth 2 -maxdepth 2 -type d -print
```

That depth is appropriate for a root target but excludes direct exercise
directories below a chapter target.

## Implementation

The non-single-target branch now finds `main.go` and `*_test.go` files under
the target, maps each to its containing directory, de-duplicates directories,
and orders them numerically by the topic prefix and exercise numeric suffix.
The existing direct-package branch is unchanged.

The TODO gate, canonical `--verify-starter` target check, exact
`grep -F -x` `race.list` matching, `--race-all`, commands, per-directory
output, and failure handling are unchanged. In particular, `--race` does not
select the goroutines or channels chapters because neither is listed in the
manifest.

## Verification

| Check | Result |
| --- | --- |
| `sh check.sh solutions/12_goroutines --run-all` | PASS — `10/10 passed` |
| `sh check.sh solutions/13_channels --run-all` | PASS — `50/50 passed` |
| `sh check.sh solutions/12_goroutines --run-all --race` | PASS — `10/10 passed`; manifest selected no chapter package |
| `sh check.sh solutions/13_channels --run-all --race` | PASS — `50/50 passed`; manifest selected no chapter package |
| `sh check.sh solutions --run-all` | PASS — `153/153 passed` |
| `sh check.sh solutions/12_goroutines --run-all --race-all` | PASS — `10/10 passed` |
| Direct solution target `solutions/12_goroutines/goroutines1` | PASS — `1/1 passed` |
| Direct starter target `exercises/12_goroutines/goroutines1` | Expected nonzero — rejected by the TODO gate before execution |
| Canonical starter target `./exercises/12_goroutines/goroutines1 --verify-starter` | Expected focused-test failure, no usage error; confirms valid canonical path acceptance |
| Solution target with `--verify-starter` | Expected usage exit 2 |
| `sh -n check.sh` | PASS |
| `git diff --check` | PASS |

No exercise or learner documentation content was modified. The pre-existing,
intentional deletions of the Task 3 and Task 5 implementer reports remain
preserved.

## Concerns

None. The selective race manifest intentionally has no entries for chapters
12 and 13, so those chapter checks remain normal checks when invoked with
`--race`.
