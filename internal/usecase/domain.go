package usecase

import (
	"context"
	"fmt"
)

// TranslationUseCase -.
type DomainsUseCase struct {
	repo Domain
}

// New -.
func NewDomainsUseCase(r Domain) *DomainsUseCase {
	return &DomainsUseCase{
		repo: r,
	}
}

// History - getting translate history from store.
func (uc *DomainsUseCase) GetCount(ctx context.Context, tenantId string) (int, error) {
	count, err := uc.repo.GetCount(ctx, tenantId)
	if err != nil {
		return 0, fmt.Errorf("DomainsUseCase - Count - s.repo.GetCount: %w", err)
	}

	return count, nil
}
