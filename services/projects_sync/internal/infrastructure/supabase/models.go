package supabase

import "time"

type Project struct {
	Uuid        string     `json:"uuid" db:"uuid"`
	Name        string     `json:"name" db:"name"`
	Description *string    `json:"description" db:"description"`
	Url         *string    `json:"url" db:"url"`
	StartDate   time.Time  `json:"start_date" db:"start_date"`
	EndDate     *time.Time `json:"end_date" db:"end_date"`
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" db:"updated_at"`
}

type Skill struct {
	Uuid      string    `json:"uuid" db:"uuid"`
	Label     string    `json:"label" db:"label"`
	Url       *string   `json:"url" db:"url"`
	ImagePath *string   `json:"image_path" db:"image_path"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type ProjectSkill struct {
	ProjectId string    `json:"project_id" db:"project_id"`
	SkillId   string    `json:"skill_id" db:"skill_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
