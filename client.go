package octo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

// Client represents the main API client for interacting with Mailed.
type Client struct {
	BaseURL    string       // BaseURL is the base URL for the Mailed Backend.
	HTTPClient *http.Client // HTTPClient is the HTTP client used to make requests.
	Token      string       // Token is the API token used to authenticate requests.
}

// Response represents a response from the Mailed Backend.
type Response struct {
	Code    int         `json:"code"`    // Code contains the status code of the response.
	Data    interface{} `json:"data"`    // Data contains the returned data from the Backend.
	Message string      `json:"message"` // Message contains any informative message returned by the Backend.
}

// NewClient creates a new instance of the Mailed Backend client.
func NewClient(baseURL string, token string) *Client {
	return &Client{
		BaseURL:    baseURL,
		HTTPClient: http.DefaultClient,
		Token:      token,
	}
}

// DoRequest makes an HTTP request to the Mailed Backend and returns the Backend response or an error.
// method: HTTP method (GET, POST, PUT, DELETE, etc.)
// endpoint: Backend endpoint relative to the base URL.
// body: Optional. Body data to send with the request (typically for POST and PUT requests).
func (c *Client) DoRequest(method, endpoint string, body interface{}) (*Response, error) {
	url := fmt.Sprintf("%s%s", c.BaseURL, endpoint)
	var req *http.Request
	var err error

	if body != nil {
		var jsonBody []byte
		jsonBody, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
		req, err = http.NewRequest(method, url, bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return nil, err
	}
	req.Header.Set("X-API-KEY", c.Token)
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		if _, ok := err.(net.Error); ok {
			return nil, fmt.Errorf("backend unvailable! report to the administrator")
		}

		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var jsonResponse Response
	err = json.Unmarshal(responseData, &jsonResponse)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, jsonResponse.Message)
	}

	return &jsonResponse, nil
}
