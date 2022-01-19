package handlers

import (
	"encoding/json"
	"fmt"
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

	if err := validate(&wallet); err != nil {
		log.Println(err)

		writter.Header().Add("Content-Type", "application/json")
		writter.WriteHeader(http.StatusBadRequest)
	} else {
		h.walletRepository.Add(&wallet)

		writter.Header().Add("Content-Type", "application/json")
		writter.WriteHeader(http.StatusCreated)
		json.NewEncoder(writter).Encode("Created")
	}
}

func validate(wallet *models.Wallet) error {
	if wallet.Id > 0 || wallet.Name == "" {
		return fmt.Errorf("invalid body for wallet creation endpoint: wallet[%d, %s]", wallet.Id, wallet.Name)
	}

	return nil
}
