package main

import (
	"github.com/bmizerany/pat"
	"github.com/tanc7/go-course/pkg/config"
	"github.com/tanc7/go-course/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux
}
