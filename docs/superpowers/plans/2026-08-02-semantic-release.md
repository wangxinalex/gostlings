# semantic-release Support Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (- [ ] syntax) for tracking.

**Goal:** Add reproducible semantic-release automation that updates the canonical VERSION file from Conventional Commits, generates a changelog, and creates GitHub releases from successful main builds.

**Architecture:** Keep the Go project independent from Node release tooling. A private Node package owns pinned semantic-release dependencies and a CommonJS configuration; the existing CI workflow gets a release job that depends on the two existing validation jobs.

**Tech Stack:** Go 1.23, GitHub Actions, Node.js 24, npm lockfile, semantic-release, Conventional Commits, @semantic-release/changelog, @semantic-release/exec, @semantic-release/github, and @semantic-release/git.

## Global Constraints

- Never commit on main; all implementation commits go to agent/add-semantic-release.
- VERSION is the canonical repository version and starts at exactly 0.1.0.
- Only main push events may publish; pull requests and non-main pushes must never create releases.
- The release job must depend on the existing format-and-layout and solutions jobs.
- Do not run go test ./... over exercises; exercises can be intentionally incomplete.
- Do not publish an npm package; package.json is private release tooling metadata only.
- The remote baseline tag v0.1.0 must exist before the first automatic release, otherwise an untagged first release would be 1.0.0.

---

### Task 1: Add reproducible Node release tooling

**Files:**
- Create: package.json
- Create: package-lock.json

**Interfaces:**
- Produces npm ci, npm run release, and npm run release:dry-run for the workflow and verification.

- [ ] Step 1: Create private package metadata

Create package.json:

~~~json
{
  "name": "gostlings-release-tools",
  "version": "0.0.0",
  "private": true,
  "description": "Private release tooling for the gostlings Go exercises",
  "engines": {
    "node": ">=22.14.0"
  },
  "scripts": {
    "release": "semantic-release",
    "release:dry-run": "semantic-release --dry-run --no-ci"
  }
}
~~~

- [ ] Step 2: Install exact release dependencies

Run:

~~~bash
npm install --save-dev --save-exact \
  semantic-release \
  @semantic-release/commit-analyzer \
  @semantic-release/release-notes-generator \
  @semantic-release/changelog \
  @semantic-release/exec \
  @semantic-release/github \
  @semantic-release/git \
  conventional-changelog-conventionalcommits
~~~

Expected: npm writes exact versions to devDependencies, creates package-lock.json, and exits 0.

- [ ] Step 3: Verify the locked install

Run:

~~~bash
rm -rf node_modules
npm ci
npm exec semantic-release -- --help
~~~

Expected: npm ci installs from the lockfile and semantic-release prints help. Do not commit node_modules.

- [ ] Step 4: Commit the tooling

~~~bash
git add package.json package-lock.json
git commit -m "build: add semantic-release tooling"
~~~

### Task 2: Add version writing and semantic-release configuration

**Files:**
- Create: VERSION
- Create: CHANGELOG.md
- Create: scripts/set-version.cjs
- Create: release.config.cjs

**Interfaces:**
- node scripts/set-version.cjs 1.2.3 writes a validated SemVer to VERSION; VERSION_FILE may override the path for isolated tests.
- release.config.cjs restricts releases to main, uses v${version} tags, and orders changelog/version preparation before the git plugin.

- [ ] Step 1: Define the writer verification cases

Run this before implementing the writer:

~~~bash
VERSION_FILE="$(mktemp)" node scripts/set-version.cjs 1.2.3
test "$(cat "$VERSION_FILE")" = "1.2.3"
if VERSION_FILE="$VERSION_FILE" node scripts/set-version.cjs 1.2; then
  echo "invalid version was accepted" >&2
  exit 1
fi
~~~

Expected before implementation: it fails because scripts/set-version.cjs does not exist.

- [ ] Step 2: Implement the writer

Create scripts/set-version.cjs:

~~~js
const fs = require('node:fs');
const path = require('node:path');

const version = process.argv[2];
const versionPattern = /^\d+\.\d+\.\d+$/;

if (!versionPattern.test(version || '')) {
  console.error('invalid release version: ' + (version || '<missing>'));
  process.exit(1);
}

const versionFile = process.env.VERSION_FILE || path.join(__dirname, '..', 'VERSION');
fs.writeFileSync(versionFile, version + '\n');
~~~

- [ ] Step 3: Rerun the writer verification

Run the commands from Step 1. Expected: 1.2.3 is written with one trailing newline and 1.2 is rejected.

- [ ] Step 4: Add the initial version and changelog

Create VERSION with exactly 0.1.0. Create CHANGELOG.md:

~~~markdown
# Changelog

All notable changes to this project will be documented in this file.

## [0.1.0] - 2026-08-02

- Initial baseline release.
~~~

- [ ] Step 5: Configure semantic-release

Create release.config.cjs:

~~~js
module.exports = {
  branches: ['main'],
  tagFormat: 'v${version}',
  plugins: [
    ['@semantic-release/commit-analyzer', {preset: 'conventionalcommits'}],
    ['@semantic-release/release-notes-generator', {preset: 'conventionalcommits'}],
    [
      '@semantic-release/changelog',
      {changelogFile: 'CHANGELOG.md', changelogTitle: '# Changelog'},
    ],
    [
      '@semantic-release/exec',
      {prepareCmd: 'node scripts/set-version.cjs ${nextRelease.version}'},
    ],
    '@semantic-release/github',
    [
      '@semantic-release/git',
      {
        assets: ['VERSION', 'CHANGELOG.md'],
        message: 'chore(release): ${nextRelease.version} [skip ci]',
      },
    ],
  ],
};
~~~

