package main

import (
	"crypto-price-system/bot"
	"crypto-price-system/collector"
	"crypto-price-system/storage"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	apiUrl := "https://api.wallex.ir/v1/markets"

	go bot.StartBot("OUR_TELEGRAM_BOT_TOKEN")

	// get data every 1 minutes
	for {
		prices, err := collector.FetchPrices(apiUrl)
		if err != nil {
			log.Println("error while getting prices:", err)
			continue
		}

		for _, price := range prices {
			err := storage.SaveToRedis(rdb, price.Currency, price.Price)
			if err != nil {
				log.Println("error:", err)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
