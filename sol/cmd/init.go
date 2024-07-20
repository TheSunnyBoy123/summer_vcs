/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	solMainDir      = ".sol"
	solCommits      = ".sol/commits"
	solBranches     = ".sol/branches"
	solObjects      = ".sol/objects"
	solRefs         = ".sol/refs"
	solHead         = ".sol/HEAD"
	excessArgsError = "Too many arguments\n"
)

var ignore bool

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [<directory>]",
	Short: "Initialises sol within the provided directory",
	Long:  `Creates .sol directory as well as the metadata files required for sol to function.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// switch len(args) {
		// 	case 0:
		// 		//check no parent dir has .sol dir
		// 		wd, _ := os.Getwd()
		// 		if notInitialisedRepo(wd) {
		// 			initializeDirs([]string{solMainDir, solCommits, solBranches, solObjects})
		// 			createFiles([]string{stagePath, solHead, solRefs})
		// 			fmt.Println("Repository initialised")
		// 		} else {
		// 			fmt.Println("Repository already exists for this directory")
		// 			for _, dir := range []string{solMainDir, solCommits, solBranches, solObjects, solRefs} {
		// 				if dirExists(dir) {
		// 					fmt.Println(dir, " already exists")
		// 				}
		// 			}
		// 		}
		// 		if ignore {
		// 			createFiles([]string{solignorePath})
		// 		}
		// 	default:
		// 		fmt.Print(excessArgsError)
		// 	get working directory
		// }

		root_path := "."
		fullPath := getFullPath(root_path)

		if notInitialisedRepo(root_path) {
			initializeDirs([]string{solPath, refsPath, objectsPath})
			// createFiles([]string{stagePath, solHead, solRefs})
			fmt.Println("Repository initialised in ", fullPath)
		} else {
			fmt.Println("Repository already exists for this directory")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&ignore, "ignore", "i", false, "Initialise with a .solignore file")
}
