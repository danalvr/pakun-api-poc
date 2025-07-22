package routes

import (
	"net/http"
	"pakun-api-poc/services"

	"github.com/gin-gonic/gin"
)

func GetTransactionHistory(c *gin.Context) {
	identifier := c.Query("identifier")
	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing identifier"})
		return
	}

	println("sebelum", identifier)
	identifier = identifier + "@s.whatsapp.net"
	println("sesudah", identifier)


	transactions, err := services.GetTransactions(identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}