package main

import (
	"net/http"

	"github.com/bhuvaneshsaha/website-demo/web/handlers"
	"github.com/go-chi/chi"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	// Middleware
	// mux.Use(middleware.Recoverer)
	// mux.Use(WriteToConsole)
	// mux.Use(NoSurf)
	// mux.Use(SessionLoad)

	fs := http.FileServer(http.Dir("./web/public"))
	mux.Handle("/static/*", http.StripPrefix("/static/", fs))

	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)
	mux.Get("/service", handlers.ServiceHandler)

	return mux
}
