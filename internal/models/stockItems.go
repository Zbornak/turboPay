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
	TrackListing string
	Created      time.Time
	Expires      time.Time
	Format       string
	Price        int
	ReleaseDate  string
}

// wraps a sql.DB connection pool
type StockItemModel struct {
	DB *sql.DB
}

// insert new stock item
func (m *StockItemModel) Insert(title string, artist string, trackListing string, expires int, format string, price int, releaseDate string) (int, error) {
	stmt := `INSERT INTO stockItems (title, artist, trackListing, created, expires, format, price, releaseDate)
    VALUES(?, ?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY), ?, ?, ?)`

	result, err := m.DB.Exec(stmt, title, artist, trackListing, expires, format, price, releaseDate)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// return a requested stock item using id
func (m *StockItemModel) Get(id int) (*StockItem, error) {
	return nil, nil
}

// return the 10 most recent stock items
func (m *StockItemModel) Latest() ([]*StockItem, error) {
	return nil, nil
}
