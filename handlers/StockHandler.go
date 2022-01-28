package handlers

import (
	"github.com/pedrovitorlima/stock-wallet-backend/database/repository"
)

type StockHandler struct {
	stockRepository repository.StockRepository
}

func NewStockHandler(repository repository.StockRepository) StockHandler {
	return StockHandler{stockRepository: repository}
}
