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

func NewRepository(db bun.IDB) *Repository {
	return &Repository{
		userRepo:    repository.NewUserRepo(db),
		orderRepo:   repository.NewOrderRepo(db),
		productRepo: repository.NewProductRepo(db),
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
