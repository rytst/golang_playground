package main

import (
	"fmt"
	"time"
)

func downloadJSON(u string) {
	// download JSON
	println(u)
	time.Sleep(time.Second * 6)
}

func main() {
	before := time.Now()
	for i := 0; i < 100; i++ {
		u := fmt.Sprintf("http://example.com/api/users?id=%d", i)
		downloadJSON(u)
	}

	fmt.Printf("%v\n", time.Since(before))
}
