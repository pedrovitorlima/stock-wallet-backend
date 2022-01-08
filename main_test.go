package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	walletAPI "github.com/pedrovitorlima/stock-wallet-backend/api"
	"github.com/steinfletcher/apitest"
)

func Test_CreateWallet(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/wallet", walletAPI.Create).Methods("POST")
	router.HandleFunc("/wallet", walletAPI.Get).Methods("GET")

	server := httptest.NewServer(router)
	defer server.Close()

	apitest.New().
		Handler(router).
		Post("/wallet").
		JSON(`{"id": 10, "name": "WalletTest"}`).
		Expect(t).
		Status(http.StatusOK).
		End()

}
