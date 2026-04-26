package mockdata

import "example.com/product-api/model"

func GetProviderCData() []model.ProviderCProduct {
	// Mock data for ProviderC
	return []model.ProviderCProduct{
		{Ref: "C1", Product: "Gadget X", AmountUSD: 15.99, Currency: "USD"},
		{Ref: "C2", Product: "Gadget Y", AmountUSD: 25.99, Currency: "USD"},
		{Ref: "C3", Product: "Gadget Z", AmountUSD: 35.99, Currency: "USD"},
		{Ref: "C4", Product: "Product A", AmountUSD: 35.99, Currency: "USD"},
	}
}