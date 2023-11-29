package models

import (
	"database/sql"
	"time"
)

// model for mySQL stock item
type StockItem struct {
	ID           int
	Title        string
	Artist       string
	ReleaseDate  time.Time
	TrackListing string
	Created      time.Time
	Expires      time.Time
	Format       string
	Price        int
}

// wraps a sql.DB connection pool
type StockItemModel struct {
	DB *sql.DB
}

// insert new stock item
func (m *StockItemModel) Insert(title string, artist string, releaseDate int, trackListing string, expires int, format string, price int) (int, error) {
	return 0, nil
}

// return a requested stock item using id
func (m *StockItemModel) Get(id int) (*StockItem, error) {
	return nil, nil
}

// return the 10 most recent stock items
func (m *StockItemModel) Latest() ([]*StockItem, error) {
	return nil, nil
}
