package model

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Address   string `json:"address"`
	Status    string `json:"status"`
}

type GetUser struct {
	ID uint `json:"id"`
}
