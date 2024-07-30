/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		root_path := "."
		db_path := objectsPath

		workspace := NewWorkspace(root_path)
		database := NewDatabase(db_path)
		// refs := NewRefs(solPath)
		// fmt.Println("reached line 35")

		listFiles, _ := workspace.ListFiles("")

		// to store entries
		var entries []*Entry

		// fmt.Println("reached line")
		// iterate over files in the root directory
		for _, file := range listFiles {
			// contents
			data := readFile(file)
			// create a blob object
			blob := NewBlob(data)
			// store this blob object
			database.Store(blob)

			// stat is whether file is executable or not
			stat := workspace.StatFile(file)
			// fmt.Println("Stat in main = ", stat)
			// create entry object and add to list
			entry := NewEntry(file, blob.OID, stat)
			entries = append(entries, entry)
		}
		root := NewTree()
		root.Build(entries)
		// root.Traverse(entries)

	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
