package handlers

import (
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
			http.NotFound(w, r)
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
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = ts.ExecuteTemplate(w, "base", nil)
		if err != nil {
			app.ErrorLog.Println("Error occurred while executing templates...", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		app.InfoLog.Println("rendered home page successfully...")
	}
}

func SnippetViewHandler(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil || id < 1 {
			http.NotFound(w, r)
			return
		}
		bytesWritten, err := fmt.Fprintf(w, "Displaying a specific snippet with id: %d", id)
		app.InfoLog.Println("Written ", bytesWritten, " Bytes to response...")
		if err != nil {
			app.ErrorLog.Println("Error occurred while returning response...", err)
			return
		}

	}
}

func SnippetCreateHandler(app *config.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.Header().Set("Allow", "POST")
			app.ErrorLog.Println("Method apart from post not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		bytesWritten, err := fmt.Fprintf(w, "Created a new snippet...")
		log.Println("Written ", bytesWritten, " Bytes to response...")
		if err != nil {
			return
		}
	}
}
