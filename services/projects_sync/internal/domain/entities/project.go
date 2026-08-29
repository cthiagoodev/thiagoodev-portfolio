package entities

import "time"

type Project struct {
	Uuid        string    `json:"uuid"`
	ExternalId  string    `json:"external_id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Url         string    `json:"url"`
	Skills      []Skill   `json:"skills"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
