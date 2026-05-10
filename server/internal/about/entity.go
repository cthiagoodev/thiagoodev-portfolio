package about

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
}

type Technology struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
