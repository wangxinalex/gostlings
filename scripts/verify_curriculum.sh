#!/bin/sh
# Verify curriculum invariants that the exercise checker cannot see.
#
#   1. the topic table in README.md matches the exercise tree
#   2. every exercise still has a TODO seam
#   3. every exercise ships a focused test
#   4. hints in chapters 12+ stay clues rather than copyable answers
#   5. no CJK text outside the untracked working-progress worktree
#
# Read-only and dependency-free; exits non-zero on the first batch of
# violations it finds (it always reports all of them).
#
# Usage: sh scripts/verify_curriculum.sh

set -u

script_dir=$(CDPATH= cd "$(dirname "$0")" && pwd -P)
repo_root=$(CDPATH= cd "$script_dir/.." && pwd -P)
cd "$repo_root" || exit 2

fail=0
report() {
  printf 'FAIL: %s\n' "$1" >&2
  fail=$((fail+1))
}

table=$(mktemp "${TMPDIR:-/tmp}/gostlings-curriculum.XXXXXX")
trap 'rm -f "$table"' EXIT HUP INT TERM

awk -F'|' '/^\|[[:space:]]*[0-9][0-9]_[a-z_]+[[:space:]]*\|/ {
  topic = $2
  count = $3
  gsub(/[[:space:]]/, "", topic)
  gsub(/[[:space:]]/, "", count)
  print topic, count
}' README.md > "$table"

# --- 1. the documented topic table matches the tree -------------------------

while read -r topic count; do
  [ -n "$topic" ] || continue
  dir="exercises/$topic"
  if [ ! -d "$dir" ]; then
    report "README.md lists $topic but $dir does not exist"
    continue
  fi
  actual=$(find "$dir" -mindepth 1 -maxdepth 1 -type d | wc -l | tr -d ' ')
  if [ "$actual" != "$count" ]; then
    report "$topic: README.md documents $count exercises, the tree has $actual"
  fi
done < "$table"

for dir in exercises/*/; do
  topic=$(basename "$dir")
  if ! grep -q "^$topic " "$table"; then
    report "$topic has exercises but no row in the README.md topic table"
  fi
done

# --- 2. and 3. every exercise stays an unsolved, testable starter -----------
#
# Whether a starter was solved in place is decided behaviorally by
# scripts/verify_exercise_starters.sh: a solved exercise passes its focused
# test, which that audit rejects. Comparing files against solutions/ cannot
# tell an answer from a fixture, because the 14_testing helpers and some
# *_test.go files are intentionally identical in both trees.

for dir in $(find exercises -mindepth 2 -maxdepth 2 -type d | sort); do
  if ! grep -lq 'TODO' "$dir"/*.go 2>/dev/null; then
    report "$dir has no TODO seam (the starter looks solved)"
  fi

  if ! ls "$dir"/*_test.go >/dev/null 2>&1; then
    report "$dir has no focused test file"
  fi
done

# --- 4. hints stay clues, not answers ---------------------------------------
#
# Chapters 12 and up are past the syntax tour, so their Hint and Stuck? blocks
# must not carry a copyable call, declaration, or template action. Naming an API
# is allowed; handing over the expression is not. Chapters 00-11 keep explicit
# hints on purpose, because beginners still need the exact form.

hint_files=$(find exercises -mindepth 3 -maxdepth 3 -name '*.go' ! -name '*_test.go' | sort)
hint_hits=$(awk '
  function spoil(text) {
    # A call that passes arguments hands over the expression, unless it is the
    # channel protocol notation close(name) that the channel ledgers also use.
    if (text ~ /[A-Za-z_][A-Za-z0-9_.]*\([^)]/) {
      if (text !~ /close\(/) return "a call with arguments"
    }
    if (text ~ /:=/) return "a declaration or assignment"
    if (text ~ /\{\{/) return "template syntax"
    return ""
  }
  FNR == 1 {
    count = split(FILENAME, parts, "/")
    topic = parts[count - 2]
    exempt = (topic ~ /^0[0-9]_/ || topic ~ /^1[01]_/)
    inblock = 0
  }
  /^\/\/[[:space:]]*(Hint|Stuck\?):/ { inblock = 1 }
  inblock && !exempt {
    why = spoil($0)
    if (why != "") printf "%s:%d: %s\n", FILENAME, FNR, why
  }
  /^\/\/[[:space:]]*[A-Z][A-Za-z ]*:/ {
    if ($0 !~ /^\/\/[[:space:]]*(Hint|Stuck\?):/) inblock = 0
  }
  !/^\/\// { inblock = 0 }
' $hint_files)
if [ -n "$hint_hits" ]; then
  old_ifs=$IFS
  IFS='
'
  for hit in $hint_hits; do
    report "hint gives away the answer: $hit"
  done
  IFS=$old_ifs
fi

# --- 5. no CJK text outside working-progress --------------------------------
#
# working-progress is git-ignored, so listing tracked plus untracked but
# non-ignored files exempts it without naming it. [\344-\351] matches the
# leading UTF-8 byte of CJK ideographs, which leaves the repository's existing
# en dashes, curly quotes, and arrows alone. LC_ALL=C keeps grep in byte mode so
# the raw byte pattern is accepted on both BSD and GNU grep.

cjk=$(printf '[\344-\351]')
cjk_hits=$(git ls-files --cached --others --exclude-standard -z |
  LC_ALL=C xargs -0 grep -lI "$cjk" 2>/dev/null || true)
if [ -n "$cjk_hits" ]; then
  old_ifs=$IFS
  IFS='
'
  for hit in $cjk_hits; do
    report "$hit contains CJK text; keep Chinese in working-progress only"
  done
  IFS=$old_ifs
fi

# --- summary ----------------------------------------------------------------

if [ "$fail" -ne 0 ]; then
  printf '\n%d curriculum invariant violation(s)\n' "$fail" >&2
  exit 1
fi

printf 'Curriculum invariants hold: %s exercises across %s topics\n' \
  "$(find exercises -mindepth 2 -maxdepth 2 -type d | wc -l | tr -d ' ')" \
  "$(find exercises -mindepth 1 -maxdepth 1 -type d | wc -l | tr -d ' ')"
