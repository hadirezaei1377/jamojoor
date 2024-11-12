package booksearch

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
)

const elasticURL = "http://localhost:9200"

func main() {
	// connect to elk
	client, err := elastic.NewClient(elastic.SetURL(elasticURL), elastic.SetSniff(false))
	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %v", err)
	}
	fmt.Println("Connected to Elasticsearch")

	createIndex(client)
}

func createIndex(client *elastic.Client) {
	ctx := context.Background()

	exists, err := client.IndexExists("books").Do(ctx)
	if err != nil {
		log.Fatalf("Error checking if index exists: %v", err)
	}

	if !exists {
		// create new index
		mapping := `{
            "mappings": {
                "properties": {
                    "title": { "type": "text" },
                    "author": { "type": "text" },
                    "genre": { "type": "keyword" },
                    "publish_date": { "type": "date" },
                    "description": { "type": "text" }
                }
            }
        }`
		_, err := client.CreateIndex("books").BodyString(mapping).Do(ctx)
		if err != nil {
			log.Fatalf("Error creating index: %v", err)
		}
		fmt.Println("Index created successfully")
	} else {
		fmt.Println("Index already exists")
	}
}

type Book struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	Genre       string `json:"genre"`
	PublishDate string `json:"publish_date"`
	Description string `json:"description"`
}

func addBook(client *elastic.Client, book Book) {
	ctx := context.Background()

	_, err := client.Index().
		Index("books").
		BodyJson(book).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error adding book: %v", err)
	}
	fmt.Println("Book added successfully")
}

func searchBooks(client *elastic.Client, query string) {
	ctx := context.Background()

	searchResult, err := client.Search().
		Index("books").
		Query(elastic.NewMultiMatchQuery(query, "title", "author", "description")).
		Do(ctx)
	if err != nil {
		log.Fatalf("Error searching books: %v", err)
	}

	fmt.Printf("Found %d books\n", searchResult.TotalHits())
	for _, hit := range searchResult.Hits.Hits {
		var book Book
		err := json.Unmarshal(hit.Source, &book)
		if err != nil {
			log.Fatalf("Error unmarshaling book: %v", err)
		}
		fmt.Printf("Title: %s, Author: %s\n", book.Title, book.Author)
	}
}
