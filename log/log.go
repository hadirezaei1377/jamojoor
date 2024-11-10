package log

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// logger config
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	// log info
	logger.Println("application is started")

	// simulate an operation
	result, err := performOperation()
	if err != nil {
		// error log
		logger.SetPrefix("ERROR: ")
		logger.Println("error while applying:", err)
		return
	}

	// result log
	logger.SetPrefix("INFO: ")
	logger.Println("operation result log:", result)
}

func performOperation() (int, error) {
	return 0, fmt.Errorf("operation was not successfully!")
}
