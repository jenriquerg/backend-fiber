package models

import "time"

type RequestLog struct {
	ID         uint      `gorm:"primaryKey"`
	Method     string
	Path       string
	StatusCode int
	DurationMs int64
	IPAddress  string
	UserID     string
	IntCode    string
	CreatedAt  time.Time
}
