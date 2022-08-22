package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/grimyx/bookings/pkg/config"
	"github.com/grimyx/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	// creates new mux
	mux := chi.NewRouter()

	// Add middleware to mux
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	// Adds routes and handlers to mux
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.MajorsSuite)
	mux.Get("/search-now", handlers.Repo.SearchAvailability)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	// creates handler for static files
	fileServer := http.FileServer(http.Dir("./static/"))

	// add handler for static files to mux
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
