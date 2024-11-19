package handlers

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func AddScore(client *redis.Client, user string, score float64) error {
	return client.ZAdd(context.Background(), "leaderboard", &redis.Z{
		Score:  score,
		Member: user,
	}).Err()
}

func GetTopUsers(client *redis.Client, top int64) ([]redis.Z, error) {
	return client.ZRevRangeWithScores(context.Background(), "leaderboard", 0, top-1).Result()
}
