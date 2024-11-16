package handlers

import (
	"encoding/json"
	"net/http"
	"product-inventory-management/config"
	"product-inventory-management/internal/models"

	"github.com/gorilla/mux"
)

func RegisterProductRoutes(r *mux.Router) {
	r.HandleFunc("/products", GetAllProducts).Methods("GET")
	r.HandleFunc("/products", CreateProduct).Methods("POST")
}

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	config.DB.Find(&products)
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)
	config.DB.Create(&product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}
