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
	CreateUser(ctx context.Context, req *model.UserRequest) (uint, error)
	Auth(ctx context.Context, req *model.Auth) (bool, uint, error)
	GetUser(ctx context.Context, req *model.GetUserRequest) ([]*model.GetUserResp, error)
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
	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		log.Println("Ошибка при хешировании пароля:", err)
		return 0, err
	}

	timeNow := time.Now()
	userRepo := &repository.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  hashedPassword,
		CreatedAt: &timeNow,
	}

	id, err := m.user.CreateUser(ctx, userRepo)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (m *manager) GetUser(ctx context.Context, req *model.GetUserRequest) ([]*model.GetUserResp, error) {

	listUser, err := m.user.ListUserNoSecure(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.GetUserResp, len(listUser))

	for i, user := range listUser {
		resp[i] = &model.GetUserResp{
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
		}
	}

	return resp, nil

}

func (m *manager) Auth(ctx context.Context, req *model.Auth) (bool, uint, error) {
	isHash, id, err := m.user.Auth(ctx, req.Email, req.Password)
	if err != nil {
		return false, 0, err
	}

	return isHash, id, nil
}
