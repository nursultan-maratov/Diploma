package service

import (
	"github.com/nursultan-maratov/Diploma.git/internal/manager/user"
	"github.com/uptrace/bun"
)

type Service struct {
	userManager user.ManagerSDK
	repository  *Repository
}

type ServiceSDK interface {
	GetUserManager() user.ManagerSDK
	GetRepository() Repository
}

func NewService(db bun.IDB) *Service {
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

func (s *Service) GetRepository() *Repository {
	return s.repository
}
