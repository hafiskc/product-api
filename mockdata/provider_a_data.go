package mockdata

import "example.com/product-api/model"

var providerAData = []model.ProviderAProduct{
    {Code: "A1", Title: "Wireless Mouse", Cost: 20, Currency: "USD"},
    {Code: "A2", Title: "Keyboard", Cost: 30, Currency: "USD"},
}
func GetProviderAData() []model.ProviderAProduct {
	return providerAData
}
	