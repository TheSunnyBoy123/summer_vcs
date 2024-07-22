package cmd

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	// "github.com/spf13/cobra"
)

const (
	solignorePath = ".solignore"
	solPath       = ".sol"
	objectsPath   = ".sol/objects/"
	stagePath     = ".sol/stagedChanges"
	refsPath      = ".sol/refs/"
	indexPath     = ".sol/index"
	headPath      = ".sol/HEAD"
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

func readFile(dir string) string {
	file, err := os.Open(dir)
	if err != nil {
		fmt.Printf("Failed opening file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()
	contents, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("Failed reading file: %s\n", err)
		os.Exit(1)
	}
	return string(contents)
}

func writeToFile(dir string, contents string) {
	file, err := os.OpenFile(dir, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed opening file: %s", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(contents + "\n")
	if err != nil {
		fmt.Printf("Failed writing to file: %s", err)
		os.Exit(1)
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

// misc functions

func notInitialisedRepo(dir string) bool {
	// check if current directory has .sol directory
	return !dirExists(dir + "/" + solPath)
}

func contains(list []string, element string) bool {
	for _, i := range list {
		if i == element {
			return true
		}
	}
	return false
}

func contentsObject(hash string) string {
	contents := decompress(readFile(objectsPath + hash[:2] + "/" + hash[2:]))
	lines := bytes.Split([]byte(contents), []byte("\x00"))
	return string(lines[0])
}

func getFullPath(dir string) string {
	wd, _ := os.Getwd()
	return wd + "/"
}
func getAuthorEnv() (string, string, error) {
	author_name := ""
	author_email := ""

	// Get the home directory
	homeDir, _ := os.UserHomeDir()

	configPath := homeDir + "/.solconfig"
	contents := readFile(configPath)
	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		if strings.Contains(line, "SOL_AUTHOR_NAME") {
			author_name = strings.Split(line, "=")[1]
			author_name = strings.Trim(author_name, " ")
		}
		if strings.Contains(line, "SOL_AUTHOR_EMAIL") {
			author_email = strings.Split(line, "=")[1]
			author_email = strings.Trim(author_email, " ")
		}
	}

	if author_name == "" || author_email == "" {
		return "", "", fmt.Errorf("Author name or email not found")
	} else {
		return author_name, author_email, nil
	}
}
