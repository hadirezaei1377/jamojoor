package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"spam-filter/bloom"
)

var filter = bloom.NewBloomFilter(1000, 5) // hash

type Request struct {
	Item string `json:"item"`
}

func AddSpam(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	filter.Add(req.Item)
	fmt.Fprintf(w, "Item '%s' added to spam list", req.Item)
}

func CheckSpam(w http.ResponseWriter, r *http.Request) {
	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if filter.Check(req.Item) {
		fmt.Fprintf(w, "Item '%s' is spam", req.Item)
	} else {
		fmt.Fprintf(w, "Item '%s' is not spam", req.Item)
	}
}
