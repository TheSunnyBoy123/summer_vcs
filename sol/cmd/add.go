package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
)

func hashChild(dir string) error { 
	return nil
}

//function makes the file content, saves, then return hash
func hashDir(dir string) (string, error) { 
	// recursive function that first goes through all children directories and files, saving each file's hash to a list so that later we can create the object file for each directory
	// structure for a hashDir object file:
	// Tree <length of file>\0
	// ChildType <obj_sha>\0
	// ChildType <obj_sha>\0
	// ...

	entries, _ := ioutil.ReadDir(dir)
	for _, entry := range entries {
		fullPath := filepath.Join(dir, entry.Name())
		lines := []string{}
		contents := "Tree "
		if entry.IsDir() { //this is a directory
			// need to create a tree object
			if entry.Name() == ".sol" {
				continue
			}
			fmt.Println("Hashing child directory: ", fullPath)
			objHash, _ := hashDir(fullPath)
			//add objhash to lines
			lines = append(lines, "Tree " + objHash + "\x00")
		} else {
			// If the entry is a file
			fmt.Println("Hashing child file: ", fullPath)
			objHash := hashFile(fullPath)
			lines = append(lines, "Blob " + objHash + "\x00")
		}
		fmt.Println(lines)
		toAdd := ""
		// go through each item in lines
		for _, item := range lines {
			toAdd += item
		}
		fmt.Println("To add is: %s", toAdd)
		size := len(toAdd)
		contents += string(size) + "\x00" + toAdd
		fmt.Println("Contents is: %s", contents)
		hash := hashContents(contents)
		fmt.Println("Hash is: %s", hash)
		return hash, nil

	}
}

func hashFile(dir string) (string, error) {
	// fmt.Println("Hashing file: ", dir)
	
	return nil
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
		hashDir(currentDir)
	
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
