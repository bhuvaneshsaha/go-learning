package main

import (
    "net/http"

    "github.com/bhuvaneshsaha/web-app-1/web/handlers"
)

func main() {
    http.HandleFunc("/", handlers.HomeHandler)
    http.HandleFunc("/about", handlers.AboutHandler)
    http.ListenAndServe(":8080", nil)
}
