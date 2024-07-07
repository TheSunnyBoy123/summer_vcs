package cmd

import (
	"crypto/sha1"
    "fmt"
	"path/filepath"
    "io/ioutil"
    "log"
    "os"
)


func fileExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func notInitialisedRepo(dir string) bool {
	for {
		dir = filepath.Dir(dir)
		if _, err := os.Stat(filepath.Join(dir, ".sol")); err == nil {
			// .sol directory exists in this directory
			return false
		}
		if dir == "/" || dir == "." {
			//reached root directory
			break
		}
	}
	return true
}

func createFiles(fileDir []string) {
	for _, path := range fileDir {
		_, err := os.Create(path)
		if err != nil {
			log.Fatalf("Failed creating file: %s", err)
		}
	}
}

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



func writeFile(dir string, contents string) {
	file, err := os.Create(dir)
	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}
	defer file.Close()
	_, err = file.WriteString(contents)
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}
}

func readFile (dir string) string{
	file, err := os.Open(dir)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed reading file: %s", err)
	}
	return string(contents)
}

func hashContents(contents string) string {
	// hash contents using sha1 library
	hash := sha1.New()
	hash.Write([]byte(contents))
	return fmt.Sprintf("%x", hash.Sum(nil))
}



