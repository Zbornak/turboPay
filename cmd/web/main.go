/*
A web point-of-sales application
Author: Mark Strijdom (zbornak)
Date: 27/11/2023
*/

package main

import (
	"log"
	"net/http"
)

func main() {
	// router
	mux := http.NewServeMux()

	// create fileserver for static directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// register fileserver for all /static/ paths
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register all other paths
	mux.HandleFunc("/", home)
	mux.HandleFunc("/stockItem/view", itemView)
	mux.HandleFunc("/stockItem/create", itemCreate)

	log.Print("Starting server on :4000...")

	// start new web server on port 4000
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
