package service

import (
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"github.com/uptrace/bun"
)

type Repository struct {
	userRepo repository.UserSDK
}

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		userRepo: repository.NewUserRepo(db),
	}
}

func (r Repository) GetUserRepo() repository.UserSDK {
	return r.userRepo
}
