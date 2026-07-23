#!/bin/sh
# Run every solution; list failures and exit non-zero if any fail
fail=0
for dir in solutions/*/*/; do
  [ -d "$dir" ] || continue
  d="./${dir%/}"
  case "$dir" in
    solutions/14_testing/*) cmd="go test" ;;
    *) cmd="go run" ;;
  esac
  if ! $cmd "$d" >/dev/null 2>&1; then
    echo "FAIL: $dir"
    fail=1
  fi
done
if [ "$fail" -eq 0 ]; then
  echo "All solutions pass ✓"
else
  exit 1
fi
