package orderservice

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

	http.HandleFunc("/order", func(w http.ResponseWriter, r *http.Request) {
		orderID := "order-1"

		// add order to queue
		err := rdb.LPush(ctx, "order_queue", orderID).Err()
		if err != nil {
			log.Fatalf("Could not add order to queue: %v", err)
		}
		fmt.Fprintf(w, "Order added to queue: %s", orderID)
	})

	http.HandleFunc("/process-orders", func(w http.ResponseWriter, r *http.Request) {

		for {
			orderID, err := rdb.LPop(ctx, "order_queue").Result()
			if err == redis.Nil {
				break // queue is empty
			} else if err != nil {
				fmt.Fprintf(w, "Error retrieving order: %v", err)
				return
			}
			fmt.Fprintf(w, "Processing order: %s\n", orderID)
		}
	})

	log.Fatal(http.ListenAndServe(":8082", nil))
}
