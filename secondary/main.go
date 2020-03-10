package main

import (
    "os"
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", HelloServer)
    fmt.Println("Server Online")
    http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s! %s\n", r.URL.Path[1:], os.Getenv("MY_POD_NAME"))
}
