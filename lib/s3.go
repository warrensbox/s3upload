package lib

import (
	"bytes"
	"fmt"
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

	excludes := strings.Split(id.Exclude, ",")

	set := make(map[string]struct{}, len(excludes))
	for _, s := range excludes {
		set[s] = struct{}{}
	}

	err := filepath.Walk(id.Directory, visit(&files, set))
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

		fmt.Printf("Files %s\n", fil)

		if !id.IncludeBase {
			fil = removeBaseDir(fil, "/")
		}

		go pushingToS3(file, uploader, id.Bucket, fil, ch)

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

func pushingToS3(file *os.File, uploader *s3manager.Uploader, bucket string, key string, ch chan<- string) {
	defer wg.Done()
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	//os.Exit(0)
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
