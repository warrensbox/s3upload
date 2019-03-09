package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var wg = sync.WaitGroup{}

func (id *Constructor) PushToS3() error {

	uploader := s3manager.NewUploader(id.Session)

	var files []string

	err := filepath.Walk(id.Directory, visit(&files))
	if err != nil {
		panic(err)
	}
	ch := make(chan string)

	for _, fil := range files {
		wg.Add(1)
		file, err := os.Open(fil)
		if err != nil {
			return err
		}
		//defer file.Close()

		fmt.Printf("Files %s\n", fil)

		if !id.IncludeBase {
			fil = removeBaseDir(fil, "/")
		}

		go pushingToS3(file, uploader, id.Bucket, fil, ch)

		// go func() {
		// 	pushingToS3(file, uploader, id.MyBucket, fil)
		// 	wg.Done()
		// }()

		//fmt.Printf("run %s\n", fil)

		//Get file size and read the file content into a buffer
		// fileInfo, _ := file.Stat()
		// size := fileInfo.Size()
		// buffer := make([]byte, size)
		// file.Read(buffer)

		// // // Upload the file to S3.
		// result, errS3 := uploader.Upload(&s3manager.UploadInput{
		// 	Bucket: aws.String(id.MyBucket),
		// 	Key:    aws.String(fil),
		// 	Body:   bytes.NewReader(buffer),
		// })
		// if errS3 != nil {
		// 	fmt.Printf("failed to upload file, %v", errS3)
		// }
		// fmt.Printf("file uploaded to, %s\n", result.Location)

	}

	go func(ch chan<- string) {
		defer close(ch)
		wg.Wait()
	}(ch)

	for i := range ch {
		fmt.Println(i)
	}

	wg.Wait()

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

func removeBaseDir(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[pos+1:]
}

func pushingToS3(file *os.File, uploader *s3manager.Uploader, bucket string, key string, ch chan<- string) {
	defer wg.Done()
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// // Upload the file to S3.
	result, errS3 := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer),
	})
	if errS3 != nil {
		fmt.Printf("failed to upload file, %v", errS3)
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)

	ch <- "word"

}
