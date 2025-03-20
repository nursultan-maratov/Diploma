package order

import (
	"context"
	"fmt"
	"github.com/nursultan-maratov/Diploma.git/internal/model"
	"github.com/nursultan-maratov/Diploma.git/internal/repository"
	"log"
	"time"
)

type ManagerSDK interface {
	BuyProduct(ctx context.Context, req *model.BuyProduct) (int, error)
	ListOrder(ctx context.Context, req int) ([]*model.OrderResp, error)
	ListProduct(ctx context.Context) ([]*model.ProductResp, error)
}

type manager struct {
	orderRepo   repository.OrderSDK
	productRepo repository.ProductSDK
}

func NewManager(orderRepo repository.OrderSDK, productRepo repository.ProductSDK) ManagerSDK {
	return &manager{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

func (m *manager) BuyProduct(ctx context.Context, req *model.BuyProduct) (int, error) {
	timeNow := time.Now()
	product, err := m.productRepo.GetProductByID(ctx, req.ProductID)
	if err != nil {
		return 0, err
	}

	id, err := m.orderRepo.BuyProduct(ctx, repository.Order{
		UserID:    int(req.UserID),
		ProductID: product.ID,
		CreatedAt: timeNow,
	})
	log.Println(err)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (m *manager) ListOrder(ctx context.Context, req int) ([]*model.OrderResp, error) {
	listOrders, err := m.orderRepo.ListOrder(ctx, req)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	if listOrders == nil {
		return nil, nil
	}

	listIDS := make([]int, len(listOrders))
	for i, v := range listOrders {
		listIDS[i] = v.ProductID
	}

	listProduct, err := m.productRepo.ListImgByIDS(ctx, listIDS)
	if err != nil {
		return nil, err
	}

	resp := make([]*model.OrderResp, len(listOrders))
	for index, order := range listOrders {
		img := ""
		description := ""
		for _, product := range listProduct {
			if product.ID == order.ProductID {
				img = product.Img
				description = product.Description
			}
		}

		resp[index] = &model.OrderResp{
			ID:          order.ID,
			ProductID:   order.ProductID,
			CreatedAt:   order.CreatedAt,
			Img:         img,
			Description: description,
		}
	}

	return resp, nil
}

func (m *manager) ListProduct(ctx context.Context) ([]*model.ProductResp, error) {
	listProduct, err := m.productRepo.ListProduct(ctx)
	if err != nil {
		return nil, err
	}

	if listProduct == nil {
		return nil, nil
	}

	resp := make([]*model.ProductResp, len(listProduct))
	for index, product := range listProduct {
		resp[index] = &model.ProductResp{
			ID:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			StockQuantity: product.StockQuantity,
			Price:         product.Price,
			CompanyID:     product.CompanyID,
			Status:        product.Status,
			CreatedAt:     product.CreatedAt,
			UpdatedAt:     product.UpdatedAt,
			DeletedAt:     product.DeletedAt,
			Img:           product.Img,
		}
	}

	return resp, nil
}