- [ ] Step 6: Validate configuration without publishing

Run:

~~~bash
npm exec semantic-release -- --version
node -e "const config = require('./release.config.cjs'); if (config.branches[0] !== 'main' || config.tagFormat !== 'v${version}') process.exit(1)"
~~~

Expected: both commands exit 0 and no release/tag or version mutation occurs.

- [ ] Step 7: Commit the release configuration

~~~bash
git add VERSION CHANGELOG.md scripts/set-version.cjs release.config.cjs
git commit -m "feat: configure semantic-release"
~~~

### Task 3: Gate publishing behind existing CI

**Files:**
- Modify: .github/workflows/ci.yml

**Interfaces:**
- Produces a release job that runs only for push to refs/heads/main, after format-and-layout and solutions, and invokes npm run release with GITHUB_TOKEN.

- [ ] Step 1: Add the release job

Append this job:

~~~yaml
  release:
    name: Release
    if: github.event_name == 'push' && github.ref == 'refs/heads/main'
    needs:
      - format-and-layout
      - solutions
    runs-on: ubuntu-latest
    permissions:
      contents: write
      issues: write
      pull-requests: write
    steps:
      - name: Check out repository
        uses: actions/checkout@v6
        with:
          fetch-depth: 0

      - name: Set up Node.js
        uses: actions/setup-node@v6
        with:
          node-version: 24
          cache: npm

      - name: Install release tooling
        run: npm ci

      - name: Verify the release baseline tag
        shell: bash
        run: git show-ref --tags --verify --quiet refs/tags/v0.1.0

      - name: Publish release
        run: npm run release
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
~~~

- [ ] Step 2: Parse and inspect the workflow

Run:

~~~bash
ruby -e "require 'yaml'; YAML.load_file('.github/workflows/ci.yml'); puts 'valid YAML'"
rg -n "release:|refs/heads/main|format-and-layout|solutions|fetch-depth: 0|contents: write|npm run release" .github/workflows/ci.yml
~~~

Expected: YAML parses, and all release guards are present; existing exercise-aware checks are unchanged.

- [ ] Step 3: Commit the workflow

~~~bash
git add .github/workflows/ci.yml
git commit -m "ci: publish semantic releases from main"
~~~

### Task 4: Document release conventions and bootstrap

**Files:**
- Modify: README.md

**Interfaces:**
- Documents the version source, commit-to-version mapping, main-only publication, and one-time v0.1.0 tag setup.

- [ ] Step 1: Add the release section

Add a section stating: VERSION is canonical; successful main pushes update CHANGELOG.md, the GitHub Release, and vX.Y.Z; fix:/perf: are patch; feat: is minor; ! or a BREAKING CHANGE: footer is major; docs:/test:/ci:/chore:/refactor: do not release alone.

Also document the one-time post-merge bootstrap:

~~~bash
git tag -a v0.1.0 "$(git rev-parse HEAD)" -m "chore: mark 0.1.0 baseline"
git push origin v0.1.0
~~~

- [ ] Step 2: Check documentation/config consistency

Confirm README and release.config.cjs agree on VERSION, CHANGELOG.md, main, v0.1.0, and the release categories.

- [ ] Step 3: Commit the documentation

~~~bash
git add README.md
git commit -m "docs: document semantic release workflow"
~~~

### Task 5: Run the verification matrix

**Files:**
- Test: all release files, workflow, and existing Go checks

**Interfaces:**
- Verifies intended changes only and preserves the existing exercise-aware CI behavior.

- [ ] Step 1: Verify version and formatting

~~~bash
test "$(cat VERSION)" = "0.1.0"
gofmt -d $(git ls-files '*.go')
~~~

Expected: the version assertion succeeds and gofmt prints no diff.

- [ ] Step 2: Run existing Go checks

~~~bash
GOCACHE=/tmp/gostlings-go-cache go test ./solutions/...
GOCACHE=/tmp/gostlings-go-cache go vet ./solutions/...
GOCACHE=/tmp/gostlings-go-cache sh check.sh solutions --run-all
GOCACHE=/tmp/gostlings-go-cache sh check.sh solutions --run-all --race
~~~

Expected: every command exits 0; no command compiles the incomplete exercises tree as a whole.

- [ ] Step 3: Verify clean lockfile installation

~~~bash
rm -rf node_modules
npm ci
~~~

Expected: install succeeds from package-lock.json.

- [ ] Step 4: Verify a patch dry-run in an isolated clone

Run:

~~~bash
tmp_dir="$(mktemp -d)"
git clone . "$tmp_dir/repo"
git -C "$tmp_dir/repo" switch -c main
git -C "$tmp_dir/repo" tag v0.1.0
git -C "$tmp_dir/repo" commit --allow-empty -m "fix: verify patch release"
(cd "$tmp_dir/repo" && npm ci && npm run release:dry-run)
~~~

Expected: semantic-release recognizes main, reads v0.1.0, reports next version 0.1.1, and does not create a tag, GitHub Release, or committed file change in the temporary clone.

- [ ] Step 5: Verify branch and final diff

~~~bash
git status -sb
git diff main...HEAD --check
git log --oneline --decorate main..HEAD
~~~

Expected: current branch is agent/add-semantic-release, the diff has no whitespace errors, and every implementation commit is ahead of main.

- [ ] Step 6: Commit only after all checks pass

If a check fails, fix the specific failure and rerun it. Before handoff, confirm the final implementation commits are on agent/add-semantic-release and no implementation commit exists on main.
