package repository

import (
	"github.com/pedrovitorlima/stock-wallet-backend/models"
	"gorm.io/gorm"
)

type StockRepository interface {
	Add(stock *models.Stock)
}

type StockRepositoryImpl struct {
	Database *gorm.DB
}

func (w StockRepositoryImpl) Add(stock *models.Stock) {
	w.Database.Create(stock)
}
