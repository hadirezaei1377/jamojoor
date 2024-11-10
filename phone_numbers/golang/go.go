package golang

import (
	"encoding/csv"
	"fmt"
	"os"
)

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
		err := sendSMS(phoneNumber)
		if err != nil {
			fmt.Println("Error sending SMS:", err)
		}
	}
}

func sendSMS(phoneNumber string) error {
	fmt.Printf("Sending SMS to %s\n", phoneNumber)
	return nil
}
