package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"github.com/spf13/cobra"
)

func hashDir(dir string) error {
	fmt.Println("Hashing directory: ", dir)
	return nil
}

func hashFile(dir string) error {
	// create blob of file
	// get 
	fmt.Println("Hashing file: ", dir)
	return nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds all modified files to the staging area",
	Long: `Adds all modified files to the staging area. This command is used to stage all the files that have been modified in the working directory.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		// Assuming the current directory is the target, but you can adjust the path as needed
		currentDir := "."

		// Read the directory contents
		entries, err := ioutil.ReadDir(currentDir)
		if err != nil {
			return fmt.Errorf("failed to read directory: %w", err)
		}
	
		for _, entry := range entries {
			// Construct the full path of the entry
			fullPath := filepath.Join(currentDir, entry.Name())
	
			if entry.IsDir() {
				// If the entry is a directory
				err := hashDir(fullPath) // Assuming hashDir is implemented elsewhere
				if err != nil {
					return fmt.Errorf("failed to hash directory '%s': %w", fullPath, err)
				}
			} else {
				// If the entry is a file
				err := hashFile(fullPath) // Assuming hashFile is implemented elsewhere
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
