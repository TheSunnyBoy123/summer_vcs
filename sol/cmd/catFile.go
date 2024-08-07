package cmd

import (
	"bytes"
	"fmt"

	"github.com/spf13/cobra"
)

var t bool
var p bool
var s bool

var catFileCmd = &cobra.Command{
	Use:   "cat-file",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println("Usage: cat-file <hash>")
			return
		}
		// fmt.Println("catFile called with hash: ", args[0])
		hash := args[0]
		// TODO: add ability to access file just from 5 characters of hash
		filePath := ".sol/objects/" + hash[:2] + "/" + hash[2:]
		if fileExists(filePath) {
			contents := decompress(readFile(filePath))
			// fmt.Println("contents: ", string(contents))
			lines := bytes.Split([]byte(contents), []byte("\x00"))
			firstLine := lines[0]
			// fmt.Println("firstLine: ", string(firstLine))
			words := bytes.Split(firstLine, []byte(" "))
			typeObject, size := string(words[0]), string(words[1])
			if t { // type case
				fmt.Println(typeObject)
			}
			if p { // pretty print case
				for _, line := range lines[1:] {
					fmt.Println(string(line))
				}
			}
			if s { // size case
				fmt.Println("size: ", size)
			}

		} else {
			fmt.Println("File does not exist")
		}

	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)

	catFileCmd.Flags().BoolVarP(&t, "type", "t", false, "Prints the type of the object")
	catFileCmd.Flags().BoolVarP(&p, "pretty", "p", false, "Pretty print the contents of the object")
	catFileCmd.Flags().BoolVarP(&s, "size", "s", false, "Prints the size of the object")

}
