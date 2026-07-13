package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var createBranch bool

func readCommitFromHash(hash string) (string, string, error) {
	path := objectsPath + hash[:2] + "/" + hash[2:]
	if !fileExists(path) {
		return "", "", fmt.Errorf("object not found: %s", hash)
	}
	contents := decompress(readFile(path))
	parts := strings.SplitN(contents, "\x00", 2)
	if len(parts) < 2 {
		return "", "", fmt.Errorf("invalid object: %s", hash)
	}
	return parts[0], parts[1], nil
}

func parseCommitBody(body string) map[string]string {
	result := make(map[string]string)
	parts := strings.SplitN(body, "\n\n", 2)
	if len(parts) > 0 {
		for _, line := range strings.Split(parts[0], "\n") {
			if idx := strings.Index(line, " "); idx > 0 {
				key := line[:idx]
				val := strings.TrimSpace(line[idx+1:])
				result[key] = val
			}
		}
		if len(parts) > 1 {
			result["message"] = strings.TrimSpace(parts[1])
		}
	}
	return result
}

func parseTreeEntries(oid string) ([]string, error) {
	path := objectsPath + oid[:2] + "/" + oid[2:]
	if !fileExists(path) {
		return nil, fmt.Errorf("tree not found: %s", oid)
	}
	contents := decompress(readFile(path))
	parts := strings.SplitN(contents, "\x00", 2)
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid tree: %s", oid)
	}
	entries := strings.Split(parts[1], "\x00")
	var result []string
	for _, e := range entries {
		if e != "" {
			result = append(result, e)
		}
	}
	return result, nil
}

func checkoutTree(treeOID, basePath string) error {
	entries, err := parseTreeEntries(treeOID)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fields := strings.SplitN(entry, " ", 3)
		if len(fields) < 3 {
			continue
		}
		mode := fields[0]
		oid := fields[1]
		name := fields[2]
		fullPath := filepath.Join(basePath, name)
		if mode == "40000" {
			if !dirExists(fullPath) {
				createDir(fullPath)
			}
			checkoutTree(oid, fullPath)
		} else {
			if !dirExists(filepath.Dir(fullPath)) {
				createDir(filepath.Dir(fullPath))
			}
			raw := decompress(readFile(objectsPath + oid[:2] + "/" + oid[2:]))
			blobParts := strings.SplitN(raw, "\x00", 2)
			if len(blobParts) >= 2 {
				writeFile(fullPath, blobParts[1])
			}
		}
	}
	return nil
}

var checkoutCmd = &cobra.Command{
	Use:   "checkout <branch-or-commit>",
	Short: "Switch branches or restore working tree files",
	Long: `Switches to a branch or commit and restores the working tree.
Use -b to create and switch to a new branch.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("usage: sol checkout [-b] <branch-or-commit>")
		}
		target := args[0]
		refs := NewRefs(solPath)

		if createBranch {
			head := refs.ReadHead()
			if head == "" {
				return fmt.Errorf("no commits yet")
			}
			if refs.ReadBranch(target) != "" {
				return fmt.Errorf("branch '%s' already exists", target)
			}
			refs.UpdateBranch(target, head)
			refs.SetHeadToBranch(target)
			fmt.Println("Switched to new branch " + target)
			return nil
		}

		commitHash := refs.ReadBranch(target)
		isBranch := commitHash != ""
		if commitHash == "" {
			commitHash = target
		}

		header, body, err := readCommitFromHash(commitHash)
		if err != nil {
			return fmt.Errorf("could not find commit '%s'", target)
		}
		if !strings.HasPrefix(header, "commit") {
			return fmt.Errorf("object '%s' is not a commit", target)
		}

		meta := parseCommitBody(body)
		treeOID := meta["tree"]

		checkoutTree(treeOID, ".")

		if isBranch {
			refs.SetHeadToBranch(target)
			fmt.Println("Switched to branch " + target)
		} else {
			refs.UpdateHead(commitHash)
			fmt.Println("HEAD is now at " + commitHash[:8])
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkoutCmd)
	checkoutCmd.Flags().BoolVarP(&createBranch, "branch", "b", false, "Create a new branch and switch to it")
}
