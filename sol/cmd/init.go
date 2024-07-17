/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

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

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init [<directory>]",
	Short: "Initialises sol within the provided directory",
	Long:  `Creates .sol directory as well as the metadata files required for sol to function.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			//check no parent dir has .sol dir
			wd, _ := os.Getwd()
			if notInitialisedRepo(wd) {
				initializeDirs([]string{solMainDir, solCommits, solBranches, solObjects})
				createFiles([]string{stagePath, solHead})
				fmt.Println("Repository initialised")
			} else {
				fmt.Println("Repository already exists for this directory")
			}
		default:
			fmt.Print(excessArgsError)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
