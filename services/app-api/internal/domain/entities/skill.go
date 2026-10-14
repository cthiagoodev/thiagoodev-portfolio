package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type Skill struct {
	UUID      string
	Label     string
	URL       string
	ImagePath string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSkill(
	uuid string,
	label string,
	url string,
	imagePath string,
	createdAt time.Time,
	updatedAt time.Time,
) (*Skill, error) {
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	skill := &Skill{
		UUID:      strings.TrimSpace(uuid),
		Label:     strings.TrimSpace(label),
		URL:       strings.TrimSpace(url),
		ImagePath: strings.TrimSpace(imagePath),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if err := skill.Validate(); err != nil {
		return nil, err
	}

	return skill, nil
}

func (s *Skill) Validate() error {
	if strings.TrimSpace(s.Label) == "" {
		return fmt.Errorf("%w: label is required", errors.ErrDomainEntityValidation)
	}
	return nil
}
