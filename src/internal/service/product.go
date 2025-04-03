package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"src/internal/repository"
	"src/internal/repository/model"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) CreateProduct(ctx context.Context, product model.Product) (uuid.UUID, error) {
	id, err := s.repo.CreateProduct(ctx, product)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении товара: %w", err)
	}

	return id, nil
}

func (s *ProductService) ReduceStock(ctx context.Context, productID uuid.UUID, quantity int) error {
	return s.repo.ReduceStock(ctx, productID, quantity)
}

func (s *ProductService) GetProductById(ctx context.Context, productID uuid.UUID) (model.Product, error) {
	product, err := s.repo.GetProductById(ctx, productID)
	if err != nil {
		return model.Product{}, fmt.Errorf("ошибка при получении товара: %w", err)
	}

	return product, nil
}

func (s *ProductService) GetProductList(ctx context.Context) ([]model.Product, error) {
	users, err := s.repo.GetProductList(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка товаров: %w", err)
	}

	return users, nil
}

func (s *ProductService) RemoveProduct(ctx context.Context, productID uuid.UUID) error {
	err := s.repo.DeleteProduct(ctx, productID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении товара: %w", err)
	}

	return nil
}
