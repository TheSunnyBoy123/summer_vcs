package cmd

import (
	"fmt"
	"os"
)

func initializeDirs(dirs []string) {
	for _, dir := range dirs {
		if !dirExists(dir) {
			if err := createDir(dir); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func dirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDir(dir string) error {
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}
	return nil
}

func deleteDir(dir string) error {
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	return nil
}