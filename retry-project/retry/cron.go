package retry

import (
	"fmt"
	"log"
	"retry-project/models"

	"github.com/robfig/cron/v3"
)

func StartCronJob() {
	c := cron.New()
	c.AddFunc("@every 10m", func() {
		fmt.Println("Running cron job for retry...")
		rows, err := db.Query("SELECT id, endpoint, payload, retry_count, last_tried, status FROM requests WHERE status = 'pending'")
		if err != nil {
			log.Println("Error fetching pending requests:", err)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var req models.Request
			if err := rows.Scan(&req.ID, &req.Endpoint, &req.Payload, &req.RetryCount, &req.LastTried, &req.Status); err != nil {
				log.Println("Error scanning request:", err)
				continue
			}
			RetryRequest(req)
		}
	})
	c.Start()

	// Wait indefinitely
	select {}
}
