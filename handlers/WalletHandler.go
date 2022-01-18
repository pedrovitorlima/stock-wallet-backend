package handlers

import (
	"github.com/pedrovitorlima/stock-wallet-backend/database/repository"
)

type WalletHandler struct {
	walletRepository repository.WalletRepository
}

func NewWalletHandler(repository repository.WalletRepository) WalletHandler {
	return WalletHandler{walletRepository: repository}
}
