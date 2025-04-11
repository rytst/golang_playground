package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf(" ユーザー ID: %d\n", os.Getuid())
	fmt.Printf(" グループ ID: %d\n", os.Getgid())
	groups, _ := os.Getgroups()
	fmt.Printf(" サブグループ ID: %v\n", groups)
}
