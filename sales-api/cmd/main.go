package main

import (
	"log"
	"sales-api/internal/db"
	"sales-api/internal/handlers"
	"sales-api/internal/refresh"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func main() {
	db.InitDB()
	r := gin.Default()

	r.GET("/api/revenue/total", handlers.TotalRevenue)
	r.POST("/api/refresh", handlers.TriggerRefresh)

	c := cron.New()
	c.AddFunc("@daily", func() {
		log.Println("Running daily refresh")
		refresh.LoadCSV("data/sales_data.csv")
	})
	c.Start()

	r.Run()
}
