package main

import (
	"github.com/tanc7/go-course/pkg/config"
	"github.com/tanc7/go-course/pkg/handlers"
	"github.com/tanc7/go-course/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

//Home is / path

//main wtf bro.
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
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
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
