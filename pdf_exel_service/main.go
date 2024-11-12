package pdfexelservice

import (
	"log"
	"net/http"
)

func main() {
	InitDB()
	http.HandleFunc("/register", CreateUser)
	http.HandleFunc("/generate-file-sync", GenerateFileSync)
	http.HandleFunc("/generate-file-async", GenerateFileAsync)

	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
