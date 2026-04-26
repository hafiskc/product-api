package model


type SearchResponse struct {
    Query           string    `json:"query"`
    Products        []Product `json:"products"`
    FailedProviders []string  `json:"failed_providers"`
}