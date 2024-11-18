package spamfilter

import (
	"log"
	"net/http"
	"spam-filter/handlers"
)

func main() {
	http.HandleFunc("/add", handlers.AddSpam)
	http.HandleFunc("/check", handlers.CheckSpam)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
