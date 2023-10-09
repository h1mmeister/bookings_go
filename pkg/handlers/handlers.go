package handlers

import (
	"net/http"

	"github.com/h1mmeister/bookings_go/pkg/config"
	"github.com/h1mmeister/bookings_go/pkg/models"
	"github.com/h1mmeister/bookings_go/pkg/render"
)

var Repo *Repository


// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(appConfig *config.AppConfig) *Repository {
	return &Repository{
		App: appConfig,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home page handler
func (m *Repository) Home(writer http.ResponseWriter, request *http.Request) {
	remoteIP := request.RemoteAddr
	m.App.Session.Put(request.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(writer, "home.page.tmpl", &models.TemplateData{})
}

// About page handler
func (m *Repository) About(writer http.ResponseWriter, request *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello"

	remoteIP := m.App.Session.GetString(request.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(writer, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
