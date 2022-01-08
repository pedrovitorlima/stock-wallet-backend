package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	walletdb "github.com/pedrovitorlima/stock-wallet-backend/database/store"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/wallet", CreateWallet).Methods("POST")
	router.HandleFunc("/wallet", GetWallet).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}

func CreateWallet(w http.ResponseWriter, request *http.Request) {

}

func GetWallet(w http.ResponseWriter, request *http.Request) {
	wallets := walletdb.FindAll()
	json.NewEncoder(w).Encode(wallets)

}
