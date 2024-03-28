package postgresdb

import "github.com/rsdmike/rps/pkg/postgres"

// DomainRepo -.
type CIRARepo struct {
	*postgres.DB
}

// New -.
func NewCIRARepo(pg *postgres.DB) *CIRARepo {
	return &CIRARepo{pg}
}
