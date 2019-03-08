package lib

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (id *Constructor) PushToS3() error {

	uploader := s3manager.NewUploader(id.Session)

	//svc := s3.New(id.Session)

	// Open the file for use
	// file, err := os.Open(id.Directory)
	// if err != nil {
	// 	return err
	// }
	// defer file.Close()

	var files []string

	err := filepath.Walk(id.Directory, visit(&files))
	if err != nil {
		panic(err)
	}

	for _, fil := range files {
		fmt.Println("fil")
		fmt.Println(fil)
		file, err := os.Open(fil)
		if err != nil {
			return err
		}
		//defer file.Close()

		//Get file size and read the file content into a buffer
		fileInfo, _ := file.Stat()
		var size int64 = fileInfo.Size()
		buffer := make([]byte, size)
		file.Read(buffer)

		// // Upload the file to S3.
		result, errS3 := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(id.MyBucket),
			Key:    aws.String(fil),
			Body:   bytes.NewReader(buffer),
		})
		if errS3 != nil {
			fmt.Printf("failed to upload file, %v", errS3)
		}
		fmt.Printf("file uploaded to, %s\n", result.Location)

		// _, err = svc.PutObject(&s3.PutObjectInput{
		// 	Bucket:             aws.String(id.MyBucket),
		// 	Key:                aws.String(id.MyKey),
		// 	ACL:                aws.String("private"),
		// 	Body:               bytes.NewReader(buffer),
		// 	ContentLength:      aws.Int64(size),
		// 	ContentType:        aws.String(http.DetectContentType(buffer)),
		// 	ContentDisposition: aws.String("attachment"),
		// 	// ServerSideEncryption: aws.String("AES256"),
		// })

		// fmt.Printf("ERROR, %s\n", err)

	}

	// Get file size and read the file content into a buffer
	// fileInfo, _ := file.Stat()
	// var size int64 = fileInfo.Size()
	// buffer := make([]byte, size)
	// file.Read(buffer)

	// // Upload the file to S3.
	// result, errS3 := uploader.Upload(&s3manager.UploadInput{
	// 	Bucket: aws.String(id.MyBucket),
	// 	Key:    aws.String(id.MyKey),
	// 	Body:   file,
	// })
	// if errS3 != nil {
	// 	fmt.Printf("failed to upload file, %v", errS3)
	// }
	// fmt.Printf("file uploaded to, %s\n", result.Location)

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
