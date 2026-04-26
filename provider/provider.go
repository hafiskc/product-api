package provider

import "context"

type Provider interface {
	Search(ctx context.Context, query string) (interface{}, error)
	Name() string
}