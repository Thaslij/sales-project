package main

import (
	"log"
	"sales-api/internal/db"
	"sales-api/internal/refresh"
)

func main() {
	db.InitDB()
	err := refresh.LoadCSV("data/sales_data.csv")
	if err != nil {
		log.Fatal("Failed to load CSV:", err)
	}
}
