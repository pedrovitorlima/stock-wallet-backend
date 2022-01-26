package models

import (
	"github.com/pedrovitorlima/stock-wallet-backend/errors"
)

type Wallet struct {
	Id   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

func (w *Wallet) Validate() *[]errors.ValidationErrors {
	errs := []errors.ValidationErrors{}
	if w.Name == "" {
		err := errors.ValidationErrors{"name", "Name cannot be empty"}
		errs = append(errs, err)
	}

	if len(w.Name) > 200 {
		err := errors.ValidationErrors{"name", "Name size should not be bigger than 200"}
		errs = append(errs, err)
	}

	return &errs
}
