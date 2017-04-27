package simple_test

import (
	"testing"
)

// $ go test -v
// $ go test -v -run="French|Canal"

// Coverage
// $ go tool cover
// Usage of 'go tool cover':
// Given a coverage profile produced by 'go test':
//     go test -coverprofile=c.out

// Open a web browser displaying annotated source code:
//     go tool cover -html=c.out

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
