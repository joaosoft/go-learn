package main

import (
	"github.com/labstack/gommon/log"
	common "golang-learn/examples/51_elastic_search/common/config"
	"golang-learn/examples/51_elastic_search/config"
	"golang-learn/examples/51_elastic_search/domain"
	"os"
	"golang-learn/examples/51_elastic_search/controllers"
	"golang-learn/examples/51_elastic_search/interactors"
	"golang-learn/examples/51_elastic_search/repositories"
	"strconv"
	"encoding/json"
)


var _configuration config.Configuration

func init() {
	if err := common.LoadConfigFromFile("config", &_configuration); err != nil {
		log.Error("Error loading config: ", err)
		os.Exit(0)
	}
}

func main() {
	log.Infof("JOB START")

	//controller := controllers.Controller {
	//	Interactor: interactors.Interactor{
	//		Repository: repositories.Repository {
	//			Configuration: _configuration,
	//		},
	//	},
	//}


	repository := repositories.NewRepository(_configuration)
	interactor := interactors.NewInteractor(repository)
	controller := controllers.Controller {
		Interactor: *interactor,

	}

	controller.CreateIndex("dummy")

	bulkInsert := [10]domain.Something {}
	type Teste struct {
		Fruit    string `json:"fruit"`
		Category string `json:"category"`
	}

	for i := 0 ; i < 10 ; i++ {
		teste := Teste { Fruit: "banana", Category: "natureza"}
		bytes, _ := json.Marshal(teste)
		bulkInsert[i] = domain.Something {
			TYPE: "Teste " + strconv.Itoa(i),
			ID: strconv.Itoa(i),
			DATA: bytes,
		}
	}

	controller.Insert(bulkInsert[:])

	log.Infof("JOB END")
}
