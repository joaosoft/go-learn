package original

import (
	"github.com/labstack/gommon/log"
	"golang-learn/examples/51_elastic_search/config"
	"golang-learn/examples/51_elastic_search/domain"
    "github.com/elastic/go-elasticsearch/client"
	"fmt"
)

type Repository struct {
	Configuration config.Configuration
}

func NewRepository(config config.Configuration) *Repository {
	repository := Repository{
		Configuration: config,
	}

	return &repository
}

func (repository *Repository) CreateIndex(index string, mapping map[string]interface{}) error {


	client, err := client.New(client.WithHost(repository.Configuration.MyIndexer.Hosts[0]))

	if err != nil {
		return err
	}

	//body, err := json.Marshal(mapping)
	if err != nil {
		return err
	}

	resp, err := client.Indices.Create(index)

	if err != nil {
		return err
	}

	fmt.Println("RESPONSE:", resp)

	log.Infof("New index created [name:%s]", index)

	return nil
}

func (repository *Repository) DeleteIndex(index string) error {
	return nil
}

func (repository *Repository) Insert(data []domain.Something) error {

	client, err := client.New(client.WithHost(repository.Configuration.MyIndexer.Hosts[0]))
	body := map[string]interface{}{
		"query": map[string]interface{}{
			"term": map[string]interface{}{
				"user": "kimchy",
			},
		},
	}
	resp, err := client.Bulk(body)

	if err != nil {
		return err
	}

	fmt.Println("RESPONSE:", resp)

	return nil
}

