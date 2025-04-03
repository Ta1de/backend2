package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"src/internal/repository/model"
)

type ImagePostgres struct {
	db *pgx.Conn
}

func NewImagePostgres(db *pgx.Conn) *ImagePostgres {
	return &ImagePostgres{db: db}
}

func (r *ImagePostgres) AddImage(ctx context.Context, image model.Image) (uuid.UUID, error) {
	query := `
	INSERT INTO images (image) 
	VALUES ($1) 
	RETURNING id;
	`

	var imageID uuid.UUID
	err := r.db.QueryRow(ctx, query, image.Image).Scan(&imageID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении изображения: %w", err)
	}

	return imageID, nil
}

func (r *ImagePostgres) AddImageToProduct(ctx context.Context, productID uuid.UUID, imageID uuid.UUID) error {
	query := `
		UPDATE product 
		SET image_id = $1 
		WHERE id = $2;
	`

	_, err := r.db.Exec(ctx, query, imageID, productID)
	if err != nil {
		return fmt.Errorf("ошибка при добавлении изображения к продукту: %w", err)
	}

	return nil
}

func (r *ImagePostgres) UploadImage(ctx context.Context, image model.Image, imageID uuid.UUID) error {
	query := `
		UPDATE images 
		SET image = $1 
		WHERE id = $2;
	`

	_, err := r.db.Exec(ctx, query, image.Image, imageID)
	if err != nil {
		return fmt.Errorf("ошибка при изменении изображения: %w", err)
	}

	return nil
}

func (r *ImagePostgres) DeleteImage(ctx context.Context, imageID uuid.UUID) error {
	query := `
		DELETE FROM images 
		WHERE id = $1;
	`
	_, err := r.db.Exec(ctx, query, imageID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении изображения: %w", err)
	}

	return nil
}

func (r *ImagePostgres) DeleteImageIdFromProduct(ctx context.Context, imageID uuid.UUID) error {
	query := `
		UPDATE product 
		SET image_id = NULL
		WHERE image_id = $1;
	`
	_, err := r.db.Exec(ctx, query, imageID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении image_id из product: %w", err)
	}

	return nil
}

func (r *ImagePostgres) GetImageIdByProductId(ctx context.Context, productId uuid.UUID) (uuid.UUID, error) {
	query := `
		SELECT image_id FROM product 
		WHERE id = $1;
	`
	var imageID uuid.UUID

	err := r.db.QueryRow(ctx, query, productId).Scan(&imageID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при получении image_id: %w", err)
	}

	return imageID, nil
}

func (r *ImagePostgres) GetImageByProductId(ctx context.Context, productID uuid.UUID) (model.Image, error) {
	query := `
		SELECT * FROM images 
		WHERE id = $1;
	`

	var image model.Image
	err := r.db.QueryRow(ctx, query, productID).Scan(&image.ID, &image.Image)
	if err != nil {
		return model.Image{}, fmt.Errorf("ошибка при получении изображения: %w", err)
	}

	return image, nil
}

func (r *ImagePostgres) GetImageById(ctx context.Context, imageID uuid.UUID) (model.Image, error) {
	query := `
		SELECT * FROM images 
		WHERE id = $1;
	`

	var image model.Image
	err := r.db.QueryRow(ctx, query, imageID).Scan(&image.ID, &image.Image)
	if err != nil {
		return model.Image{}, fmt.Errorf("ошибка при получении изображения: %w", err)
	}

	return image, nil
}
