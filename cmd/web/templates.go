package main

import (
	"html/template"
	"path/filepath"
	"turboPay/internal/models"
)

type templateData struct {
	StockItem  *models.StockItem
	StockItems []*models.StockItem
}

// map to cache parsed templates
func newTemplateCache() (map[string]*template.Template, error) {
	// initialise new map
	cache := map[string]*template.Template{}

	// slice of all filepaths for page templates
	pages, err := filepath.Glob("./ui/html/pages/*.html")
	if err != nil {
		return nil, err
	}

	// loop through filepaths
	for _, page := range pages {
		name := filepath.Base(page)
		files := []string{
			"./ui/html/base.tmpl.html",
			"./ui/html/partials/nav.tmpl.html",
			page,
		}

		// parse files into a template set
		ts, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		// add template set to the map
		cache[name] = ts
	}

	// return the map
	return cache, nil
}
