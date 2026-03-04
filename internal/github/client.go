package github

import (
	"fmt"
	"net/http"
)

type Client struct { //deffining a custom type client (github API client)
	httpClient *http.Client
}

func NewClient() *Client { //creating a new client, wrapping it inside Client
	return &Client{ //returning a pointer to it
		httpClient: &http.Client{},
	}
}

func (c *Client) SearchGoodFirstIssues() error { // method on the client
	url := `https://api.github.com/search/issues?q=label:"good first issue"+language:go`

	resp, err := c.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("faile to make a request: %w", err)
	}
	defer resp.Body.Close()

	fmt.Println("Github API response status:", resp.Status)

	return nil


}