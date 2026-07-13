package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func diffTrees(parentFiles, currentFiles map[string]string) (added, modified, deleted []string) {
	for path := range currentFiles {
		if _, ok := parentFiles[path]; !ok {
			added = append(added, path)
		} else if parentFiles[path] != currentFiles[path] {
			modified = append(modified, path)
		}
	}
	for path := range parentFiles {
		if _, ok := currentFiles[path]; !ok {
			deleted = append(deleted, path)
		}
	}
	sort.Strings(added)
	sort.Strings(modified)
	sort.Strings(deleted)
	return
}

var showCmd = &cobra.Command{
	Use:   "show [object]",
	Short: "Show commit details with file changes",
	Long:  `Display the commit metadata and list of files changed relative to its parent.`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		refs := NewRefs(solPath)
		hash := refs.ReadHead()
		if len(args) > 0 {
			hash = resolveRevision(args[0])
			if hash == "" {
				hash = args[0]
			}
		}
		if hash == "" {
			fmt.Println("No commits yet")
			return
		}

		header, body, err := readCommitFromHash(hash)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !strings.HasPrefix(header, "commit") {
			fmt.Println("object is not a commit")
			return
		}

		meta := parseCommitBody(body)
		treeOID := meta["tree"]
		parentOID := meta["parent"]
		author := meta["author"]
		message := meta["message"]

		fmt.Println("commit " + hash)
		fmt.Println("Author: " + author)
		fmt.Println("Tree:   " + treeOID)
		if parentOID != "" {
			fmt.Println("Parent: " + parentOID)
		}
		fmt.Println()
		for _, l := range strings.Split(message, "\n") {
			fmt.Println("    " + l)
		}

		currentFiles := TreeFileMap(treeOID)

		if parentOID == "" {
			fmt.Println(fmt.Sprintf("\n%d files changed:", len(currentFiles)))
			var names []string
			for p := range currentFiles {
				names = append(names, p)
			}
			sort.Strings(names)
			for _, p := range names {
				fmt.Println("  create mode " + currentFiles[p][:8] + " " + p)
			}
		} else {
			parentMeta := parseCommitBody(func() string {
				_, b, e := readCommitFromHash(parentOID)
				if e != nil {
					return ""
				}
				return b
			}())
			parentFiles := TreeFileMap(parentMeta["tree"])

			added, modified, deleted := diffTrees(parentFiles, currentFiles)

			total := len(added) + len(modified) + len(deleted)
			fmt.Println(fmt.Sprintf("\n%d files changed:", total))

			for _, p := range added {
				fmt.Println("  create mode " + currentFiles[p][:8] + " " + p)
			}
			for _, p := range modified {
				fmt.Println("  modified       " + p)
			}
			for _, p := range deleted {
				fmt.Println("  deleted        " + p)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
