package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	dbconfig "github.com/pedrovitorlima/stock-wallet-backend/database"
	"github.com/pedrovitorlima/stock-wallet-backend/database/repository"
	"github.com/pedrovitorlima/stock-wallet-backend/handlers"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// Set headers
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		fmt.Println("ok")

		// Next
		next.ServeHTTP(w, r)
	})
}

func main() {
	DB := dbconfig.Init()
	repository := repository.WalletRepositoryImpl{Database: DB}
	walletHandler := handlers.NewWalletHandler(repository)
	router := mux.NewRouter()

	router.Use(CORS)

	router.HandleFunc("/wallet", walletHandler.CreateWallet).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc("/wallet/stocks", walletHandler.CreateWallet).Methods(http.MethodPost, http.MethodOptions)

	err := http.ListenAndServe(":4000", router)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("API is running!")
}
