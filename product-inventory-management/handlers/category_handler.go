package handlers

import (
	"encoding/json"
	"net/http"
	"product-inventory-management/config"
	"product-inventory-management/internal/models"

	"github.com/gorilla/mux"
)

func RegisterCategoryRoutes(r *mux.Router) {
	r.HandleFunc("/categories", GetAllCategories).Methods("GET")
	r.HandleFunc("/categories", CreateCategory).Methods("POST")
}

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	var categories []models.Category
	config.DB.Find(&categories)
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	json.NewDecoder(r.Body).Decode(&category)
	config.DB.Create(&category)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(category)
}
