package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Directory string
var targetFiles []string



// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <directory>",
	Short: "Initialises sol within the provided directory",
	Long: "Creates .sol directory as well as the metadata files required for sol to function properly.",
	Args: cobra.MinimumNArgs(1),
	RunE: func(_ *cobra.Command, _ []string) error {
		return runInitCommand()
	},
}


func runInitCommand() error {
	if exists, err := checkPathExists(vxRootDirName); err != nil {
		return err
	} else if exists {
		return errors.New("vx root directory is already exist")
	}

	if err := createDirectories(vxRootDirName, vxCommitDirPath, vxCheckoutDirPath); err != nil {
		return err
	}

	if err := createFile(vxStatusFilePath); err != nil {
		return err
	}

	if err := createFile(vxStagingFilePath); err != nil {
		return err
	}

	color.Green("All files are initialized within .vx/ directory!")

	return nil
}