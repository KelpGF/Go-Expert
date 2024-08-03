package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string

	wg sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String("you-region"),
			Credentials: credentials.NewStaticCredentials(
				"you-id",
				"you-secret",
				"",
			),
		},
	)

	if err != nil {
		panic(err)
	}

	s3Client = s3.New(sess)
	s3Bucket = "you-bucket"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	errorControl := make(chan string, 10)

	go uploadRetry(
		uploadControl,
		errorControl,
		&wg,
	)

	for {
		files, err := dir.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Printf("Error reading directory: %v\n", err)
			continue
		}

		uploadFileAsync(files[0].Name(), uploadControl, errorControl, &wg)
	}

	wg.Wait()
}

func uploadFileAsync(
	filename string,
	uploadControl chan struct{},
	errorControl chan string,
	wg *sync.WaitGroup,
) {
	wg.Add(1)
	uploadControl <- struct{}{}

	go func() {
		defer wg.Done()

		fmt.Printf("Uploading file %s\n", filename)

		err := uploadFile(filename)

		if err != nil {
			fmt.Printf("Error uploading file %s: %v\n", filename, err)
			errorControl <- err.Error()
		}

		fmt.Printf("File %s uploaded\n", filename)
		<-uploadControl
	}()
}

func uploadFile(filename string) error {
	completeFileName := fmt.Sprintf("./tmp/%s", filename)
	file, err := os.Open(completeFileName)

	if err != nil {
		errMessage := fmt.Sprintf("Error opening file %s: %v", completeFileName, err)

		return errors.New(errMessage)
	}
	defer file.Close()

	_, err = s3Client.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(s3Bucket),
			Key:    aws.String(filename),
			Body:   file,
		},
	)
	if err != nil {
		errMessage := fmt.Sprintf("Error uploading file %s: %v", completeFileName, err)

		return errors.New(errMessage)
	}

	return nil
}

func uploadRetry(
	uploadControl chan struct{},
	errorControl chan string,
	wg *sync.WaitGroup,
) {
	for {
		select {
		case filename := <-errorControl:
			fmt.Printf("Retrying upload: %s\n", filename)
			wg.Add(1)
			uploadControl <- struct{}{}
			go uploadFileAsync(filename, uploadControl, errorControl, wg)
		}
	}
}
