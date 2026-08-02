# semantic-release 支持设计

## 背景

这是一个 Go 练习项目。仓库目前没有版本文件、Node 发布工具配置或 Git release tag；现有 GitHub Actions 负责格式检查、solutions 测试、vet 和练习脚本验证。`exercises` 中存在按设计不能整体编译的代码，因此发布流程不能把整个仓库作为一个 Go 可编译模块处理。

## 目标

1. 根据提交信息自动计算 SemVer 版本。
2. 在仓库中保留当前对外版本号，使用根目录 `VERSION` 文件作为唯一展示来源。
3. 自动生成 `CHANGELOG.md`，创建 GitHub Release 和 `vX.Y.Z` tag。
4. 发布只在 `main` 分支的 push 上执行，并与现有 CI 分工。
5. 使用锁定的 Node 依赖，确保本地 dry-run 与 GitHub Actions 可复现。

## 非目标

- 不发布 npm 包；Node 依赖只用于运行 semantic-release。
- 不改变 Go module 名称、练习目录结构或现有 `check.sh` 行为。
- 不让现有 CI 对 `exercises` 执行整体 `go test ./...` 或整体编译。
- 不在 pull request 上创建正式 release。

## 方案

### 版本来源与初始化

- 新增根目录 `VERSION`，初始内容为 `0.1.0`。
- 新增 `CHANGELOG.md`，初始内容保留项目尚未发布的说明；后续由 changelog plugin 在 release 的 `prepare` 阶段更新。
- 由于 semantic-release 通过 Git tag 判断上一次 release，发布能力启用前需要把当前基线 commit 标记为 `v0.1.0` 并推送到远端。这样首次自动发布从 `0.1.0` 之后计算，而不会意外生成 `1.0.0`。
- semantic-release 生成的版本通过 `@semantic-release/exec` 写入 `VERSION`，不依赖 `package.json` 的版本字段。

### Commit 分析规则

使用 `conventional-changelog-conventionalcommits` preset：

| 提交格式 | 版本变化 |
| --- | --- |
| `fix:`、`perf:` | patch |
| `feat:` | minor |
| `feat!:`、任意类型带 `!`、正文包含 `BREAKING CHANGE:` | major |
| `docs:`、`test:`、`ci:`、`chore:`、`refactor:` | 不发布，除非包含 breaking change |

发布文档中明确说明以上规则，并建议后续提交遵守 Conventional Commits。

### 发布工作流

在现有 `.github/workflows/ci.yml` 中新增 `release` job，而不是重复创建一套发布 workflow：

1. 整个 workflow 仍响应 `push` 和 pull request；`release` job 只在 `push` 到 `main` 且 `format-and-layout`、`solutions` 两个 job 成功后执行。
2. checkout 时必须保留完整 Git 历史和 tags，供 semantic-release 分析提交。
3. 使用 Node.js 与 lockfile 安装发布工具。
4. 复用现有 CI 的成功结果后再执行 semantic-release，发布失败不能创建半成品 release。
5. 使用 GitHub Actions 内置 `GITHUB_TOKEN`，在 `release` job 上声明创建 tag、release 和更新 release commit 所需权限。
6. release commit 使用 `chore(release): <version> [skip ci]`，避免自动生成的版本提交再次触发不必要的 CI。

插件顺序固定为：

1. `@semantic-release/commit-analyzer`
2. `@semantic-release/release-notes-generator`
3. `@semantic-release/changelog`
4. `@semantic-release/exec`
5. `@semantic-release/github`
6. `@semantic-release/git`

changelog 和 VERSION 的更新必须发生在 git plugin 提交文件之前。

### 工具依赖

- 用私有 `package.json` 管理 semantic-release 运行时依赖，不将项目当作 npm 包发布。
- 提交 `package-lock.json`，发布 workflow 使用 `npm ci`。
- semantic-release 及插件使用精确版本；升级发布工具需要单独的依赖更新提交。

## 验证标准

- `VERSION` 只包含合法的 `MAJOR.MINOR.PATCH` 字符串。
- `gofmt -d $(git ls-files '*.go')` 无输出。
- 现有 solutions 测试、vet、`check.sh solutions --run-all` 与 race 检查继续通过。
- `npm ci` 能根据 lockfile 完成安装。
- `npx semantic-release --dry-run --no-ci` 能读取配置和当前 Git 历史；没有满足发布条件时不修改文件。
- 使用临时 Git 仓库或 semantic-release dry-run 验证以下映射：`fix` 为 patch、`feat` 为 minor、breaking change 为 major、docs/ci/chore 不发布。
- 工作流 YAML 可被解析，release job 仅允许 `main` push 且依赖现有验证 job 成功后触发。

## 风险与回滚

- GitHub Actions 的 `GITHUB_TOKEN` 权限不足会导致 tag 或 release 创建失败；workflow 明确声明 `contents: write`，并在发布前检查权限。
- release commit 回写仓库会触发现有 CI；`[skip ci]` 只用于避免重复流水线，不影响 GitHub Release 创建。
- 若发布配置有误，删除错误的远端 release/tag 可能造成历史不一致；优先修复配置并按 semantic-release 的 tag 历史继续发布，不重写已有 release 历史。
- 如需停用自动发布，只需暂时禁用 `release` job；`VERSION`、CHANGELOG 和历史 tag 保留。
