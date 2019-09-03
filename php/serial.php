<?php

require 'vendor/autoload.php';

use Aws\S3\S3Client;

$s3 = new S3Client([
  'region' => 'us-east-1',
  'version' => 'latest'
]);

$bucket = $argv[1];
$inputFile = $argv[2];
$outputDir = $argv[3];

$files = [];
$handle = fopen($inputFile, 'r');
while($line = trim(fgets($handle))) {
  array_push($files, $line);
}

// Download the contents of the object.
foreach ($files as $file) {
  $result = $s3->getObject([
    'Bucket' => $bucket,
    'Key' => $file
  ]);
  file_put_contents("$outputDir/$file", $result['Body']);
}
