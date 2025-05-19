package models

type OrderItem struct {
	ID           string `gorm:"primaryKey"`
	OrderID      string
	ProductID    string
	Quantity     int
	UnitPrice    float64
	Discount     float64
	ShippingCost float64
}
