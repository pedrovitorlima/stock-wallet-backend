package models

import (
	"testing"

	"github.com/pedrovitorlima/stock-wallet-backend/test_utils"
)

func Test_ShouldReturnErrorGivenNameIsTooBig(t *testing.T) {
	bigName := "Lorem ipsum dolor sit, amet consectetur adipisicing elit. Neque fuga ducimus dignissimos laborum sed, porro, nemo accusamus non nihil itaque ipsum recusandae nulla eius officia dolorem. Quisquam autem placeat quae fuga quasi, corrupti atque totam qui unde amet et at sint pariatur nam ipsam distinctio quo laudantium voluptates nobis mollitia vitae. Quaerat itaque quibusdam est amet eum excepturi, tempora, eaque distinctio, dolorem quo pariatur consectetur dolorum sed animi expedita placeat aliquid labore quae voluptatum dicta doloremque fugiat! Obcaecati accusantium nemo facere quam distinctio quibusdam, praesentium blanditiis, eius amet iste, doloribus at. Dolorem beatae consequuntur totam aliquid, temporibus officiis nihil, mollitia id pariatur non, expedita nisi nemo corrupti aperiam. Qui voluptate temporibus adipisci voluptatem illum magni nisi ratione at, culpa minus dolore, excepturi fugit corrupti deserunt quo totam ducimus, vero ex doloribus ad quos! Dolore voluptatum unde quia deserunt dolorum. Sit facere accusamus nihil, quas ratione magni obcaecati, atque ab, commodi earum doloremque vitae suscipit! Illo eaque assumenda qui quia, amet quae voluptatem debitis. Sed architecto sint porro ad dolorem eligendi placeat voluptatum quo asperiores perspiciatis, maxime fugiat, voluptatem facere soluta et alias ipsam. Exercitationem numquam illo ratione rerum vero eum, molestiae odit non tempora, repellendus ea consequuntur fugiat. Expedita, dolorem quos."
	wallet := Wallet{1, bigName}
	expectedDescription := "Name size should not be bigger than 200"

	expectErrorToExist(wallet, "name", expectedDescription, t)
}

func Test_ShouldReturnErrorGivenNameIsNotPresent(t *testing.T) {
	wallet := Wallet{1, ""}
	expectedDescription := "Name cannot be empty"

	expectErrorToExist(wallet, "name", expectedDescription, t)

}

func Test_ShouldReturnErrorGivenIdIsFilledWithNumberDifferentFromZero(t *testing.T) {
	wallet := Wallet{1, ""}
	expectedDescription := "Id cannot have value for Create"

	expectErrorToExist(wallet, "id", expectedDescription, t)
}

func expectErrorToExist(wallet Wallet, field string, expectedError string, t *testing.T) {

	errors := wallet.ValidateToCreate()

	if len(*errors) == 0 {
		t.Error("Should have errors returned")
	}

	if test_utils.NotContainError(*errors, field, expectedError) {
		t.Errorf("Should contain error {%s, %s}", field, expectedError)
	}
}
