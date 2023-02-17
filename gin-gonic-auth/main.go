package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/kwamekyeimonies/gin-gonic-authetication/routes"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env Files")
	}

	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("No available port funning")
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET(
		"/api",
		func(c *gin.Context) {
			c.JSON(200, gin.H{"success": "access granted for api"})
		},
	)

	router.Run(":" + port)
}
