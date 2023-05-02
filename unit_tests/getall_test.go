package unit_tests

import (
	"reflect"
	"std/entities"
	"std/model"
	"testing"
)

func TestGetAllInDb(t *testing.T) {

	var err error = nil

	cryptos := []entities.Crypto{}

	expected, err := cryptos, err

	if err != nil {
		return
	}

	got, err := model.GetAllInDB()

	equal := reflect.DeepEqual(cryptos, got)

	if equal != true {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
