package repository

import (
	"context"
	"fmt"
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
	secureDB bun.IDB
	unsafeDB bun.IDB
}

type UserSDK interface {
	CreateUser(ctx context.Context, user *User) (uint, error)
	GetUserByID(ctx context.Context, id uint) (*User, error)
	Auth(ctx context.Context, email, password string) (bool, uint, error)
	CreateUserNoSecure(ctx context.Context, user *User) error
	ListUserNoSecure(ctx context.Context, email string) ([]*User, error)
}

func NewUserRepo(secureDB bun.IDB, unsafeDB bun.IDB) UserSDK {
	return &userRepo{
		secureDB: secureDB,
		unsafeDB: unsafeDB,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, user *User) (uint, error) {
	timeNow := time.Now()
	user.CreatedAt = &timeNow

	_, err := u.secureDB.NewInsert().Model(user).Exec(ctx)
	return user.ID, err
}

func (u *userRepo) GetUserByID(ctx context.Context, id uint) (*User, error) {
	var user User

	err := u.secureDB.NewSelect().
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

	err := u.secureDB.NewSelect().
		Model(&user).
		Where("email = ?", email).
		Scan(ctx)
	if err != nil {
		return false, 0, err
	}
	isHash := security.CheckPasswordHash(password, user.Password)

	return isHash, user.ID, nil
}

func (u *userRepo) CreateUserNoSecure(ctx context.Context, user *User) error {
	timeNow := time.Now()
	user.CreatedAt = &timeNow

	query := fmt.Sprintf(
		`INSERT INTO users 
		(first_name, last_name,  password,email) 
		VALUES ('%s', '%s', '%s', '%s')`,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Email,
	)

	_, err := u.unsafeDB.ExecContext(ctx, query)
	return err
}

func (u *userRepo) ListUserNoSecure(ctx context.Context, email string) ([]*User, error) {
	query := fmt.Sprintf(`
		SELECT id, first_name, last_name, email, password, created_at 
		FROM users WHERE email = '%s'`, email)

	var listUser []*User
	rows, err := u.unsafeDB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(User)
		if err = rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		listUser = append(listUser, user)
	}
	return listUser, nil
}
