package pm

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"golang-learn/examples/64_manager/elastic"
)

// -------------- ELASTIC --------------
// NewElasticClient ... creates a new elastic client
func (instance *manager) NewElasticClient(config *elastic.Config) *elastic.ElasticController {
	log.Infof(fmt.Sprintf("elastic, creating elastic"))
	return elastic.NewElastic(config)
}

// -------------- METHODS --------------
// GetElasticClient ... get elastic client by key
func (instance *manager) GetElasticClient(key string) (*elastic.ElasticController, error) {
	return instance.ElasticController[key], nil
}

// AddElasticClient, add a new elastic client
func (instance *manager) AddElasticClient(key string, elasticClient *elastic.ElasticController) error {
	log.Infof(fmt.Sprintf("elastic, add a new elastic client '%s'", key))
	instance.ElasticController[key] = elasticClient
	return nil
}

// RemElasticClient, remove the elastic client by key
func (instance *manager) RemElasticClient(key string) (*elastic.ElasticController, error) {
	log.Infof(fmt.Sprintf("elastic, remove the elastic client '%s'", key))

	// get gateway
	controller := instance.ElasticController[key]

	// delete gateway
	delete(instance.ElasticController, key)
	log.Infof(fmt.Sprintf("elastic, elastic client '%s' removed", key))

	return controller, nil
}
