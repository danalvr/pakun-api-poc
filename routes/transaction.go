package routes

import (
	"net/http"
	"pakun-api-poc/services"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTransactionHistory(c *gin.Context) {
	identifier := c.Query("identifier")
	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing identifier"})
		return
	}

	identifier = identifier + "@s.whatsapp.net"


	transactions, err := services.GetTransactions(identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

func GetFinanceSummary(c *gin.Context) {
	identifier := c.Query("identifier")
	fromStr := c.Query("from")
	toStr := c.Query("to")

	if identifier == "" || fromStr == "" || toStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameters"})
		return
	}

	from, err := time.Parse("2006-01", fromStr)
	to, err2 := time.Parse("2006-01", toStr)
	if err != nil || err2 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM format."})
		return
	}

	identifier += "@s.whatsapp.net"

	summary, err := services.GetMonthlySummary(identifier, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch summary"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"summary": summary,
	})
}