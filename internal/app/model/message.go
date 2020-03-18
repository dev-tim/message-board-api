package model

import "time"

type Message struct {
	Id           string     `csv:"id"`
	Name         string     `csv:"name"`
	Email        string     `csv:"email"`
	Text         string     `csv:"text"`
	CreationTime *time.Time `csv:"creation_time"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}
