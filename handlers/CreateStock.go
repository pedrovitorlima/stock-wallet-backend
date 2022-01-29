package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pedrovitorlima/stock-wallet-backend/errors"
	"github.com/pedrovitorlima/stock-wallet-backend/models"
)

func (handler StockHandler) CreateStock(writter http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var stock models.Stock
	json.Unmarshal(body, &stock)

	if errs := *stock.ValidateToCreate(); len(errs) != 0 {
		writter.WriteHeader(http.StatusBadRequest)

		apiErrors := errors.ApiRequestErrors{errs}
		json.NewEncoder(writter).Encode(apiErrors)
		return
	}

	handler.stockRepository.Add(&stock)

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusCreated)
}
