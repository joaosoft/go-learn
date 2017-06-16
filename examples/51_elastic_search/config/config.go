package config

type Configuration struct {
	MyIndexer struct {
		ElasticSearch `json:"elasticsearch"`
	} `json:"my_indexer"`
}

type ElasticSearch struct {
	Index string `json:"index"`
	Hosts []string `json:"hosts"`
}