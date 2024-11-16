package cmd

import (
	"jamojoor/product-inventory-management/internal/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	config.InitDB()

	config.DB.AutoMigrate(&models.Product{}, &models.Category{})

	handlers.RegisterProductRoutes(r)
	handlers.RegisterCategoryRoutes(r)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
