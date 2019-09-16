<?php

use GuzzleHttp\Promise;

require './shared.php';

const NUM_WORKERS = 60;

[
  'bucket' => $bucket,
  'inputFile' => $inputFile,
  'outputDir' => $outputDir
] = getParams();

$files = readFilenams($inputFile);

function worker() {
  global $files, $bucket, $outputDir;
  if (empty($files)) {
    return;
  }

  $file = array_pop($files);
  return fetchFileAsync($bucket, $file)->then(
    function ($contents) use ($file, $outputDir) {
      file_put_contents("$outputDir/$file", $contents);
      return worker();
    }
  );
}

$workers = [];
for ($i = 0; $i < NUM_WORKERS; $i++) {
  array_push($workers, worker());
}

// Now block on all the workers finishing
Promise\all($workers)->wait();

