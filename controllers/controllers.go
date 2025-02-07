package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"oms/database"
	"os"
	"strings"

	orders "oms/services"

	"github.com/gin-gonic/gin"
	"github.com/omniful/go_commons/csv"
)

// OrderRequest represents the API request payload
type OrderRequest struct {
	FilePath string `json:"file_path"`
}

// CreateOrder handles incoming order requests, validates them, and stores them in the database
func CreateOrder(c *gin.Context) {
	var req OrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Fatal("Error", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		// return
	}

	// Validate csv
	if err := validateCSVFilePath(req.FilePath); err != nil {
		log.Fatal("Invalid Path")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// Push event in Create Bulk Order Queue - SQS
	orders.SetProducer(c, database.Queue, req.FilePath)

	// Parse the csv
	records, err := getCsvData(req.FilePath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(records)
}

// View orders according to filters
func ViewOrders(c *gin.Context) {
	// Create a map to hold the filter parameters
	filters := make(map[string]string)

	// Bind query parameters to the filter map
	tenantID := c.DefaultQuery("tenant_id", "")
	if tenantID != "" {
		filters["tenant_id"] = tenantID
	}
	sellerID := c.DefaultQuery("seller_id", "")
	if sellerID != "" {
		filters["seller_id"] = sellerID
	}
	status := c.DefaultQuery("status", "")
	if status != "" {
		filters["status"] = status
	}
	startDate := c.DefaultQuery("start_date", "")
	if startDate != "" {
		filters["start_date"] = startDate
	}
	endDate := c.DefaultQuery("end_date", "")
	if endDate != "" {
		filters["end_date"] = endDate
	}

	// Fetch orders based on the filters
	orders, err := database.GetFilteredOrders(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the results
	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

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
