# Concurrency Playground

This is a small test of various languauges and their ability to download files from s3 using serial, concurrent and parallel structures. Findings will be captured in this file.

## Running the code

Each language will have up to 3 entry points, one for serialized s3/filesystem access, one for concurrent s3/filesystem access (if language supports it), and one for parallel s3/filesystem access.

Generally code is run from the root project directory like this:
```bash
js/run serial MyS3BucketName ListOfS3Files.txt outputDir
go/run concurr MyS3BucketName ListOfS3Files.txt outputDir
```
Notice we use either `serial` or `concurr` as the fisrt argument, to specify which download approach.

## Concurrent from SageMaker Notebook instance
lanugage,real,user,system,cpu
js,11.05,9.98,1.91,107.64
python,123.71,133.03,9.50,115.21
php,9.62,8.15,1.26,97.87
go,5.62,6.71,1.97,154.35

## Serial from SageMaker Notebook instance

# Node
1000 files 12060kb / ~45.33 = 244.03 kb/s

real 49.95
user 5.34
sys 0.49

real 44.15
user 5.32
sys 0.63

real 41.90
user 5.07
sys 0.62

# Python

1000 files 12060kb / ~82.39 = 146.376 kb/s

real 84.02
user 14.60
sys 1.60

real 82.05
user 14.74
sys 1.53

real 81.11
user 14.81
sys 1.51

# PHP
1000 files 12060kb / ~38.21s = 315.624 kb/s

real    0m36.113s
user    0m1.191s
sys     0m0.214s

real    0m42.675s
user    0m1.248s
sys     0m0.219s

real    0m35.864s
user    0m1.145s
sys     0m0.282s

# Go
1000 files 12060kb / ~39.04s = 308.91 kb/s

real    0m45.753s
user    0m0.710s
sys     0m0.159s

real    0m34.184s
user    0m0.668s
sys     0m0.185s

real 37.18
user 0.68
sys 0.16


## From home

# Node
### Serial
100 files - `1.11s user 0.17s system 2% cpu 53.618 total`
1088k / 53.618s = `20.29 k/s`

### Concurrent
1000 files - `2.57s user 0.38s system 44% cpu 6.711 total`
12064k / 6.711s = `1797.65 k/s`

# Python
### Serial
100 files - `2.49s user 0.12s system 7% cpu 33.700 total`
1099k / 33.700s = `32.28 k/s`

### Concurrent
1000 files - `16.62s user 4.44s system 163% cpu 12.915 total`
12064k / 12.915 = `934.11 k/s`


# PHP
### Serial
100 files - `0.31s user 0.06s system 2% cpu 17.368 total`
1099k / 17.368 = `63.28 k/s`

### Concurrent
1000 files = `1.51s user 0.50s system 43% cpu 4.593 total`
12064k / 4.593 = `2626.61 k/s`

# Go
### Serial
100 files - `0.15s user 0.07s system 1% cpu 17.585 total`
1099k / 16.536 = `66.46 k/s`

### Concurrent
1000 files - `0.71s user 0.28s system 19% cpu 5.143 total`
12064k / 5.143 = `2345.71 k/s`
