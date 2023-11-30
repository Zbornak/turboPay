package main

import "turboPay/internal/models"

type templateData struct {
	StockItem  *models.StockItem
	StockItems []*models.StockItem
}
