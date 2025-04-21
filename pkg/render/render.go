package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/watcharaY/building-modern-webapp-with-go/pkg/config"
	"github.com/watcharaY/building-modern-webapp-with-go/pkg/models"
)

var app *config.AppConfig

// NewTemplates sets the config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		log.Println("get template from cache")
		tc = app.TemplateCache
	} else {
		// create template cache
		tc, _ = CreateTemplateCache()
	}
	// get requested template from cache
	t, ok := tc[tmpl]
	if !ok {
		log.Fatalln("error getting template from cache")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, data)
	if err != nil {
		log.Fatalln(err)
	}

	// render template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with .page.tmpl
	for _, page := range pages {
		log.Println("iteration over page:", page)
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
		log.Println("added template to cache:", name)
	}
	return myCache, nil
}
