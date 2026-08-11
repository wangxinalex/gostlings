#!/bin/sh
# Run exercises in order; report PASS/FAIL per exercise.
# Usage: ./check.sh [exercises|solutions|path/to/exercise] [--run-all] [--race] [--race-all] [--verify-starter]
# Default: stop at first FAIL and show its output. --run-all: continue and print a summary.
target="exercises"
run_all=0
race=0
race_all=0
verify_starter=0

usage() {
  printf "Usage: sh check.sh <target> [--run-all] [--race] [--race-all] [--verify-starter]\n" >&2
}

for arg in "$@"; do
  case "$arg" in
    --run-all) run_all=1 ;;
    --race) race=1 ;;
    --race-all) race_all=1 ;;
    --verify-starter) verify_starter=1 ;;
    *) target="$arg" ;;
  esac
done

if [ "$race" -eq 1 ] && [ "$race_all" -eq 1 ]; then
  usage
  exit 2
fi

fail=0; total=0
if [ ! -d "$target" ]; then
  printf "target directory does not exist: %s\n" "$target" >&2
  exit 2
fi

exercise_root=""
if [ -d exercises ]; then
  exercise_root=$(cd exercises && pwd -P)
fi
solution_root=""
if [ -d solutions ]; then
  solution_root=$(cd solutions && pwd -P)
fi
target_path=$(cd "$target" && pwd -P)

if [ "$verify_starter" -eq 1 ]; then
  if [ -z "$exercise_root" ]; then
    usage
    exit 2
  fi
  case "$target_path/" in
    "$exercise_root/"*) ;;
    *)
      usage
      exit 2
      ;;
  esac
fi

if [ -f "$target/main.go" ] || ls "$target"/*_test.go >/dev/null 2>&1; then
  dirs="${target%/}/"
else
  dirs=$(
    find "$target" -type f \( -name main.go -o -name '*_test.go' \) -print |
      sed 's#/[^/]*$##' |
      sort -u |
      awk -F/ '
        {
          topic = $(NF - 1)
          exercise = $NF
          if (match(topic, /^[0-9]+/) == 0) {
            next
          }
          topic_number = substr(topic, RSTART, RLENGTH)
          if (match(exercise, /[0-9]+$/) == 0) {
            next
          }
          exercise_number = substr(exercise, RSTART, RLENGTH)
          printf "%09d\t%09d\t%s/\n", topic_number, exercise_number, $0
        }
      ' |
      sort -k1,1n -k2,2n -k3,3 |
      cut -f3-
  )
fi

while IFS= read -r dir; do
  [ -n "$dir" ] || continue
  [ -d "$dir" ] || continue
  case "$dir" in
    /*) d="${dir%/}" ;;
    *) d="./${dir%/}" ;;
  esac
  dir_path=$(cd "$d" && pwd -P)
  exercise_path=""
  if [ -n "$exercise_root" ]; then
    case "$dir_path/" in
      "$exercise_root/"*) exercise_path=${dir_path#"$exercise_root"/} ;;
    esac
  fi
  if [ -z "$exercise_path" ] && [ -n "$solution_root" ]; then
    case "$dir_path/" in
      "$solution_root/"*) exercise_path=${dir_path#"$solution_root"/} ;;
    esac
  fi

  if [ -n "$exercise_path" ] && [ -f "$d/main.go" ] &&
    grep -q 'TODO:' "$d/main.go"; then
    if [ "$verify_starter" -eq 0 ]; then
      total=$((total+1))
      fail=$((fail+1))
      printf "FAIL: %s (starter still contains TODO)\n" "$dir"
      if [ "$run_all" -eq 0 ]; then
        exit 1
      fi
      continue
    fi
  fi

  run_race=0
  if [ "$race_all" -eq 1 ]; then
    run_race=1
  elif [ "$race" -eq 1 ] && [ -n "$exercise_path" ] &&
    [ -f race.list ] && grep -F -x "$exercise_path" race.list >/dev/null 2>&1; then
    run_race=1
  fi

  if ls "$d"/*_test.go >/dev/null 2>&1; then
    if grep -q '^[[:space:]]*func Benchmark' "$d"/*_test.go 2>/dev/null; then
      if [ "$run_race" -eq 1 ]; then
        cmd="go test -race -bench=. -timeout 5s"
      else
        cmd="go test -bench=. -timeout 5s"
      fi
    else
      if [ "$run_race" -eq 1 ]; then
        cmd="go test -race -timeout 5s"
      else
        cmd="go test -timeout 5s"
      fi
    fi
  else
    if [ "$run_race" -eq 1 ]; then
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
