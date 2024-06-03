package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

const (
	solMainDir      = ".sol"
	solMetaDir      = ".sol/meta"
	solCheckout     = ".sol/checkout"
	solBranches     = ".sol/branches"
	excessArgsError = "Too many arguments\n"
)

func checkParentRepos(dir string) bool {
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

var createCmd = &cobra.Command{
	Use:   "create [<directory>]",
	Short: "Initialises sol within the provided directory",
	Long:  `Creates .sol directory as well as the metadata files required for sol to function.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			//check no parent dir has .sol dir
			wd, _ := os.Getwd()
			if checkParentRepos(wd) {
				initializeDirs([]string{solMainDir, solMetaDir, solCheckout, solBranches})
			} else {
				fmt.Println("Repository already exists for this directory")
			}

		case 1:
			if dirExists(args[0]) {
				initializeDirs([]string{args[0] + "/" + solMainDir, args[0] + "/" + solMetaDir, args[0] + "/" + solCheckout, args[0] + "/" + solBranches})
			} else {
				fmt.Println("Directory does not exist")
			}
		default:
			fmt.Print(excessArgsError)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
