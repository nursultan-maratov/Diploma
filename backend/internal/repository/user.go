package repository

import (
	"context"
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
	db *bun.DB
}

type UserSDK interface {
	CreateUser(ctx context.Context, user *User) (uint, error)
	GetUser(ctx context.Context, ID uint) (*User, error)
	UpdateUser(ctx context.Context, user *User) error
	DeleteUser(ctx context.Context, ID uint) error
	ListUsers(ctx context.Context) ([]*User, error)
}

func NewUserRepo(db *bun.DB) UserSDK {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(ctx context.Context, user *User) (uint, error) {
	now := time.Now()
	user.CreatedAt = &now

	_, err := u.db.NewInsert().Model(user).Exec(ctx)
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (u *userRepo) GetUser(ctx context.Context, ID uint) (*User, error) {
	user := new(User)
	err := u.db.NewSelect().Model(user).Where("id = ?", ID).Scan(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, user *User) error {
	user.UpdatedAt = new(time.Time)
	*user.UpdatedAt = time.Now()

	_, err := u.db.NewUpdate().Model(user).Where("id = ?", user.ID).Exec(ctx)
	return err
}

func (u *userRepo) DeleteUser(ctx context.Context, ID uint) error {
	_, err := u.db.NewDelete().Model((*User)(nil)).Where("id = ?", ID).Exec(ctx)
	return err
}

func (u *userRepo) ListUsers(ctx context.Context) ([]*User, error) {
	var users []*User
	err := u.db.NewSelect().Model(&users).Scan(ctx)
	return users, err
}
