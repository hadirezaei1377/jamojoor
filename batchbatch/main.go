package batchbatch

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

type Record struct {
	ID    int
	Name  string
	Email string
}

func main() {
	inputFile := "input.csv"
	outputFile := "output.csv"
	errorLogFile := "errors.log"

	file, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer file.Close()

	output, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer output.Close()

	errorLog, err := os.Create(errorLogFile)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer errorLog.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	batchSize := 1000
	totalRecords := len(records)
	var wg sync.WaitGroup
	successCount := 0
	errorCount := 0

	writer := csv.NewWriter(output)
	defer writer.Flush()

	for i := 0; i < totalRecords; i += batchSize {
		wg.Add(1)
		go func(start int) {
			defer wg.Done()
			end := start + batchSize
			if end > totalRecords {
				end = totalRecords
			}

			for _, row := range records[start:end] {
				record, err := processRecord(row)
				if err != nil {
					logError(errorLog, row, err)
					errorCount++
					continue
				}

				err = writer.Write([]string{
					strconv.Itoa(record.ID),
					record.Name,
					record.Email,
				})
				if err != nil {
					logError(errorLog, row, err)
					errorCount++
					continue
				}
				successCount++
			}
		}(i)
	}

	wg.Wait()

	fmt.Printf("number of successful records: %d\n", successCount)
	fmt.Printf("number of unsuccessful records: %d\n", errorCount)
}

func processRecord(row []string) (*Record, error) {
	if len(row) < 3 {
		return nil, fmt.Errorf("imperfect data")
	}

	id, err := strconv.Atoi(row[0])
	if err != nil {
		return nil, fmt.Errorf("invalid id: %v", row[0])
	}

	name := row[1]
	email := row[2]
	if name == "" || email == "" {
		return nil, fmt.Errorf("invalid name or email")
	}

	return &Record{
		ID:    id,
		Name:  name,
		Email: email,
	}, nil
}

func logError(file *os.File, row []string, err error) {
	logMsg := fmt.Sprintf("error: %v: %v\n", row, err)
	file.WriteString(logMsg)
}
