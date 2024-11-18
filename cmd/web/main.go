package main

import (
	"log"
	"net/http"
)

const PORT = ":3000"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Printf("Starting server on %v\n", PORT)
	err := http.ListenAndServe(PORT, mux)
	log.Fatal("Couldn't start Server: ", err)
}
