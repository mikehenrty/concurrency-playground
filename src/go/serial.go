// +build serial

package main

import (
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/mikehenrty/concurrency-playground/shared"
)

func main() {
	bucket := os.Args[1]
	input_file := os.Args[2]
	output_dir := os.Args[3]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	shared.Check(err)

	downloader := s3manager.NewDownloader(sess)
	downloader.Concurrency = 1

	files := shared.Load(input_file)
	for _, filename := range files {
		outfile, err := os.Create(filepath.Join(output_dir, filename))
		shared.Check(err)
		defer outfile.Close()

		_, err = downloader.Download(outfile,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})

		shared.Check(err)
	}
}
