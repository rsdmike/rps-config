package repo

import (
	"context"
	"fmt"

	"github.com/rsdmike/rps/pkg/postgres"
)

// DomainRepo -.
type DomainRepo struct {
	*postgres.Postgres
}

// New -.
func NewDomainRepo(pg *postgres.Postgres) *DomainRepo {
	return &DomainRepo{pg}
}

// GetCount -.
func (r *DomainRepo) GetCount(ctx context.Context, tenantId string) (int, error) {
	sql, _, err := r.Builder.
		Select("COUNT(*) OVER() AS total_count").
		From("domains").
		Where("tenant_id = ?").
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("DomainRepo - GetCount - r.Builder: %w", err)
	}

	var count int
	err = r.Pool.QueryRow(ctx, sql, tenantId).Scan(&count)
	if err.Error() == "no rows in result set" {
		return 0, nil
	} else if err != nil {
		return 0, fmt.Errorf("DomainRepo - GetCount - r.Pool.QueryRow: %w", err)
	}

	return count, nil
}
