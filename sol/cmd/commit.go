/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

type Workspace struct {
	Pathname string
}

func NewWorkspace(pathname string) *Workspace {
	return &Workspace{
		Pathname: pathname,
	}
}

func (ws *Workspace) ListFiles() ([]string, error) {
	ignore := map[string]bool{
		".":    true,
		"..":   true,
		".sol": true,
	}

	files, err := ioutil.ReadDir(ws.Pathname)
	if err != nil {
		return nil, err
	}

	var result []string
	for _, file := range files {
		if !ignore[file.Name()] {
			result = append(result, file.Name())
		}
	}

	return result, nil
}

func (ws *Workspace) ReadFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		root_path := "."
		db_path := objectsPath
		// sol_path := solPath

		workspace := NewWorkspace(root_path)
		database := NewDatabase(db_path)
		listFiles, _ := workspace.ListFiles()

		var entries []*Entry
		for _, file := range listFiles {
			data := readFile(file)
			blob := NewBlob(data)
			database.Store(blob)

			entry := NewEntry(file, blob.OID)
			entries = append(entries, entry)
		}

		tree := NewTree(entries)
		database.Store(tree)

		fmt.Println("tree: ", tree.OID)
		return nil
	},
}

func init() {}
func init() {
	rootCmd.AddCommand(commitCmd)
}
