package api

import (
	"encoding/json"
	"log"
	"net/http"

	walletdb "github.com/pedrovitorlima/stock-wallet-backend/database/repository"
)

func Create(w http.ResponseWriter, request *http.Request) {
	var wallet walletdb.Wallet
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&wallet); err != nil {
		log.Fatalf("Error while decoding wallet from request: %v", err)
	}

	defer request.Body.Close()

	if err := walletdb.Create(&wallet); err != nil {
		log.Fatalf("Error while creating wallet: %v", err)
	}
}

func Get(w http.ResponseWriter, request *http.Request) {
	wallets := walletdb.FindAll()
	json.NewEncoder(w).Encode(wallets)
}
