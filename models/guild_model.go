package models

import "time"

type Guild struct {
	Guild              string    `json:"guild"`               // Guild ID
	AccessRoles        []string  `json:"access_roles"`        // Staff roles (allowed to use the bot)
	PingRoles          []string  `json:"ping_roles"`          // Roles to ping when a ticket is created
	BlacklistedMembers []string  `json:"blacklisted_members"` // Blacklisted members (cannot create tickets)
	WelcomeText        string    `json:"welcome_text"`        // Welcome text for created tickets
	GoodbyeText        string    `json:"goodbye_text"`        // Goodbye text for closed tickets
	Anonymous          bool      `json:"anonymous"`           // Anonymous tickets (hide the staff member who claimed the ticket)
	LoggingChannel     string    `json:"logging_channel"`     // Logging channel for ticket creations and deletions
	CreatedAt          time.Time `json:"created_at"`          // Unix timestamp of when the guild was created
}
