package unit_tests

import (
	"std/model"
	"testing"
)

func TestClean(t *testing.T) {

	var err error = nil

	expected, err := int64(100), err

	if err != nil {
		return
	}

	got, err := model.Clean()

	if got != expected {
		t.Errorf("expected %v, got %v", expected, got)
	}

}
