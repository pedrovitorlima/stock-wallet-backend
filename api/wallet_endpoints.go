package api

import (
	"encoding/json"
	"log"
	"net/http"

	walletdb "github.com/pedrovitorlima/stock-wallet-backend/database/repository"
)

type WalletEndpoint struct {
	repository walletdb.WalletRepository
}

func (endpoint WalletEndpoint) Create(w http.ResponseWriter, request *http.Request) {
	var wallet walletdb.Wallet
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&wallet); err != nil {
		log.Fatalf("Error while decoding wallet from request: %v", err)
	}

	defer request.Body.Close()

	if _, err := endpoint.repository.Create(&wallet); err != nil {
		log.Fatalf("Error while creating wallet: %v", err)
	}
}
