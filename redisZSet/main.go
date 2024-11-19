package rediszset

import (
	"fmt"
	"rediszset/handlers"
	"rediszset/redis"
)

func main() {
	client := redis.Connect()
	defer client.Close()

	handlers.AddScore(client, "Ali", 100)
	handlers.AddScore(client, "Boby", 150)

	topUsers, _ := handlers.GetTopUsers(client, 2)
	for _, user := range topUsers {
		fmt.Printf("User: %v, Score: %v\n", user.Member, user.Score)
	}
}
