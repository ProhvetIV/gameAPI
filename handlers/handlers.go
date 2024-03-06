package handlers

import "net/http"

func HandleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("starts here"))
}
