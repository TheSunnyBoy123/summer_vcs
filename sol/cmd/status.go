package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show working tree status",
	Long:  `Displays the status of the working directory relative to the current commit.`,
	Run: func(cmd *cobra.Command, args []string) {
		refs := NewRefs(solPath)
		branch := refs.CurrentBranch()
		commitHash := refs.ReadHead()

		if branch != "" {
			fmt.Println("On branch " + branch)
		} else if commitHash != "" {
			fmt.Println("HEAD detached at " + commitHash[:8])
		} else {
			fmt.Println("No commits yet")
		}

		workspace := NewWorkspace(".")
		wsFiles, _ := workspace.ListFiles("")

		tracked := make(map[string]string)
		if commitHash != "" {
			_, body, err := readCommitFromHash(commitHash)
			if err == nil {
				meta := parseCommitBody(body)
				if treeOID := meta["tree"]; treeOID != "" {
					BuildFileMap(treeOID, "", tracked)
				}
			}
		}

		wsSet := make(map[string]bool)
		for _, f := range wsFiles {
			wsSet[f] = true
		}

		var modified, added, deleted, unmodified []string

		for path, trackedOID := range tracked {
			if !wsSet[path] {
				deleted = append(deleted, path)
				continue
			}
			data := readFile(path)
			currentOID := hashContents(fmt.Sprintf("blob %d\x00%s", len(data), data))
			if currentOID == trackedOID {
				unmodified = append(unmodified, path)
			} else {
				modified = append(modified, path)
			}
		}

		for _, f := range wsFiles {
			if _, ok := tracked[f]; !ok {
				added = append(added, f)
			}
		}

		sort.Strings(modified)
		sort.Strings(added)
		sort.Strings(deleted)

		if len(modified) > 0 {
			fmt.Println("\nChanges not staged for commit:")
			for _, f := range modified {
				fmt.Println("  modified: " + f)
			}
		}

		if len(added) > 0 {
			fmt.Println("\nUntracked files:")
			for _, f := range added {
				fmt.Println("  " + f)
			}
		}

		if len(deleted) > 0 {
			fmt.Println("\nDeleted files:")
			for _, f := range deleted {
				fmt.Println("  deleted: " + f)
			}
		}

		if len(modified) == 0 && len(added) == 0 && len(deleted) == 0 && commitHash != "" {
			fmt.Println("\nnothing to commit, working tree clean")
		}
	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
}
