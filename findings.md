# Concurrency Playground

This is a small test of various languauges and their ability to download files from s3 using serial, concurrent and parallel structures. Findings will be captured in this file.

## Running the code

Each language will have up to 3 entry points, one for serialized s3/filesystem access, one for concurrent s3/filesystem access (if language supports it), and one for parallel s3/filesystem access.

Generally code is run from the root project directory like this:
```bash
node js/serial.js MyS3BucketName ListOfS3Files.txt outputDir
```

## Serial

* Node
100 files - `1.11s user 0.17s system 2% cpu 53.618 total`
1088k / 53.618s = `20.29 k/s`

* Python
100 files - `2.49s user 0.12s system 7% cpu 33.700 total`
1099k / 33.700s = `32.28 k/s`

* PHP
100 files - `0.31s user 0.06s system 2% cpu 17.368 total`
1099k / 17.368 = `63.28 k/s`
