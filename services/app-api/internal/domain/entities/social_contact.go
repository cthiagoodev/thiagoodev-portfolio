package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type SocialContact struct {
	UUID      string
	Label     string
	URL       string
	ImagePath string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewSocialContact(
	uuid string,
	label string,
	url string,
	imagePath string,
	createdAt time.Time,
	updatedAt time.Time,
) (*SocialContact, error) {
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	contact := &SocialContact{
		UUID:      strings.TrimSpace(uuid),
		Label:     strings.TrimSpace(label),
		URL:       strings.TrimSpace(url),
		ImagePath: strings.TrimSpace(imagePath),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	if err := contact.Validate(); err != nil {
		return nil, err
	}

	return contact, nil
}

func (s *SocialContact) Validate() error {
	if strings.TrimSpace(s.Label) == "" {
		return fmt.Errorf("%w: label is required", errors.ErrDomainEntityValidation)
	}

	if strings.TrimSpace(s.URL) == "" {
		return fmt.Errorf("%w: url is required", errors.ErrDomainEntityValidation)
	}

	return nil
}
