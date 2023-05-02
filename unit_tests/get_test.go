package unit_tests

import (
	"reflect"
	"std/entities"
	"std/model"
	"testing"
)

func TestGetOneDb(t *testing.T) {

	var err error = nil

	crypto := &entities.Crypto{}

	expected, err := crypto, err

	if err != nil {
		return
	}

	got, err := model.GetOneDb("1")

	equal := reflect.DeepEqual(crypto, got)

	if equal != true {
		t.Errorf("expected %v, got %v", expected, got)
	}

	// if got != expected {
	// 	t.Errorf("expected %v, got %v", expected, got)
	// }

}
