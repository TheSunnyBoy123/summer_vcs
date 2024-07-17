/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)



var ignoreCmd = &cobra.Command{
	Use:   "ignore",
	Short: "Adds files and directories to the ignore list",
	Long: `Adds files and directories to the ignore list, stored in .solignore file. 
	When the add argument is used, the .solignore file is checked for all dirs and files to be ignored. 
	When the rm argument is used, the file is removed from the ignore list.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !fileExists(solignorePath) {
			createFiles([]string{solignorePath})
		}
		for _, arg := range args {
			if fileExists(arg){
				writeToFile(solignorePath, arg)
			}
		}
	},
}


var addIgnoreCmd = &cobra.Command{
	Use:  "add",
	Short: "Add files and directories to the .solignore file",
	Long: `Add files and directories to the .solignore file, to be ignored by the add command.
	All passed arguments along with this will be added to the .solignore at one go.`,
	Run: func(cmd *cobra.Command, args []string) {
		if !fileExists(solignorePath) {
			createFiles([]string{solignorePath})
		}
		for _, arg := range args {
			writeToFile(solignorePath, arg)
		}
	},
}

var clearIgnoreCmd = &cobra.Command{
	Use:  "clear",
	Short: "Clear the .solignore file",
	Long: `Clear the .solignore file, removing all the files and directories to be ignored.`,
	Run: func(cmd *cobra.Command, args []string) {
		if fileExists(solignorePath) {
			writeFile(solignorePath, "")
			fmt.Println(".solignore file cleared")
		}
	},
}

var rmCmd = &cobra.Command{
	Use:  "rm",
	Short: "Remove the .solignore file",
	Long: `Remove the .solignore file.`,
	Run: func(cmd *cobra.Command, args []string) {
		if fileExists(solignorePath) {
			err := deleteFile(solignorePath)
			if err != nil {
				fmt.Println("Failed to remove .solignore file")
			} else {
				fmt.Println(".solignore file removed")
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(ignoreCmd)
	ignoreCmd.AddCommand(rmCmd)
	ignoreCmd.AddCommand(addIgnoreCmd)
	ignoreCmd.AddCommand(clearIgnoreCmd)
}
