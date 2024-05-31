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
	solMainDir = ".sol"
	solMetaDir = ".sol/meta"
	solCheckout = ".sol/checkout"
	solBranches = ".sol/branches"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
    Use:   "create [<directory>]",
    Short: "Initialises sol within the provided directory",
    Long: `Creates .sol directory as well as the metadata files required for sol to function.`,
    RunE: func(cmd *cobra.Command, args []string) error {
        switch len(args) {
            case 0:
                initializeDirs([]string{solMainDir, solMetaDir, solCheckout, solBranches})
            case 1:
                if dirExists(args[0]) {
                    initializeDirs([]string{args[0] + "/" + solMainDir, args[0] + "/" + solMetaDir, args[0] + "/" + solCheckout, args[0] + "/" + solBranches})
                } else {
                    fmt.Println("Directory does not exist")
                }	
            default:
                fmt.Println("Too many arguments")
        }
		return nil
    },
}

func initializeDirs(dirs []string) {
	for _, dir := range dirs {
		if !dirExists(dir) {
			if err := createDir(dir); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func dirExists(dir string) bool {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return false
	}
	return true
}

func createDir(dir string) error {
	if err := os.Mkdir(dir, 0755); err != nil {
		return err
	}
	return nil
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
