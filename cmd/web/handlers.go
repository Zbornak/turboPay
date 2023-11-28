package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// handlers
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 404 if URL path doesn't match "/"
	if r.URL.Path != "/" {
		app.notFound(w) // helper
		return
	}

	// slice containing path to base and home templates
	files := []string{
		"./ui/html/base.tmpl.html",
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	// read HTML template
	ts, err := template.ParseFiles(files...) // ...variadic
	if err != nil {
		app.serverError(w, err) // helper
		return
	}

	// write content of base HTML template
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, err) // helper
	}
}

// view stock item
func (app *application) itemView(w http.ResponseWriter, r *http.Request) {
	// allow for user id item query, checking to make sure user enters a valid uint
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // helper
		return
	}

	fmt.Fprintf(w, "Display a specific stock item with ID %d...", id)
}

// create new stock item
func (app *application) itemCreate(w http.ResponseWriter, r *http.Request) {
	// method not allowed (405) if request method isn't POST
	if r.Method != http.MethodPost {
		// add 'Allow:POST' to response header map to let user know what request is allowed
		w.Header().Set("Allow", http.MethodPost)

		app.clientError(w, http.StatusMethodNotAllowed) // helper
		return
	}

	w.Write([]byte("Create a new stock item..."))
}
