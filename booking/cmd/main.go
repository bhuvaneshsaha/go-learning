package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bhuvaneshsaha/booking/web/config"
)

var app *config.AppConfig

func main() {
	// http.HandleFunc("/", handlers.HomeHandler)
	// http.HandleFunc("/about", handlers.AboutHandler)
	// http.HandleFunc("/product", handlers.ProductsHandler)

	// http.ListenAndServe(":8080", nil)

	err := config.LoadConfig("cmd/config.json")
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Get the configuration object
	app = config.GetConfig()

	log.Println("UseCache:", app.UseCache)
	log.Println("InfoLog:", app.InfoLog)
	log.Println("InProduction:", app.InProduction)
	log.Println("Sessions:", app.Sessions)

	app.Sessions = scs.New()
	app.Sessions.Lifetime = 24 * time.Hour
	app.Sessions.Cookie.Persist = true
	app.Sessions.Cookie.SameSite = http.SameSiteLaxMode
	app.Sessions.Cookie.Secure = app.InProduction

	serve := &http.Server{
		Addr:    ":8080",
		Handler: routes(app),
	}

	erro := serve.ListenAndServe()

	log.Fatal(erro)

}
