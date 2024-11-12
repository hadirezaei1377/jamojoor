package pdfexelservice

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jung-kurt/gofpdf"
	"github.com/xuri/excelize/v2"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = db.NamedExec("INSERT INTO users (name, email) VALUES (:name, :email)", &user)
	if err != nil {
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func GeneratePDF(user User) string {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, fmt.Sprintf("User: %s", user.Name))
	pdf.Cell(0, 10, fmt.Sprintf("Email: %s", user.Email))

	fileName := fmt.Sprintf("%s.pdf", user.Name)
	err := pdf.OutputFileAndClose(fileName)
	if err != nil {
		log.Println("Failed to generate PDF:", err)
	}
	return fileName
}

func GenerateExcel(user User) string {
	f := excelize.NewFile()
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Email")
	f.SetCellValue("Sheet1", "A2", user.Name)
	f.SetCellValue("Sheet1", "B2", user.Email)

	fileName := fmt.Sprintf("%s.xlsx", user.Name)
	if err := f.SaveAs(fileName); err != nil {
		log.Println("Failed to generate Excel:", err)
	}
	return fileName
}

// GenerateFileSync generates a file (PDF or Excel) synchronously
func GenerateFileSync(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("id")
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = ?", userID)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	fileType := r.URL.Query().Get("type")
	var fileName string
	if fileType == "pdf" {
		fileName = GeneratePDF(user)
	} else {
		fileName = GenerateExcel(user)
	}

	http.ServeFile(w, r, fileName)
}

// GenerateFileAsync generates a file asynchronously
func GenerateFileAsync(w http.ResponseWriter, r *http.Request) {
	go GenerateFileSync(w, r)
	w.Write([]byte("Request received, file is being generated"))
}
