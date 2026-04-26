package model

type Product struct {
    SKU      string  `json:"sku"`
    Name     string  `json:"name"`
    Price    float64 `json:"price"`
    Currency string  `json:"currency"`
    Provider string  `json:"provider"`
}