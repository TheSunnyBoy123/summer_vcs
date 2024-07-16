package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var write bool

// hashObjectCmd represents the hashObject command
var hashObjectCmd = &cobra.Command{
	Use:   "hash-file",
	Short: "Displays the SHA hash of the created object for the provided file",
	Long: `Displays the SHA hash of the created object for the provided file. 
	When the write flag is used, the object is written to the objects directory
	
	The content for the object is read from the provided file, the object file is formatted as "ObjectType <size>\00<contents>".
	The object file is then compressed and the SHA hash is calculated for the compressed content. 
	Which is then saved at .sol/objects/<hash[:2]>/<hash[2:]>; hash is the SHA hash of the compressed content.`,
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
