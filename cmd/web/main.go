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

// struct to hold app-wide dependencies
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// flag allowing user to decide port (:4000 default)
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// logger for informational messages
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

	// logger for error messages
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// initialise new instance of application struct
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// router
	mux := http.NewServeMux()

	// create fileserver for static directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// register fileserver for all /static/ paths
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// register all other paths
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/stockItem/view", app.itemView)
	mux.HandleFunc("/stockItem/create", app.itemCreate)

	// new server to log errors with errorlog instead of default logger
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	// infoLog handles info message
	infoLog.Printf("Starting server on port %s...", *addr)

	// start new web server on port 4000
	err := srv.ListenAndServe()

	//errorLog handles error message (if any)
	errorLog.Fatal(err)
}
