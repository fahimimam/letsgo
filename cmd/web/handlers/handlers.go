package handlers

import (
	"errors"
	"fmt"
	"github.com/fahimimam/letsgo/cmd/web/config"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func HomeHandler(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			app.NotFound(w)
			return
		}
		files := []string{
			"./ui/html/base.gohtml",
			"./ui/html/partials/nav.gohtml",
			"./ui/html/pages/home.gohtml",
		}
		ts, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println("Error occurred while parsing templates...", err)
			app.ServerError(w, err, http.StatusInternalServerError)
			return
		}

		if ts == nil {
			app.ServerError(w, errors.New("templates not found"), http.StatusInternalServerError)
			return
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.ErrorLog.Println("Error occurred while executing templates...", err)
			app.ServerError(w, err, http.StatusInternalServerError)
		}
		app.InfoLog.Println("rendered home page successfully...")
	}
}

func SnippetViewHandler(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			app.NotFound(w)
			return
		}
		bytesWritten, err := fmt.Fprintf(w, "Displaying a specific snippet with id: %d", id)
		app.InfoLog.Println("Written ", bytesWritten, " Bytes to response...")
		if err != nil {
			app.ErrorLog.Println("Error occurred while returning response...", err)
			app.ServerError(w, err, http.StatusInternalServerError)
			return
		}

	}
}

func SnippetCreateHandler(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			app.ServerError(w, errors.New("method apart from post not allowed"), http.StatusMethodNotAllowed)
			return
		}

		bytesWritten, err := fmt.Fprintf(w, "Created a new snippet...")
		log.Println("Written ", bytesWritten, " Bytes to response...")
		if err != nil {
			app.ErrorLog.Println("Error occurred while returning response...", err)
			app.ServerError(w, err, http.StatusInternalServerError)
			return
		}
	}
}
