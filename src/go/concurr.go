package main

import (
	"os"
	"path/filepath"
	"sync"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/mikehenrty/concurrency-playground/shared"
)

const NUM_WORKERS = 50
const SINGLE_FILE_CONCURRENCY = 1

var	bucket = os.Args[1]
var	input_file = os.Args[2]
var	output_dir = os.Args[3]

var files []string

func worker(downloader *s3manager.Downloader, wg *sync.WaitGroup) {
	var filename string
	for len(files) != 0 {
		// Pop from the front, (or shift)
		filename, files = files[0], files[1:]

		outfile, err := os.Create(filepath.Join(output_dir, filename))
		shared.Check(err)

		_, err = downloader.Download(outfile, &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(filename),
			})

		outfile.Close()
		shared.Check(err)
	}

	wg.Done()
}

func getNumWorkers() int {
	if len(os.Args) > 4 {
		numWorkers, err := strconv.Atoi(os.Args[4])
		shared.Check(err)
		return numWorkers
	}

	return NUM_WORKERS
}

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	shared.Check(err)

	downloader := s3manager.NewDownloader(sess)
	downloader.Concurrency = SINGLE_FILE_CONCURRENCY

	files = shared.Load(input_file)

	var wg sync.WaitGroup
	for i := 1; i < getNumWorkers(); i++ {
		wg.Add(1)
		go worker(downloader, &wg)
	}

	wg.Wait()
}
