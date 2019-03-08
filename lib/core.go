package lib

import (
	"log"

	session "github.com/aws/aws-sdk-go/aws/session"
)

//Constructor : struct for session
type Constructor struct {
	Session   *session.Session
	Directory string
	MyBucket  string
	MyKey     string
}

//NewConstructor :validate object
func NewConstructor(attr *Constructor) (*Constructor, error) {

	if attr.Session == nil {
		log.Panic("Unable to obtain AWS credentials")
	}

	if attr.Directory == "" {
		attr.Directory = "./test/"
	}

	if attr.MyBucket == "" {
		attr.MyBucket = "simple.s3.uploader"
	}

	if attr.MyKey == "" {
		attr.MyKey = "/"
	}

	return attr, nil
}
