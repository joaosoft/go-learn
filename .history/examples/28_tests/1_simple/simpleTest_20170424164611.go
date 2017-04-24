package simple_test

import (
	"testing"
)

func TestDummyError(t *testing.T) {
	id := "1"
	e := "2"

	if e != id {
		t.Error("erro")
	}
}

func TestDummySuccess(t *testing.T) {
	id := "1"
	e := id

	if e != id {
		t.Error("erro")
	}
}
