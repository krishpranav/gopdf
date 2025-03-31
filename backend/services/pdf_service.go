package services

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

func ExtractTextFromPDF(filePath string) (string, error) {
	tempDir, err := ioutil.TempDir("", "pdfcpu_extract")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir) 

	err = api.ExtractContentFile(filePath, tempDir, nil, model.NewDefaultConfiguration())
	if err != nil {
		return "", fmt.Errorf("failed to extract text: %v", err)
	}

	var extractedText string
	files, err := ioutil.ReadDir(tempDir)
	if err != nil {
		return "", fmt.Errorf("failed to read extracted files: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(tempDir, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			return "", fmt.Errorf("failed to read extracted text file: %v", err)
		}
		extractedText += string(content) + "\n"
	}

	return extractedText, nil
}
