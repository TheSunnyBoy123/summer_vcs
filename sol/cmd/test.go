package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test command for debugging",
	Run: func(cmd *cobra.Command, args []string) {
		workspace := NewWorkspace(".")
		database := NewDatabase(objectsPath)

		listFiles, _ := workspace.ListFiles("")

		var entries []*Entry

		for _, file := range listFiles {
			data := readFile(file)
			blob := NewBlob(data)
			database.Store(blob)

			stat := workspace.StatFile(file)
			entry := NewEntry(file, blob.OID, stat)
			entries = append(entries, entry)
		}
		root := BuildTree(entries)
		database.Store(root)
		fmt.Println("Root tree OID:", root.GetOID())
		fmt.Println("Entries:", len(entries))
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
