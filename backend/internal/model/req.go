package model

type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
}

type GetUser struct {
	ID uint `json:"id"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BuyProduct struct {
	ProductID int  `json:"product_id"`
	UserID    uint `json:"user_id"`
}

type GetUserRequest struct {
	Email string `json:"email"`
}
