/*
A web point-of-sales application
Author: Mark Strijdom (zbornak)
Date: 27/11/2023
*/

package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// flag allowing user to decide port (:4000 default)
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

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

	log.Printf("Starting server on port %s...", *addr)

	// start new web server on port 4000
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
