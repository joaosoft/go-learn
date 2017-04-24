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

	if e.Error() != "Unknown Channel "+e.ChannelID.String() {
		t.Error("Error() not returing correct format")
	}
}

func TestSuccess(t *testing.T) {
	whouseID := "My Warehouse"
	e := NewDomainErrorUnknownWarehouse(whouseID)

	if e.WarehouseID != whouseID {
		t.Error("Warehouse not set")
	}

	if e.Error() != "Unknown warehouseID "+e.WarehouseID {
		t.Error("Error() not returing correct format")
	}
}
