package main

import (
	"log"
	"os"

	"gopdf/handlers"
	"gopdf/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	utils.CreateUploadDir()

	r := gin.Default()

	r.POST("/upload", handlers.UploadPDF)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("🚀 Server running on port", port)
	r.Run(":" + port)
}
