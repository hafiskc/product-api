package model

import "time"

type SearchHistory struct {
	ID              string    `json:"id"`
	Query           string    `json:"query"`
	Timestamp       time.Time `json:"timestamp"`
	ResultCount     int       `json:"result_count"`
	FailedProviders []string  `json:"failed_providers"`
}