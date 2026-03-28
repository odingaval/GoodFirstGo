package models

type Issue struct {
	ID      int    `json:"id"`
	Number  int    `json:"number"`
	Title   string `json:"title"`
	HTMLURL string `json:"html_url"`
	Repo    struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
}
