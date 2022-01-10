package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	walletAPI "github.com/pedrovitorlima/stock-wallet-backend/api"
	wallet "github.com/pedrovitorlima/stock-wallet-backend/database/repository"
	"github.com/steinfletcher/apitest"
)

func Test_CreateWallet(t *testing.T) {
	router, server := initServerAndRouter()

	expectedWallet := wallet.Wallet{
		Id:   101,
		Name: "WalletTest",
	}
	callCreateEndpoint(router, t)
	wallets := callGetAllEndpoint(server, t)

	if !containsWallet(wallets, expectedWallet) {
		t.Error("Saved the wallet but dindnt find there after")
	}
}

func callGetAllEndpoint(server *httptest.Server, t *testing.T) []wallet.Wallet {
	res, err := http.Get(server.URL + "/wallet")
	if err != nil {
		t.Errorf("Failled to get wallets: %v", err)
	}

	var wallets []wallet.Wallet
	json.NewDecoder(res.Body).Decode(&wallets)
	return wallets
}

func callCreateEndpoint(router *mux.Router, t *testing.T) {
	apitest.New().
		Handler(router).
		Post("/wallet").
		JSON(`{"id": 101, "name": "WalletTest"}`).
		Expect(t).
		Status(http.StatusOK).
		End()
}

func initServerAndRouter() (*mux.Router, *httptest.Server) {
	router := mux.NewRouter()
	router.HandleFunc("/wallet", walletAPI.Create).Methods("POST")
	router.HandleFunc("/wallet", walletAPI.Get).Methods("GET")

	server := httptest.NewServer(router)
	defer server.Close()
	return router, server
}

func containsWallet(wallets []wallet.Wallet, walletToBeFound wallet.Wallet) bool {
	for _, walletEl := range wallets {
		if walletEl == walletToBeFound {
			log.Printf("Found wallet %s in wallets {%s}",
				walletsToString([]wallet.Wallet{walletToBeFound}),
				walletsToString(wallets))

			return true
		}
	}

	return false
}

func walletsToString(wallets []wallet.Wallet) string {
	var wallet_tostring string
	for _, wallet := range wallets {
		wallet_tostring = wallet_tostring + " " + fmt.Sprintf("[%b, %s]", wallet.Id, wallet.Name)
	}

	return wallet_tostring
}
