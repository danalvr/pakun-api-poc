package routes

import (
	"net/http"
	"pakun-api-poc/services"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {
	identifier := c.Query("identifier")
	if identifier == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing identifier"})
		return
	}

	identifier = identifier + "@s.whatsapp.net"

	data, err := services.GetDashboardData(identifier)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get dashboard data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total_income": data.TotalIncome,
		"total_expense": data.TotalExpense,
	})
}