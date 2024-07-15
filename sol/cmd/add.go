package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
	"strings"
)

var filesToIgnore = []string{".sol", ".sol/.solignore"}
var force bool

//function makes the file content, saves, then return hash
func hashDir(dir string) (string, error) { 
	// recursive function that first goes through all children directories and files, saving each file's hash to a list so that later we can create the object file for each directory
	// structure for a hashDir object file:
	// Tree <length of file>\0
	// Tree <obj_sha> name\0
	// Blob <obj_sha> name\0
	// ...


	// files to ignore are stored in .solignore
	// contents_solignore := readFile(".sol/.solignore")

	entries, _ := ioutil.ReadDir(dir)
	lines := []string{}
	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())

		if contains(filesToIgnore, entry.Name()){ //skip the sol directory
			// fmt.Println("Skipping directory: ", fullPath)
			continue
		}
		
		if entry.IsDir() { //this is a directory
			// need to create a tree object

			// fmt.Println("Hashing child directory: ", fullPath) //debug line
			objHash, _ := hashDir(fullPath) // get this tree obj created and saved + objHash returned
			//add objhash to lines
			//name of dir
			nameDir := entry.Name()
			lines = append(lines, "Tree " + objHash + " " + nameDir +"\x00")
		} else {
			// If the entry is a file
			// fmt.Println("Hashing child file: ", fullPath)
			objHash, _ := hashFile(fullPath) //get object
			fileName := entry.Name()
			lines = append(lines, "Blob " + objHash + " " + fileName + "\x00")
		}
	}
	toAdd := ""
	// go through each item in lines
	for _, item := range lines {
		toAdd += item
	}
	// fmt.Println("To add is: %s", toAdd)
		
	size := fmt.Sprintf("%d", len(toAdd))
	// fmt.Println("Size is: ", size)
	// fmt.Println("To add is: ", toAdd)
	contents := "Tree " + size + "\x00" + toAdd
	// fmt.Println("For directory: ", dir, " contents is: ", contents)

	// fmt.Println("Contents is: ", contents)
	
	contents = compress(contents)
	hash := hashContents(contents)
	
	// fmt.Println("Hash is: %s", hash)

	createDir(".sol/objects/" + hash[:2])
	writeFile(".sol/objects/" + hash[:2] + "/" + hash[2:], contents)
	fmt.Println(dir, " hash: ", hash)

	return hash, nil
}

func hashFile(dir string) (string, error) {
	contents := readFile(dir)
	size := len(contents)

	contents = "Blob " + string(size) + "\x00" + contents
	contents = compress(contents)

	hash := hashContents(contents)


	if dirExists(".sol/objects/" + hash[:2]) {
		writeFile(".sol/objects/" + hash[:2] + "/" + hash[2:], contents)
	} else {
		createDir(".sol/objects/" + hash[:2])
		writeFile(".sol/objects/" + hash[:2] + "/" + hash[2:], contents)
	}

	fmt.Println(dir, " hash: ", hash)
	return hash, nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds all modified files to the staging area",
	Long: `Adds all modified files to the staging area. This command is used to stage all the files that have been modified in the working directory.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		currentDir := "."


		// entries, err := ioutil.ReadDir(currentDir)
		// if err != nil {
		// 	return fmt.Errorf("failed to read directory: %w", err)
		// }
		
		// for _, entry := range entries {
		// 	fullPath := filepath.Join(currentDir, entry.Name())
	
		// 	if entry.IsDir() {
		// 		if entry.Name() == ".sol" {
		// 			continue
		// 		}

		// 		err := hashDir(fullPath)
		// 		if err != nil {
		// 			return fmt.Errorf("failed to hash directory '%s': %w", fullPath, err)
		// 		}
		// 	} else {
		// 		// If the entry is a file
		// 		err := hashFile(fullPath) 
		// 		if err != nil {
		// 			return fmt.Errorf("failed to hash file '%s': %w", fullPath, err)
		// 		}
		// 	}
		// }

		// we have to save this with the header commit:
		//if no args, then we are in the root directory

		filesToIgnore := ""

		if fileExists(solignorePath) {
			filesToIgnore = readFile(solignorePath)
			filesToIgnore := strings.Split(filesToIgnore, "\n")
			fmt.Println("Files to ignore: ", filesToIgnore)
		}
		if filesToIgnore == "" {
			filesToIgnore := []string{".sol", ".sol/.solignore"}
		} else {
			filesToIgnore = append(filesToIgnore, ".sol", ".sol/.solignore")
		}
		
		if len(args) == 0 {
			// fmt.Println("Hashing root directory")
			hash, _ := hashDir(currentDir)
			writeFile(".sol/stagedChanges", "Tree " + hash +  " " + currentDir + "\n")
			return nil
		} else {
			// stagingContents := ""
			// all the args are treated as as directories and items to be added
			for _, arg := range args { //grab each arg
				if fileExists(arg) {
					hash, _ := hashFile(arg)
					writeToFile(".sol/stagedChanges", "Blob " + hash + " " + arg)
				} else if dirExists(arg) {
					hash, _ := hashDir(arg)
					writeToFile(".sol/stagedChanges", "Tree " + hash + " " + arg)
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
	rootCmd.Flags().BoolVarP(&force, "force", "f", false, "force add otherwise ignored files")
}
