package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"gopdf/handlers"
	"gopdf/utils"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

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
