package queue

import (
	"encoding/csv"
	"fmt"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	file, err := os.Open("phones.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	for _, record := range records {
		phoneNumber := record[0]
		wg.Add(1)
		go sendSMS(phoneNumber)
	}
	wg.Wait()
}

func sendSMS(phoneNumber string) {
	defer wg.Done()
	fmt.Printf("Sending SMS to %s\n", phoneNumber)
}
