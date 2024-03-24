package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, HTTPS world!")
}

func main() {
    http.HandleFunc("/", handler)

    // HTTPSサーバーを起動
    fmt.Println("Starting HTTPS server on :443...")
    err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    if err != nil {
        fmt.Println("Error starting server: ", err)
    }
}
