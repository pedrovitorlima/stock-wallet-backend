package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dbconfig "github.com/pedrovitorlima/stock-wallet-backend/database"
	"github.com/pedrovitorlima/stock-wallet-backend/database/repository"
	"github.com/pedrovitorlima/stock-wallet-backend/handlers"
)

func main() {
	DB := dbconfig.Init()
	repository := repository.WalletRepositoryImpl{Database: DB}
	walletHandler := handlers.NewWalletHandler(repository)
	router := mux.NewRouter()

	router.HandleFunc("/wallet", walletHandler.CreateWallet).Methods(http.MethodPost)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)

}
