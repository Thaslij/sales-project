package models

type Customer struct {
	ID      string `gorm:"primaryKey"`
	Name    string
	Email   string
	Address string
}
