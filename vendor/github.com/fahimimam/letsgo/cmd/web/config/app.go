package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

type App struct {
	ErrorLog         *log.Logger
	InfoLog          *log.Logger
	EnableStackTrace bool
}

var app *App

func init() {
	initApp()
}

func GetApp() *App {
	if app == nil {
		initApp()
	}
	return app
}

func initApp() *App {
	// Configure Custom Logger for the project.
	infoLog := log.New(os.Stdout, "[INFO]\t", log.LstdFlags)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize app for dependency injection - (For now it's only the logger)
	app = &App{
		ErrorLog:         errorLog,
		InfoLog:          infoLog,
		EnableStackTrace: os.Getenv("STACK_TRACE") == "true",
	}
	if app.EnableStackTrace {
		infoLog.Println("Trace is online")
	}
	return app
}

func (app *App) ServerError(w http.ResponseWriter, err error, code int) {
	if app.EnableStackTrace {
		trace := fmt.Sprintf("Error: %s\nTrace: \n%s", err.Error(), debug.Stack())
		printErr := app.ErrorLog.Output(2, trace)
		if printErr != nil {
			app.ErrorLog.Println(trace)
			return
		}
	} else {
		printErr := app.ErrorLog.Output(2, err.Error())
		if printErr != nil {
			app.ErrorLog.Println("Error: ", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
	}
	if code >= 400 && code < 500 {
		http.Error(w, http.StatusText(code), code)
		return
	} else {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func (app *App) ClientError(w http.ResponseWriter, code int) {
	http.Error(w, http.StatusText(code), code)

}

func (app *App) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
