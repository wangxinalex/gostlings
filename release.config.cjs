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
