package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"src/internal/repository"
	"src/internal/repository/model"
)

type SupplierService struct {
	repoSupplier repository.Supplier
	repoAddress  repository.Address
}

func NewSupplierService(repoSupplier repository.Supplier, repoAddress repository.Address) *SupplierService {
	return &SupplierService{
		repoSupplier: repoSupplier,
		repoAddress:  repoAddress,
	}
}

func (s *SupplierService) AddSupplier(ctx context.Context, supplier model.Supplier, address model.Address) (uuid.UUID, error) {
	addressID, err := s.repoAddress.CreateAddress(ctx, address)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении адреса: %w", err)
	}

	supplier.AddressID = addressID

	id, err := s.repoSupplier.AddSupplier(ctx, supplier)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении поставщика: %w", err)
	}

	return id, nil
}

func (s *SupplierService) UpdateSupplierAddress(ctx context.Context, SupplierID uuid.UUID, address model.Address) error {
	addressID, err := s.repoSupplier.GetAddressIDBySupplierID(ctx, SupplierID)
	if err != nil {
		return fmt.Errorf("ошибка при получении адреса поставщика: %w", err)
	}

	address.ID = addressID

	err = s.repoAddress.UpdateAddress(ctx, address)
	if err != nil {
		return fmt.Errorf("ошибка при изменении адреса поставщика: %w", err)
	}

	return nil
}

func (s *SupplierService) RemoveSupplier(ctx context.Context, SupplierID uuid.UUID) error {
	addressID, err := s.repoSupplier.GetAddressIDBySupplierID(ctx, SupplierID)
	if err != nil {
		return fmt.Errorf("ошибка при получении адреса поставщика: %w", err)
	}

	err = s.repoSupplier.DeleteSupplier(ctx, SupplierID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении поставщика: %w", err)
	}

	err = s.repoAddress.DeleteAddress(ctx, addressID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении адреса: %w", err)
	}

	return nil
}

func (s *SupplierService) GetSuppliersList(ctx context.Context) ([]model.Supplier, error) {
	users, err := s.repoSupplier.GetSupplierList(ctx)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка поставщиков: %w", err)
	}

	return users, nil
}

func (s *SupplierService) GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (model.Supplier, error) {
	supplier, err := s.repoSupplier.GetSupplierByID(ctx, supplierID)
	if err != nil {
		return model.Supplier{}, fmt.Errorf("ошибка при получении поставщика: %w", err)
	}

	return supplier, nil
}
