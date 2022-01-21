package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	dbconfig "github.com/pedrovitorlima/stock-wallet-backend/database"
	"github.com/pedrovitorlima/stock-wallet-backend/database/repository"
	"github.com/pedrovitorlima/stock-wallet-backend/handlers"
	"github.com/pedrovitorlima/stock-wallet-backend/models"
	"github.com/steinfletcher/apitest"
)

func Test_CreateWalletShouldReturnCreatedAsStatusCodeGivenValidBody(t *testing.T) {
	router, _ := initServerAndRouter()

	wallet := models.Wallet{
		Id:   0,
		Name: "WalletTest",
	}

	requestBody, _ := json.Marshal(wallet)

	apitest.New().
		Handler(router).
		Post("/wallet").
		JSON(requestBody).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func Test_CreateWalletShouldReturnBadRequestGivenInvalidBody(t *testing.T) {
	router, _ := initServerAndRouter()

	//It should not accept a request to create an object with id <> 0
	bodyRequest := `{"Id": 1, "Nameaa": "invalid name"}`

	apitest.New().
		Handler(router).
		Post("/wallet").
		JSON(bodyRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		End()

}

func initServerAndRouter() (*mux.Router, *httptest.Server) {
	router := mux.NewRouter()

	DB := dbconfig.InitSqlLiteForTest()
	repository := repository.WalletRepositoryImpl{Database: DB}

	walletHandler := handlers.NewWalletHandler(repository)
	router.HandleFunc("/wallet", walletHandler.CreateWallet).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()
	return router, server
}

func containsWallet(wallets []models.Wallet, walletToBeFound models.Wallet) bool {
	for _, walletEl := range wallets {
		if walletEl == walletToBeFound {
			log.Printf("Found wallet %s in wallets {%s}",
				walletsToString([]models.Wallet{walletToBeFound}),
				walletsToString(wallets))

			return true
		}
	}

	return false
}

func walletsToString(wallets []models.Wallet) string {
	var wallet_tostring string
	for _, wallet := range wallets {
		wallet_tostring = wallet_tostring + " " + fmt.Sprintf("[%b, %s]", wallet.Id, wallet.Name)
	}

	return wallet_tostring
}
