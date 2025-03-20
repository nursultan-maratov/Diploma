package repository

import (
	"context"
	"github.com/nursultan-maratov/Diploma.git/internal/security"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	ID        uint       `json:"id" bun:",pk,autoincrement"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email" bun:",unique"`
	Password  string     `json:"password"`
	CreatedAt *time.Time `json:"created_at" bun:",nullzero,default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" bun:",nullzero"`
	DeletedAt *time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}

type userRepo struct {
	db bun.IDB
}

type UserSDK interface {
	CreateUser(ctx context.Context, user *User) (uint, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	Auth(ctx context.Context, email, password string) (bool, uint, error)
}

func NewUserRepo(db bun.IDB) UserSDK {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(ctx context.Context, user *User) (uint, error) {
	timeNow := time.Now()
	user.CreatedAt = &timeNow

	_, err := u.db.NewInsert().Model(user).Exec(ctx)
	return user.ID, err
}

func (u *userRepo) GetUserByID(ctx context.Context, id uint) (*User, error) {
	var user User

	err := u.db.NewSelect().
		Model(&user).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *userRepo) Auth(ctx context.Context, email, password string) (bool, uint, error) {
	var user User

	err := u.db.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return false, 0, err
	}
	isHash := security.CheckPasswordHash(password, user.Password)

	return isHash, user.ID, nil
}
