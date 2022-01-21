package dbconfig

import (
	"fmt"
	"log"

	"github.com/pedrovitorlima/stock-wallet-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const PostgresDriver = "postgres"
const SqlLiteDriver = "sqllite"
const User = "pgtest"
const Host = "localhost"

const Port = "15432"

const Password = "1234"

const DbName = "stock-wallet-db"

var DataSourceNamePostgres = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)

func Init() *gorm.DB {
	db, err := gorm.Open(postgres.Open(DataSourceNamePostgres), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Wallet{})

	return db
}

func InitSqlLiteForTest() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Wallet{})

	return db
}
