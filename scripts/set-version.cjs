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
