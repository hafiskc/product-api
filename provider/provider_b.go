package provider

import (
	"context"
	"errors"
	"math/rand"
	"strings"
	"time"

	"example.com/product-api/mockdata"
	"example.com/product-api/model"
)

type ProviderB struct{}

func (p *ProviderB) Name() string {
	return "provider_b"
}

func (p *ProviderB) Search(ctx context.Context, query string) (interface{}, error) {

	 // ✅ 1. Simulate latency (100ms–1200ms)
    delay := time.Duration(100+rand.Intn(1100)) * time.Millisecond

    timer := time.NewTimer(delay)
    defer timer.Stop()

    select {
    case <-timer.C:
        // continue after delay
    case <-ctx.Done():
        return nil, ctx.Err() // timeout/cancel
    }

    // ✅ 2. Simulate occasional failure (~20%)
    if rand.Intn(10) < 2 {
        return nil, errors.New("provider B failed")
    }
	
	data := mockdata.GetProviderBData()

	var result []model.ProviderBProduct

	for _, d := range data {
		if strings.Contains(strings.ToLower(d.Name), strings.ToLower(query)) {
			result = append(result, d)
		}
	}

	return result, nil
}
