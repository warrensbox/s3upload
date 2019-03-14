package lib

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func Visit(files *[]string, exclude map[string]struct{}) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		if err != nil {
			log.Fatal(err)
		}

		if _, ok := exclude[info.Name()]; ok == true {
			fmt.Printf("SKIPPING FILE : %s\n", info.Name())
			return nil
		}

		if info.IsDir() && info.Name() == ".git" {
			fmt.Println("SKIPPING GIT DIRECTORY")
			return filepath.SkipDir
		}

		if info.IsDir() {
			return nil
		}

		*files = append(*files, path)
		return nil
	}
}

func RemoveBaseDir(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[pos+1:]
}

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
