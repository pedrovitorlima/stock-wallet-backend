package main

import (
	"encoding/json"
	"fmt"
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

func Test_CreateStockShouldReturnCreatedAsStatusCodeGivenValidBody(t *testing.T) {
	router, _ := initServerAndRouterWithStockEndpoints()

	stock := models.Stock{
		Id:       0,
		Quantity: 10,
		IdWallet: 1,
	}

	requestBody, _ := json.Marshal(stock)

	apitest.New().
		Handler(router).
		Post("/wallet/stock").
		JSON(requestBody).
		Expect(t).
		Status(http.StatusCreated).
		End()
}

func Test_CreateStockShouldReturnErrorsGivenInvalidRequest(t *testing.T) {
	router, _ := initServerAndRouterWithStockEndpoints()

	bodyRequest := fmt.Sprintf(`{"id": 0, "quantity": %d, "idWallet": %d}`, -10, 0)

	expectedBody := `{"field":"quantity","description":"-10 is not a valid number for quantity"},`
	expectedBody += `{"field":"idWallet","description":"No wallet id supplied"}`

	apitest.New().
		Handler(router).
		Post("/wallet/stock").
		JSON(bodyRequest).
		Expect(t).
		Status(http.StatusBadRequest).
		Bodyf(`{"errors":[%s]}`, expectedBody).
		End()
}

func initServerAndRouterWithStockEndpoints() (*mux.Router, *httptest.Server) {
	router := mux.NewRouter()

	DB := dbconfig.InitSqlLiteForTest()
	stockRepository := repository.StockRepositoryImpl{Database: DB}

	stockHandler := handlers.NewStockHandler(stockRepository)

	router.HandleFunc("/wallet/stock", stockHandler.CreateStock).Methods("POST")

	server := httptest.NewServer(router)
	defer server.Close()
	return router, server
}
