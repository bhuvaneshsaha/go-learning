package handlers

import (
	"net/http"

	"github.com/bhuvaneshsaha/booking/web/config"
	"github.com/bhuvaneshsaha/booking/web/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	appConfig := config.GetConfig()

	remoteIP := r.RemoteAddr

	appConfig.Sessions.Put(r.Context(), "remote_ip", remoteIP)

	name := r.URL.Query().Get("name")
	title := r.URL.Query().Get("title")
	message := r.URL.Query().Get("message")

	// Define the data that you want to pass to the template
	data := struct {
		Title   string
		Message string
		Name    string
	}{
		Title:   title,
		Message: message,
		Name:    name,
	}

	render.RenderTemplate(w, data, "web/templates/home.page.tmpl", "web/templates/base.layout.tmpl")
}
