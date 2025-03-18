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
	GetUser(ctx context.Context, ID uint) (*repository.User, error)
	UpdateUser(ctx context.Context, updateUser *model.UserRequest) error
	DeleteUser(ctx context.Context, ID uint) error
	ListUsers(ctx context.Context) ([]*repository.User, error)
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

	id, err := m.user.CreateUser(ctx, userRepo)
	log.Println(err)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *manager) GetUser(ctx context.Context, ID uint) (*repository.User, error) {
	user, err := m.user.GetUser(ctx, ID)
	if err != nil {
		log.Println("Ошибка получения пользователя:", err)
		return nil, err
	}

	return &repository.User{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		Address:   user.Address,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (m *manager) UpdateUser(ctx context.Context, updateUser *model.UserRequest) error {
	hashPassword, err := security.HashPassword(updateUser.Password + model.CreateUserSalt)
	if err != nil {
		log.Println("Ошибка хеширования пароля:", err)
		return err
	}

	userRepo := &repository.User{
		FirstName: updateUser.FirstName,
		LastName:  updateUser.LastName,
		Email:     updateUser.Email,
		Password:  hashPassword,
		Phone:     updateUser.Phone,
		Address:   updateUser.Address,
		Status:    updateUser.Status,
	}

	err = m.user.UpdateUser(ctx, userRepo)
	if err != nil {
		log.Println("Ошибка обновления пользователя:", err)
	}

	return err
}

func (m *manager) DeleteUser(ctx context.Context, ID uint) error {
	err := m.user.DeleteUser(ctx, ID)
	if err != nil {
		log.Println("Ошибка удаления пользователя:", err)
	}
	return err
}

func (m *manager) ListUsers(ctx context.Context) ([]*repository.User, error) {
	users, err := m.user.ListUsers(ctx)
	if err != nil {
		log.Println("Ошибка получения списка пользователей:", err)
		return nil, err
	}

	var usersResponse []*repository.User
	for _, u := range users {
		usersResponse = append(usersResponse, &repository.User{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Email:     u.Email,
			Phone:     u.Phone,
			Address:   u.Address,
			Status:    u.Status,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		})
	}

	return usersResponse, nil
}
