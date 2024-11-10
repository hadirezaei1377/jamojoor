package userservice

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

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		sessionID := "user-session-id" // random session
		err := rdb.Set(ctx, sessionID, "user123", 0).Err()
		if err != nil {
			log.Fatalf("Could not set session: %v", err)
		}
		fmt.Fprintf(w, "Logged in with session: %s", sessionID)
	})

	http.HandleFunc("/session", func(w http.ResponseWriter, r *http.Request) {
		sessionID := "user-session-id" // random session
		userID, err := rdb.Get(ctx, sessionID).Result()
		if err == redis.Nil {
			fmt.Fprintln(w, "Session not found")
		} else if err != nil {
			fmt.Fprintf(w, "Error retrieving session: %v", err)
		} else {
			fmt.Fprintf(w, "User ID: %s", userID)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
