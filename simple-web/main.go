package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
)

const portNumber = "localhost:8080"

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.page.tmpl")
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divide(100, 0)
	if err != nil {
		fmt.Fprintf(w, "Cannot divided by zero")
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("Result %f", f))
}

func divide(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("cannot divided by zero")
		return 0, err
	}

	return x / y, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
}

func main() {

	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	_ = http.ListenAndServe(portNumber, nil)
}
