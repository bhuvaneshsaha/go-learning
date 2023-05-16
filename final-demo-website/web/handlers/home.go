package handlers

import (
	"net/http"

	"github.com/bhuvaneshsaha/website-demo/web/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// Define the data that you want to pass to the template
	data := struct {
		Title string
	}{
		Title: "Home",
	}

	render.RenderTemplate(w, data, "web/templates/home.page.html", "web/templates/base.layout.html")
}
