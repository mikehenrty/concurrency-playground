const fs = require('fs');
const path = require('path');
const S3 = require('aws-sdk/clients/s3');

const s3 = new S3();

const bucket = process.argv[2];
const inputFile = process.argv[3];
const output = process.argv[4];

const contents = fs.readFileSync(inputFile, 'utf-8');
const filenames = contents.split('\n').filter(f => f);

function fetchFile(Bucket, Key, callback) {
  const params = { Bucket, Key };
  s3.getObject(params, function(err, data) {
    if (err) {
      console.error('Error fetching file', key, err);
      callback(err);
      return;
    }

    const { Body } = data;
    callback(null, Body);
  });
}

function getNext(filenames) {
  if (filenames.length === 0) {
    return;
  }

  filename = filenames.pop();
  fetchFile(bucket, filename, (err, contents) => {
    if (err) {
      console.error('Count not fetch file', filename, err);
      return;
    }

    fs.writeFile(path.join(output, filename), contents, (err) => {
      if (err) {
        console.error('Could not write file', filename, err);
        return;
      }

      getNext(filenames);
    });
  });
}


if (require.main === module) {
  getNext(filenames);
}
