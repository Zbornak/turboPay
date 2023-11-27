package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

// handlers
func home(w http.ResponseWriter, r *http.Request) {
	// 404 if URL path doesn't match "/"
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// read HTML template
	ts, err := template.ParseFiles("./ui/html/pages/home.tmpl.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// write HTML template
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

// view stock item
func itemView(w http.ResponseWriter, r *http.Request) {
	// allow for user id item query, checking to make sure user enters a valid uint
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific stock item with ID %d...", id)
}

// create new stock item
func itemCreate(w http.ResponseWriter, r *http.Request) {
	// method not allowed (405) if request method isn't POST
	if r.Method != http.MethodPost {
		// add 'Allow:POST' to response header map to let user know what request is allowed
		w.Header().Set("Allow", "POST")

		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Create a new stock item..."))
}
