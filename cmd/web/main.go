package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/watcharaY/building-modern-webapp-with-go/pkg/config"
	"github.com/watcharaY/building-modern-webapp-with-go/pkg/handlers"
	"github.com/watcharaY/building-modern-webapp-with-go/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

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
