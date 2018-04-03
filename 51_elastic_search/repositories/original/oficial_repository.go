package original

import (
	"fmt"
	"go-learn/51_elastic_search/config"
	"go-learn/51_elastic_search/domain"

	"github.com/elastic/go-elasticsearch/client"
	"github.com/labstack/gommon/log"
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
