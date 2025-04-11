package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf(" プロセス ID: %d\n", os.Getpid())

	// Parent Process ID
	fmt.Printf(" 親プロセス ID: %d\n", os.Getppid())
}
