package main

import (
	"log"
	"net/http"

	"github.com/bhuvaneshsaha/web-app-1/web/config"
)

func main() {
	// http.HandleFunc("/", handlers.HomeHandler)
	// http.HandleFunc("/about", handlers.AboutHandler)
	// http.HandleFunc("/product", handlers.ProductsHandler)

	// http.ListenAndServe(":8080", nil)

	var app config.AppConfig

	serve := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}

	erro := serve.ListenAndServe()

	log.Fatal(erro)

}
