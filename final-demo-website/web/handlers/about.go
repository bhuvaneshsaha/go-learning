package handlers

import (
	"net/http"

	"github.com/bhuvaneshsaha/website-demo/web/render"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	// Define the data that you want to pass to the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "About Page",
		Message: "Learn more about me!",
	}

	render.RenderTemplate(w, data, "web/templates/about.page.html", "web/templates/base.layout.html")
}
