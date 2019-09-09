package main

import (
	"bufio"
	"os"
	"path/filepath"


	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func load(filename string) (files []string) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	reader:= bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)

	files = make([]string, 0, 100000)

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		files = append(files, scanner.Text())
	}

	return files
}

func main() {
	bucket := os.Args[1]
	input_file := os.Args[2]
	output_dir := os.Args[3]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)
	check(err)

	downloader := s3manager.NewDownloader(sess)

	files := load(input_file)
	for _, filename := range files {
		outfile, err := os.Create(filepath.Join(output_dir, filename))
		check(err)
		defer outfile.Close()

		_, err = downloader.Download(outfile,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(filename),
		})

		check(err)
	}
}
