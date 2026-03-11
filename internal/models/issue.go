package models

type Issue struct {
	Title string `json:"title"` // Maps JSON field to Go fields
	URL   string `json:"html_url"`
}

type SearchResponses struct {
	Items []Issue `json:"items"`
}
