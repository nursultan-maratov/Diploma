package service

import (
	"database/sql"
	"github.com/nursultan-maratov/Diploma.git/internal/manager/user"
)

type Service struct {
	userManager user.ManagerSDK
	repository  *Repository
}

type ServiceSDK interface {
	GetUserManager() user.ManagerSDK
}

func NewService(db *sql.DB) *Service {
	repository := NewRepository(db)

	userManager := user.NewManager(repository.GetUserRepo())
	return &Service{
		userManager: userManager,
		repository:  repository,
	}
}

func (s *Service) GetUserManager() user.ManagerSDK {
	return s.userManager
}
