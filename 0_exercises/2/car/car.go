package car

import (
	"encoding/json"
	"fmt"
	"log"
)

const (
	BrandBMW  Brand = "BMW"
	BrandVW   Brand = "VW"
	BrandAudi Brand = "Audi"
)

var (
	ErrorInvalidVersion = fmt.Errorf("invalid version")
	Defects             []Defect
)

type Brand string

type Car struct {
	Model             CarModel
	ManufacturingYear int
	EngineSerial      string
}

type CarModel struct {
	Name    string
	Version int
	Brand   Brand
}

type Defect struct {
	Model         CarModel
	AffectedYears []int
	Code          string
}

func (c *Car) Validate() (errs []error) {
	if c.Model.Version < 1 {
		errs = append(errs, ErrorInvalidVersion)
	}
	return errs
}

func (c *CarModel) String() string {
	return fmt.Sprintf("Name: %s, Version: %d, Brand: %s", c.Name, c.Version, c.Brand)
}

func (c *CarModel) StringForLog() string {
	bytes, err := json.Marshal(c)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func (c *CarModel) IncrementVersion() {
	c.Version++
}

func GetCarsDefects(imp CarDefecter, cars ...*Car) (defects []Defect) {
	if len(cars) == 0 {
		return defects
	}

	var models []CarModel

	for _, c := range cars {
		models = append(models, c.Model)
	}

	// second step
	modelsDefects := imp.GetCarModelsDefects(models...)

	for _, c := range cars {
		for _, d := range modelsDefects {
			if c.Model.Version == d.Model.Version &&
				c.Model.Brand == d.Model.Brand &&
				c.Model.Name == d.Model.Name {

				for _, y := range d.AffectedYears {
					if c.ManufacturingYear == y {
						defects = append(defects, d)
					}
				}
			}
		}
	}

	return defects
}

func CarDefectFactory() CarDefecter {
	return &CarModelDefectImp{}
}

type CarDefecter interface {
	GetCarModelsDefects(cars ...CarModel) (defects []Defect)
}

type CarModelDefectImp struct{}

func (c *CarModelDefectImp) GetCarModelsDefects(cars ...CarModel) (defects []Defect) {
	return []Defect{}
}
