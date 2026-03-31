package github

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/odingaval/GoodFirstGo/internal/models"
)

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	transport := &http.Transport{
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		ForceAttemptHTTP2:     true,
	}

	return &Client{
		httpClient: &http.Client{
			Timeout:   30 * time.Second,
			Transport: transport,
		},
	}
}

func (c *Client) SearchIssues(query string, limit int) ([]models.Issue, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	perPage := strconv.Itoa(limit)
	encodedQuery := url.QueryEscape(query)
	searchUrl := fmt.Sprintf(`https://api.github.com/search/issues?q=%s&per_page=%s`, encodedQuery, perPage)

	req, err := http.NewRequestWithContext(ctx, "GET", searchUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	token := os.Getenv("GITHUB_TOKEN")
	if token != "" {
		req.Header.Add("Authorization", "token "+token)
	}
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	req.Header.Add("User-Agent", "GoodFirstGo/1.0")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	type SearchResponse struct {
		Items []models.Issue `json:"items"`
	}

	var searchResp SearchResponse
	if err := json.Unmarshal(body, &searchResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return searchResp.Items, nil
}
