package service

import (
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"github.com/uptrace/bun"
)

type Repository struct {
	userRepo    repository.UserSDK
	orderRepo   repository.OrderSDK
	productRepo repository.ProductSDK
}

func NewRepository(secureDB bun.IDB, unsafeDB bun.IDB) *Repository {
	return &Repository{
		userRepo:    repository.NewUserRepo(secureDB, unsafeDB),
		orderRepo:   repository.NewOrderRepo(secureDB),
		productRepo: repository.NewProductRepo(secureDB),
	}
}

func (r Repository) GetUserRepo() repository.UserSDK {
	return r.userRepo
}

func (r Repository) GetOrderRepo() repository.OrderSDK {
	return r.orderRepo
}

func (r Repository) GetProductRepo() repository.ProductSDK {
	return r.productRepo
}
