package cmd

import (
	"fmt"
)

func main() {
	test_entry := NewEntry("bin/test", "1234", nil)
	fmt.Println("Parent directories: ", test_entry.ParentDirectories())
}
