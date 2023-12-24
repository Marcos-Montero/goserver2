package server

import "github.com/MarcosMRod/goserver2/internal/database"

type ApiConfig struct {
	DB *database.Queries
}
func NewApiConfig(db *database.Queries) ApiConfig {
	apiCfg := ApiConfig{ DB: db }
	return apiCfg
}

