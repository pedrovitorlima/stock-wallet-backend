package models

import (
	"testing"

	"github.com/pedrovitorlima/stock-wallet-backend/test_utils"
)

func Test_ShouldNotReturnErrorsGivenValidStock(t *testing.T) {
	stock := Stock{
		Id:       0,
		Quantity: 10,
		IdWallet: 1,
	}

	validationErrors := stock.ValidateToCreate()

	if len(*validationErrors) > 0 {
		t.Error("Should not return errors")
	}
}

func Test_ShouldReturnErrorGivenIdIsFilled(t *testing.T) {
	stock := Stock{
		Id: 1,
	}

	validationErrors := stock.ValidateToCreate()
	expectedError := "Id should not has value"

	if test_utils.NotContainError(*validationErrors, "id", expectedError) {
		t.Errorf("expected to find an error like [%s]", expectedError)
	}

}

func Test_ShouldReturnErrorGivenQuantityIsLessOrEqualsToZero(t *testing.T) {
	stock := Stock{
		Quantity: 0,
		IdWallet: 1,
	}

	validationErrors := stock.ValidateToCreate()
	expectedError := "0 is not a valid number for quantity"

	if test_utils.NotContainError(*validationErrors, "quantity", expectedError) {
		t.Errorf("expected to find an error like [%s]", expectedError)
	}
}

func Test_ShouldReturnErrorGivenQuantityIsNotSupplied(t *testing.T) {
	stock := Stock{
		Quantity: 0,
	}

	validationErrors := stock.ValidateToCreate()
	expectedError := "No wallet id supplied"

	if test_utils.NotContainError(*validationErrors, "idWallet", expectedError) {
		t.Errorf("expected to find an error like [%s]", expectedError)
	}
}
