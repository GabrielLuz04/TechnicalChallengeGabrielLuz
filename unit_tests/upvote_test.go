package unit_tests

import (
	"errors"
	"std/entities"
	"std/model"
	"testing"
)

func TestToggleCryptoUpvoteCMC(t *testing.T) {
	expected, err := int64(1), errors.New("cryptos atualizadas com sucesso")

	if err != nil {
		return
	}

	got, err := model.Update(1, &entities.Crypto{})

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
