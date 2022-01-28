package models

import "github.com/pedrovitorlima/stock-wallet-backend/errors"

type Stock struct {
	Id       int    `json:"id" gorm:"primaryKey"`
	Quantity int    `json:"quantity"`
	IdWallet int    `json:"idWallet"`
	wallet   Wallet `json:"wallet" gorm:"foreignKey:idWallet"`
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

	if stock.IdWallet <= 0 {
		err := errors.ValidationErrors{"idWallet", "No wallet id supplied"}
		errs = append(errs, err)
	}

	return &errs
}
