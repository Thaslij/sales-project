package models

import "time"

type Order struct {
	ID            string `gorm:"primaryKey"`
	CustomerID    string
	Region        string
	DateOfSale    time.Time
	PaymentMethod string
}
