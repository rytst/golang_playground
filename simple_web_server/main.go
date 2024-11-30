package main

import (
	"io"
	"net/http"
	"os"
)

func MyHandler(w http.ResponseWriter /* io.Writer */, r *http.Request) {
	f, err := os.Open("./main.go")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()
	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":8080", nil)
}
