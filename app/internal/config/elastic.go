package config

import "github.com/elastic/go-elasticsearch"

type ElasticSearch struct {
	Client *elasticsearch.Client
}

func NewElastic(addresses []string) (*ElasticSearch, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}

	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &ElasticSearch{
		Client: client,
	}, nil
}
