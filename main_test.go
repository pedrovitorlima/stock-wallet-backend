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

func Test_CreateWalletShouldReturnErrorsGivenNameSizeBiggerThan200(t *testing.T) {
	router, _ := initServerAndRouter()

	bigName := "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Neque fuga ducimus dignissimos laborum sed, porro, nemo accusamus non nihil itaque ipsum recusandae nulla eius officia dolorem. Quisquam autem placeat quae fuga quasi, corrupti atque totam qui unde amet et at sint pariatur nam ipsam distinctio quo laudantium voluptates nobis mollitia vitae. Quaerat itaque quibusdam est amet eum excepturi, tempora, eaque distinctio, dolorem quo pariatur consectetur dolorum sed animi expedita placeat aliquid labore quae voluptatum dicta doloremque fugiat! Obcaecati accusantium nemo facere quam distinctio quibusdam, praesentium blanditiis, eius amet iste, doloribus at. Dolorem beatae consequuntur totam aliquid, temporibus officiis nihil, mollitia id pariatur non, expedita nisi nemo corrupti aperiam. Qui voluptate temporibus adipisci voluptatem illum magni nisi ratione at, culpa minus dolore, excepturi fugit corrupti deserunt quo totam ducimus, vero ex doloribus ad quos! Dolore voluptatum unde quia deserunt dolorum. Sit facere accusamus nihil, quas ratione magni obcaecati, atque ab, commodi earum doloremque vitae suscipit! Illo eaque assumenda qui quia, amet quae voluptatem debitis. Sed architecto sint porro ad dolorem eligendi placeat voluptatum quo asperiores perspiciatis, maxime fugiat, voluptatem facere soluta et alias ipsam. Exercitationem numquam illo ratione rerum vero eum, molestiae odit non tempora, repellendus ea consequuntur fugiat. Expedita, dolorem quos."
	bodyRequest := fmt.Sprintf(`{"Id": 0, "Name": "%s"}`, bigName)
	expectedBody := fmt.Sprintf(`{"errors": [{"field":"name", "description": "Name size should not be bigger than 200"}]}`)

	apitest.New().
		Handler(router).
		Post("/wallet").
		JSON(bodyRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		Body(expectedBody).
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
