/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// stageCmd represents the stage command
var stageCmd = &cobra.Command{
	Use:   "stage [files] [<directories>]",
	Short: "Adds files to the staging area",
	Long: `The stage command adds files to the staging area. The staging area is a temporary storage area where you can add files before committing them to the repository. This allows you to review the changes before committing them.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stage called")
	},
}

func init() {
	rootCmd.AddCommand(stageCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stageCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// stageCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
