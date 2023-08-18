package routes

import (
	"net/http"
	"search/internal/handlers"

	"github.com/gorilla/mux"
)

func SetRouter() {
	// Create a new mux router
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/users", handlers.UsersSearch).Queries("query", "{query}")
	r.HandleFunc("/hashtags", handlers.HashtagsSearch).Queries("query", "{query}")
	r.HandleFunc("/fuzzy", handlers.FuzzySearch).Queries("query", "{query}")

	http.Handle("/", r)
}
