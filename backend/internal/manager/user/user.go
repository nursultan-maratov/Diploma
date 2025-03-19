package user

import (
	"context"
	"github.com/nursultan-maratov/Diploma.git/internal/model"
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"log"
	"time"
)

type ManagerSDK interface {
	CreateUser(ctx context.Context, user *model.UserRequest) (uint, error)
}

type manager struct {
	user repository.UserSDK
}

func NewManager(user repository.UserSDK) ManagerSDK {
	return &manager{
		user: user,
	}
}

func (m *manager) CreateUser(ctx context.Context, user *model.UserRequest) (uint, error) {

	timeNow := time.Now()
	userRepo := &repository.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		Status:    user.Status,
		CreatedAt: &timeNow,
	}

	id, err := m.user.CreateUser(ctx, userRepo)
	log.Println(err)
	if err != nil {
		return 0, err
	}

	return id, nil
}
