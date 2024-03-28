// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/rsdmike/rps/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_test.go -package=usecase_test

type (
	// Translation -.
	Domain interface {
		GetCount(context.Context, string) (int, error)
		Get(ctx context.Context, top, skip int, tenantID string) ([]entity.Domain, error)
		GetDomainByDomainSuffix(ctx context.Context, domainSuffix, tenantID string) (*entity.Domain, error)
		GetByName(ctx context.Context, domainName, tenantID string) (*entity.Domain, error)
		Delete(ctx context.Context, domainName, tenantID string) (bool, error)
		Update(ctx context.Context, d *entity.Domain) (bool, error)
		Insert(ctx context.Context, d *entity.Domain) (string, error)
	}
)
