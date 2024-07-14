package models

import "time"

type Tag struct {
	Guild     int       `json:"guild"`      // Guild ID
	Name      string    `json:"name"`       // Tag name
	Text      string    `json:"text"`       // Tag text
	CreatedAt time.Time `json:"created_at"` // Unix timestamp of when the tag was created
}
