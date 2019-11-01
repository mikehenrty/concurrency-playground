// +build !serial

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

var files chan string

func worker(downloader *s3manager.Downloader) {
	for filename := range files {
		outfile, err := os.Create(filepath.Join(output_dir, filename))
		shared.Check(err)

		_, err = downloader.Download(outfile, &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(filename),
			})

		shared.Check(err)
		outfile.Close()
	}
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

	//create a channel to hold up to our number of workers files
	files = make(chan string, getNumWorkers())

	var wg sync.WaitGroup
	fileString := shared.Load(input_file)
	wg.Add(1)
	//fill the channel
	go func(){
		defer wg.Done()
		defer close(files)
		for _, f := range fileString {
			files <- f
		}
	}()

	for i := 1; i < getNumWorkers(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(downloader)
		}()
	}

	wg.Wait()
}
