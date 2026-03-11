package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"GoodFirstGo/internal/models"
)

type Client struct { // deffining a custom type client (github API client)
	httpClient *http.Client
}

func NewClient() *Client { // creating a new client, wrapping it inside Client
	return &Client{ // returning a pointer to it
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) SearchGoodFirstIssues() ([]models.Issue, error) { // method on the client
	url := `https://api.github.com/search/issues?q=label:"good-first-issue"+language:go:page=5`

	fmt.Println("Making a request to GitHub API...")

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make a request: %w", err)
	}
	defer resp.Body.Close()

	var result models.SearchResponses

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return result.Items, nil
}
