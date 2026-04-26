package service

import (
	"context"
	"errors"
	"testing"

	"example.com/product-api/model"
	"example.com/product-api/provider"
	"example.com/product-api/repository"
)

// ✅ Mock Provider (reusable)
type MockProvider struct {
	name string
	data []model.ProviderAProduct
	err  error
}

func (m *MockProvider) Name() string {
	return m.name
}

func (m *MockProvider) Search(ctx context.Context, query string) (interface{}, error) {
	return m.data, m.err
}

// ✅ Test 1: Basic success case
func TestSearch_ReturnsProducts(t *testing.T) {

	providers := []provider.Provider{
		&MockProvider{
			name: "provider_a",
			data: []model.ProviderAProduct{
				{Code: "A1", Title: "Mouse", Cost: 10, Currency: "USD"},
			},
		},
	}

	repo := repository.NewHistoryRepository()
	svc := NewProductService(providers, repo)

	result := svc.Search("mouse")

	if len(result.Products) == 0 {
		t.Error("expected products, got empty")
	}
}

// ✅ Test 2: Deduplication (keep lowest price)
func TestSearch_Deduplicate(t *testing.T) {

	providers := []provider.Provider{
		&MockProvider{
			name: "provider_a",
			data: []model.ProviderAProduct{
				{Code: "A1", Title: "Mouse", Cost: 20, Currency: "USD"},
				{Code: "A1", Title: "Mouse", Cost: 10, Currency: "USD"},
			},
		},
	}

	repo := repository.NewHistoryRepository()
	svc := NewProductService(providers, repo)

	result := svc.Search("mouse")

	if len(result.Products) != 1 {
		t.Errorf("expected 1 product after deduplication, got %d", len(result.Products))
	}

	if result.Products[0].Price != 10 {
		t.Errorf("expected lowest price 10, got %f", result.Products[0].Price)
	}
}

// ✅ Test 3: Failed provider handling
func TestSearch_FailedProvider(t *testing.T) {

	providers := []provider.Provider{
		&MockProvider{
			name: "provider_a",
			err:  errors.New("failed"),
		},
	}

	repo := repository.NewHistoryRepository()
	svc := NewProductService(providers, repo)

	result := svc.Search("mouse")

	if len(result.FailedProviders) != 1 {
		t.Error("expected 1 failed provider")
	}
}

