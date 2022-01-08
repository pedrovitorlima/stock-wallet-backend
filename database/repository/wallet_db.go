package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	dbconfig "github.com/pedrovitorlima/stock-wallet-backend/database"
)

type Wallet struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var db *sql.DB
var err error

func FindAll() []Wallet {

	db, err = dbconfig.GetConnection()
	checkError(err)

	sql, err := db.Query("SELECT * FROM public.wallet")
	checkError(err)

	var wallets []Wallet

	for sql.Next() {
		var wallet Wallet

		err = sql.Scan(&wallet.Id, &wallet.Name)
		checkError(err)

		wallets = append(wallets, wallet)
	}

	return wallets
}

func Create(wallet *Wallet) error {
	log.Println(fmt.Sprintf(`
		Adding a new wallet into the database 
		[idWallet=%d, walletName=%s]
	`, wallet.Id, wallet.Name))

	db, err = dbconfig.GetConnection()
	if err != nil {
		return err
	}

	sqlStatement := "INSERT INTO public.wallet(idWallet, name) VALUES($1, $2)"
	sql, err := db.Prepare(sqlStatement)
	if err != nil {
		return err
	}

	result, err := sql.Exec(wallet.Id, wallet.Name)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Println(affect)
	if err != nil {
		return err
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("Error while fetching Wallet information from database: %v", err)
	}
}
