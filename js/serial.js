const fs = require('fs');
const path = require('path');
const shared = require('./shared');

const { bucket, inputFile, output } = shared.getCLIParameters();
const filenames = shared.readFilenames(inputFile);

async function fetchAllSerial(filenames) {
  while (filenames.length > 0) {
    filename = filenames.pop();
    const contents = await shared.fetchFile(bucket, filename);
    await shared.writeFile(path.join(output, filename), contents);
  }
}

if (require.main === module) {
  fetchAllSerial(filenames);
}
