package main

import (
	"fmt"
	"os"
)


func initialize() {
	_, err := os.Stat(".vcs")
	if os.IsNotExist(err) {
		os.Mkdir(".vcs", 0755)
		file, err := os.Create(".vcs/updates.txt")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		fmt.Println("Initialized")

	} else {
		fmt.Println("Already initialized")
	}
}

func add(files []string) {
	fmt.Println(files)
}

//function commit
func commit(args []string) {
	fmt.Println(args)
}


func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "init" {
			initialize()
		}
		if os.Args[1] == "add" {
			add(os.Args[2:])
		}
		if os.Args[1] == "commit" {
			commit(os.Args[2:])
		}
	} else {
		fmt.Println("Please provide a command")
	}
}


