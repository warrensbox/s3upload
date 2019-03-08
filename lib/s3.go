package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (id *Constructor) PushToS3() error {

	uploader := s3manager.NewUploader(id.Session)

	var files []string

	err := filepath.Walk(id.Directory, visit(&files))
	if err != nil {
		panic(err)
	}

	for _, fil := range files {

		fmt.Printf("fil %s\n", fil)
		file, err := os.Open(fil)
		if err != nil {
			return err
		}
		//defer file.Close()

		run := before(fil, "/")

		fmt.Printf("run %s\n", run)

		//Get file size and read the file content into a buffer
		fileInfo, _ := file.Stat()
		var size int64 = fileInfo.Size()
		buffer := make([]byte, size)
		file.Read(buffer)

		//os.Exit(0)
		// // Upload the file to S3.
		result, errS3 := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(id.MyBucket),
			Key:    aws.String(run),
			Body:   bytes.NewReader(buffer),
		})
		if errS3 != nil {
			fmt.Printf("failed to upload file, %v", errS3)
		}
		fmt.Printf("file uploaded to, %s\n", result.Location)

	}

	return nil
}

func visit(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if info.IsDir() {
			return nil
		}
		*files = append(*files, path)
		return nil
	}
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[pos+1:]
}
