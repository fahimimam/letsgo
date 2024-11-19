package main

import (
	"flag"
	"fmt"
	"github.com/fahimimam/letsgo/cmd/web/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const PORT = "3000"

type config struct {
	addr      string
	staticDir string
}

var cfg config

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	cfg = config{
		addr:      os.Getenv("STATIC_DIR"),
		staticDir: os.Getenv("ADDR"),
	}
	flag.StringVar(&cfg.addr, "addr", cfg.addr, "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", cfg.staticDir, "Static directory")
	flag.Parse()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/snippet/view", handlers.SnippetView)
	mux.HandleFunc("/snippet/create", handlers.SnippetCreate)

	if cfg.addr == "" {
		cfg.addr = PORT
	}
	log.Printf("Starting server on %v\n", cfg.addr)
	err = http.ListenAndServe(fmt.Sprintf(":%v", cfg.addr), mux)
	if err != nil {
		log.Fatal("Couldn't start Server: ", err)
	}

}
