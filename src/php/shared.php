<?php

require 'vendor/autoload.php';

use Aws\S3\S3Client;

$s3 = new S3Client([
  'region' => 'us-east-1',
  'version' => 'latest'
]);

function getParams() {
  global $argv;
  return [
    'bucket' => $argv[1],
    'inputFile' => $argv[2],
    'outputDir' => $argv[3],
    'workers' => $argv[4]
  ];
}

function readFilenams($inputFile) {
  $files = [];
  $handle = fopen($inputFile, 'r');
  while($line = trim(fgets($handle))) {
    array_push($files, $line);
  }
  return $files;
}

function fetchFile($bucket, $key) {
  global $s3;
  $result = $s3->getObject([
    'Bucket' => $bucket,
    'Key' => $key
  ]);
  return $result['Body'];
}

function fetchFileAsync($bucket, $key) {
  global $s3;
  return $s3->getObjectAsync([
    'Bucket' => $bucket,
    'Key' => $key
  ])->then(function ($result) {
    return $result['Body'];
  });
}
