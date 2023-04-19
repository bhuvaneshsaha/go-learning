package handlers

import (
    "html/template"
    "net/http"
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

    // Parse the HTML template file
    tmpl, err := template.ParseFiles("web/templates/home.html")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Render the template with the data
    err = tmpl.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
