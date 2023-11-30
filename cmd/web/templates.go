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
		ts, err := template.ParseFiles("./ui/html/base.tmpl.html")
		if err != nil {
			return nil, err
		}

		// add any partials
		ts, err = ts.ParseGlob("./ui/html/partials/*.tmpl.html")
		if err != nil {
			return nil, err
		}

		// parse files into a template set
		ts, err = ts.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		// add template set to the map
		cache[name] = ts
	}

	// return the map
	return cache, nil
}
