package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gopdf/services"
	"gopdf/utils"
)

func UploadPDF(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, "No file uploaded", err)
		return
	}

	uploadDir := os.Getenv("UPLOAD_DIR")
	uniqueFileName := uuid.New().String() + filepath.Ext(file.Filename)
	filePath := filepath.Join(uploadDir, uniqueFileName)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to save file", err)
		return
	}

	extractedText, err := services.ExtractTextFromPDF(filePath)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to extract text", err)
		return
	}

	chatGPTResponse, err := services.ProcessTextWithChatGPT(extractedText)
	if err != nil {
		utils.HandleError(c, http.StatusInternalServerError, "Failed to process text with ChatGPT", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"extracted_text":  extractedText,
		"chatgpt_summary": chatGPTResponse,
	})
}
