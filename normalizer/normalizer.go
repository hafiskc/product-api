package normalizer

import (
	"example.com/product-api/model"
)

func Normalize(providerName string, data interface{}) []model.Product {

	var products []model.Product

	switch providerName {

	case "provider_a":
		items := data.([]model.ProviderAProduct)
		for _, d := range items {
			products = append(products, model.Product{
				SKU:      d.Code,
				Name:     d.Title,
				Price:    d.Cost,
				Currency: d.Currency,
				Provider: providerName,
			})
		}

	case "provider_b":
		items, ok := data.([]model.ProviderBProduct)
		if !ok {
			return products
		}

		for _, d := range items {
			products = append(products, model.Product{
				SKU:      d.ID,
				Name:     d.Name,
				Price:    d.Price,
				Currency: d.Currency,
				Provider: providerName,
			})
		}

	// 🔴 Provider C
	case "provider_c":
		items, ok := data.([]model.ProviderCProduct)
		if !ok {
			return products
		}

		for _, d := range items {
			products = append(products, model.Product{
				SKU:      d.Ref,
				Name:     d.Product,
				Price:    d.AmountUSD,
				Currency: d.Currency,
				Provider: providerName,
			})
		}

	}

	return products

}
