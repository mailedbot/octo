package endpoints

import (
	"fmt"

	"github.com/mailedbot/octo"
	"github.com/mailedbot/octo/models"
	"github.com/mailedbot/octo/parser"
)

// TagsEndpoint handles the tags endpoint.
type TagsEndpoint struct {
	client *octo.Client // The client used to make requests
}

// NewTagsEndpoint creates a new instance of the TagsEndpoint.
func NewTagsEndpoint(client *octo.Client) *TagsEndpoint {
	return &TagsEndpoint{client: client}
}

// CreateTag creates a new tag.
func (t *TagsEndpoint) CreateTag(tag *models.Tag) error {
	endpoint := "/v1/tags"
	_, err := t.client.DoRequest("POST", endpoint, tag)
	return err
}

// DeleteTag deletes a tag.
func (t *TagsEndpoint) DeleteTag(guildID int, tagName string) error {
	endpoint := fmt.Sprintf("/v1/tags/%d/%s", guildID, tagName)
	_, err := t.client.DoRequest("DELETE", endpoint, nil)
	return err
}

// GetTag gets a tag by name.
func (t *TagsEndpoint) GetTag(guildID int, tagName string) (*models.Tag, error) {
	endpoint := fmt.Sprintf("/v1/tags/%d/%s", guildID, tagName)
	resp, err := t.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var tag models.Tag
	err = parser.ParseDataToType(resp.Data, &tag)
	if err != nil {
		return nil, err
	}

	return &tag, nil
}

// GetTags gets all tags.
func (t *TagsEndpoint) GetTags(guildID int) ([]models.Tag, error) {
	endpoint := fmt.Sprintf("/v1/tags/%d", guildID)
	resp, err := t.client.DoRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var tags []models.Tag
	err = parser.ParseDataToType(resp.Data, &tags)
	if err != nil {
		return nil, err
	}

	return tags, nil
}

// UpdateTag updates a tag.
func (t *TagsEndpoint) UpdateTag(tag *models.Tag) error {
	endpoint := "/v1/tags"
	_, err := t.client.DoRequest("PUT", endpoint, tag)
	return err
}
