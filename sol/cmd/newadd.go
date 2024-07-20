/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newhashDir(dir string) string {
	return ""
}

// newaddCmd represents the newadd command
var newaddCmd = &cobra.Command{
	Use:   "newadd",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		refs := readFile(refsPath)
		refsList := strings.Split(refs, "\x00")
		// split each element in refs then make it a map of string to string
		// for each element in refs, split it by space and add it to the map
		refMap := make(map[string]string)
		for _, ref := range refsList {
			refList := strings.Split(ref, " ")
			refMap[refList[0]] = refList[1]
		}
		fmt.Println(refMap)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(newaddCmd)
}
