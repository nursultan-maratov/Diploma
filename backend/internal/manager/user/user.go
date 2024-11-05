package user

import (
	"context"
	"github.com/nursultan-maratov/Diploma.git/internal/model"
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"github.com/nursultan-maratov/Diploma.git/internal/security"
	"log"
	"time"
)

type ManagerSDK interface {
	CreateUser(ctx context.Context, user *model.UserRequest) (uint, error)
}

type Manager struct {
	user repository.UserSDK
}

func NewManager(user repository.UserSDK) *Manager {
	return &Manager{
		user: user,
	}
}

func (m *Manager) CreateUser(ctx context.Context, user *model.UserRequest) (uint, error) {
	hashPassword, err := security.HashPassword(user.Password + model.CreateUserSalt)
	if err != nil {
		return 0, err
	}

	timeNow := time.Now()
	userRepo := &repository.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashPassword,
		Phone:     user.Phone,
		Address:   user.Address,
		Status:    user.Status,
		CreatedAt: &timeNow,
	}

	id, err := m.user.CreateUser(userRepo)
	log.Println(err)
	if err != nil {
		return 0, err
	}

	return id, nil
}
