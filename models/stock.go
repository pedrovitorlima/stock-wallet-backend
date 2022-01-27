package models

import "github.com/pedrovitorlima/stock-wallet-backend/errors"

type Stock struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Quantity int    `json:"quantity"`
	IdWallet int    `json:"idWallet"`
	wallet   Wallet `json:"id" gorm:"primaryKey"`
}

func (stock *Stock) ValidateToCreate() *[]errors.ValidationErrors {
	errs := []errors.ValidationErrors{}

	if stock.Id != 0 {
		err := errors.ValidationErrors{"id", "Id should not has value"}
		errs = append(errs, err)
	}

	if stock.Quantity <= 0 {
		err := errors.ValidationErrors{"quantity", "Quantity needs to be a positive number"}
		errs = append(errs, err)
	}

	return &errs
}
