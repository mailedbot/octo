package models

import "time"

type Premium struct {
	ID             string    `json:"id"`              // The id from the user who bought the premium
	Guilds         []string  `json:"guild"`           // The guilds who the user has assigned the premium to
	PremiumExpires time.Time `json:"premium_expires"` // The time when the premium expires
	CreatedAt      time.Time `json:"created_at"`      // Unix timestamp of when the premium was created
}

// IsPremium checks if a guild has premium
func (p *Premium) IsPremium(guildID string) bool {
	for _, v := range p.Guilds {
		if v == guildID {
			return true
		}
	}
	return false
}

// HasPremium checks if the user has premium
func (p *Premium) HasPremium() bool {
	if p.PremiumExpires.IsZero() { //If time is January 1, year 1, 00:00:00 UTC.
		return true
	}

	return time.Now().Before(p.PremiumExpires)
}
