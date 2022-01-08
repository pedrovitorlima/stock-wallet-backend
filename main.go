package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	walletAPI "github.com/pedrovitorlima/stock-wallet-backend/api"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/wallet", walletAPI.Create).Methods("POST")
	router.HandleFunc("/wallet", walletAPI.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))

}
