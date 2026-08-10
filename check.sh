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
dirs=$(
  find "$target" -mindepth 2 -maxdepth 2 -type d -print |
    awk -F/ '
      {
        topic = $(NF - 1)
        exercise = $NF
        if (match(exercise, /[0-9]+$/) == 0) {
          next
        }
        number = substr(exercise, RSTART, RLENGTH)
        printf "%s\t%09d\t%s/\n", topic, number, $0
      }
    ' |
    sort -k1,1 -k2,2n |
    cut -f3-
)

while IFS= read -r dir; do
  [ -n "$dir" ] || continue
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
done <<EOF
$dirs
EOF
echo "----"
echo "$((total - fail))/$total passed"
[ "$fail" -eq 0 ] && echo "All $target pass ✓" || exit 1
