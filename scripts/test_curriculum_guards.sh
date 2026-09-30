#!/bin/sh
# Prove that scripts/verify_curriculum.sh rejects every violation it claims to
# catch. Each case builds a throwaway fixture repository, mutates exactly one
# invariant, and asserts the guard's verdict. Without this the guard could rot
# into a script that reports success unconditionally.
#
# Usage: sh scripts/test_curriculum_guards.sh

set -eu

script_dir=$(CDPATH= cd "$(dirname "$0")" && pwd -P)
repo_root=$(CDPATH= cd "$script_dir/.." && pwd -P)
tmp=$(mktemp -d "${TMPDIR:-/tmp}/gostlings-guards.XXXXXX")
trap 'rm -rf "$tmp"' EXIT HUP INT TERM

fixture() {
  rm -rf "$tmp/repo"
  mkdir -p "$tmp/repo/scripts" "$tmp/repo/exercises/00_intro/intro1" \
    "$tmp/repo/exercises/03_control_flow/flow1" "$tmp/repo/exercises/27_strings/strings1" \
    "$tmp/repo/solutions/00_intro/intro1" "$tmp/repo/working-progress"
  cp "$repo_root/scripts/verify_curriculum.sh" "$tmp/repo/scripts/"

  cat > "$tmp/repo/README.md" <<'EOF'
| Topic | # | Go Tour section |
|---|---|---|
| 00_intro | 1 | Basics 1 |
| 03_control_flow | 1 | Flowcontrol |
| 27_strings | 1 | strings |
EOF

  cat > "$tmp/repo/exercises/00_intro/intro1/main.go" <<'EOF'
// Concept: fixture
// Task: make the fixture pass
package main

func main() {
	// TODO: print the expected value.
}
EOF

  cat > "$tmp/repo/exercises/00_intro/intro1/main_test.go" <<'EOF'
package main

import "testing"

func TestOutput(t *testing.T) {
	if false {
		t.Fatal("fixture starter must fail until solved")
	}
}
EOF

  cat > "$tmp/repo/exercises/03_control_flow/flow1/main.go" <<'EOF'
// Concept: fixture
// Task: make the fixture pass
// Hint: compare the loop index with the slice length.
package main

func main() {
	// TODO: print the expected value.
}
EOF

  cp "$tmp/repo/exercises/03_control_flow/flow1/main.go" \
    "$tmp/repo/exercises/03_control_flow/flow1/main.go.tmp"
  cat > "$tmp/repo/exercises/03_control_flow/flow1/main_test.go" <<'EOF'
package main

import "testing"

func TestOutput(t *testing.T) {
	if false {
		t.Fatal("fixture starter must fail until solved")
	}
}
EOF
  rm -f "$tmp/repo/exercises/03_control_flow/flow1/main.go.tmp"

  cat > "$tmp/repo/exercises/27_strings/strings1/main.go" <<'EOF'
// Concept: fixture
// Task: make the fixture pass
// Hint: the strings package has a predicate for substring membership.
package main

func main() {
	// TODO: print the expected value.
}
EOF

  cp "$tmp/repo/exercises/03_control_flow/flow1/main_test.go" \
    "$tmp/repo/exercises/27_strings/strings1/main_test.go"

  cat > "$tmp/repo/solutions/00_intro/intro1/main.go" <<'EOF'
package main

func main() {}
EOF

  printf 'working-progress/\n' > "$tmp/repo/.gitignore"
  printf 'progress notes\n' > "$tmp/repo/working-progress/notes.md"
  git -C "$tmp/repo" init -q
}

guard() {
  sh "$tmp/repo/scripts/verify_curriculum.sh" 2>&1
}

expect_pass() {
  if ! out=$(guard); then
    printf 'FAIL: guard rejected a valid fixture: %s\n' "$1" >&2
    printf '%s\n' "$out" >&2
    exit 1
  fi
  printf 'PASS: %s\n' "$1"
}

expect_fail() {
  if out=$(guard); then
    printf 'FAIL: guard accepted a violation: %s\n' "$1" >&2
    printf '%s\n' "$out" >&2
    exit 1
  fi
  printf 'PASS: %s\n' "$1"
}

fixture
expect_pass "a consistent fixture repository"

fixture
sed -i.bak 's/| 00_intro | 1 |/| 00_intro | 2 |/' "$tmp/repo/README.md"
expect_fail "a documented count that disagrees with the tree"

fixture
printf '| 99_extra | 1 | nothing |\n' >> "$tmp/repo/README.md"
expect_fail "a topic that exists only in the table"

fixture
sed -i.bak '/TODO/d' "$tmp/repo/exercises/00_intro/intro1/main.go"
expect_fail "a starter with no TODO seam"

fixture
rm "$tmp/repo/exercises/00_intro/intro1/main_test.go"
expect_fail "an exercise with no focused test"

fixture
sed -i.bak 's/the strings package has a predicate for substring membership/strings.Contains(text, substr) is the answer/' \
  "$tmp/repo/exercises/27_strings/strings1/main.go"
expect_fail "an answer-carrying hint in a chapter past the core track"

fixture
sed -i.bak 's/compare the loop index with the slice length/strings.Contains(text, substr) is fine here/' \
  "$tmp/repo/exercises/03_control_flow/flow1/main.go"
expect_pass "a core chapter that keeps its explicit hint"

fixture
# CJK bytes are written as octal escapes on purpose: the guard needs CJK input
# to be tested, and a literal here would itself violate the rule it checks.
printf '// \350\247\243\346\263\225\n' >> "$tmp/repo/exercises/00_intro/intro1/main.go"
expect_fail "CJK text in an exercise file"

fixture
mkdir -p "$tmp/repo/exercises/01_extra/extra1"
cp "$tmp/repo/exercises/00_intro/intro1/main.go" "$tmp/repo/exercises/01_extra/extra1/main.go"
cp "$tmp/repo/exercises/00_intro/intro1/main_test.go" "$tmp/repo/exercises/01_extra/extra1/main_test.go"
expect_fail "a topic directory that the table omits"

fixture
printf '\350\277\233\345\272\246\357\274\232\345\267\262\345\256\214\346\210\220\n' > "$tmp/repo/working-progress/notes.md"
expect_pass "CJK text inside the ignored working-progress worktree"

printf '\nall curriculum guard cases behave as specified\n'
