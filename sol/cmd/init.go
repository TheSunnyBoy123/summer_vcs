/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var verbose bool

const (
	solMainDir      = ".sol"
	solCommits      = ".sol/commits"
	solBranches     = ".sol/branches"
	solObjects      = ".sol/objects"
	solRefs         = ".sol/refs"
	excessArgsError = "Too many arguments\n"
)


// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [<directory>]",
	Short: "Initialises sol within the provided directory",
	Long: `Creates .sol directory as well as the metadata files required for sol to function.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if verbose {
			fmt.Println("verbose asked for")
			return nil
		}
		switch len(args) {
		case 0:
			//check no parent dir has .sol dir
			wd, _ := os.Getwd()
			if notInitialisedRepo(wd) {
				initializeDirs([]string{solMainDir, solCommits, solBranches, solObjects})
				createFiles([]string{"./.sol/stagedChanges.txt"})
				fmt.Println("Repository initialised")
			} else {
				fmt.Println("Repository already exists for this directory")
			}
		// case 1:
		// 	if dirExists(args[0]) && notInitialisedRepo(args[0]) {
		// 		initializeDirs([]string{args[0] + "/" + solMainDir, args[0] + "/" + solCommits, args[0] + "/" + solBranches, args[0] + "/" + solObjects})
		// 		createFiles([]string{args[0] + "/.sol/stagedChanges.txt"})
		// 		fmt.Println("Repository initialised")
		// 	} else {
		// 		fmt.Println("Directory does not exist")
		// 	}
		default:
			fmt.Print(excessArgsError)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Print verbose output")
	rootCmd.AddCommand(initCmd)
}
