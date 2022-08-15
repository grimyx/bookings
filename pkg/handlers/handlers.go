package handlers

import (
	"net/http"

	"github.com/grimyx/bookings/models"
	"github.com/grimyx/bookings/pkg/config"
	"github.com/grimyx/bookings/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// Creates a new repo
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// Sets repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// cuva adresu sa koje je korisnik pristupio
	remoteIP := r.RemoteAddr

	// kreira polje remote_ip u session cookiu i daje mu vrednost remoteIP
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	// vadi vrednost za remote_ip iz cokisa
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}
