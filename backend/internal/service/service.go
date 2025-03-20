package service

import (
	"github.com/nursultan-maratov/Diploma.git/internal/manager/order"
	"github.com/nursultan-maratov/Diploma.git/internal/manager/user"
	"github.com/uptrace/bun"
)

type Service struct {
	userManager  user.ManagerSDK
	orderManager order.ManagerSDK
	repository   *Repository
}

type ServiceSDK interface {
	GetUserManager() user.ManagerSDK
	GetProductManager() order.ManagerSDK
	GetRepository() Repository
}

func NewService(db bun.IDB) *Service {
	repository := NewRepository(db)

	userManager := user.NewManager(repository.GetUserRepo())
	orderManager := order.NewManager(repository.GetOrderRepo(), repository.GetProductRepo())
	return &Service{
		userManager:  userManager,
		orderManager: orderManager,
		repository:   repository,
	}
}

func (s *Service) GetUserManager() user.ManagerSDK {
	return s.userManager
}

func (s *Service) GetProductManager() order.ManagerSDK {
	return s.orderManager
}

func (s *Service) GetRepository() *Repository {
	return s.repository
}
