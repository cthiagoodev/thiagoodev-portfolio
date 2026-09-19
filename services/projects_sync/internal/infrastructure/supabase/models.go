package supabase

import "time"

type Project struct {
	Uuid        string     `json:"uuid"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Url         *string    `json:"url"`
	StartDate   time.Time  `json:"start_date"`
	EndDate     *time.Time `json:"end_date"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type Skill struct {
	Uuid      string    `json:"uuid"`
	Label     string    `json:"label"`
	Url       *string   `json:"url"`
	ImagePath *string   `json:"image_path"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProjectSkill struct {
	ProjectId string    `json:"project_id"`
	SkillId   string    `json:"skill_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
