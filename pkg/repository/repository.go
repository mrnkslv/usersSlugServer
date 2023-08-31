package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrnkslv/user-segmentation-service/models"
)

type Slugger interface {
	CreateSlug(slug models.Slug) (int64, error)
	DeleteSlug(slug models.Slug) (int64, error)
}

type User interface {
	AddUserToSlug(data models.AddSlugstoUser) (int64, error)
	GetActiveSlugsByID(userId int64) ([]models.Slug, error)
}

type Repository struct {
	Slugger
	User
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Slugger: NewSlugPostgres(db),
		User:    NewUserPostgres(db),
	}
}