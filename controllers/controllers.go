package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"oms/models"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/csv"
)

// CreateOrder handles incoming order requests, validates them, and stores them in the database
func CreateOrder(c *gin.Context) {
	var orderReq models.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		log.Fatal("Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}

	// Validate csv
	if err := validateCSVFilePath(orderReq.Path); err != nil {
		log.Fatal("Invalid Path")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Push event in Create Bulk Order Queue - SQS (in future)

	// Parse the csv
	records, err := getCsvData(orderReq.Path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)
}

// func ViewOrder(c *gin.Context) {

// }

func getCsvData(path string) ([]csv.Records, error) {
	fmt.Println(path)
	fmt.Println("path")
	Csv, err := csv.NewCommonCSV(
		csv.WithBatchSize(100),
		csv.WithSource(csv.Local),
		csv.WithLocalFileInfo(path),
		csv.WithHeaderSanitizers(csv.SanitizeAsterisks, csv.SanitizeToLower),
		csv.WithDataRowSanitizers(csv.SanitizeSpace, csv.SanitizeToLower),
	)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = Csv.InitializeReader(context.TODO())
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var records []csv.Records
	for !Csv.IsEOF() {
		var row csv.Records
		row, err := Csv.ReadNextBatch()
		if err != nil {
			log.Fatal(err)
		}
		// Process the records row-wise
		records = append(records, row)
		fmt.Println(row)
	}

	return records, nil

}

// Validate CSV Path
func validateCSVFilePath(filePath string) error {
	// Extension check
	if !strings.HasSuffix(filePath, ".csv") {
		return fmt.Errorf("invalid file: not a CSV file")
	}

	// File Existence check
	info, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("file does not exist")
	}
	if err != nil {
		return fmt.Errorf("error accessing file: %v", err)
	}

	// Regular file check
	if info.IsDir() {
		return fmt.Errorf("invalid file: path points to a directory, not a CSV file")
	}

	return nil
}
