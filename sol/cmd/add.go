package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [files...]",
	Short: "Hash files and store them as objects",
	Long:  `Computes the SHA-1 hash of each file, stores it as a blob object, and records it in the staging area.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		workspace := NewWorkspace(".")
		database := NewDatabase(objectsPath)

		if len(args) == 0 {
			listFiles, err := workspace.ListFiles("")
			if err != nil {
				return err
			}
			args = listFiles
		}

		var entries []*Entry

		for _, file := range args {
			if !fileExists(file) {
				fmt.Printf("fatal: '%s' does not exist\n", file)
				return nil
			}

			data := readFile(file)
			blob := NewBlob(data)
			database.Store(blob)

			stat := workspace.StatFile(file)
			entry := NewEntry(file, blob.OID, stat)
			entries = append(entries, entry)

			fmt.Printf("add '%s' [%s]\n", file, blob.OID[:8])
		}

		if len(entries) > 0 {
			root := BuildTree(entries)
			database.Store(root)
			writeFile(stagePath, root.GetOID())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
