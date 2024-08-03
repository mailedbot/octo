package endpoints

import (
	"fmt"

	"github.com/mailedbot/octo"
	"github.com/mailedbot/octo/models"
	"github.com/mailedbot/octo/parser"
)

// GuildsEndpoint handles the guilds endpoint.
type GuildsEndpoint struct {
	client *octo.Client // The client used to make requests
}

// NewGuildsEndpoint creates a new instance of the GuildsEndpoint.
func NewGuildsEndpoint(client *octo.Client) *GuildsEndpoint {
	return &GuildsEndpoint{client: client}
}

// CreateGuild creates a new guild.
func (g *GuildsEndpoint) CreateGuild(guild *models.Guild) error {
	endpoint := "/v1/guilds"
	_, err := g.client.DoRequest("POST", endpoint, guild)
	return err
}

// DeleteGuild deletes a guild by ID.
func (g *GuildsEndpoint) DeleteGuildByID(guildID string) error {
	endpoint := fmt.Sprintf("/v1/guilds?guild_id=%s", guildID)
	_, err := g.client.DoRequest("DELETE", endpoint, nil)
	return err
}

// GetGuild gets a guild by ID.
func (g *GuildsEndpoint) GetGuildByID(guildID string) (*models.Guild, error) {
	endpoint := fmt.Sprintf("/v1/guilds?guild_id=%s", guildID)
	resp, err := g.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var guild models.Guild
	err = parser.ParseDataToType(resp.Data, &guild)
	if err != nil {
		return nil, err
	}

	return &guild, nil
}

// GetGuilds gets all guilds.
func (g *GuildsEndpoint) GetGuilds() ([]models.Guild, error) {
	endpoint := "/v1/guilds"
	resp, err := g.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var guilds []models.Guild
	err = parser.ParseDataToType(resp.Data, &guilds)
	if err != nil {
		return nil, err
	}

	return guilds, nil
}

// UpdateGuild updates a guild by ID.
func (g *GuildsEndpoint) UpdateGuildByID(guildID string, guild *models.Guild) error {
	endpoint := fmt.Sprintf("/v1/guilds?guild_id=%s", guildID)
	_, err := g.client.DoRequest("PUT", endpoint, guild)
	return err
}
