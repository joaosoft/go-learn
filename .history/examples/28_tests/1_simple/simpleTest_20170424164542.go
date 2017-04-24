package simple_test

import (
	"testing"
)

func TestError(t *testing.T) {
	id := "1"
	e := "2"

	if e != id {
		t.Error("erro")
	}
}

func TestSuccess(t *testing.T) {
	id := "1"
	e := id

	if e != id {
		t.Error("erro")
	}
}
