package main

import (
	// "fmt"
	// "log"
	// "net/http"
	// "net/http/httptest"
	"testing"
	// "github.com/gorilla/mux"
	// walletAPI "github.com/pedrovitorlima/stock-wallet-backend/api"
	// wallet "github.com/pedrovitorlima/stock-wallet-backend/database/repository"
	// "github.com/steinfletcher/apitest"
)

func Test_CreateWallet(t *testing.T) {
	// router, _ := initServerAndRouter()

	// expectedWallet := wallet.Wallet{
	// 	Id:   101,
	// 	Name: "WalletTest",
	// }
	// callCreateEndpointAndAssertResult(router, t, &expectedWallet)
}

// func callCreateEndpointAndAssertResult(router *mux.Router, t *testing.T, wallet *wallet.Wallet) {
// 	apitest.New().
// 		Handler(router).
// 		Post("/wallet").
// 		JSON(`{"id": 101, "name": "WalletTest"}`).
// 		Expect(t).
// 		Status(http.StatusOK).
// 		End()
// }

// func initServerAndRouter() (*mux.Router, *httptest.Server) {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/wallet", walletAPI.Create).Methods("POST")

// 	server := httptest.NewServer(router)
// 	defer server.Close()
// 	return router, server
// }

// func containsWallet(wallets []wallet.Wallet, walletToBeFound wallet.Wallet) bool {
// 	for _, walletEl := range wallets {
// 		if walletEl == walletToBeFound {
// 			log.Printf("Found wallet %s in wallets {%s}",
// 				walletsToString([]wallet.Wallet{walletToBeFound}),
// 				walletsToString(wallets))

// 			return true
// 		}
// 	}

// 	return false
// }

// func walletsToString(wallets []wallet.Wallet) string {
// 	var wallet_tostring string
// 	for _, wallet := range wallets {
// 		wallet_tostring = wallet_tostring + " " + fmt.Sprintf("[%b, %s]", wallet.Id, wallet.Name)
// 	}

// 	return wallet_tostring
// }
