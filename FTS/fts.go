package main

import (
	"context"
	"fmt"
	"log"

	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

func main() {
	// connect to Elasticsearch
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	// adding data
	documents := []string{
		`{"content": "Full Text Search in Golang with Elasticsearch."}`,
		`{"content": "Elasticsearch is powerful and scalable."}`,
	}
	for i, doc := range documents {
		req := esapi.IndexRequest{
			Index:      "documents",
			DocumentID: fmt.Sprintf("%d", i+1),
			Body:       strings.NewReader(doc),
			Refresh:    "true",
		}
		res, err := req.Do(context.Background(), es)
		if err != nil {
			log.Fatalf("Error indexing document %d: %s", i+1, err)
		}
		res.Body.Close()
	}

	// FTS
	query := `{
		"query": {
			"match": {
				"content": "Elasticsearch"
			}
		}
	}`
	res, err := es.Search(
		es.Search.WithIndex("documents"),
		es.Search.WithBody(strings.NewReader(query)),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error searching the index: %s", err)
	}
	defer res.Body.Close()

	fmt.Println(res)
}
