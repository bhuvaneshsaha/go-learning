package main

import (
	"net/http"

	"github.com/bhuvaneshsaha/web-app-1/web/config"
	"github.com/bhuvaneshsaha/web-app-1/web/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// PAT
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.HomeHandler))
	// mux.Get("/about", http.HandlerFunc(handlers.AboutHandler))
	// mux.Get("/product", http.HandlerFunc(handlers.ProductsHandler))

	// return mux

	// Chi
	mux := chi.NewRouter()

	// Middleware
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	mux.Get("/", handlers.HomeHandler)
	mux.Get("/about", handlers.AboutHandler)
	mux.Get("/product", handlers.ProductsHandler)

	return mux
}
