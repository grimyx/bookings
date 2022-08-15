package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/grimyx/bookings/pkg/config"
	"github.com/grimyx/bookings/pkg/handlers"
	"github.com/grimyx/bookings/pkg/render"
)

const portNumber string = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// change this to true when in production
	app.InProduction = false

	// kreira novu sesiju
	session = scs.New()

	// podesava trajanje sesije, u ovom slucaju 24 sata
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("can not create template cache", err)
	}

	app.TemplateCache = tc
	render.NewTemplate(&app)

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	fmt.Printf("Starting server at port %s ", portNumber)
	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}
