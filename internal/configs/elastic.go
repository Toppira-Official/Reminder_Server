package configs

import (
	"log"

	elasticsearch "github.com/elastic/go-elasticsearch/v9"
)

func GetElastic(envs Environments) *elasticsearch.TypedClient {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://" + envs.ELASTIC_HOST.String() +
				":" + envs.ELASTIC_PORT.String(),
		},
		Username: "elastic",
		Password: envs.ELASTIC_PASSWORD.String(),
	}

	es, err := elasticsearch.NewTypedClient(cfg)
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
	}

	return es
}
