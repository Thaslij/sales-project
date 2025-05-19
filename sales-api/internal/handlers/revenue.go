package handlers

import (
	"net/http"
	"sales-api/internal/db"
	"sales-api/internal/refresh"

	"github.com/gin-gonic/gin"
)

func TotalRevenue(c *gin.Context) {
	start := c.Query("start")
	end := c.Query("end")

	var total float64
	db.DB.Raw(`
        SELECT SUM((unit_price * quantity) - discount)
        FROM order_items
        JOIN orders ON orders.id = order_items.order_id
        WHERE orders.date_of_sale BETWEEN ? AND ?
    `, start, end).Scan(&total)

	c.JSON(http.StatusOK, gin.H{"total_revenue": total})
}

func TriggerRefresh(c *gin.Context) {
	err := refresh.LoadCSV("data/sales_data.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "success"})
	}
}
