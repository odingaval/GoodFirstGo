package github

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct { //deffining a custom type client (github API client)
	httpClient *http.Client
}

func NewClient() *Client { //creating a new client instance, wrapping it inside Client
	return &Client{ //returning a pointer to it
		httpClient: &http.Client{},
	}
}

func (c *Client) SearchGoodFirstIssues() error { // method on the client
	url := `https://api.github.com/search/issues?q=label:"good-first-issue"+language:go&per_page=5`

	resp, err := c.httpClient.Get(url) //making a get request using embedded http.Client
	if err != nil {
		return fmt.Errorf("feild to make a request: %w", err)
	}
	defer resp.Body.Close() //closing the response body after the function returns to prevent resource leaks

	fmt.Println("Github API response status:", resp.Status)

	body, err := io.ReadAll(resp.Body) //Reading the response body
	if err != nil {
		return err
	}

	fmt.Println(string(body))

	return nil

}
