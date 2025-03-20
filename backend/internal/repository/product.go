package repository

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

type Product struct {
	ID            int        `bun:"id,pk,autoincrement"`
	Name          string     `bun:"name"`
	Description   string     `bun:"description"`
	StockQuantity int        `bun:"stock_quantity"`
	Price         float64    `bun:"price"`
	CompanyID     int        `bun:"company_id"`
	Status        string     `bun:"status"`
	CreatedAt     time.Time  `bun:"created_at"`
	UpdatedAt     time.Time  `bun:"updated_at"`
	DeletedAt     *time.Time `bun:"deleted_at,soft_delete,nullzero"`
	Img           string     `json:"img"`
}

type productRepo struct {
	db bun.IDB
}

type ProductSDK interface {
	ListProduct(ctx context.Context) ([]*Product, error)
	GetProductByID(ctx context.Context, id int) (*Product, error)
	ListImgByIDS(ctx context.Context, listIDS []int) ([]*Product, error)
}

func NewProductRepo(db bun.IDB) ProductSDK {
	return &productRepo{db: db}
}

func (p *productRepo) ListProduct(ctx context.Context) ([]*Product, error) {
	var listProduct []*Product

	err := p.db.NewSelect().
		Model(&listProduct).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return listProduct, nil
}

func (p *productRepo) GetProductByID(ctx context.Context, id int) (*Product, error) {
	var product Product

	err := p.db.NewSelect().
		Model(&product).
		Where("id = ?", id).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepo) ListImgByIDS(ctx context.Context, listIDS []int) ([]*Product, error) {
	var products []*Product

	err := p.db.NewSelect().
		Model(&products).
		Where("id IN (?)", bun.In(listIDS)).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	return products, nil
}
