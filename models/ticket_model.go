package models

import "time"

type Ticket struct {
	Guild     string    `json:"guild"`      // Guild ID
	TicketID  string    `json:"ticket_id"`  // Ticket ID
	CreatedBy string    `json:"created_by"` // User ID of the ticket creator
	Channel   string    `json:"channel"`    // Channel ID of the ticket
	CreatedAt time.Time `json:"created_at"` // Unix timestamp of when the ticket was created
}
