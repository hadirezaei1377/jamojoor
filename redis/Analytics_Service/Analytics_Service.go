package analyticsservice

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

	http.HandleFunc("/log-event", func(w http.ResponseWriter, r *http.Request) {
		event := "User Login Event"

		// save event
		err := rdb.LPush(ctx, "event_log", event).Err()
		if err != nil {
			log.Fatalf("Could not log event: %v", err)
		}
		fmt.Fprintf(w, "Logged event: %s", event)
	})

	http.HandleFunc("/get-events", func(w http.ResponseWriter, r *http.Request) {
		events, err := rdb.LRange(ctx, "event_log", 0, -1).Result()
		if err == redis.Nil {
			fmt.Fprintln(w, "No events found")
		} else if err != nil {
			fmt.Fprintf(w, "Error retrieving events: %v", err)
		} else {
			for _, event := range events {
				fmt.Fprintf(w, "Event: %s\n", event)
			}
		}
	})

	log.Fatal(http.ListenAndServe(":8083", nil))
}
