package main

import "github.com/musictopeople/go-postgres-service/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}
