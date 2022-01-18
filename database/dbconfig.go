package dbconfig

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pedrovitorlima/stock-wallet-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const PostgresDriver = "postgres"
const User = "pgtest"
const Host = "localhost"

const Port = "15432"

const Password = "1234"

const DbName = "stock-wallet-db"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

func GetConnection() (*sql.DB, error) {
	return sql.Open(PostgresDriver, DataSourceName)
}

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DataSourceName), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Wallet{})

	return db
}
