package service

import (
	"database/sql"
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
)

type Repository struct {
	userRepo repository.UserSDK
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		userRepo: repository.NewUserRepo(db),
	}
}

func (r Repository) GetUserRepo() repository.UserSDK {
	return r.userRepo
}
