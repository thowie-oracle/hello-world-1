package main

import (
    "net/http"
    "github.com/gorilla/handlers"
    "os"
    "fmt"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, World!"))
}

func main() {
    fmt.Printf("Starting Hello World application... \n")
    http.HandleFunc("/", HelloWorld)
    http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
