package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/bhuvaneshsaha/final-demo-website/app/services/emailService"
	"github.com/bhuvaneshsaha/final-demo-website/web/render"
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

func FormHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data that you want to pass to the template
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message := r.FormValue("message")
	fullName := r.FormValue("fullname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	subject := r.FormValue("subject")

	fmt.Println(message)
	fmt.Println(fullName)
	fmt.Println(email)
	fmt.Println(phone)
	fmt.Println(subject)

	emailService.SendEmail("")

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
