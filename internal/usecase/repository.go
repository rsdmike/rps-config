package usecase

import (
	"github.com/rsdmike/rps/internal/usecase/postgresdb"
	"github.com/rsdmike/rps/pkg/postgres"
)

// Repositories -.
type Repositories struct {
	Domains Domain
}

// New -.
func New(pg *postgres.DB) *Repositories {
	return &Repositories{
		Domains: postgresdb.NewDomainRepo(pg),
	}
}
