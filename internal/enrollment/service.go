package enrollment

import (
	"context"
	"log"

	"github.com/ncostamagna/g_ms_domain_ex/domain"
)

type (
	Filters struct {
		UserID   string
		CourseID string
	}

	Service interface {
		Create(ctx context.Context, userID, courseID string) (*domain.Enrollment, error)
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

func (s service) Create(ctx context.Context, userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	if err := s.repo.Create(ctx, enroll); err != nil {
		return nil, err
	}

	return enroll, nil
}
