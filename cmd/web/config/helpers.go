package config

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *App) ServerError(w http.ResponseWriter, err error, code int) {
	if app.EnableStackTrace {
		trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
		printErr := app.ErrorLog.Output(2, trace)
		if printErr != nil {
			app.ErrorLog.Println(trace)
			return
		}
	} else {
		printErr := app.ErrorLog.Output(2, err.Error())
		if printErr != nil {
			app.ErrorLog.Println(err.Error())
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
