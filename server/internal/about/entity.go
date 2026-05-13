package about

import "time"

type About struct {
	Uuid       string       `json:"uuid"`
	Name       string       `json:"name"`
	Bio        string       `json:"bio"`
	Photo      string       `json:"photo"`
	Curriculum string       `json:"curriculum"`
	Linkedin   string       `json:"linkedin"`
	GitHub     string       `json:"github"`
	Stack      []Technology `json:"stack"`
	City       string       `json:"city"`
	State      string       `json:"state"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  *time.Time   `json:"updated_at"`
}

func (a *About) IsValid() bool {
	return a.Name != "" &&
		a.Bio != "" &&
		a.Photo != "" &&
		a.Linkedin != "" &&
		a.GitHub != "" &&
		a.City != "" &&
		a.State != ""
}

func (a *About) HasStack() bool {
	return len(a.Stack) > 0
}

type Technology struct {
	Uuid string `json:"uuid"`
	Name string `json:"name"`
}
