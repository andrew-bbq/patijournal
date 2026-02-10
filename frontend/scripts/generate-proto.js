const { execSync } = require('child_process');
const path = require('path');

const protocGenTsProto = path.join(__dirname, '../node_modules/.bin/protoc-gen-ts_proto.cmd');
const outputDir = path.join(__dirname, '../src/generated');
const protoPath = path.join(__dirname, '../../proto');
const protoFile = path.join(protoPath, 'entry.proto');

const command = `protoc --plugin=protoc-gen-ts_proto="${protocGenTsProto}" --ts_proto_out="${outputDir}" --proto_path="${protoPath}" "${protoFile}"`;

try {
  execSync(command, { stdio: 'inherit' });
  console.log('Proto files generated successfully!');
} catch (error) {
  console.error('Failed to generate proto files');
  process.exit(1);
}
