package main

import "time"

type Event struct {
	ID              int64           `gorm:"column:id;primary_key" json:"id"`
	Title           string          `gorm:"column:title" json:"title"`
	Description     string          `gorm:"column:description" json:"description"`
	PublishedStatus PublishedStatus `gorm:"column:published_status" json:"published_status"`
	CreatedAt       time.Time       `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"column:updated_at" json:"updated_at"`
}

// PublishedStatus identifies event published status.
type PublishedStatus int32

const (
	// Draft ...
	Draft PublishedStatus = 0
	// Published ...
	Published PublishedStatus = 1
)
