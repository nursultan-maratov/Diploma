package repository

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"time"
)

type User struct {
	ID        uint       `json:"id" bun:",pk,autoincrement"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email" bun:",unique"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at" bun:",nullzero,default:current_timestamp"`
	UpdatedAt *time.Time `json:"updated_at" bun:",nullzero"`
	DeletedAt *time.Time `json:"deleted_at" bun:",soft_delete,nullzero"`
}

type userRepo struct {
	db bun.IDB
}

type UserSDK interface {
	CreateUser(ctx context.Context, user *User) (uint, error)
}

func NewUserRepo(db bun.IDB) UserSDK {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(ctx context.Context, user *User) (uint, error) {
	now := time.Now()
	user.CreatedAt = &now

	query := fmt.Sprintf(
		`INSERT INTO users (first_name, last_name, email, password, phone, address, status, created_at)
		 VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s') RETURNING id`,
		user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.Address, user.Status, user.CreatedAt.Format(time.RFC3339),
	)

	var id uint
	err := u.db.QueryRowContext(ctx, query).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
