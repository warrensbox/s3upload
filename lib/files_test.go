package lib_test

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/warrensbox/s3upload/lib"
)

func TestVisit(t *testing.T) {

	var files []string

	list := "main,README.md"

	excludes := strings.Split(list, ",")

	set := make(map[string]struct{}, len(excludes))
	for _, s := range excludes {
		set[s] = struct{}{}
	}

	err := filepath.Walk("../test", lib.Visit(&files, set))

	if err != nil {
		t.Errorf("Error walking through path is not a pointer %v [unexpected]\n", err)
	}

	count := 0

	for _, doc := range files {

		if doc == "../test/goat/kid.html" {
			t.Logf("Found : %s\n", doc)
			count++
		}
		if doc == "../test/index.html" {
			t.Logf("Found : %s\n", doc)
			count++
		}
		if doc == "../test/secondary.index" {
			t.Logf("Found : %s\n", doc)
			count++
		}
		if doc == "../test/sparrow/additional/file.txt" {
			t.Logf("Found : %s\n", doc)
			count++
		}
	}

	if count == 4 {
		t.Logf("All files found : %v\n", count)
	} else {
		t.Errorf("Not all files found : %v\n", count)
	}
}

func TestRemoveBaseDir(t *testing.T) {

	IncludeBase := false

	doc := "test/secondary.index"

	if !IncludeBase {
		doc = lib.RemoveBaseDir(doc, "/")
	}

	if doc == "secondary.index" {
		t.Logf("Found %s, [expected]\n", doc)
	} else {
		t.Errorf("No file found : %v, [unexpected]\n", doc)
	}

}

func TestExists(t *testing.T) {

	path := "../test_config/s3config.json"

	exist, err := lib.Exists(path)

	if err != nil {
		t.Errorf("Error found : %v, [unexpected]\n", err)
	}

	if exist == true {
		t.Logf("File : %v, [expected]\n", exist)
	} else {
		t.Errorf("No file found : %v, [unexpected]\n", exist)
	}

}
