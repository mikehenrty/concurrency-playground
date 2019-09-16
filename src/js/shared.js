const fs = require('fs');
const util = require('util');
const S3 = require('aws-sdk/clients/s3');
const s3 = new S3();

module.exports = {
  getCLIParameters: function() {
    return {
      bucket: process.argv[2],
      inputFile: process.argv[3],
      output: process.argv[4],
    };
  },

  readFilenames: function(filename) {
    const contents = fs.readFileSync(filename, 'utf-8');
    return contents.split('\n').filter(f => f);
  },

  fetchFile: async function(Bucket, Key) {
    const params = { Bucket, Key };
    const { Body } = await s3.getObject(params).promise();
    return Body;
  },

  writeFile: util.promisify(fs.writeFile),
}


