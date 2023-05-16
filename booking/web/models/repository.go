package models

import "github.com/bhuvaneshsaha/booking/web/config"

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

// NewRepo creates a ne repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// New handlers sets the repository for the Handlers
func NewHandlers(r *Repository) {
	Repo = r
}
