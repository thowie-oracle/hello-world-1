package main

import (
	"fmt"
	"github.com/gorilla/handlers"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello again, Open World peeps!!!!!"))
}

func main() {
	fmt.Printf("Welcome to the Hello World application... \n")
	http.HandleFunc("/", HelloWorld)
	http.ListenAndServe(":8080", handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
}
