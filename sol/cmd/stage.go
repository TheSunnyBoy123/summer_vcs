/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var view bool

var stageCmd = &cobra.Command{
	Use:   "stage [files] [<directories>]",
	Short: "Adds files to the staging area",
	Long:  `The stage command adds files to the staging area. The staging area is a temporary storage area where you can add files before committing them to the repository. This allows you to review the changes before committing them.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		validFiles := []string{}
		for _, arg := range args {
			if _, err := os.Stat(arg); err == nil {
				if notInitialisedRepo(arg) { //only allowing staging from working directory and repo should be initialised
					fmt.Println("Sol was not initialised in the current directory.")
					return
				} else {
					validFiles = append(validFiles, arg)
				}
			} else {
				fmt.Printf("%s not added to staged changes as it was not found\n", arg)
			}
		}
		if len(validFiles) > 0 {
			for _, file := range validFiles {
				//write filename, current date and time to ./sol/stagedChanges.txt

				fmt.Printf("Added %v to staged changes\n", file)
			}
		} else {
			fmt.Println("No files added to staged changes")
		}
		if view {
			fmt.Println("Files in the staging area")
		}

	},
}

func init() {
	stageCmd.Flags().BoolVarP(&view, "view", "v", false, "View the files in the staging area")
	rootCmd.AddCommand(stageCmd)
}
