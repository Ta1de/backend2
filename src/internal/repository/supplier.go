package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"src/internal/repository/model"
)

type SupplierPostgres struct {
	db *pgx.Conn
}

func NewSupplierPostgres(db *pgx.Conn) *SupplierPostgres {
	return &SupplierPostgres{db: db}
}

func (r *SupplierPostgres) AddSupplier(ctx context.Context, supplier model.Supplier) (uuid.UUID, error) {
	var id uuid.UUID

	query := `
		INSERT INTO supplier (name, address_id, phone_number) 
		VALUES ($1, $2, $3)
		RETURNING id;
	`

	err := r.db.QueryRow(ctx, query, supplier.Name, supplier.AddressID, supplier.PhoneNumber).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении поставщика: %w", err)
	}

	return id, nil
}

func (r *SupplierPostgres) GetAddressIDBySupplierID(ctx context.Context, supplierID uuid.UUID) (uuid.UUID, error) {
	var addressID uuid.UUID
	query := `SELECT address_id FROM supplier WHERE id = $1;`
	err := r.db.QueryRow(ctx, query, supplierID).Scan(&addressID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при получении address_id поставщика: %w", err)
	}
	return addressID, nil
}

func (r *SupplierPostgres) DeleteSupplier(ctx context.Context, supplierID uuid.UUID) error {
	query := `DELETE FROM supplier WHERE id = $1;`
	_, err := r.db.Exec(ctx, query, supplierID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении поставщика: %w", err)
	}
	return nil
}

func (r *SupplierPostgres) GetSupplierList(ctx context.Context) ([]model.Supplier, error) {
	query := `SELECT * FROM supplier;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении поставщиков: %w", err)
	}
	defer rows.Close()

	var suppliers []model.Supplier
	for rows.Next() {
		var supplier model.Supplier
		if err := rows.Scan(
			&supplier.ID, &supplier.Name, &supplier.AddressID, &supplier.PhoneNumber,
		); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		suppliers = append(suppliers, supplier)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return suppliers, nil
}

func (r *SupplierPostgres) GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (model.Supplier, error) {
	query := `SELECT * FROM supplier WHERE id = $1;`

	var supplier model.Supplier
	err := r.db.QueryRow(ctx, query, supplierID).Scan(&supplier.ID, &supplier.Name, &supplier.AddressID, &supplier.PhoneNumber)
	if err != nil {
		return model.Supplier{}, fmt.Errorf("ошибка при получении поставщика: %w", err)
	}

	return supplier, nil
}
