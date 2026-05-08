package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request: " + r.URL.Path)
	http.ServeFile(w, r, "public/index.html")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting server at port 8080")
	fmt.Println("Serving public/index.html")
	http.ListenAndServe(":8080", nil)
}
