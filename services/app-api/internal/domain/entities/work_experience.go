package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type WorkExperience struct {
	UUID        string
	Role        string
	Description string
	Company     string
	StartDate   time.Time
	EndDate     *time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewWorkExperience(
	uuid string,
	role string,
	description string,
	company string,
	startDate time.Time,
	endDate *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*WorkExperience, error) {
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	workExperience := &WorkExperience{
		UUID:        strings.TrimSpace(uuid),
		Role:        strings.TrimSpace(role),
		Description: strings.TrimSpace(description),
		Company:     strings.TrimSpace(company),
		StartDate:   startDate,
		EndDate:     endDate,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if err := workExperience.Validate(); err != nil {
		return nil, err
	}

	return workExperience, nil
}

func (w *WorkExperience) Validate() error {
	if strings.TrimSpace(w.Role) == "" {
		return fmt.Errorf("%w: role is required", errors.ErrDomainEntityValidation)
	}

	if strings.TrimSpace(w.Description) == "" {
		return fmt.Errorf("%w: description is required", errors.ErrDomainEntityValidation)
	}

	if strings.TrimSpace(w.Company) == "" {
		return fmt.Errorf("%w: company is required", errors.ErrDomainEntityValidation)
	}

	if w.StartDate.IsZero() {
		return fmt.Errorf("%w: start date is required", errors.ErrDomainEntityValidation)
	}

	if w.EndDate != nil && w.EndDate.Before(w.StartDate) {
		return fmt.Errorf("%w: end date cannot be before start date", errors.ErrDomainEntityValidation)
	}

	return nil
}
