package config

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const PORT = "3000"

type ServerConfig struct {
	Addr      string
	StaticDir string
}

var srvConfig ServerConfig

func init() {
	loadEnv()
}

func InitServerConfig() ServerConfig {
	// Configure Environment variables for the project.
	srvConfig = ServerConfig{
		Addr:      os.Getenv("PORT"),
		StaticDir: os.Getenv("STATIC_DIR"),
	}
	flag.StringVar(&srvConfig.Addr, "addr", srvConfig.Addr, "HTTP network address")
	flag.StringVar(&srvConfig.StaticDir, "static-dir", srvConfig.StaticDir, "Static directory")
	flag.Parse()

	if srvConfig.Addr == "" {
		srvConfig.Addr = PORT
	}

	return srvConfig
}

func (srvConfig ServerConfig) StartServer(handler *http.ServeMux) {
	srv := http.Server{
		Addr:     fmt.Sprintf(":%v", srvConfig.Addr),
		Handler:  handler,
		ErrorLog: app.ErrorLog,
	}
	app.InfoLog.Printf("Starting server on %v\n", srvConfig.Addr)
	err := srv.ListenAndServe()
	app.ErrorLog.Fatal("Couldn't start Server: ", err)
}
