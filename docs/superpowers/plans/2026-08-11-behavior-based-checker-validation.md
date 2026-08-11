# Behavior-Based Checker Validation Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Make exercise checking independent of comment text while preserving a behavior-based guard against untouched starters passing.

**Architecture:** `check.sh` will only execute the selected exercise and report its compiler/test/runtime result. `scripts/verify_exercise_starters.sh` will run untouched starters and require those executions to fail; neither script will inspect comments. `intro1` will be corrected so its starter genuinely fails to compile until the learner fixes it.

**Tech Stack:** POSIX shell, Go tests, existing `check.sh` runner, Go 1.26.5.

## Global Constraints

- A `TODO:` string in any Go comment must never determine exercise pass/fail status.
- Untouched starters must fail their focused behavioral check.
- Solutions must continue to pass the complete solutions suite.
- Do not add dependencies.

---

### Task 1: Add a regression test for comment-independent checking

**Files:**
- Create: `scripts/test_check_behavior.sh`
- Create: `scripts/testdata/check_behavior/exercises/99_checker_behavior/comment_only/main.go`
- Create: `scripts/testdata/check_behavior/exercises/99_checker_behavior/comment_only/main_test.go`

**Interfaces:**
- Consumes: the repository's `check.sh` command.
- Produces: an executable shell regression test that exits non-zero if a passing exercise with a `TODO:` comment is rejected.

- [ ] **Step 1: Write the failing regression test**

Create a temporary mini-repository containing `go.mod` and copy the tracked fixture files into it:

```sh
#!/bin/sh
set -eu

script_dir=$(CDPATH= cd "$(dirname "$0")" && pwd -P)
repo_root=$(CDPATH= cd "$script_dir/.." && pwd -P)
tmp_dir=$(mktemp -d)
trap 'rm -rf "$tmp_dir"' EXIT HUP INT TERM

cp "$repo_root/check.sh" "$tmp_dir/check.sh"
cp "$repo_root/go.mod" "$tmp_dir/go.mod"
cp -R "$script_dir/testdata/check_behavior/." "$tmp_dir/"

cd "$tmp_dir"
sh check.sh exercises/99_checker_behavior/comment_only
```

- [ ] **Step 2: Run the regression test and verify the expected failure**

Run:

```sh
sh scripts/test_check_behavior.sh
```

Expected before the checker fix: failure containing `starter still contains TODO`. This proves the test detects the current bug rather than passing immediately.

### Task 2: Remove comment-based gates and repair the invalid starter

**Files:**
- Modify: `check.sh`
- Modify: `scripts/verify_exercise_starters.sh`
- Modify: `exercises/00_intro/intro1/main.go`
- Modify: `exercises/00_intro/intro2/main.go`
- Modify: `exercises/13_channels/channels1/main.go`
- Modify: `exercises/13_channels/channels2/main.go`
- Modify: `exercises/13_channels/channels3/main.go`
- Modify: `exercises/13_channels/channels4/main.go`

**Interfaces:**
- Consumes: the existing exercise discovery, selective race mapping, and focused test execution.
- Produces: behavior-only checking and behavior-only starter verification.

- [ ] **Step 1: Remove the TODO branch and obsolete starter mode from `check.sh`**

Delete the branch that greps `main.go` for `TODO:` and emits `starter still contains TODO`. Remove the `--verify-starter` option, its parsing, and its path validation because it has no behavior beyond bypassing the comment gate; retain exercise-path discovery needed by selective race checks.

- [ ] **Step 2: Remove the TODO requirement from starter verification**

Delete the `grep -q 'TODO:'` requirement in `scripts/verify_exercise_starters.sh`, and change its invocation from `sh check.sh "$exercise" --verify-starter` to `sh check.sh "$exercise"`. The script must fail only when the untouched starter command succeeds.

- [ ] **Step 3: Make `intro1` genuinely incomplete**

Change only each already-complete starter seam so the starter is genuinely incomplete:

```go
fmt.PrintLn("Hello, gostlings!")
```

For `intro2`, remove the completed `main` function. Apply these exact channel starter changes:

```text
channels1: replace `go func() { ch <- "hi" }()` with `ch <- "hi"`.
channels2: replace `make(chan int, 2)` with `make(chan int)`.
channels3: remove the producer's `close(ch)` call.
channels4: replace `value, ok := <-ch; return value, ok` with `value := <-ch; return value, true`.
```

Keep explanatory comments and solution implementations unchanged. The starter verifier must fail because of compilation or focused behavior, never because of comment text.

### Task 3: Run the green test cycle

**Files:**
- Modify: none

- [ ] **Step 1: Re-run the comment regression test**

Run:

```sh
sh scripts/test_check_behavior.sh
```

Expected: exit 0; the temporary exercise passes despite its `TODO:` comment.

- [ ] **Step 2: Verify the repaired starter fails behaviorally**

Run:

```sh
sh scripts/verify_exercise_starters.sh exercises/00_intro
```

Expected: exit 0 for the verifier, with `intro1` rejected internally because compilation fails, not because of comment text.

- [ ] **Step 3: Check the corresponding solution**

Run:

```sh
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH \
GOCACHE=/private/tmp/gostlings-go-cache \
sh check.sh solutions/00_intro/intro1 --run-all
```

Expected: `1/1 passed`.

### Task 4: Verify the repository and publish the fix

**Files:**
- Modify: only the files from Tasks 1–2.

- [ ] **Step 1: Run the complete solutions suite**

```sh
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH \
GOCACHE=/private/tmp/gostlings-go-cache \
sh check.sh solutions --run-all
```

Expected: all solutions pass.

- [ ] **Step 2: Run formatting and diff checks**

```sh
git diff --check
PATH=/Users/wangxinalex/.goenv/versions/1.26.5/bin:$PATH \
gofmt -d $(git ls-files '*.go')
```

Expected: both commands exit 0 and `gofmt -d` prints no diff.

- [ ] **Step 3: Commit the focused changes**

```sh
git add check.sh scripts/verify_exercise_starters.sh scripts/test_check_behavior.sh scripts/testdata/check_behavior exercises/00_intro/intro1/main.go exercises/00_intro/intro2/main.go exercises/13_channels/channels1/main.go exercises/13_channels/channels2/main.go exercises/13_channels/channels3/main.go exercises/13_channels/channels4/main.go
git commit -m "fix: make exercise checks behavior based"
```

- [ ] **Step 4: Push and open/update the draft PR**

Push the current branch with tracking, then create or update a draft PR targeting `main`. The PR description must state that TODO comment detection was removed, starter validation is behavior-based, and the complete solutions check was run.
