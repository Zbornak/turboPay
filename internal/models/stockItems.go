package models

import (
	"database/sql"
	"errors"
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
	// ? -> placeholder for values
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
	stmt := `SELECT id, title, artist, trackListing, created, expires, format, price, releaseDate FROM stockItems
    WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// execute stmt, passing in id and returning a pointer to a sql.Row object
	row := m.DB.QueryRow(stmt, id)

	// initlialise pointer to a new blank StockItem struct
	s := &StockItem{}

	err := row.Scan(&s.ID, &s.Title, &s.Artist, &s.TrackListing, &s.Created, &s.Expires, &s.Format, &s.Price, &s.ReleaseDate)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

// return the 10 most recent stock items
func (m *StockItemModel) Latest() ([]*StockItem, error) {
	stmt := `SELECT id, title, artist, trackListing, created, expires, format, price, releaseDate FROM stockItems
    WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// return sql.Rows resultset from query
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// close resultset before Latest() returns
	// if not closed the underlying db connection will stay open
	defer rows.Close()

	//empty slice to contain stockItem structs
	stockItems := []*StockItem{}

	// iterate over resultset
	for rows.Next() {
		s := &StockItem{}
		err = rows.Scan(&s.ID, &s.Title, &s.Artist, &s.TrackListing, &s.Created, &s.Expires, &s.Format, &s.Price, &s.ReleaseDate)
		if err != nil {
			return nil, err
		}

		// append to slice
		stockItems = append(stockItems, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return stockItems, nil
}
