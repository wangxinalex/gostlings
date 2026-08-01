#!/bin/sh
# Run exercises in order; report PASS/FAIL per exercise.
# Usage: ./check.sh [exercises|solutions] [--run-all] [--race]
# Default: stop at first FAIL and show its output. --run-all: continue and print a summary.
target="exercises"
run_all=0
race=0
for arg in "$@"; do
  case "$arg" in
    --run-all) run_all=1 ;;
    --race) race=1 ;;
    *) target="$arg" ;;
  esac
done
fail=0; total=0
for dir in "$target"/*/*/; do
  [ -d "$dir" ] || continue
  d="./${dir%/}"
  if ls "$d"/*_test.go >/dev/null 2>&1; then
    if [ "$race" -eq 1 ]; then
      cmd="go test -race -timeout 5s"
    else
      cmd="go test -timeout 5s"
    fi
  else
    if [ "$race" -eq 1 ]; then
      cmd="go run -race"
    else
      cmd="go run"
    fi
  fi
  total=$((total+1))
  out=$($cmd "$d" 2>&1)
  if [ $? -eq 0 ]; then
    printf "PASS: %s\n" "$dir"
  else
    printf "FAIL: %s\n" "$dir"
    fail=$((fail+1))
    if [ "$run_all" -eq 0 ]; then
      printf "%s\n" "$out"
      exit 1
    fi
  fi
done
echo "----"
echo "$((total - fail))/$total passed"
[ "$fail" -eq 0 ] && echo "All $target pass ✓" || exit 1
