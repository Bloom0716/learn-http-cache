package handler

import "github.com/Bloom0716/learn-http-cache/internal/database"

type Provider struct {
	Repository database.Repository
}

func NewProvider(repository database.Repository) *Provider {
	return &Provider{Repository: repository}
}
