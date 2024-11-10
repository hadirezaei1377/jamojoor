package log

import (
	"log"
	"os"
)

func main2() {
	file, err := os.Open("nonexistent.txt")
	if err != nil {
		log.Fatal(err) // This will log the error and exit the program
	}
	defer file.Close()

	// If the file is opened successfully, continue with the logic
}
