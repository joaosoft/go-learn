package domain

import (
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestError(t *testing.T) {
	id := uuid.NewV4()
	e := 1

	if e.id != id {
		t.Error("erro")
	}

	if e.Error() != "erro "+e.id.String() {
		t.Error("Error() not returing correct format")
	}
}

func TestSuccess(t *testing.T) {
	id := uuid.NewV4()
	e := id

	if e.id != id {
		t.Error("erro")
	}

	if e.Error() != "erro "+e.id.String() {
		t.Error("Error() not returing correct format")
	}
}
