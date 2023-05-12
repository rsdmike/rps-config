package usecase

import (
	"context"
	"fmt"

	"github.com/rsdmike/rps/internal/entity"
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
func (uc *DomainsUseCase) GetCount(ctx context.Context, tenantID string) (int, error) {
	count, err := uc.repo.GetCount(ctx, tenantID)
	if err != nil {
		return 0, fmt.Errorf("DomainsUseCase - Count - s.repo.GetCount: %w", err)
	}

	return count, nil
}

func (uc *DomainsUseCase) Get(ctx context.Context, top, skip int, tenantID string) ([]entity.Domain, error) {
	data, err := uc.repo.Get(ctx, top, skip, tenantID)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - Get - s.repo.Get: %w", err)
	}

	return data, nil
}

func (uc *DomainsUseCase) GetDomainByDomainSuffix(ctx context.Context, domainSuffix, tenantID string) (*entity.Domain, error) {
	data, err := uc.repo.GetDomainByDomainSuffix(ctx, domainSuffix, tenantID)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - GetDomainByDomainSuffix - s.repo.GetDomainByDomainSuffix: %w", err)
	}

	return data, nil
}

func (uc *DomainsUseCase) GetByName(ctx context.Context, domainName, tenantID string) (*entity.Domain, error) {
	data, err := uc.repo.GetByName(ctx, domainName, tenantID)
	if err != nil {
		return nil, fmt.Errorf("DomainsUseCase - GetByName - s.repo.GetByName: %w", err)
	}

	return data, nil
}

func (uc *DomainsUseCase) Delete(ctx context.Context, domainName, tenantID string) (bool, error) {
	data, err := uc.repo.Delete(ctx, domainName, tenantID)
	if err != nil {
		return false, fmt.Errorf("DomainsUseCase - Delete - s.repo.Delete: %w", err)
	}

	return data, nil
}

func (uc *DomainsUseCase) Update(ctx context.Context, d *entity.Domain) (bool, error) {
	data, err := uc.repo.Update(ctx, d)
	if err != nil {
		return false, fmt.Errorf("DomainsUseCase - Update - s.repo.Update: %w", err)
	}

	return data, nil
}

func (uc *DomainsUseCase) Insert(ctx context.Context, d *entity.Domain) (string, error) {
	data, err := uc.repo.Insert(ctx, d)
	if err != nil {
		return "", fmt.Errorf("DomainsUseCase - Insert - s.repo.Insert: %w", err)
	}

	return data, nil
}
