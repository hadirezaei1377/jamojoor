package models

import "time"

type Request struct {
	ID         int       `json:"id"`
	Endpoint   string    `json:"endpoint"`
	Payload    string    `json:"payload"`
	RetryCount int       `json:"retry_count"`
	LastTried  time.Time `json:"last_tried"`
	Status     string    `json:"status"` // values: "pending", "failed", "successful"
}
