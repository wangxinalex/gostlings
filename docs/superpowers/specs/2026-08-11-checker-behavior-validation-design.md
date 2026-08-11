# Behavior-Based Exercise Checker Design

## Goal

Ensure exercise checking is based on compilation, tests, and runtime behavior rather than whether a starter comment contains `TODO:`.

## Problem and Root Cause

`check.sh` currently treats any `TODO:` in an exercise `main.go` as an automatic failure. This makes comment text part of the correctness contract. It also masks a separate problem: `exercises/00_intro/intro1` currently passes its focused test without modification.

## Design

1. Remove every comment-content check from `check.sh`. A normal exercise check must report only the result of the exercise's compile, test, benchmark, or run command; a `TODO:` string in any comment is inert.
2. Remove every comment-content check from `scripts/verify_exercise_starters.sh`. Starter validation must not require or reject any particular comment text.
3. Keep starter validation behavior-based: for each untouched starter, run `check.sh <exercise>` and require a non-zero result. The verifier has no comment-specific execution mode.
4. Make the already-complete `intro1`, `intro2`, and `channels1`–`channels4` starters genuinely incomplete while keeping their solutions unchanged. Their starter checks must fail because of the program or focused behavior, and their solution checks must pass.
5. Add a checker regression test that uses a temporary exercise containing a `TODO:` comment and a passing test, proving that `check.sh` does not inspect or fail on comment text. Keep the existing starter verifier as the regression test for untouched starters.

## Alternatives Considered

- Comparing starter and solution source files was rejected because it would couple correctness to formatting, comments, and one particular implementation.
- Only removing the `TODO:` gate was rejected because an unchanged `intro1` starter would then pass.

## Verification

- The temporary TODO regression test passes.
- `scripts/verify_exercise_starters.sh` rejects the untouched `intro1` starter for its actual behavior, not for its comment text.
- The complete solutions suite passes.
- `git diff --check` and the repository formatting/layout checks pass.
