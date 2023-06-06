package main

import (
	"log"
	"net/http"

	"github.com/bhuvaneshsaha/final-demo-website/web/render"
)

func main() {

	render.SetUseCache(false)
	serve := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	err := serve.ListenAndServe()
	log.Fatal(err)
}
