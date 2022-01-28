package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pedrovitorlima/stock-wallet-backend/errors"
	"github.com/pedrovitorlima/stock-wallet-backend/models"
)

func (h WalletHandler) CreateWallet(writter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var wallet models.Wallet
	json.Unmarshal(body, &wallet)

	writter.Header().Add("Content-Type", "application/json")

	if validationErrors := wallet.ValidateToCreate(); len(*wallet.ValidateToCreate()) > 0 {
		apiErrors := errors.ApiRequestErrors{*validationErrors}
		writter.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(writter).Encode(apiErrors)
	} else {
		h.walletRepository.Add(&wallet)
		writter.WriteHeader(http.StatusCreated)
		json.NewEncoder(writter).Encode("Created")
		log.Printf("Created anoter wallet[%d, %s]", wallet.Id, wallet.Name)
	}
}
