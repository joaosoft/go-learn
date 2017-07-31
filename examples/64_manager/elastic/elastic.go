package elastic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ElasticController ... elastic search controller structure
type ElasticController struct {
	config *Config
	client *http.Client
}

// NewElastic ... creates a new elastic search connector
func NewElastic(config *Config) *ElasticController {
	elasticSearch := &ElasticController{
		config: config,
		client: &http.Client{},
	}

	return elasticSearch
}

// Search ... search method in elastic search
func (instance *ElasticController) Search(index string, docType string, query string) (*Response, error) {
	var elasticResponse Response
	var httpResponse *http.Response
	var body []byte
	var endpoint string
	var err error

	// search by ...
	if docType != "" {
		endpoint = fmt.Sprintf("/%s/%s/_search", index, docType)
	} else {
		endpoint = fmt.Sprintf("/%s/_search", index)
	}

	// Call to http client
	if httpResponse, err = instance.client.Post(instance.config.Host+endpoint, "application/json", strings.NewReader(query)); err != nil {
		return nil, fmt.Errorf("unexpected error when executing searching", err)
	}

	switch httpResponse.StatusCode {
	case 200:
		// Read body response
		if body, err = ioutil.ReadAll(httpResponse.Body); err != nil {
			return nil, fmt.Errorf("unexpected error when reading response body", err)
		}

		// Unmarshal response to elastic search struct
		if err := json.Unmarshal(body, &elasticResponse); err != nil {
			return nil, fmt.Errorf("error unmarshalling elasticsearch response", err)
		}
		return &elasticResponse, nil
	default:
		return nil, fmt.Errorf(fmt.Sprintf("unexpected response when executing search [code:%d][message:%s]", httpResponse.Status, string(body)), nil)
	}

}
