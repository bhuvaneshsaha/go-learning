package handlers

import (
	"html/template"
	"net/http"

	"github.com/bhuvaneshsaha/website-demo/web/render"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	// Define the data that you want to pass to the template
	data := struct {
		Title   string
		Message string
	}{
		Title:   "About Page",
		Message: "Learn more about me!",
	}

	render.RenderTemplate(w, data, "web/templates/home.page.html", "web/templates/base.layout.html")
}

func FeatureHandler(w http.ResponseWriter, r *http.Request) {

	// Define the data that you want to pass to the template
	data := struct {
		Title  string
		Script template.HTML
	}{
		Title: "Feature",
		Script: template.HTML(`
            <script>
				document.addEventListener('DOMContentLoaded', function() { 
				const element = document.getElementById('features');
				
				if (element) {
				  setTimeout(() => {
					element.scrollIntoView();
				  }, 1)
					
				}
			   })
            </script>
        `),
	}

	render.RenderTemplate(w, data, "web/templates/home.page.html", "web/templates/base.layout.html")
}

func SupportHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data that you want to pass to the template
	data := struct {
		Title  string
		Script template.HTML
	}{
		Title: "Feature",
		Script: template.HTML(`
            <script>
				document.addEventListener('DOMContentLoaded', function() { 
				const element = document.getElementById('support');
				
				if (element) {
				  setTimeout(() => {
					element.scrollIntoView();
				  }, 1)
					
				}
			   })
            </script>
        `),
	}

	// Render the template with the data
	render.RenderTemplate(w, data, "web/templates/home.page.html", "web/templates/base.layout.html")
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data that you want to pass to the template
	data := struct {
		Title string
	}{
		Title: "Feature",
	}

	// Render the template with the data
	render.RenderTemplate(w, data, "web/templates/signin.html", "web/templates/base.layout.html")
}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data that you want to pass to the template
	data := struct {
		Title string
	}{
		Title: "Feature",
	}

	// Render the template with the data
	render.RenderTemplate(w, data, "web/templates/signup.html", "web/templates/base.layout.html")
}
