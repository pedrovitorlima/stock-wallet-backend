package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pedrovitorlima/stock-wallet-backend/models"
)

type StockRepositoryMock struct {
	AddFn func(stock *models.Stock)
}

func (mock StockRepositoryMock) Add(stock *models.Stock) {
	mock.AddFn(stock)
}

func Test_CreateStockShouldReturnHttpCreatedGivenInputIsValid(t *testing.T) {
	stock := models.Stock{
		Quantity: 10,
		IdWallet: 1,
	}

	writter := callCreateEndpoint(&stock, t)

	if writter.Code != http.StatusCreated {
		t.Errorf("Expected status Created (%d) but got (%d)", http.StatusCreated, writter.Code)
	}
}

func Test_CreateStockShouldReturnHttpBadRequestGivenInputIsInvalid(t *testing.T) {
	stock := models.Stock{
		Quantity: 0,
		IdWallet: 1,
	}

	writter := callCreateEndpoint(&stock, t)

	if writter.Code != http.StatusBadRequest {
		t.Errorf("Expected status Bad Request (%d) but got (%d)", http.StatusCreated, writter.Code)
	}
}

func Test_CreateStockShouldCallRepositoryGivenInputIsValid(t *testing.T) {
	stock := models.Stock{
		Quantity: 1,
		IdWallet: 1,
	}

	mockCalled := false
	repository := StockRepositoryMock{
		AddFn: func(expectedStock *models.Stock) {
			if expectedStock.Quantity == stock.Quantity && expectedStock.IdWallet == stock.IdWallet {
				mockCalled = true
			}
		},
	}

	createStockHandler := NewStockHandler(repository)
	req, writter := createStockEndpointRequestObjects(&stock, t)

	createStockHandler.CreateStock(writter, req)

	if !mockCalled {
		t.Error("Expected to call the repository given inputs are valid but it didnt")
	}
}

func Test_CreateStockShouldNotCallRepositoryGivenInputIsInvalid(t *testing.T) {
	stock := models.Stock{
		Quantity: 1,
	}

	mockCalled := false
	repository := StockRepositoryMock{
		AddFn: func(expectedStock *models.Stock) {
			if expectedStock.Quantity == stock.Quantity && expectedStock.IdWallet == stock.IdWallet {
				mockCalled = true
			}
		},
	}

	callCreateEndpointWithMock(&stock, t, repository)

	if mockCalled {
		t.Error("Expected to not call the repository given inputs are valid but it was")
	}
}

func callCreateEndpointWithMock(stock *models.Stock, t *testing.T, mock StockRepositoryMock) httptest.ResponseRecorder {
	createStockHandler := NewStockHandler(mock)

	req, writter := createStockEndpointRequestObjects(stock, t)

	createStockHandler.CreateStock(writter, req)
	return *writter
}

func callCreateEndpoint(stock *models.Stock, t *testing.T) httptest.ResponseRecorder {
	mock := StockRepositoryMock{
		AddFn: func(stock *models.Stock) {},
	}

	return callCreateEndpointWithMock(stock, t, mock)
}

func createStockEndpointRequestObjects(objectToSave *models.Stock, t *testing.T) (*http.Request, *httptest.ResponseRecorder) {
	walletJson, err := json.Marshal(objectToSave)
	readerWithWalletAsBody := strings.NewReader(string(walletJson))

	if err != nil {
		t.Error("Error marshalling stock")
	}

	req := httptest.NewRequest(http.MethodPost, "/wallet/stock", readerWithWalletAsBody)
	writter := httptest.NewRecorder()
	return req, writter
}
