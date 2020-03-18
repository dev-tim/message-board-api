package model

import "time"

type Message struct {
	Id           string
	Name         string
	Email        string
	Text         string
	CreationTime *time.Time
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
