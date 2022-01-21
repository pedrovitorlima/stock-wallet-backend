package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pedrovitorlima/stock-wallet-backend/models"
)

type WalletRepositoryMock struct {
	AddFn func(wallet *models.Wallet)
}

func (d WalletRepositoryMock) Add(wallet *models.Wallet) {
	d.AddFn(wallet)
}

func Test_CreateWalletShouldNotCallRepositoryGivenBodyIsInvalid(t *testing.T) {

	called := false
	walletRepositoryMock := WalletRepositoryMock{
		AddFn: func(wallet *models.Wallet) {
			called = true
		},
	}

	walletHandler := NewWalletHandler(walletRepositoryMock)

	walletBody := `{ "Ida": 1, "Namea": "wallet" }`
	readerWithWalletAsBody := strings.NewReader(string(walletBody))
	req := httptest.NewRequest(http.MethodPost, "/wallet", readerWithWalletAsBody)
	writter := httptest.NewRecorder()

	walletHandler.CreateWallet(writter, req)

	if called {
		t.Error("expect to not call the repository when there are errors in the body request")
	}
}

func Test_CreateWalletShouldCallRepositoryWithExpectedWallet(t *testing.T) {
	walletToSave := models.Wallet{
		Id:   0,
		Name: "Testing wallet",
	}

	called := false
	walletRepositoryMock := WalletRepositoryMock{
		AddFn: func(wallet *models.Wallet) {
			called = true
			if wallet.Id != walletToSave.Id || wallet.Name != walletToSave.Name {
				t.Errorf("Expected wallet[%d, %s] but got wallet[%d, %s]",
					walletToSave.Id,
					walletToSave.Name,
					wallet.Id,
					wallet.Name,
				)
			}
		},
	}

	walletHandler := NewWalletHandler(walletRepositoryMock)
	req, writter := createEndpointRequestObjects(&walletToSave, t)
	walletHandler.CreateWallet(writter, req)

	if !called {
		t.Error("Expected to call the repository")
	}
}

func createEndpointRequestObjects(walletToSave *models.Wallet, t *testing.T) (*http.Request, *httptest.ResponseRecorder) {
	walletJson, err := json.Marshal(walletToSave)
	readerWithWalletAsBody := strings.NewReader(string(walletJson))

	if err != nil {
		t.Error("Error marshalling wallet")
	}

	req := httptest.NewRequest(http.MethodPost, "/wallet", readerWithWalletAsBody)
	writter := httptest.NewRecorder()
	return req, writter
}
