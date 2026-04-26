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
type ProviderC struct{}

func (p *ProviderC) Name() string {
	return "provider_c"
}

func (p *ProviderC) Search(ctx context.Context, query string) (interface{}, error) {

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
        return nil, errors.New("provider C failed")
    }
	
	data := mockdata.GetProviderCData()
	
	var result []model.ProviderCProduct
	for _, d := range data {
		if strings.Contains(strings.ToLower(d.Product), strings.ToLower(query)) {
			result = append(result, d)
		}
	}

	return result, nil
}
