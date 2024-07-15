package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var write bool

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// exactly 1 arg allowed
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		objHash := args[0]
		dir := ".sol/objects/" + objHash[:2] + "/" + objHash[2:]
		
		contents := readFile(dir)
		size := len(contents)

		contents = "Blob " + string(size) + "\x00" + contents
		contents = compress(contents)

		hash := hashContents(contents)
		fmt.Println(hash)

		if write {
			if dirExists(".sol/objects/" + hash[:2]) {
				writeFile(".sol/objects/" + hash[:2] + "/" + hash[2:], contents)
			} else {
				createDir(".sol/objects/" + hash[:2])
				writeFile(".sol/objects/" + hash[:2] + "/" + hash[2:], contents)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(hashObjectCmd)
	hashObjectCmd.Flags().BoolVarP(&write, "write", "w", false, "Write the object to objects directory")
}
