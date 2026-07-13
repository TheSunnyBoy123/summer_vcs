package cmd

import (
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

func buildFileMap(treeOID, prefix string, fileMap map[string]string) {
	entries, err := parseTreeEntries(treeOID)
	if err != nil {
		return
	}
	for _, entry := range entries {
		fields := strings.SplitN(entry, " ", 3)
		if len(fields) < 3 {
			continue
		}
		mode := fields[0]
		oid := fields[1]
		name := fields[2]
		path := name
		if prefix != "" {
			path = prefix + "/" + name
		}
		if mode == "40000" {
			buildFileMap(oid, path, fileMap)
		} else {
			fileMap[path] = oid
		}
	}
}

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
					buildFileMap(treeOID, "", tracked)
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
