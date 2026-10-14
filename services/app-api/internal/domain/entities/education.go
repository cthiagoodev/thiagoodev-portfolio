package entities

import (
	"fmt"
	"strings"
	"time"

	"github.com/cthiagoodev/portfolio/services/app-api/internal/domain/errors"
)

type Education struct {
	UUID                   string
	Course                 string
	Description            string
	EducationalInstitution string
	StartDate              time.Time
	EndDate                *time.Time
	CreatedAt              time.Time
	UpdatedAt              time.Time
}

func NewEducation(
	uuid string,
	course string,
	description string,
	educationalInstitution string,
	startDate time.Time,
	endDate *time.Time,
	createdAt time.Time,
	updatedAt time.Time,
) (*Education, error) {
	if createdAt.IsZero() {
		createdAt = time.Now()
	}

	if updatedAt.IsZero() {
		updatedAt = time.Now()
	}

	education := &Education{
		UUID:                   strings.TrimSpace(uuid),
		Course:                 strings.TrimSpace(course),
		Description:            strings.TrimSpace(description),
		EducationalInstitution: strings.TrimSpace(educationalInstitution),
		StartDate:              startDate,
		EndDate:                endDate,
		CreatedAt:              createdAt,
		UpdatedAt:              updatedAt,
	}

	if err := education.Validate(); err != nil {
		return nil, err
	}

	return education, nil
}

func (e *Education) Validate() error {
	if strings.TrimSpace(e.Course) == "" {
		return fmt.Errorf("%w: course is required", errors.ErrDomainEntityValidation)
	}

	if strings.TrimSpace(e.Description) == "" {
		return fmt.Errorf("%w: description is required", errors.ErrDomainEntityValidation)
	}

	if strings.TrimSpace(e.EducationalInstitution) == "" {
		return fmt.Errorf("%w: educational institution is required", errors.ErrDomainEntityValidation)
	}

	if e.StartDate.IsZero() {
		return fmt.Errorf("%w: start date is required", errors.ErrDomainEntityValidation)
	}

	if e.EndDate != nil && e.EndDate.Before(e.StartDate) {
		return fmt.Errorf("%w: end date cannot be before start date", errors.ErrDomainEntityValidation)
	}

	return nil
}
