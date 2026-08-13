#!/bin/sh
# Verify every reference solution against its matching exercise tests.
#
# The exercise tree owns the *_test.go files. This script overlays them into
# the matching solution directories temporarily, runs `go test` against all
# solutions, and then removes the overlay. Solution directories that already
# ship a *_test.go are skipped: the 14_testing chapter commits its tests
# because the test file is the deliverable there, and a couple of other
# directories commit copies so `check.sh solutions` also runs their tests.
# Everywhere else, the exercise tests stay the single source of truth.

set -eu

script_dir=$(CDPATH= cd "$(dirname "$0")" && pwd -P)
repo_root=$(CDPATH= cd "$script_dir/.." && pwd -P)
cd "$repo_root"

overlay_list=$(mktemp "${TMPDIR:-/tmp}/gostlings-solutions.XXXXXX")

cleanup() {
  if [ -f "$overlay_list" ]; then
    while IFS= read -r d; do
      [ -n "$d" ] || continue
      find "$d" -maxdepth 1 -name '*_test.go' -delete
    done < "$overlay_list"
    rm -f "$overlay_list"
  fi
}
trap cleanup EXIT

for d in $(find solutions -mindepth 2 -maxdepth 2 -type d | sort); do
  if find "$d" -maxdepth 1 -name '*_test.go' -print -quit | grep -q .; then
    continue
  fi
  rel=${d#solutions/}
  src="exercises/$rel"
  if [ -d "$src" ] && find "$src" -maxdepth 1 -name '*_test.go' -print -quit | grep -q .; then
    find "$src" -maxdepth 1 -name '*_test.go' -exec cp {} "$d/" \;
    printf '%s\n' "$d" >> "$overlay_list"
  fi
done

go test ./solutions/...
