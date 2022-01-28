package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

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

	if len(*stock.ValidateToCreate()) != 0 {
		writter.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.stockRepository.Add(&stock)

	writter.Header().Add("Content-Type", "application/json")
	writter.WriteHeader(http.StatusCreated)
}
