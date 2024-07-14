/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// dissolveCmd represents the dissolve command
var dissolveCmd = &cobra.Command{
	Use:   "dissolve [<directory>]",
	Short: "Dissolve the entire repository",
	Long: `Dissolve the entire repository. 
	This will delete all the files and directories created by sol. 
	This action is irreversible and will delete all the files and directories created by sol.`,
	Run: func(cmd *cobra.Command, args []string) {
		switch len(args) {
		case 0:
			if dirExists("./.sol") {
				if err := deleteDir("./.sol"); err != nil {
					fmt.Println(err)
				}
				if fileExists(".solignore") {
					if err := deleteFile(".solignore"); err != nil {
						fmt.Println(err)
					}
				}
				fmt.Println("Repository dissolved")
			} else {
				fmt.Println("Sol was not initialised in the current directory.")
			}
		// case 1:
		// 	if dirExists(args[0] + "/.sol") {
		// 		fmt.Println("Dissolving the provided directory")
		// 		if err := deleteDir(args[0] + "/.sol"); err != nil {
		// 			fmt.Println(err)
		// 		}
		// 		fmt.Println("Repository dissolved")
		// 	} else {
		// 		fmt.Println("Directory does not exist")
		// 	}
		default:
			fmt.Print(excessArgsError)
			fmt.Print("Usage: " + cmd.Use + "\n")
		}
	},
}

func init() {

	rootCmd.AddCommand(dissolveCmd)

}
