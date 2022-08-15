package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grimyx/bookings/pkg/config"
	"github.com/grimyx/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// kreira novi mux koji sadrzi hendlere za rute
	mux := chi.NewRouter()

	// ovde se dodaje middleware
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// dodaju se rute i njihovi hendleri
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// vraca se mux
	return mux
}
