package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func SaveToRedis(client *redis.Client, currency string, price float64) error {
	score := float64(time.Now().Unix())
	_, err := client.ZAdd(ctx, "crypto_prices", &redis.Z{Score: score,
		Member: fmt.Sprintf("%s:%f", currency, price)}).Result()
	return err
}
