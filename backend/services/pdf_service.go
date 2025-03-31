package services

import (
	"bytes"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func ExtractTextFromPDF(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		return "", err
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		return "", err
	}

	var extractedText bytes.Buffer
	for i := 1; i <= numPages; i++ {
		page, err := pdfReader.GetPage(i)
		if err != nil {
			return "", err
		}
		ex, err := extractor.New(page)
		if err != nil {
			return "", err
		}
		pageText, err := ex.ExtractText()
		if err != nil {
			return "", err
		}
		extractedText.WriteString(pageText + "\n")
	}

	return extractedText.String(), nil
}
