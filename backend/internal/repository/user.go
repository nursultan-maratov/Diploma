package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type User struct {
	ID        uint       `json:"id"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	Phone     string     `json:"phone"`
	Address   string     `json:"address"`
	Status    string     `json:"status"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type userRepo struct {
	db *sql.DB
}

type UserSDK interface {
	CreateUser(user *User) (uint, error)
}

func NewUserRepo(db *sql.DB) UserSDK {
	return &userRepo{
		db: db,
	}
}

func (u userRepo) CreateUser(user *User) (uint, error) {
	var ID uint
	sqlStatement := fmt.Sprintf(`INSERT INTO users (first_name, last_name, email, password, phone, address, status)
VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s') RETURNING id`,
		user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.Address, user.Status)

	err := u.db.QueryRow(sqlStatement).Scan(&ID)
	if err != nil {
		return 0, err
	}

	return ID, nil
}
