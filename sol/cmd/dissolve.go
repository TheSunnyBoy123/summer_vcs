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
			} else {
				fmt.Println("Sol was not initialised in the current directory.")
			}
		case 1:
			if dirExists(args[0] + "/.sol") {
				fmt.Println("Dissolving the provided directory")
				if err := deleteDir(args[0] + "/.sol"); err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Directory does not exist")
			}
		default:
			//Print ExcessArgsError + Use string defined above
			fmt.Print(excessArgsError)
			fmt.Print("Usage: " + cmd.Use + "\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(dissolveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dissolveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dissolveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
