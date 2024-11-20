package routes

import (
	"github.com/fahimimam/letsgo/cmd/web/config"
	"github.com/fahimimam/letsgo/cmd/web/handlers"
	"net/http"
)

func AppRoutes(app *config.App) *http.ServeMux {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.Handle("/", handlers.HomeHandler(app))
	mux.Handle("/snippet/view", handlers.SnippetViewHandler(app))
	mux.Handle("/snippet/create", handlers.SnippetCreateHandler(app))

	return mux
}
