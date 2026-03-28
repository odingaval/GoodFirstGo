package github

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"GoodFirstGo/internal/models"
)

type Client struct { // defining a custom type client (github API client)
	httpClient *http.Client
}

func NewClient() *Client { // creating a new client instance, wrapping it inside Client
	return &Client{ // returning a pointer to it
		httpClient: &http.Client{},
	}
}

func (c *Client) SearchIssues(query string, limit int) ([]models.Issue, error) {
	perPage := strconv.Itoa(limit)
	url := fmt.Sprintf(`https://api.github.com/search/issues?q=%s&per_page=%s`, query, perPage)

	resp, err := c.httpClient.Get(url) // making a get request using embedded http.Client
	if err != nil {
		return nil, fmt.Errorf("failed to make a request: %w", err)
	}
	defer resp.Body.Close() // closing the response body after the function returns to prevent resource leaks

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body) // Reading the response body
	if err != nil {
		return nil, err
	}

	type SearchResponse struct {
		Items []models.Issue `json:"items"`
	}

	var searchResp SearchResponse
	// Unmarshal the JSON response from GitHub API into our Issue structs
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return searchResp.Items, nil
}
