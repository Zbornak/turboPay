package main

import "net/http"

// returns servemux containing app routes
func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/stockItem/view", app.itemView)
	mux.HandleFunc("/stockItem/create", app.itemCreate)

	return mux
}
