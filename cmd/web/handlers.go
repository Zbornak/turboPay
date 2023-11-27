package main

import (
	"fmt"
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

	w.Write([]byte("Welcome to turboPay"))
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
