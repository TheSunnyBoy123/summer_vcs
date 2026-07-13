package cmd

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

var deleteTag bool

var tagCmd = &cobra.Command{
	Use:   "tag [name]",
	Short: "List, create, or delete tags",
	Long: `With no arguments, list existing tags.
With a name, create a lightweight tag at the current commit.
With -d, delete the named tag.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		refs := NewRefs(solPath)

		if deleteTag {
			if len(args) != 1 {
				return fmt.Errorf("usage: sol tag -d <name>")
			}
			name := args[0]
			if refs.ReadTag(name) == "" {
				return fmt.Errorf("tag '%s' not found", name)
			}
			refs.DeleteTag(name)
			fmt.Println("Deleted tag " + name)
			return nil
		}

		if len(args) == 0 {
			tags := refs.ListTags()
			sort.Strings(tags)
			for _, t := range tags {
				commit := refs.ReadTag(t)
				short := ""
				if len(commit) >= 8 {
					short = commit[:8]
				}
				fmt.Printf("%s %s\n", t, short)
			}
			return nil
		}

		name := args[0]
		if refs.ReadTag(name) != "" {
			return fmt.Errorf("tag '%s' already exists", name)
		}

		head := refs.ReadHead()
		if head == "" {
			return fmt.Errorf("no commits yet")
		}

		refs.UpdateTag(name, head)
		fmt.Println("Created tag " + name)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
	tagCmd.Flags().BoolVarP(&deleteTag, "delete", "d", false, "Delete a tag")
}
