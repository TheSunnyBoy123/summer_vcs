package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Record changes to the repository",
	Long:  `Creates a new commit containing the current state of the working directory.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		database := NewDatabase(objectsPath)
		refs := NewRefs(solPath)

		var root *Tree

		if fileExists(stagePath) {
			stagedOID := readFile(stagePath)
			if stagedOID != "" {
				root = &Tree{OID: stagedOID}
			}
		}

		if root == nil || root.GetOID() == "" {
			workspace := NewWorkspace(".")
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

			root = BuildTree(entries)
			database.Store(root)
		}

		parent := refs.ReadHead()

		author_name, author_email, err := getAuthorEnv()
		if err != nil {
			os.Exit(1)
		}
		time_now := time.Now()
		time_now_string := time_now.Format(time.RFC3339)
		author := NewAuthor(author_name, author_email, time_now_string)
		message, _ := cmd.Flags().GetString("message")
		message = strings.Trim(message, " ")

		commit := NewCommit(parent, root, author, message)
		database.Store(commit)

		branch := refs.CurrentBranch()
		if branch != "" {
			refs.UpdateBranch(branch, commit.GetOID())
		} else {
			refs.UpdateHead(commit.GetOID())
		}

		fmt.Println("[" + commit.GetOID()[:8] + "] " + message)

		return nil
	},
}

func init() {
	commitCmd.Flags().StringP("message", "m", "", "Commit message")
	rootCmd.AddCommand(commitCmd)
}
