package github

import "time"

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	HtmlUrl     string    `json:"html_url"`
	Language    *string   `json:"language"`
	Topics      []string  `json:"topics"`
	Fork        bool      `json:"fork"`
	Archived    bool      `json:"archived"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PushedAt    time.Time `json:"pushed_at"`
}
