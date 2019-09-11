<?php

require './shared.php';

[
  'bucket' => $bucket,
  'inputFile' => $inputFile,
  'outputDir' => $outputDir
] = getParams();

$files = readFilenams($inputFile);

// Download the contents of the object.
foreach ($files as $file) {
  $fileContent = fetchFile($bucket, $file);
  file_put_contents("$outputDir/$file", $fileContent);
}
