package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/ledongthuc/pdf"
)

// Constants for file permissions and error messages
const (
	FilePermission = 0644
	UsageMessage   = "Usage: -input <file.pdf> -output <file.json>"
)

// Struct to hold the extracted PDF content
type PDFContent struct {
	Content string `json:"content"`
}

func main() {
	// Parse command-line flags
	inputPath := flag.String("input", "", "Path to the input PDF file")
	outputPath := flag.String("output", "", "Path to the output JSON file")
	flag.Parse()

	// Validate input arguments
	if err := validateArgs(*inputPath, *outputPath); err != nil {
		log.Fatalf("Error: %v\n%s", err, UsageMessage)
	}

	// Extract text content from the PDF
	content, err := extractPDFContent(*inputPath)
	if err != nil {
		log.Fatalf("Failed to extract PDF content: %v", err)
	}

	// Write the extracted content to a JSON file
	if err := writeJSONContent(*outputPath, content); err != nil {
		log.Fatalf("Failed to write JSON content: %v", err)
	}

	fmt.Printf("PDF content successfully written to %s\n", *outputPath)
}

// validateArgs validates the input and output file paths
func validateArgs(inputPath, outputPath string) error {
	if inputPath == "" || outputPath == "" {
		return fmt.Errorf("both input and output file paths must be provided")
	}

	// Check if the input file exists and is a PDF
	if _, err := os.Stat(inputPath); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", inputPath)
	}
	if filepath.Ext(inputPath) != ".pdf" {
		return fmt.Errorf("input file must be a PDF: %s", inputPath)
	}

	// Check if the output file has a .json extension
	if filepath.Ext(outputPath) != ".json" {
		return fmt.Errorf("output file must have a .json extension: %s", outputPath)
	}

	return nil
}

// extractPDFContent reads and extracts text content from a PDF file
func extractPDFContent(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("failed to get file info: %w", err)
	}

	reader, err := pdf.NewReader(file, fileStat.Size())
	if err != nil {
		return "", fmt.Errorf("failed to create PDF reader: %w", err)
	}

	var content string
	for i := 1; i <= reader.NumPage(); i++ {
		page := reader.Page(i)
		text, err := page.GetPlainText(nil)
		if err != nil {
			log.Printf("Warning: Failed to read page %d: %v", i, err)
			continue
		}
		content += text
	}

	return content, nil
}

// writeJSONContent writes the extracted content to a JSON file
func writeJSONContent(outputPath string, content string) error {
	data := PDFContent{Content: content}
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	if err := os.WriteFile(outputPath, jsonData, FilePermission); err != nil {
		return fmt.Errorf("failed to write JSON file: %w", err)
	}

	return nil
}