package main

import "snippetbox.beslan.net/internal/models"

type templateData struct {
	Snippet  models.Snippet
	Snippets []models.Snippet
}