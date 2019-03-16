package lib

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

		semverRegex := regexp.MustCompile(`^\.\/\*$|^\.$|^\.\/$`)

		if !id.IncludeBase && !semverRegex.MatchString(id.Directory) {
			doc = RemoveBaseDir(doc, "/")
		}

		if id.AddKey != "" {
			doc = fmt.Sprintf("%s/%s", id.AddKey, doc)
		}

		ct := getContentType(filepath.Ext(doc))

		go pushingToS3(file, uploader, id.Bucket, doc, id.ACL, ct, ch)
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

func getContentType(name string) string {

	switch name {
	case ".ai":
		return "application/postscript"
	case ".png":
		return "image/png"
	case ".jpeg":
		return "image/jpeg"
	case ".jpg":
		return "image/jpeg"
	case ".gif":
		return "image/gif"
	case ".zip":
		return "application/zip"
	case ".gzip":
		return "application/x-gzip"
	case ".tar.gz":
		return "application/x-compressed"
	case ".android":
		return "application/vnd.android.package-archive"
	case ".svg":
		return "image/svg+xml"
	case ".xml":
		return "image/svg+xml"
	case ".txt":
		return "text/plain"
	case ".bmp":
		return "image/bmp"
	case ".pdf":
		return "application/pdf"
	case ".rtf":
		return "text/rtf"
	case ".mov":
		return "video/quicktime"
	default:
		return "binary/octet-stream"
	}

	// 		application/msword
	// application/pdf
	// application/vnd.android.package-archive
	// application/x-compressed
	// application/x-gzip
	// application/zip
	// application/zip
	// audio/mpeg
	// binary/octet-stream
	// image/bmp
	// image/gif
	// image/jpeg
	// image/png
	// image/svg+xml
	// image/tiff
	// text/plain
	// text/rtf
}

func pushingToS3(file *os.File, uploader *s3manager.Uploader, bucket string, key string, acl string, contentType string, ch chan<- string) {
	defer wg.Done()
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	message := "Uploading " + key

	// Upload the file to S3.
	result, errS3 := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(bucket),
		Key:         aws.String(key),
		Body:        bytes.NewReader(buffer),
		ACL:         aws.String(acl),
		ContentType: aws.String(contentType),
	})
	if errS3 != nil {
		message = fmt.Sprintf("Failed to upload file, %v", errS3)
	} else {
		message = fmt.Sprintf("File uploaded to, %s", result.Location)
	}

	ch <- message

}
