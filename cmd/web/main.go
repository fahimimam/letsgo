package main

import (
	"flag"
	"fmt"
	"github.com/fahimimam/letsgo/cmd/web/config"
	"github.com/fahimimam/letsgo/cmd/web/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const PORT = "3000"

type serverConfig struct {
	addr      string
	staticDir string
}

var cfg serverConfig

func main() {
	// Configure Environment variables for the project.
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	cfg = serverConfig{
		addr:      os.Getenv("STATIC_DIR"),
		staticDir: os.Getenv("ADDR"),
	}
	flag.StringVar(&cfg.addr, "addr", cfg.addr, "HTTP network address")
	flag.StringVar(&cfg.staticDir, "static-dir", cfg.staticDir, "Static directory")
	flag.Parse()

	// Configure Custom Logger for the project.
	infoLog := log.New(os.Stdout, "[INFO]\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize app for dependency injection - (For now it's only the logger)
	stackTrace := os.Getenv("STACK_TRACE")
	app := &config.App{
		ErrorLog:         errorLog,
		InfoLog:          infoLog,
		EnableStackTrace: stackTrace == "true",
	}
	if app.EnableStackTrace {
		infoLog.Println("Trace is online")
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("/", handlers.HomeHandler(app))
	mux.Handle("/snippet/view", handlers.SnippetViewHandler(app))
	mux.Handle("/snippet/create", handlers.SnippetCreateHandler(app))

	if cfg.addr == "" {
		cfg.addr = PORT
	}
	infoLog.Printf("Starting server on %v\n", cfg.addr)
	srv := http.Server{
		Addr:     fmt.Sprintf(":%v", cfg.addr),
		Handler:  mux,
		ErrorLog: errorLog,
	}
	err = srv.ListenAndServe()
	errorLog.Fatal("Couldn't start Server: ", err)

}
