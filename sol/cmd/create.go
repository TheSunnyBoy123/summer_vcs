/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)
const (
	solMainDir = ".sol"
	solMetaDir = ".sol/meta"
	solCheckout = ".sol/checkout"
	solBranches = ".sol/branches"
	excessArgsError = "Too many arguments\n"
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
                fmt.Print(excessArgsError)
        }
		return nil
    },
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
