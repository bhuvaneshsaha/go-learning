package handlers

import (
	"net/http"

	"github.com/bhuvaneshsaha/booking/web/config"
	"github.com/bhuvaneshsaha/booking/web/render"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {

	appConfig := config.GetConfig()

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again."

	iip := appConfig.Sessions.GetString(r.Context(), "remote_ip")

	// Define the data that you want to pass to the template
	data := struct {
		Title        string
		Message      string
		InProduction bool
		RemoteIP     string
	}{
		Title:        "About Page",
		Message:      "Learn more about me!",
		InProduction: appConfig.InProduction,
		RemoteIP:     iip,
	}

	render.RenderTemplate(w, data, "web/templates/about.page.tmpl", "web/templates/base.layout.tmpl")
}
