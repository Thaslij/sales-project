package models

type Product struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Category    string
	Description string
}
