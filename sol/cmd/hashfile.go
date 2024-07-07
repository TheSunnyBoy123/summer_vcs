/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("sha1 of 'test' ", hashContents("test"))
		//only one argument allowed
		if len(args) != 1 {
			fmt.Println("Usage: hash-file <filename>")
			return
		}
		fmt.Println("hashObject called with filename: ", args[0])
		file := args[0]
		if fileExists(file) {
			contents := readFile(file)
			size := len(contents)
			newContents := "Blob " + fmt.Sprint(size) + "\x00" + contents
			// hash newcontents using sha1 library
			hash := hashContents(newContents)
			// write newcontents to .sol/objects/hash[:2]/hash[2:]
			createDir(".sol/objects/"+hash[:2])
			writeFile(".sol/objects/"+hash[:2]+"/"+hash[2:], newContents)
		} else {
			fmt.Println("File does not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)
}
