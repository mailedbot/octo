package endpoints

import (
	"fmt"

	"github.com/mailedbot/octo"
	"github.com/mailedbot/octo/models"
	"github.com/mailedbot/octo/parser"
)

// PremiumsEndpoint handles the premiums endpoint.
type PremiumsEndpoint struct {
	client *octo.Client // The client used to make requests
}

// NewPremiumsEndpoint creates a new instance of the PremiumsEndpoint.
func NewPremiumsEndpoint(client *octo.Client) *PremiumsEndpoint {
	return &PremiumsEndpoint{client: client}
}

// CreatePremium creates a new premium.
func (p *PremiumsEndpoint) CreatePremium(premium *models.Premium) error {
	endpoint := "/v1/premiums"
	_, err := p.client.DoRequest("POST", endpoint, premium)
	return err
}

// DeletePremium deletes a premium.
func (p *PremiumsEndpoint) DeletePremium(premiumID string) error {
	endpoint := fmt.Sprintf("/v1/premiums/%s", premiumID)
	_, err := p.client.DoRequest("DELETE", endpoint, nil)
	return err
}

// GetPremium gets a premium by ID.
func (p *PremiumsEndpoint) GetPremium(premiumID string) (*models.Premium, error) {
	endpoint := fmt.Sprintf("/v1/premiums/%s", premiumID)
	resp, err := p.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var premium models.Premium
	err = parser.ParseDataToType(resp.Data, &premium)
	if err != nil {
		return nil, err
	}

	return &premium, nil
}

// GetPremiums gets all premiums.
func (p *PremiumsEndpoint) GetPremiums() ([]models.Premium, error) {
	endpoint := "/v1/premiums"
	resp, err := p.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var premiums []models.Premium
	err = parser.ParseDataToType(resp.Data, &premiums)
	if err != nil {
		return nil, err
	}

	return premiums, nil
}

// UpdatePremium updates a premium.
func (p *PremiumsEndpoint) UpdatePremium(premiumID string, premium *models.Premium) error {
	endpoint := fmt.Sprintf("/v1/premiums/%s", premiumID)
	_, err := p.client.DoRequest("PUT", endpoint, premium)
	return err
}

// AssignPremium assigns a premium to a user.
func (p *PremiumsEndpoint) AssignPremium(premiumID string, guildID int) error {
	endpoint := fmt.Sprintf("/v1/premiums/%s/assign", premiumID)
	_, err := p.client.DoRequest("POST", endpoint, map[string]int{"guild": guildID})
	return err
}

// UnassignPremium unassigns a premium from a user.
func (p *PremiumsEndpoint) UnassignPremium(premiumID string, guildID int) error {
	endpoint := fmt.Sprintf("/v1/premiums/%s/unassign", premiumID)
	_, err := p.client.DoRequest("POST", endpoint, map[string]int{"guild": guildID})
	return err
}
