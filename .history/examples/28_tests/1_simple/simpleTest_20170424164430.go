package simple

import (
	"testing"
)

func TestError(t *testing.T) {
	id := "1"
	e := "2"

	if e.id != id {
		t.Error("erro")
	}

	if e.Error() != "erro "+e.id.String() {
		t.Error("Error() not returing correct format")
	}
}

func TestSuccess(t *testing.T) {
	id := ""1"
	e := id

	if e.id != id {
		t.Error("erro")
	}

	if e.Error() != "erro "+e.id.String() {
		t.Error("Error() not returing correct format")
	}
}
