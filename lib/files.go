package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Visit : walks thru directory path, ignores exclude files and .git directory
func Visit(files *[]string, exclude map[string]struct{}) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		if _, ok := exclude[info.Name()]; ok == true {
			fmt.Printf("Skipping : %s\n", info.Name())
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			fmt.Println("Skipping .git files")
			return filepath.SkipDir
		}

		if info.IsDir() {
			fmt.Println("Is dir")
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

// RemoveBaseDir : removes base directory
func RemoveBaseDir(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[pos+1:]
}

// Exists : checks if path exist
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
