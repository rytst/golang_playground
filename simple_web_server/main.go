package main


import (
    "os"
    "io"
    "net/http"
)


func main() {
    http.HandleFunc("/", func(w http.ResponseWriter /* io.Writer */, r *http.Request) {
        f, err := os.Open("./main.go")

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        defer f.Close()
        io.Copy(w, f)
    })
    http.ListenAndServe(":8080", nil)
}
