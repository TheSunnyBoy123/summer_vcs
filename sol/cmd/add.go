package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
)

func hashDir(dir string) error {
	// recursive function that first goes through all children directories and files, saving each file's hash to a list so that later we can create the object file for each directory
	// structure for a hashDir object file:
	// Tree <length of file>\0
	// ChildType <obj_sha>\0
	// ChildType <obj_sha>\0

}

func hashFile(dir string) error {
	fmt.Println("Hashing file: ", dir)
	return nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds all modified files to the staging area",
	Long: `Adds all modified files to the staging area. This command is used to stage all the files that have been modified in the working directory.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		currentDir := "."

		entries, err := ioutil.ReadDir(currentDir)
		if err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}
		
		for _, entry := range entries {
			fullPath := filepath.Join(currentDir, entry.Name())
	
			if entry.IsDir() {
				if entry.Name() == ".sol" {
					continue
				}

				err := hashDir(fullPath)
				if err != nil {
					return fmt.Errorf("failed to hash directory '%s': %w", fullPath, err)
				}
			} else {
				// If the entry is a file
				err := hashFile(fullPath) 
				if err != nil {
					return fmt.Errorf("failed to hash file '%s': %w", fullPath, err)
				}
			}
		}
	
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
