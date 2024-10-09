package repository

type User struct {
	ID           uint   `json:"id"`
	UserName     string `json:"user_name"`
	Email        string `json:"email"`
	HashPassword string `json:"hash_password"`
}
