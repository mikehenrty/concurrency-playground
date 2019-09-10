const fs = require('fs');
const path = require('path');
const shared = require('./shared');

const NUM_WORKERS = 100;

const { bucket, inputFile, output } = shared.getCLIParameters();
const filenames = shared.readFilenames(inputFile);

async function worker() {
  while (filenames.length > 0) {
    const filename = filenames.pop();
    const file = await shared.fetchFile(bucket, filename);
    await shared.writeFile(path.join(output, filename), file);
  }
}

function fetchAllConcurr(filenames) {
  for (let i = 0; i < NUM_WORKERS; i++) {
    worker();
  }
}

if (require.main === module) {
  fetchAllConcurr(filenames);
}
