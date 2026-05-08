package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request" + r.URL.Path)
	fmt.Fprintf(w, "<h1>Welcome to our server</h1>")
}
func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at port 8080")
	http.ListenAndServe(":8080", nil)
}
