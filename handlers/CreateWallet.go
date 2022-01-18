package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

	h.walletRepository.Add(&wallet)

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusCreated)
	json.NewEncoder(writter).Encode("Created")

}
