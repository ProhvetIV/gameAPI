package main

import (
	"fmt"
	"gameAPI/handlers"
	"net/http"
)

func main() {
	port := ":8000"

	mux := http.NewServeMux()

	mux.HandleFunc("/receive", handlers.HandleGet)

	fmt.Printf("Server running on port: %s\n", port)
	http.ListenAndServe(port, mux)
}
