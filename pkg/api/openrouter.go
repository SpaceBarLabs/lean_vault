package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	// BaseURL is the base URL for the OpenRouter API
	BaseURL = "https://openrouter.ai/api/v1"
)

// Client represents an OpenRouter API client
type Client struct {
	httpClient      *http.Client
	provisioningKey string
	debug           bool
}

// NewClient creates a new OpenRouter API client
func NewClient(provisioningKey string) *Client {
	return &Client{
		httpClient:      &http.Client{},
		provisioningKey: provisioningKey,
		debug:           false,
	}
}

// SetDebug enables or disables debug mode
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// CreateKeyResponse represents the response from the create key API
type CreateKeyResponse struct {
	Data struct {
		Name      string  `json:"name"`
		Label     string  `json:"label"`
		Limit     float64 `json:"limit,omitempty"`
		Disabled  bool    `json:"disabled"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		Hash      string  `json:"hash"`
		Key       string  `json:"key"`
	} `json:"data"`
}

// CreateKey creates a new OpenRouter API key
func (c *Client) CreateKey(name string) (*CreateKeyResponse, error) {
	url := fmt.Sprintf("%s/keys", BaseURL)

	// Prepare request body
	payload := map[string]interface{}{
		"name": name,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	// Create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Set headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.provisioningKey))
	req.Header.Set("Content-Type", "application/json")

	if c.debug {
		fmt.Fprintf(os.Stderr, "DEBUG: Sending request to %s\n", url)
		fmt.Fprintf(os.Stderr, "DEBUG: Request headers: %v\n", req.Header)
		fmt.Fprintf(os.Stderr, "DEBUG: Request body: %s\n", string(body))
	}

	// Send request
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if c.debug {
		fmt.Fprintf(os.Stderr, "DEBUG: Response status: %s\n", resp.Status)
		fmt.Fprintf(os.Stderr, "DEBUG: Response body: %s\n", string(respBody))
	}

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		var errorResp struct {
			Error string `json:"error"`
		}
		if err := json.Unmarshal(respBody, &errorResp); err == nil && errorResp.Error != "" {
			return nil, fmt.Errorf("API error: %s", errorResp.Error)
		}
		return nil, fmt.Errorf("API request failed with status %d: %s", resp.StatusCode, string(respBody))
	}

	// Parse response
	var result CreateKeyResponse
	if err := json.Unmarshal(respBody, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &result, nil
}
