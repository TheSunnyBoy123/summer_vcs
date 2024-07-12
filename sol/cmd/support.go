package cmd

import (
	"crypto/sha1"
    "fmt"
	"path/filepath"
    "io/ioutil"
    "log"
    "os"
	"compress/zlib"
	"bytes"
	// "github.com/spf13/cobra"
)


// file functions

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func deleteFile(path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}
	return nil
}

func createFiles(fileDir []string) {
	for _, path := range fileDir {
		_, err := os.Create(path)
		if err != nil {
			log.Fatalf("Failed creating file: %s", err)
		}
	}
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

func readFile(dir string) string{
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

func writeToFile(dir string, contents string) {
	file, err := os.OpenFile(dir, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed opening file: %s", err)
	}
	defer file.Close()
	_, err = file.WriteString(contents + "\n")
	if err != nil {
		log.Fatalf("Failed writing to file: %s", err)
	}
}


// dir functions

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


// sha functions 
func hashContents(contents string) string {
	// hash contents using sha1 library
	hash := sha1.New()
	hash.Write([]byte(contents))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func compress(contents string) string {
	// this function compresses the contents using zlib and returns the compressed contents as a string
	var compressedContent bytes.Buffer
	w := zlib.NewWriter(&compressedContent)
	w.Write([]byte(contents))
	w.Close()
	return compressedContent.String()
}

func decompress(contents string) string {
	var decompressedContent bytes.Buffer
	r, err := zlib.NewReader(bytes.NewBufferString(contents))
	if err != nil {
		log.Fatalf("Failed to decompress contents: %s", err)
	}
	defer r.Close()
	decompressedContent.ReadFrom(r)
	return decompressedContent.String()
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