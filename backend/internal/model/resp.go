package model

import "time"

type OrderResp struct {
	ID          int       `json:"id"`
	ProductID   int       `json:"product_id"`
	CreatedAt   time.Time `json:"created_at"`
	Img         string    `json:"img"`
	Description string    `json:"description"`
}

type ProductResp struct {
	ID            int        `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	StockQuantity int        `json:"stock_quantity"`
	Price         float64    `json:"price"`
	CompanyID     int        `json:"company_id"`
	Status        string     `json:"status"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
	Img           string     `json:"img"`
}

type GetUserResp struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" bun:",unique"`
	Password  string `json:"password"`
}
