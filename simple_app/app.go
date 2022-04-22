package main

import (
    "fmt"
    "net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}

func main() {
    http.HandleFunc("/hello", hello)

	fmt.Println("Starting server at :8080")
    http.ListenAndServe(":8080", nil)
}
