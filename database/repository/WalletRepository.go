package repository

import (
	"github.com/pedrovitorlima/stock-wallet-backend/models"
	"gorm.io/gorm"
)

type WalletRepository interface {
	Add(wallet *models.Wallet)
}

type WalletRepositoryImpl struct {
	Database *gorm.DB
}

func (w WalletRepositoryImpl) Add(wallet *models.Wallet) {
	w.Database.Create(wallet)
}
