package main

import (
	"blogAgg/internal/config"
	"blogAgg/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}
