package main

import (
	"log"
	"os"

	"management_buku/database"
	"management_buku/models"
	"management_buku/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
	database.DB.AutoMigrate(&models.Book{}, &models.Category{}, &models.User{})

    r := gin.Default()

    routes.SetupRoutes(r)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
   	log.Println("Server running on port:", port)
	r.Run(":" + port)
}
