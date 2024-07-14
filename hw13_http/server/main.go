package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		fmt.Fprintf(w, "Hello, this is a GET response!")
		fmt.Printf("Received GET request for %s\n", r.URL.Path)
	case "POST":
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Can't read body", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello, this is a POST response!")
		fmt.Printf("Received POST request for %s with data: %s\n", r.URL.Path, string(body))
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	addr := flag.String("addr", "localhost", "Address to bind the server")
	port := flag.Int("port", 8000, "Port to bind the server")
	flag.Parse()

	http.HandleFunc("/", handler)
	serverAddr := fmt.Sprintf("%s:%d", *addr, *port)
	fmt.Printf("Starting server on %s\n", serverAddr)
	if err := http.ListenAndServe(serverAddr, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
