package controllers

import (
	"log"
	"log-processor/payload/response"
	"log-processor/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetGroupedRowLogs handles API requests using Gin Context with sorting
func GetGroupedRowLogs(c *gin.Context) {
	sortOrder := c.DefaultQuery("sort", "desc")

	groupedLogs, err := services.GetGroupedRowLogs(sortOrder)
	if err != nil {
		log.Printf("Error fetching grouped row logs: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve grouped row logs"})
		return
	}

	responseData := response.Response[map[string][]response.RowLogResponseDTO]{
		Message: "Successfully retrieved grouped row logs",
		Data:    &groupedLogs,
	}

	c.JSON(http.StatusOK, responseData)
}
