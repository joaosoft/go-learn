package repositories

import (
	"github.com/labstack/gommon/log"
	"golang.org/x/net/context"
	"gopkg.in/olivere/elastic.v5"
	"encoding/json"
	"golang-learn/examples/51_elastic_search/config"
	"golang-learn/examples/51_elastic_search/domain"
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
	ctx, client, err := getElasticSearchContext(repository.Configuration)

	if err != nil {
		return err
	}

	body, err := json.Marshal(mapping)
	if err != nil {
		return err
	}

	result, err := client.CreateIndex(index).
		BodyString(string(body)).
		Do(ctx)
	if err != nil {
		return err
	}
	if !result.Acknowledged {
		// Not acknowledged
	}

	log.Infof("New index created [name:%s]", index)

	return nil
}

func (repository *Repository) DeleteIndex(index string) error {
	ctx, client, err := getElasticSearchContext(repository.Configuration)
	if err != nil {
		return err
	}
	result, err := client.DeleteIndex(index).Do(ctx)
	if err != nil {
		return err
	}
	if !result.Acknowledged {
		// Not acknowledged
	} else {
		log.Infof("Index deleted successfully [name:%s]", index)
	}

	return nil
}

func (repository *Repository) Insert(data []domain.Something) error {
	ctx, client, err := getElasticSearchContext(repository.Configuration)

	validateIndex(repository.Configuration, ctx, client)

	if err != nil {
		return err
	}

	service, err := client.BulkProcessor().Name("worker").Workers(2).Do(ctx)
	for _, element := range data {
		body := string(element.DATA)

		// Create a bulk index request
		r := elastic.NewBulkIndexRequest().
			Index(repository.Configuration.MyIndexer.Index).
			Type(element.TYPE).
			Id(element.ID).
			Doc(body)

		//_, err := client.Index().
		//	Index(repository.Configuration.ManagerIndexer.Index).
		//	Type(element.Type).
		//	Id(element.Id).
		//	BodyJson(body).
		//	Do(ctx)

		// Add the request r to the processor
		service.Add(r)
	}
	err = service.Flush()

	log.Infof("Bulk request finished", err)

	return nil
}

// Get elastic search context
func getElasticSearchContext(config config.Configuration) (context.Context, *elastic.Client, error) {
	// Starting with elastic.v5, you must pass a context to execute each service
	ctx := context.Background()

	client, err := elastic.NewClient(elastic.SetURL(config.MyIndexer.Hosts...))
	if err != nil {
		return nil, nil, err
	}

	return ctx, client, nil
}

func validateIndex(config config.Configuration, ctx context.Context, client *elastic.Client) (error) {
	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists(config.MyIndexer.Index).Do(ctx)
	if err != nil {
		return err
	}

	if !exists {
		// Create a worker_add index.
		if _, err := client.CreateIndex(config.MyIndexer.Index).Do(ctx); err != nil {
			return err
		}
	}

	return nil
}

