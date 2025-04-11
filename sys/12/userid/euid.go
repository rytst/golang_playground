package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf(" ユーザー ID: %d\n", os.Getuid())
	fmt.Printf(" グループ ID: %d\n", os.Getgid())
	fmt.Printf(" 実効ユーザー ID: %d\n", os.Geteuid())
	fmt.Printf(" 実効グループ ID: %d\n", os.Getegid())
}
