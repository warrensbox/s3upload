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

// PushToS3 :  Push files to s3
func (id *Constructor) PushToS3() error {

	uploader := s3manager.NewUploader(id.Session)

	var files []string

	excludes := strings.Split(id.Exclude, ",")

	set := make(map[string]struct{}, len(excludes))
	for _, s := range excludes {
		set[s] = struct{}{}
	}

	err := filepath.Walk(id.Directory, Visit(&files, set))
	if err != nil {
		panic(err)
	}
	ch := make(chan string)

	for _, doc := range files {
		wg.Add(1)
		file, err := os.Open(doc)
		if err != nil {
			return err
		}
		if !id.IncludeBase {
			doc = RemoveBaseDir(doc, "/")
		}

		if id.AddKey != "" {

			doc = fmt.Sprintf("%s/%s", id.AddKey, doc)
		}

		go pushingToS3(file, uploader, id.Bucket, doc, ch)
	}

	go func(ch chan<- string) {
		defer close(ch)
		wg.Wait()
	}(ch)

	for info := range ch {
		fmt.Println(info)
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

	message := "Uploading " + key

	// Upload the file to S3.
	result, errS3 := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   bytes.NewReader(buffer),
	})
	if errS3 != nil {
		message = fmt.Sprintf("Failed to upload file, %v", errS3)
	} else {
		message = fmt.Sprintf("File uploaded to, %s", result.Location)
	}

	ch <- message

}
