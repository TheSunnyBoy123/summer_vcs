package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Show commit history",
	Long:  `Displays the commit history starting from HEAD and following parent links.`,
	Run: func(cmd *cobra.Command, args []string) {
		refs := NewRefs(solPath)
		hash := refs.ReadHead()

		if hash == "" {
			fmt.Println("No commits yet")
			return
		}

		for hash != "" {
			objectPath := objectsPath + hash[:2] + "/" + hash[2:]
			if !fileExists(objectPath) {
				break
			}

			contents := decompress(readFile(objectPath))
			raw := strings.SplitN(contents, "\x00", 2)
			if len(raw) < 2 {
				break
			}

			header := raw[0]
			body := raw[1]

			if !strings.HasPrefix(header, "commit") {
				break
			}

			parts := strings.SplitN(body, "\n\n", 2)
			metadata := parts[0]
			message := ""
			if len(parts) > 1 {
				message = strings.TrimSpace(parts[1])
			}

			lines := strings.Split(metadata, "\n")
			treeHash := ""
			parentHash := ""
			author := ""

			for _, line := range lines {
				if strings.HasPrefix(line, "tree ") {
					treeHash = strings.TrimPrefix(line, "tree ")
				} else if strings.HasPrefix(line, "parent ") {
					parentHash = strings.TrimPrefix(line, "parent ")
				} else if strings.HasPrefix(line, "author ") {
					author = strings.TrimPrefix(line, "author ")
				}
			}

			fmt.Println("commit " + hash)
			if author != "" {
				fmt.Println("Author: " + author)
			}
			if treeHash != "" {
				fmt.Println("Tree:   " + treeHash)
			}
			if message != "" {
				for _, l := range strings.Split(message, "\n") {
					fmt.Println("    " + l)
				}
			}
			fmt.Println()

			hash = parentHash
		}
	},
}

// store a commit and read it back
func init() {
	rootCmd.AddCommand(logCmd)
}
