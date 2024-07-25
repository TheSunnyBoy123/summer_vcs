/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

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
		path := " " //placeholder

		// sol_path := solPath
		// fmt.Println("reached line 31")
		workspace := NewWorkspace(root_path)
		database := NewDatabase(db_path)
		refs := NewRefs(solPath)
		// fmt.Println("reached line 35")

		listFiles, _ := workspace.ListFiles("")

		var entries []*Entry
		for _, file := range listFiles {
			data := readFile(file)
			blob := NewBlob(data)
			database.Store(blob)

			stat := workspace.StatFile(path)
			entry := NewEntry(file, blob.OID, stat)
			entries = append(entries, entry)
		}
		fmt.Println("reached line 50")
		tree := NewTree(entries)
		database.Store(tree)

		//parent
		fmt.Println("ref pathname: ", refs.pathname)
		parent := refs.ReadHead()
		fmt.Println("Parent: ", parent)
		fmt.Println("reached line 57")

		author_name, author_email, err := getAuthorEnv()
		if err != nil {
			os.Exit(1)
		}
		time_now := time.Now()
		time_now_string := time_now.Format(time.RFC3339)
		author := NewAuthor(author_name, author_email, time_now_string)
		// fmt.Println("Author: ", author.ToString())
		message, _ := cmd.Flags().GetString("message")
		message = strings.Trim(message, " ")
		// fmt.Println("Message: ", message)

		commit := NewCommit(parent, tree, author, message)
		commit.SetOID("")
		database.Store(commit)

		// fmt.Println("Commit: ", commit.GetOID())
		writeFile(headPath, commit.GetOID())

		// fmt.Println("[(commit)] " + commit.GetOID() + " " + message)
		fmt.Println("[(commit)] " + commit.GetOID() + "\n" + message)

		// fmt.Println("Tree: ", tree.GetOID())
		return nil
	},
}

func init() {
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	rootCmd.AddCommand(commitCmd)
}
