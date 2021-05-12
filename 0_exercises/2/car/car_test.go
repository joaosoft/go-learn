package car

import (
	"testing"
)


func TestSomething(t *testing.T) {
	imp := &CarModelDefectImp{}
	cars := []*Car{
		&Car{
			Model: CarModel{
			Brand: BrandVW,
			Name: "Teste",
			Version: 1,
		},
			ManufacturingYear: 2021,
			EngineSerial: "123",
		},
	}

	defects := GetCarsDefects(imp, cars...)
	if len(defects) > 0 {
		t.Error("error: invalid car defects!")
	}
}

