package productservice

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		productID := "prod-1"
		productData := "Product Data" // it will be retrive from main database

		// data caching
		err := rdb.Set(ctx, productID, productData, 0).Err()
		if err != nil {
			log.Fatalf("Could not set product: %v", err)
		}
		fmt.Fprintf(w, "Product cached: %s", productData)
	})

	http.HandleFunc("/get-product", func(w http.ResponseWriter, r *http.Request) {
		productID := "prod-1"
		productData, err := rdb.Get(ctx, productID).Result()
		if err == redis.Nil {
			fmt.Fprintln(w, "Product not found in cache")
		} else if err != nil {
			fmt.Fprintf(w, "Error retrieving product: %v", err)
		} else {
			fmt.Fprintf(w, "Product Data: %s", productData)
		}
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
