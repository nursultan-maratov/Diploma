package repository

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Order struct {
	ID         int        `bun:"id,pk,autoincrement"`
	UserID     int        `bun:"user_id"`
	ProductID  int        `bun:"product_id"`
	TotalPrice float64    `bun:"total_price"`
	CreatedAt  time.Time  `bun:"created_at"`
	UpdatedAt  time.Time  `bun:"updated_at"`
	DeletedAt  *time.Time `bun:"deleted_at,soft_delete,nullzero"`
}

type orderRepo struct {
	db bun.IDB
}

type OrderSDK interface {
	BuyProduct(ctx context.Context, order Order) (int, error)
	ListOrder(ctx context.Context, id int) ([]*Order, error)
}

func NewOrderRepo(db bun.IDB) OrderSDK {
	return &orderRepo{db: db}
}
func (u *orderRepo) BuyProduct(ctx context.Context, order Order) (int, error) {
	err := u.db.NewInsert().
		Model(&order).
		Scan(ctx)

	if err != nil {
		return 0, err
	}

	return order.ID, nil
}

func (u *orderRepo) ListOrder(ctx context.Context, id int) ([]*Order, error) {
	var listOrder []*Order

	err := u.db.NewSelect().
		Model(&listOrder).
		Where("user_id = ?", id).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return listOrder, nil
}
