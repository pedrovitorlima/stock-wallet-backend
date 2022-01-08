package dbconfig

import "fmt"

const PostgresDriver = "postgres"
const User = "pgtest"
const Host = "localhost"

const Port = "15432"

const Password = "1234"

const DbName = "stock-wallet-db"

var DataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DbName)
