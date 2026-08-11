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
