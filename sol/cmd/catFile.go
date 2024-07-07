/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// catFileCmd represents the catFile command
var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//only one argument allowed
		if len(args) != 1 {
			fmt.Println("Usage: cat-file <hash>")
			return
		}
		fmt.Println("catFile called with hash: ", args[0])
		hash := args[0]
		if fileExists(".sol/objects/" + hash[:2] + "/" + hash[2:]) {
			contents := readFile(".sol/objects/" + hash[:2] + "/" + hash[2:])
			
			fmt.Println(contents)
		} else {
			fmt.Println("File does not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)

}
