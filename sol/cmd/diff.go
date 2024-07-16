/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sergi/go-diff/diffmatchpatch"
)

var stat bool 
var minimal bool

func diff(obj1, obj2 string) {
	obj1Content := decompress(readFile(".sol/objects/" + obj1[:2] + "/" + obj1[2:]))
	obj2Content := decompress(readFile(".sol/objects/" + obj2[:2] + "/" + obj2[2:]))

	
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(obj1Content, obj2Content, false)

	fmt.Println(dmp.DiffPrettyText(diffs))

	// if stat {
	// 	fmt.Println(diff.Stats())
	// } else {
	// 	fmt.Println(diff)
	// }
}

// diffCmd represents the diff command
var diffCmd = &cobra.Command{
	Use:   "diff",
	Short: "Shows the difference between provided two SHA objects",
	Long: `Shows the difference between provided two SHA objects using the diff algorithm.
	sol diff [<options>] <blob> <blob> to show the difference between two blobs in terms of lines
	sol diff [<options>] <commit> <commit> to show the difference between two commits in terms of objects`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		obj1 := args[0]
		obj2 := args[1]
		if fileExists(".sol/objects/" + obj1[:2] + "/" + obj1[2:]) && fileExists(".sol/objects/" + obj2[:2] + "/" + obj2[2:]) {
			diff(obj1, obj2)
		} else {
			fmt.Println("Object does not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(diffCmd)
	diffCmd.Flags().BoolVarP(&stat, "stat", "s", false, "Show the stats of the difference")
	diffCmd.Flags().BoolVarP(&minimal, "minimal", "m", false, "Show the minimal difference")
}
