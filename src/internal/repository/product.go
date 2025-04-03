package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"src/internal/repository/model"
)

type ProductPostgres struct {
	db *pgx.Conn
}

func NewProductPostgres(db *pgx.Conn) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) CreateProduct(ctx context.Context, product model.Product) (uuid.UUID, error) {
	query := `
	INSERT INTO product (name, category, price, available_stock, last_update_date, supplier_id, image_id) 
	VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5, $6) 
	RETURNING id;
	`

	var productID uuid.UUID
	err := r.db.QueryRow(ctx, query, product.Name, product.Category, product.Price,
		product.AvailableStock, product.SupplierID, product.ImageID).Scan(&productID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении товара: %w", err)
	}

	return productID, nil
}

func (r *ProductPostgres) ReduceStock(ctx context.Context, productID uuid.UUID, quantity int) error {
	query := `
		UPDATE product 
		SET available_stock = available_stock - $1,
		    last_update_date = CURRENT_TIMESTAMP
		WHERE id = $2 AND available_stock >= $1;
	`

	result, err := r.db.Exec(ctx, query, quantity, productID)
	if err != nil {
		return fmt.Errorf("ошибка при уменьшении количества товара: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("не удалось уменьшить количество товара: недостаточно на складе или неверный ID")
	}

	return nil
}

func (r *ProductPostgres) GetProductById(ctx context.Context, productID uuid.UUID) (model.Product, error) {
	query := `
		SELECT * FROM product 
		WHERE id = $1;
	`
	var product model.Product
	err := r.db.QueryRow(ctx, query, productID).Scan(&product.ID, &product.Name, &product.Category, &product.Price,
		&product.AvailableStock, &product.LastUpdateDate, &product.SupplierID, &product.ImageID)
	if err != nil {
		return model.Product{}, fmt.Errorf("ошибка при получении товара: %w", err)
	}

	return product, nil
}

func (r *ProductPostgres) GetProductList(ctx context.Context) ([]model.Product, error) {
	query := `SELECT * FROM product;`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении товаров: %w", err)
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Category, &product.Price,
			&product.AvailableStock, &product.LastUpdateDate, &product.SupplierID, &product.ImageID,
		); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return products, nil
}

func (r *ProductPostgres) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	query := `DELETE FROM product WHERE id = $1;`
	_, err := r.db.Exec(ctx, query, productID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении товара: %w", err)
	}
	return nil
}
