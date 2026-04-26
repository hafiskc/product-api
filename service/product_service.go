package service

import (
	"context"
	"sort"
	"sync"
	"time"

	"example.com/product-api/model"
	"example.com/product-api/normalizer"
	"example.com/product-api/provider"
	"example.com/product-api/repository"
	"github.com/google/uuid"
)

type ProductService struct {
	providers   []provider.Provider
	historyRepo *repository.HistoryRepository
}

func NewProductService(providers []provider.Provider, repo *repository.HistoryRepository) *ProductService {
	return &ProductService{providers: providers, historyRepo: repo}
}

func (s *ProductService) Search(query string) model.SearchResponse {

	var wg sync.WaitGroup
	resultChan := make(chan []model.Product, len(s.providers))

	var mu sync.Mutex
	var failedProviders []string

	for _, p := range s.providers {

		wg.Add(1)

		go func(pr provider.Provider) {
			defer wg.Done()

			ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
			defer cancel()

			raw, err := pr.Search(ctx, query)
			if err != nil {
				mu.Lock()
				failedProviders = append(failedProviders, pr.Name())
				mu.Unlock()
				return
			}

			normalized := normalizer.Normalize(pr.Name(), raw)
			resultChan <- normalized

		}(p)
	}

	wg.Wait()
	close(resultChan)

	var allProducts []model.Product

	for res := range resultChan {
		allProducts = append(allProducts, res...)
	}

	unique := deduplicate(allProducts)

	sort.Slice(unique, func(i, j int) bool {
		return unique[i].Price < unique[j].Price
	})

	// ✅ SAVE HISTORY (this was missing)
	s.historyRepo.Save(model.SearchHistory{
		ID:              uuid.New().String(),
		Query:           query,
		Timestamp:       time.Now(),
		ResultCount:     len(unique),
		FailedProviders: failedProviders,
	})

	return model.SearchResponse{
		Query:           query,
		Products:        unique,
		FailedProviders: failedProviders,
	}
}
func deduplicate(products []model.Product) []model.Product {

	seen := make(map[string]bool)
	var result []model.Product

	for _, p := range products {

		if !seen[p.SKU] {
			seen[p.SKU] = true
			result = append(result, p)
		}
	}

	return result
}
