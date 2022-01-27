package test_utils

import "github.com/pedrovitorlima/stock-wallet-backend/errors"

func NotContainError(validationErrors []errors.ValidationErrors, field string, description string) bool {
	for _, err := range validationErrors {
		if err.Field == field && err.Description == description {
			return false
		}
	}

	return true
}
