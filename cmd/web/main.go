package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/watcharaY/building-modern-webapp-with-go/pkg/config"
	"github.com/watcharaY/building-modern-webapp-with-go/pkg/handlers"
	"github.com/watcharaY/building-modern-webapp-with-go/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("error creating template cache:", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	log.Println(fmt.Sprintf("Starting server on port %s", portNumber))
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
