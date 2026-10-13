package github

import "time"

type Project struct {
	Id          int64      `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Language    string     `json:"language"`
	Topics      []string   `json:"topics"`
	Fork        bool       `json:"fork"`
	Homepage    *string    `json:"homepage"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	PushedAt    *time.Time `json:"pushed_at"`
}
