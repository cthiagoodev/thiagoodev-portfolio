package contact

type Contact struct {
	Uuid  string `json:"uuid"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

func (c *Contact) IsValid() bool {
	return c.Phone != "" && c.Email != ""
}
