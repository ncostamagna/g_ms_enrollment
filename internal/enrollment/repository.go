package enrollment

import (
	"context"
	"log"

	"github.com/ncostamagna/g_ms_domain_ex/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(ctx context.Context, enroll *domain.Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

//NewRepo is a repositories handler
func NewRepo(db *gorm.DB, l *log.Logger) Repository {
	return &repo{
		db:  db,
		log: l,
	}
}

func (r *repo) Create(ctx context.Context, enroll *domain.Enrollment) error {

	if err := r.db.WithContext(ctx).Create(enroll).Error; err != nil {
		r.log.Println(err)
		return err
	}
	r.log.Println("enrollment created with id: ", enroll.ID)
	return nil
}
