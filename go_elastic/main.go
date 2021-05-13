package main

import (
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/elastic/go-elasticsearch/v8/esapi"
	"log"
)

func main() {
	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		// ...
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println(err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(elasticsearch.Version)
	log.Println(es.Info())
}