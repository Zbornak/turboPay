package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

// get a stack trace for current goroutine and append to log message
func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// automatically generate text message of HTTP status code
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// sends 404 response to user
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

// helper method to render templates from cache
func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// retrieve appropriate template set based on page name
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, err)
		return
	}

	//initialise a new buffer
	buf := new(bytes.Buffer)

	// write the template to the buffer
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// write out provided HTTP status code
	w.WriteHeader(status)

	// write contents of buffer to responseWriter
	buf.WriteTo(w)
}
