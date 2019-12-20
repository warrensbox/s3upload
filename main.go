package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	session "github.com/aws/aws-sdk-go/aws/session"
	lib "github.com/warrensbox/s3upload/lib"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

/*
* Version 0.1.0
* Compatible with Mac OS X and linux OS ONLY
 */

/*** OPERATION WORKFLOW ***/
/*
* 1- SSM library gets AWS credentials from host machine
* 2- Parses command line arguments
* 3- Checks if default s3config file exist
* 4- Establishes S3 connection
* 5- Push files to S3
#redude number of args
*/

//TODO fix aws region
//upload to multiple buckets
//remove command line because you can use aws cli to do that
//benefit of this is you can upload to multiple s3 buckets
// move to multi upload
var version = "0.1.2\n"

var (
	versionFlag  *bool
	helpFlag     *bool
	includeBase  *bool
	awsRegion    *string
	directory    *string
	bucket       *string
	addkey       *string
	acl          *string
	configFile   *string
	excludeFiles *string
)

func init() {

	const (
		versionFlagDesc = "Displays the version of S3Pusher"
		skipBaseDesc    = "Skip base directory"
		awsRegionDesc   = "Provide AWS Region. Default - us-east-1"
		directoryDesc   = "Directory where files are stored. Default - current directory"
		bucketDesc      = "S3 bucket. Default bucket is in the config file"
		keyDesc         = "Append key to s3 bucket. For example: key/my.files"
		aclDesc         = "S3 ACL information"
		confDesc        = "S3 config information"
	)

	bucket = kingpin.Flag("bucket", bucketDesc).Short('b').String()
	addkey = kingpin.Flag("addkey", keyDesc).Short('k').String()
	acl = kingpin.Flag("acl", keyDesc).Short('a').String()
	configFile = kingpin.Flag("config", confDesc).Short('c').String()
	excludeFiles = kingpin.Flag("exclude", confDesc).Short('e').String()

	versionFlag = kingpin.Flag("version", versionFlagDesc).Short('v').Bool()
	includeBase = kingpin.Flag("ignoreBase", skipBaseDesc).Short('i').Bool()
	awsRegion = kingpin.Flag("region", awsRegionDesc).Short('r').String()
	directory = kingpin.Flag("directory", directoryDesc).Short('d').String()

}

//need to use runner instead
//use class type
//switch kingpin
func main() {

	kingpin.CommandLine.Interspersed(false)
	kingpin.Parse()

	if *versionFlag {
		fmt.Println(version)
	}

	config := &aws.Config{Region: aws.String(*awsRegion)}

	session := session.Must(session.NewSession(config))

	construct := &lib.Constructor{*directory, *bucket, *addkey, *includeBase, *configFile, *excludeFiles, *acl, session}
	profile := lib.NewConstructor(construct)


	err := profile.PushToS3()

	if err != nil {
		fmt.Println(err)
		//move to lib
	}

}
