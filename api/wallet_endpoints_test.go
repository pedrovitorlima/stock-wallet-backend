package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	walletdb "github.com/pedrovitorlima/stock-wallet-backend/database/repository"
	mocks "github.com/pedrovitorlima/stock-wallet-backend/mocks"
)

func Test_Create(t *testing.T) {
	walletToSave := walletdb.Wallet{Id: 1, Name: "Wallet name"}

	req, writter := createEndpointRequestObjects(walletToSave, t)

	mockedRepository := mocks.WalletRepositoryMock{
		CreateFn: func(wallet *walletdb.Wallet) (*walletdb.Wallet, error) {
			if wallet.Id != walletToSave.Id || wallet.Name != walletToSave.Name {
				t.Errorf("Expecting wallet[%d, %s] but got wallet[%d, %s]",
					wallet.Id,
					wallet.Name,
					walletToSave.Id,
					walletToSave.Name)
			}
			return wallet, nil
		},
	}

	endpoint := WalletEndpoint{
		repository: mockedRepository,
	}

	endpoint.Create(writter, req)
}

func createEndpointRequestObjects(walletToSave walletdb.Wallet, t *testing.T) (*http.Request, *httptest.ResponseRecorder) {
	walletJson, err := json.Marshal(walletToSave)
	readerWithWalletAsBody := strings.NewReader(string(walletJson))

	if err != nil {
		t.Error("Error marshalling wallet")
	}

	req := httptest.NewRequest(http.MethodPost, "/wallet", readerWithWalletAsBody)
	writter := httptest.NewRecorder()
	return req, writter
}
