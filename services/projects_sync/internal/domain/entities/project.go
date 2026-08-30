package entities

import "time"

type Project struct {
	Uuid        string    `json:"uuid"`
	ExternalId  string    `json:"external_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Url         string    `json:"url"`
	Languages   []string  `json:"languages"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
