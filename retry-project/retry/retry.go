package retry

import (
	"fmt"
	"log"
	"net/http"
	"retry-project/models"
	"time"
)

const MaxRetryCount = 5

func RetryRequest(req models.Request) {
	if req.RetryCount >= MaxRetryCount {
		fmt.Println("Max retry limit reached. Request failed.")
		UpdateRequestStatus(req.ID, "failed")
		return
	}

	resp, err := http.Get(req.Endpoint)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("Request failed. Retrying...")
		req.RetryCount++
		req.LastTried = time.Now()
		UpdateRequest(req)
		RetryRequest(req)
	} else {
		fmt.Println("Request succeeded.")
		UpdateRequestStatus(req.ID, "successful")
	}
}

func UpdateRequest(req models.Request) {
	_, err := db.Exec("UPDATE requests SET retry_count = ?, last_tried = ? WHERE id = ?", req.RetryCount, req.LastTried, req.ID)
	if err != nil {
		log.Println("Error updating request:", err)
	}
}

func UpdateRequestStatus(id int, status string) {
	_, err := db.Exec("UPDATE requests SET status = ? WHERE id = ?", status, id)
	if err != nil {
		log.Println("Error updating request status:", err)
	}
}
