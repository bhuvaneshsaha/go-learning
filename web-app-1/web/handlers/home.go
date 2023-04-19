package handlers

import (
	"net/http"

	"github.com/bhuvaneshsaha/web-app-1/web/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data that you want to pass to the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "Home Page",
		Message: "Welcome to my website!",
	}

	render.RenderTemplate(w, data, "web/templates/home.page.tmpl", "web/templates/base.layout.tmpl")
}
