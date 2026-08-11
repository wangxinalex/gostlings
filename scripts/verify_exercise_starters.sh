#!/bin/sh
# Verify that untouched exercise starters fail their focused checks.

usage() {
  printf "Usage: sh scripts/verify_exercise_starters.sh exercises/<topic> [...]\n" >&2
}

if [ "$#" -eq 0 ]; then
  usage
  exit 2
fi

script_dir=$(CDPATH= cd "$(dirname "$0")" && pwd -P)
repo_root=$(CDPATH= cd "$script_dir/.." && pwd -P)
cd "$repo_root" || exit 2

for topic in "$@"; do
  if [ ! -d "$topic" ]; then
    printf "topic directory does not exist: %s\n" "$topic" >&2
    exit 2
  fi

  dirs=$(
    find "$topic" -mindepth 1 -maxdepth 1 -type d -print |
      awk -F/ '
        {
          exercise = $NF
          if (match(exercise, /[0-9]+$/) == 0) {
            next
          }
          number = substr(exercise, RSTART, RLENGTH)
          printf "%09d\t%s\n", number, $0
        }
      ' |
      sort -n |
      cut -f2-
  )

  while IFS= read -r exercise; do
    [ -n "$exercise" ] || continue
    if [ ! -f "$exercise/main.go" ]; then
      printf "FAIL: %s (missing main.go)\n" "$exercise" >&2
      exit 1
    fi
    if [ ! -f "$exercise/main_test.go" ]; then
      printf "FAIL: %s (missing main_test.go)\n" "$exercise" >&2
      exit 1
    fi
    if ! grep -q 'TODO:' "$exercise/main.go"; then
      printf "FAIL: %s (starter is missing TODO:)\n" "$exercise" >&2
      exit 1
    fi

    if out=$(sh check.sh "$exercise" --verify-starter 2>&1); then
      printf "FAIL: %s (starter unexpectedly passed)\n" "$exercise" >&2
      printf "%s\n" "$out" >&2
      exit 1
    fi
  done <<EOF
$dirs
EOF
done
