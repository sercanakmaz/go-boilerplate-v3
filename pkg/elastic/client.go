package elastic

import (
	"github.com/elastic/go-elasticsearch/v7"
)

func NewElasticClient(nodes []string) *elasticsearch.Client {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: nodes,
	})

	if err != nil {
		panic(err)
	}

	return client
}
