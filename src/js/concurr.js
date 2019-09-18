const fs = require('fs');
const path = require('path');
const shared = require('./shared');

const NUM_WORKERS = 100;

const { bucket, inputFile, output, workers } = shared.getCLIParameters();
const filenames = shared.readFilenames(inputFile);

async function worker() {
  while (filenames.length > 0) {
    const filename = filenames.pop();
    const file = await shared.fetchFile(bucket, filename);
    await shared.writeFile(path.join(output, filename), file);
  }
}

function fetchAllConcurr(filenames) {
  const count = workers || NUM_WORKERS;
  for (let i = 0; i < count; i++) {
    worker();
  }
}

if (require.main === module) {
  fetchAllConcurr(filenames);
}
