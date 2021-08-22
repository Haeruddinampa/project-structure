package main

import (
	"fmt"
	"github.com/Budi721/Udemy_Web/pkg/config"
	"github.com/Budi721/Udemy_Web/pkg/handlers"
	"github.com/Budi721/Udemy_Web/pkg/render"
	"log"
	"net/http"
)

const (
	PORT = ":8080"
)

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot write template cache")
	}

	app.TemplateCache = tc

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("Server is listening on port %s", PORT))
	_ = http.ListenAndServe(PORT, nil)
}
