package swagger

import (
	"net/http"
	"swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
	r := gin.Default()

	// Swagger documentation setup
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// API routes
	api := r.Group("/api/v1")
	{
		api.GET("/ping", ping)
	}

	// new route
	api.GET("/user/:id", getUser)

	// Start server
	r.Run(":8080")

}

// ping example
// @Summary Ping example
// @Description Do ping
// @Tags ping
// @Success 200 {string} string "pong"
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// getUser example
// @Summary Get User
// @Description Get a user by ID
// @Tags user
// @Produce  json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /user/{id} [get]
func getUser(c *gin.Context) {
	id := c.Param("id")
	user := User{
		ID:   1,
		Name: "Ali Daei",
	}
	if id == "1" {
		c.JSON(http.StatusOK, user)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"message": "user not found"})
	}
}
