package logger

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Logger struct for handling logs
type Logger struct {
	file *os.File
}

func NewLogger() *Logger {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return &Logger{file: file}
}

// Log method to write logs to the file
func (l *Logger) Log(message string) {
	log.SetOutput(l.file)
	log.Println(message)
}

// Close method to close the log file
func (l *Logger) Close() {
	l.file.Close()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Home Page!")
	logger.Log(fmt.Sprintf("Request to Home: %s %s", r.Method, r.URL.Path))
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About Page")
	logger.Log(fmt.Sprintf("Request to About: %s %s", r.Method, r.URL.Path))
}

// ErrorHandler simulates an error for testing purposes
func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Something went wrong!", http.StatusInternalServerError)
	logger.Log(fmt.Sprintf("Error on request: %s %s - %s", r.Method, r.URL.Path,
		http.StatusText(http.StatusInternalServerError)))
}

var logger *Logger

func main() {

	logger = NewLogger()
	defer logger.Close()

	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/error", ErrorHandler)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Log("Server started on :8080")
	log.Fatal(server.ListenAndServe())
}
