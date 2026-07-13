package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var deleteBranch bool

var branchCmd = &cobra.Command{
	Use:   "branch [name]",
	Short: "List, create, or delete branches",
	Long: `With no arguments, list existing branches.
With a name, create a new branch at the current commit.
With -d, delete the named branch.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		refs := NewRefs(solPath)

		if deleteBranch {
			if len(args) != 1 {
				return fmt.Errorf("usage: sol branch -d <name>")
			}
			name := args[0]
			if refs.ReadBranch(name) == "" {
				return fmt.Errorf("branch '%s' not found", name)
			}
			current := refs.CurrentBranch()
			if name == current {
				return fmt.Errorf("cannot delete branch '%s' checked out", name)
			}
			refs.DeleteBranch(name)
			fmt.Println("Deleted branch " + name)
			return nil
		}

		if len(args) == 0 {
			branches := refs.ListBranches()
			current := refs.CurrentBranch()
			sort.Strings(branches)
			for _, b := range branches {
				prefix := "  "
				if b == current {
					prefix = "* "
				}
				commit := refs.ReadBranch(b)
				short := ""
				if len(commit) >= 8 {
					short = commit[:8]
				}
				fmt.Printf("%s%s %s\n", prefix, b, short)
			}
			return nil
		}

		name := args[0]
		if refs.ReadBranch(name) != "" {
			return fmt.Errorf("branch '%s' already exists", name)
		}

		head := refs.ReadHead()
		if head == "" {
			return fmt.Errorf("no commits yet")
		}

		refs.UpdateBranch(name, head)
		fmt.Println("Created branch " + name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(branchCmd)
	branchCmd.Flags().BoolVarP(&deleteBranch, "delete", "d", false, "Delete a branch")
}
