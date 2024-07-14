package cmd

import (
	"fmt"
	"bytes"
	"github.com/spf13/cobra"
)

var t bool
var p bool

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
		filePath := ".sol/objects/" + hash[:2] + "/" + hash[2:]
		if fileExists(filePath) {
			contents := decompress(readFile(filePath))
			elements := bytes.Split([]byte(contents), []byte("\x00"))
			if p {
				for _, element := range elements {
					fmt.Println(string(element))
				}
			}
			if t {
				fmt.Println(string(elements[0]))
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

}
