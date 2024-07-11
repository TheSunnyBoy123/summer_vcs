package cmd

import (
	"fmt"
	"bytes"
	"github.com/spf13/cobra"
)

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
		fmt.Println("catFile called with hash: ", args[0])
        hash := args[0]
        if fileExists(".sol/objects/" + hash[:2] + "/" + hash[2:]) {
            contents := readFile(".sol/objects/" + hash[:2] + "/" + hash[2:])
            contents = decompress(contents)
            // Convert contents to []byte before splitting
            elements := bytes.Split([]byte(contents), []byte("\x00"))
            // Iterate through the elements and print them
            for _, element := range elements {
                fmt.Println(string(element))
            }
        } else {
			fmt.Println("File does not exist")
		}
	},
}

func init() {
	rootCmd.AddCommand(catFileCmd)

}
