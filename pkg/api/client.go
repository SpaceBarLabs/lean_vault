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
	defaultBaseURL = "https://openrouter.ai/api/v1"
)

// Client represents the OpenRouter API client
type Client struct {
	baseURL      string
	provisionKey string
	debug        bool
	httpClient   *http.Client
}

// KeyResponse represents the response from key creation
type KeyResponse struct {
	Key  string `json:"key"`
	Data struct {
		Name      string  `json:"name"`
		Label     string  `json:"label,omitempty"`
		Limit     float64 `json:"limit,omitempty"`
		Disabled  bool    `json:"disabled"`
		CreatedAt string  `json:"created_at"`
		UpdatedAt string  `json:"updated_at"`
		Hash      string  `json:"hash"`
	} `json:"data"`
}

// NewClient creates a new OpenRouter API client
func NewClient(provisionKey string) *Client {
	return &Client{
		baseURL:      defaultBaseURL,
		provisionKey: provisionKey,
		httpClient:   &http.Client{},
	}
}

// SetDebug enables or disables debug mode
func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

// CreateKey creates a new API key
func (c *Client) CreateKey(name string) (*KeyResponse, error) {
	url := fmt.Sprintf("%s/keys", c.baseURL)

	payload := map[string]interface{}{
		"name": name,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.provisionKey)
	req.Header.Set("Content-Type", "application/json")

	if c.debug {
		fmt.Fprintf(os.Stderr, "DEBUG: Creating key with name: %s\n", name)
		fmt.Fprintf(os.Stderr, "DEBUG: URL: %s\n", url)
		fmt.Fprintf(os.Stderr, "DEBUG: Request body: %s\n", string(jsonData))
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	if c.debug {
		fmt.Fprintf(os.Stderr, "DEBUG: Response status: %s\n", resp.Status)
		fmt.Fprintf(os.Stderr, "DEBUG: Response body: %s\n", string(body))
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("API error: %s", string(body))
	}

	var response KeyResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("failed to parse response: %w", err)
	}

	return &response, nil
}

// RevokeKey revokes an API key
func (c *Client) RevokeKey(keyID string) error {
	url := fmt.Sprintf("%s/keys/%s", c.baseURL, keyID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+c.provisionKey)

	if c.debug {
		fmt.Fprintf(os.Stderr, "DEBUG: Revoking key with ID: %s\n", keyID)
		fmt.Fprintf(os.Stderr, "DEBUG: URL: %s\n", url)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response: %w", err)
	}

	if c.debug {
		fmt.Fprintf(os.Stderr, "DEBUG: Response status: %s\n", resp.Status)
		fmt.Fprintf(os.Stderr, "DEBUG: Response body: %s\n", string(body))
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("API error: %s", string(body))
	}

	return nil
}
