package endpoints

import (
	"fmt"

	"github.com/mailedbot/octo"
	"github.com/mailedbot/octo/models"
	"github.com/mailedbot/octo/parser"
)

// TicketsEndpoint handles the tickets endpoint.
type TicketsEndpoint struct {
	client *octo.Client // The client used to make requests
}

// NewTicketsEndpoint creates a new instance of the TicketsEndpoint.
func NewTicketsEndpoint(client *octo.Client) *TicketsEndpoint {
	return &TicketsEndpoint{client: client}
}

// CreateTicket creates a new ticket.
func (t *TicketsEndpoint) CreateTicket(ticketModel *models.Ticket) (*models.Ticket, error) {
	endpoint := "/v1/tickets"
	resp, err := t.client.DoRequest("POST", endpoint, ticketModel)
	if err != nil {
		return nil, err
	}

	var ticket models.Ticket
	err = parser.ParseDataToType(resp.Data, &ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

// DeleteTicket deletes a ticket.
func (t *TicketsEndpoint) DeleteTicket(guild int, channel int, ticketID string) error {
	endpoint := fmt.Sprintf("/v1/tickets/%d/%d/%s", guild, channel, ticketID)
	_, err := t.client.DoRequest("DELETE", endpoint, nil)
	return err
}

// GetTicket gets a ticket by ID.
func (t *TicketsEndpoint) GetTicket(guild int, ticketID string) (*models.Ticket, error) {
	endpoint := fmt.Sprintf("/v1/tickets/%d/%s", guild, ticketID)
	resp, err := t.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var ticket models.Ticket
	err = parser.ParseDataToType(resp.Data, &ticket)
	if err != nil {
		return nil, err
	}

	return &ticket, nil
}

// GetTickets gets all tickets.
func (t *TicketsEndpoint) GetTickets(guild int) ([]models.Ticket, error) {
	endpoint := fmt.Sprintf("/v1/tickets/%d", guild)
	resp, err := t.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var tickets []models.Ticket
	err = parser.ParseDataToType(resp.Data, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// GetTicketsByUserId gets all tickets by user ID.
func (t *TicketsEndpoint) GetTicketsByUserId(userID string) ([]models.Ticket, error) {
	endpoint := fmt.Sprintf("/v1/tickets/user/%s", userID)
	resp, err := t.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var tickets []models.Ticket
	err = parser.ParseDataToType(resp.Data, &tickets)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

// UpdateTicket updates a ticket.
func (t *TicketsEndpoint) UpdateTicket(ticket *models.Ticket) error {
	endpoint := "/v1/tickets"
	_, err := t.client.DoRequest("PUT", endpoint, ticket)
	return err
}
