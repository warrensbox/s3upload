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
