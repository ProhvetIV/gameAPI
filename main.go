package main

import "net/http"

func main() {
	port := ":8000"

	mux := http.NewServeMux()

	http.ListenAndServe(port, mux)
}
