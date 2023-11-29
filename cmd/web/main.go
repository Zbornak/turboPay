/*
A web point-of-sales application
Author: Mark Strijdom (zbornak)
Date: 27/11/2023
*/

package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"turboPay/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

// struct to hold app-wide dependencies
type application struct {
	errorLog   *log.Logger
	infoLog    *log.Logger
	stockItems *models.StockItemModel
}

func main() {
	// flag allowing user to decide port (:4000 default)
	addr := flag.String("addr", ":4000", "HTTP network address")
	// flag for mySQL DSN string
	dsn := flag.String("dsn", "web:zb0rnak_7137_@/turboPay?parseTime=true", "MySQL data source name")
	flag.Parse()

	// logger for informational messages
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)

	// logger for error messages
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// pass openDB() the DSN from flag
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// close connection pool before exiting main function
	defer db.Close()

	// initialise new instance of application struct
	app := &application{
		errorLog:   errorLog,
		infoLog:    infoLog,
		stockItems: &models.StockItemModel{DB: db},
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
		Handler:  app.routes(), // routes
	}

	// infoLog handles info message
	infoLog.Printf("Starting server on port %s...", *addr)

	// start new web server on port 4000
	err = srv.ListenAndServe()

	//errorLog handles error message (if any)
	errorLog.Fatal(err)
}

// wrap sql.Open and return a swl.DB connection pool for supplied DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
