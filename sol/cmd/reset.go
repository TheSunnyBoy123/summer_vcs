package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var hardReset bool

var resetCmd = &cobra.Command{
	Use:   "reset <commit>",
	Short: "Reset current HEAD to a specified commit",
	Long: `Moves the current branch (or HEAD) to the specified commit.
With --hard, also restores the working tree to match.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		target := resolveRevision(args[0])
		if target == "" {
			target = args[0]
		}
		refs := NewRefs(solPath)

		header, body, err := readCommitFromHash(target)
		if err != nil {
			return fmt.Errorf("invalid commit: %s", args[0])
		}
		if !strings.HasPrefix(header, "commit") {
			return fmt.Errorf("object is not a commit: %s", target)
		}

		if hardReset {
			meta := parseCommitBody(body)
			treeOID := meta["tree"]
			if treeOID != "" {
				checkoutTree(treeOID, ".")
			}
		}

		branch := refs.CurrentBranch()
		if branch != "" {
			refs.UpdateBranch(branch, target)
			fmt.Println("Branch " + branch + " is now at " + target[:8])
		} else {
			refs.UpdateHead(target)
			fmt.Println("HEAD is now at " + target[:8])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
	resetCmd.Flags().BoolVarP(&hardReset, "hard", "", false, "Reset working tree to match")
}
