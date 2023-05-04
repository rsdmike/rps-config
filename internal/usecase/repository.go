package usecase

import (
	"github.com/rsdmike/rps/internal/usecase/repo"
	"github.com/rsdmike/rps/pkg/postgres"
)

// Repositories -.
type Repositories struct {
	Domains Domain
}

// New -.
func New(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		Domains: repo.NewDomainRepo(pg),
	}
}
