package unit_tests

import (
	"errors"
	"std/entities"
	"std/model"
	"testing"
)

func TestCreate(t *testing.T) {
	expected, err := int64(0), errors.New("cryptos inseridas com sucesso")

	if err != nil {
		return
	}

	got, err := model.InsertCryptosInDB(&entities.Crypto{})

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
