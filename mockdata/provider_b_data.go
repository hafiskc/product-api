package mockdata

import "example.com/product-api/model"

func GetProviderBData() []model.ProviderBProduct {
	// Mock data for ProviderB
	return []model.ProviderBProduct{
		{ID: "1", Name: "Product A", Price: 10.99, Currency: "USD"},
		{ID: "2", Name: "Product B", Price: 20.99, Currency: "USD"},
		{ID: "3", Name: "Product C", Price: 30.99, Currency: "USD"},
	}
}