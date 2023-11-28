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
	"os"
)

func main() {
	// flag allowing user to decide port (:4000 default)
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// logger for informational messages
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

	// logger for error messages
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

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

	// infoLog handles info message
	infoLog.Printf("Starting server on port %s...", *addr)

	// start new web server on port 4000
	err := http.ListenAndServe(*addr, mux)

	//errorLog handles error message (if any)
	errorLog.Fatal(err)
}
