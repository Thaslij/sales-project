package models

import "time"

type RefreshLog struct {
	ID        uint `gorm:"primaryKey"`
	Timestamp time.Time
	Success   bool
	Message   string
}
