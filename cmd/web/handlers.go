package main

import (
	"errors"
	"fmt"

	"net/http"
	"strconv"
	"turboPay/internal/models"
)

// handlers
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// 404 if URL path doesn't match "/"
	if r.URL.Path != "/" {
		app.notFound(w) // helper
		return
	}

	stockItems, err := app.stockItems.Latest()
	if err != nil {
		app.serverError(w, err)
		return
	}

	// render helper
	app.render(w, http.StatusOK, "home.tmpl.html", &templateData{
		StockItems: stockItems,
	})
}

// view stock item
func (app *application) itemView(w http.ResponseWriter, r *http.Request) {
	// allow for user id item query, checking to make sure user enters a valid uint
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w) // helper
		return
	}

	stockItem, err := app.stockItems.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}

		return
	}

	// render helper
	app.render(w, http.StatusOK, "view.tmpl.html", &templateData{
		StockItem: stockItem,
	})
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

	// dummy data
	title := "Ten New Songs"
	artist := "Leonard Cohen"
	trackListing := "In My Secret Life\nA Thousand Kisses Deep\nThat Dont Make It Junk\nHere It Is\nLove Itself\nBy The Rivers Dark\nAlexandra Leaving\nYou Have Loved Enough\nBoogie Street\nThe Land Of Plenty"
	expires := 7
	format := "VINYL"
	price := 24
	releaseDate := "09/10/2011"

	// pass data to model
	id, err := app.stockItems.Insert(title, artist, trackListing, expires, format, price, releaseDate)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// redirect user to relevant stock item
	http.Redirect(w, r, fmt.Sprintf("/stockItem/view?id=%d", id), http.StatusSeeOther)
}
