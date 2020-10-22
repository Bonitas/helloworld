package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// YourHandler handles your stuff!
func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla!\n"))
}
func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8080", r))
}
