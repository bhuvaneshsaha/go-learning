package handlers

import (
	"net/http"

	"github.com/bhuvaneshsaha/booking/web/render"
)

type Product struct {
	Name  string
	Price float64
}

type ProductList struct {
	Title    string
	Products []Product
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Define the data that you want to pass to the template
	data := ProductList{
		Title: "Products",
		Products: []Product{
			{Name: "Product 1", Price: 9.99},
			{Name: "Product 2", Price: 19.99},
			{Name: "Product 3", Price: 29.99},
		},
	}

	// Render the template with the data
	render.RenderTemplate(w, data, "web/templates/product.page.tmpl", "web/templates/base.layout.tmpl")
}
