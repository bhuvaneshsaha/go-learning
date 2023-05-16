package main

import (
	"log"
	"net/http"

	"github.com/bhuvaneshsaha/website-demo/web/render"
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
