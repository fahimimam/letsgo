package main

import (
	"flag"
	"github.com/fahimimam/letsgo/cmd/web/handlers"
	"log"
	"net/http"
)

const PORT = ":3000"

func main() {
	addr := flag.String("addr", PORT, "HTTP network Address. (Default is set to 3000)")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)

	log.Printf("Starting server on %v\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal("Couldn't start Server: ", err)
}
