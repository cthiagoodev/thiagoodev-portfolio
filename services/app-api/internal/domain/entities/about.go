package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type About struct {
	UUID        string    `json:"uuid" db:"uuid"`
	Name        string    `json:"name" db:"name"`
	Phone       string    `json:"phone" db:"phone"`
	Description string    `json:"description" db:"description"`
	Text        string    `json:"text" db:"text"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewAbout(
	uuid string,
	name string,
	phone string,
	description string,
	text string,
	createdAt time.Time,
	updatedAt time.Time,
) (*About, error) {
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	about := &About{
		UUID:        strings.TrimSpace(uuid),
		Name:        strings.TrimSpace(name),
		Phone:       strings.TrimSpace(phone),
		Description: strings.TrimSpace(description),
		Text:        strings.TrimSpace(text),
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if err := about.Validate(); err != nil {
		return nil, err
	}

	return about, nil
}

func (a *About) Validate() error {
	if strings.TrimSpace(a.Name) == "" {
		return fmt.Errorf("%w: name is required", errors.ErrDomainEntityValidation)
	}

	if len(strings.TrimSpace(a.Phone)) > 20 {
		return fmt.Errorf("%w: phone must have at most 20 characters", errors.ErrDomainEntityValidation)
	}

	return nil
}
