package models

type Wallet struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
