//go:build exclude

package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var force bool
var filesToIgnore = []string{solPath, solignorePath}

func hashDir(dir string) (string, error) {
	// file/dir name stored in name
	name := filepath.Base(dir)
	if contains(filesToIgnore, name) || contains(filesToIgnore, dir) {
		fmt.Println("Ignoring file: ", name)
		return "", fmt.Errorf("This file to be ignored")
	}

	entries, _ := ioutil.ReadDir(dir) // the subdirectories and files in the directory
	lines := []string{}               // array to store the lines of the tree object to be made and saved
	for _, entry := range entries {   // grab each entry in the directory
		fullPath := filepath.Join(dir, entry.Name()) // get the full path of the entry
		baseName := entry.Name()                     // get the name of the entry
		// fmt.Println("Base name: ", baseName)
		// fmt.Println("Full path: ", fullPath)
		// fmt.Println(contains(filesToIgnore, baseName))
		// fmt.Println(contains(filesToIgnore, fullPath))

		// for _, i := range filesToIgnore {
		// 	fmt.Println(i)
		// }

		if contains(filesToIgnore, baseName) || contains(filesToIgnore, fullPath) {
			fmt.Println("Ignoring", baseName)
			continue
		}

		if entry.IsDir() { // if the entry is a directory
			objHash, err := hashDir(fullPath)
			if err == nil {
				nameDir := entry.Name()
				lines = append(lines, "Tree "+objHash+" "+nameDir+"\x00")
			} else {
				fmt.Println("Ignoring tree: ", fullPath)
			}
		} else {
			// this is a file
			objHash, err := hashFile(fullPath)
			if err == nil {
				fileName := entry.Name()
				lines = append(lines, "Blob "+objHash+" "+fileName+"\x00")
				// fmt.Println(entry.Name() + " hash: " + objHash)
			} else {
				fmt.Println("Ignoring file: ", fullPath)
			}
		}
	}
	toAdd := ""

	for _, item := range lines {
		toAdd += item
	}

	size := fmt.Sprintf("%d", len(toAdd))
	contents := "Tree " + size + "\x00" + toAdd
	contents = compress(contents)
	hash := hashContents(contents)

	if dirExists(objectsPath + hash[:2]) {
		writeFile(objectsPath+hash[:2]+"/"+hash[2:], contents)
	} else {
		createDir(objectsPath + hash[:2])
		writeFile(objectsPath+hash[:2]+"/"+hash[2:], contents)
	}

	fmt.Println(dir, " hash: ", hash)
	return hash, nil
}

func hashFile(dir string) (string, error) {

	fileName := filepath.Base(dir)
	if contains(filesToIgnore, fileName) || contains(filesToIgnore, dir) {
		fmt.Println("Ignoring file: ", fileName)
		return "", fmt.Errorf("this file to be ignored")
	}

	// if this file has been hashed before, get the contents from original file, make a diff and hash the diff in place of contents
	// if the file has not been hashed before, hash the file and save the contents in the objects directory

	contents := readFile(dir)
	size := len(contents)
	// fmt.Println("size when hashing file: ", size)
	sizeText := fmt.Sprintf("%d", size)
	contents = "Blob " + sizeText + "\x00" + contents
	contents = compress(contents)

	hash := hashContents(contents)

	if dirExists(objectsPath + hash[:2]) {
		writeFile(objectsPath+hash[:2]+"/"+hash[2:], contents)
	} else {
		createDir(objectsPath + hash[:2])
		writeFile(objectsPath+hash[:2]+"/"+hash[2:], contents)
	}

	fmt.Println(dir, " hash: ", hash)
	return hash, nil

}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds all modified files to the staging area",
	Long:  `Adds all modified files to the staging area. This command is used to stage all the files that have been modified in the working directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		currentDir := "."

		// add the files to staging area, with the hash of the file
		// if the file is a directory, hash the directory and add the hash to the staging area
		// check refs if the file object has been created already. if it has, make a diff of the contents and hash the diff

		// refs := readFile(refsPath)
		// if refs != "" {
		// 	// split refs
		// 	refsList := strings.Split(refs, "\n")
		// 	refMap := make(map[string]string)
		// 	for _, ref := range refsList {
		// 		// ref contains "<name> <hash>"
		// 		// create a map
		// 		words := strings.Split(ref, " ")
		// 		refMap[words[0]] = words[1]
		// 	}
		// }

		if fileExists(solignorePath) {
			contents := readFile(solignorePath)
			filesToIgnore = append(filesToIgnore, strings.Split(contents, "\n")...)
		}

		if force {
			filesToIgnore = []string{solPath}
		}
		fmt.Println("Files to ignore: ", filesToIgnore)

		if len(args) == 0 {
			// fmt.Println("Hashing root directory")
			hash, _ := hashDir(currentDir)
			writeFile(stagePath, "Tree "+hash+" "+currentDir+"\n")
			return nil
		} else {
			// stagingContents := ""
			// all the args are treated as as directories and items to be added
			for _, arg := range args { //grab each arg
				if fileExists(arg) {
					hash, _ := hashFile(arg)
					writeToFile(stagePath, "Blob "+hash+" "+arg)
				} else if dirExists(arg) {
					hash, _ := hashDir(arg)
					writeToFile(stagePath, "Tree "+hash+" "+arg)
				} else {
					fmt.Println("File or directory: ", arg, " does not exist")
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&force, "force", "f", false, "force add otherwise ignored files")
}
