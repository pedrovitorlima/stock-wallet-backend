package database

import (
	"database/sql"
	"log"

	dbconfig "github.com/pedrovitorlima/stock-wallet-backend/database"

	_ "github.com/lib/pq"
)

type Wallet struct {
	idWallet int    `json:"id,omitempty"`
	name     string `json:"name,omitempty"`
}

var db *sql.DB
var err error

func FindAll() []Wallet {

	log.Printf(dbconfig.DataSourceName)
	db, err = sql.Open(dbconfig.PostgresDriver, dbconfig.DataSourceName)

	sql, err := db.Query("SELECT * FROM public.wallet")
	checkError(err)

	var wallets []Wallet

	for sql.Next() {
		var wallet Wallet

		err = sql.Scan(&wallet.idWallet, &wallet.name)
		checkError(err)

		wallets = append(wallets, wallet)
	}

	return wallets
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error while fetching Wallet information from database: %v", err)
	}
}
