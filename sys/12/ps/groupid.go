package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, " グループ ID: %d セッション ID: %d\n",
		syscall.Getpgrp(), sid)

}
