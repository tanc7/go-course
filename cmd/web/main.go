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
	//fmt.Println("Hello world!")
	// This is basically a router.rb or route.py module
	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	//http.HandleFunc("/divide", Divide)
	//http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//	n, err := fmt.Fprintf(w, "Hello world!")
	//	if err != nil {
	//		//log.Fatal(err)
	//		fmt.Println(err)
	//	}
	//	//fmt.Sprintf("Bytes written:%d", n)
	//	fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
	//})
	_ = http.ListenAndServe(portNumber, nil)

}
