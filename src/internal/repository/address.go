package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"src/internal/repository/model"
)

type AddressPostgres struct {
	db *pgx.Conn
}

func NewAddressPostgres(db *pgx.Conn) *AddressPostgres {
	return &AddressPostgres{db: db}
}

func (r *AddressPostgres) CreateAddress(ctx context.Context, address model.Address) (uuid.UUID, error) {
	query := `
		INSERT INTO address (country, city, street) 
		VALUES ($1, $2, $3) 
		RETURNING id;
	`

	var addressID uuid.UUID
	err := r.db.QueryRow(ctx, query, address.Country, address.City, address.Street).Scan(&addressID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении адреса: %w", err)
	}

	return addressID, nil
}

func (r *AddressPostgres) DeleteAddress(ctx context.Context, addressID uuid.UUID) error {
	query := `DELETE FROM address WHERE id = $1;`
	_, err := r.db.Exec(ctx, query, addressID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении адреса: %w", err)
	}
	return nil
}

func (r *AddressPostgres) UpdateAddress(ctx context.Context, address model.Address) error {
	query := `
	UPDATE address
	SET country = $1, city = $2, street = $3
	WHERE id = $4;
	`
	_, err := r.db.Exec(ctx, query, address.Country, address.City, address.Street, address.ID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении адреса: %w", err)
	}

	return nil
}
