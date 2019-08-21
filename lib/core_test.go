package lib_test

import (
	"testing"

	"github.com/warrensbox/s3upload/lib"
)

// Directory   string
// Bucket      string
// IncludeBase bool
// ConfigFile  string
// Exclude     string
// Session     *session.Session

func TestConfiguration(t *testing.T) {

	var constructor lib.Constructor

	constructor.Directory = "../test"
	constructor.Bucket = "simple.s3.uploader"
	constructor.IncludeBase = false
	constructor.ConfigFile = "../test_config/s3config.json"
	constructor.Exclude = "main,README.md"

	profile := lib.NewConstructor(&constructor)

	if profile.IncludeBase == false {
		t.Logf("Include base: %t [expected]\n", profile.IncludeBase)
	} else {
		t.Errorf("Excluded base: %t [unexpected]\n", profile.IncludeBase)
	}

	if constructor.Directory == "../test" {
		t.Logf("Directory found: %s [expected]\n", constructor.Directory)
	} else {
		t.Errorf("Directory not found: %s [unexpected]\n", constructor.Directory)
	}

	if constructor.Bucket == "simple.s3.uploader" {
		t.Logf("Bucket found:%s [expected]\n", constructor.Bucket)
	} else {
		t.Errorf("Bucket not found: %s [unexpected]\n", constructor.Bucket)
	}

	if constructor.ConfigFile == "../test_config/s3config.json" {
		t.Logf("Configuration file found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Configuration file not found: %s [unexpected]\n", constructor.ConfigFile)
	}

	if constructor.Exclude == "main,README.md,s3upload" {
		t.Logf("Excludes found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Excludes not found: %s [unexpected]\n", constructor.ConfigFile)
	}

	//remove later
	if constructor.Exclude == "main,README.md,s3upload2" {
		t.Logf("Excludes found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Excludes not found: %s [unexpected]\n", constructor.ConfigFile)
	}

	//remove later
	if constructor.Exclude == "main,README.md,s3upload3" {
		t.Logf("Excludes found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Excludes not found: %s [unexpected]\n", constructor.ConfigFile)
	}

	//remove later
	if constructor.Exclude == "main,README.md,s3upload3, s3upload4" {
		t.Logf("Excludes found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Excludes not found: %s [unexpected]\n", constructor.ConfigFile)
	}

	//remove later
	if constructor.Exclude == "main,README.md,s3upload3, s3upload4, s3upload5" {
		t.Logf("Excludes found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Excludes not found: %s [unexpected]\n", constructor.ConfigFile)
	}

	//remove later
	if constructor.Exclude == "main,README.md,s3upload3, s3upload4, s3upload5,s3upload6" {
		t.Logf("Excludes found: %s [expected]\n", constructor.ConfigFile)
	} else {
		t.Errorf("Excludes not found: %s [unexpected]\n", constructor.ConfigFile)
	}
}
