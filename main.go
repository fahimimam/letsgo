package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	log.Println("Entered Home page function")
	w.Write([]byte("Home Page view"))
}

func view(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered view snippet function")
	w.Write([]byte("Display Snippet view"))
}

func create(w http.ResponseWriter, r *http.Request) {
	log.Println("Entered create snippet function")
	w.Write([]byte("Create Snippet view"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/view", view)
	mux.HandleFunc("/create", create)

	log.Println("Listening on :8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
