package utils

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, statusCode int, message string, err error) {
	log.Println(message, ":", err)
	c.JSON(statusCode, gin.H{"error": message})
}

func CreateUploadDir() {
	uploadDir := os.Getenv("UPLOAD_DIR")
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}
}
