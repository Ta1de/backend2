package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"src/internal/repository"
	"src/internal/repository/model"
)

type ImageService struct {
	repo repository.Image
}

func NewImageService(repo repository.Image) *ImageService {
	return &ImageService{repo: repo}
}

func (s *ImageService) CreateImage(ctx context.Context, image model.Image, productID uuid.UUID) (uuid.UUID, error) {
	id, err := s.repo.AddImage(ctx, image)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении изображения: %w", err)
	}

	err = s.repo.AddImageToProduct(ctx, productID, id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении изображения в продукт: %w", err)
	}

	return id, nil
}

func (s *ImageService) UpdateImage(ctx context.Context, image model.Image, imageID uuid.UUID) error {
	err := s.repo.UploadImage(ctx, image, imageID)
	if err != nil {
		return fmt.Errorf("ошибка при изменение изображения: %w", err)
	}

	return nil
}

func (s *ImageService) DeleteImage(ctx context.Context, imageID uuid.UUID) error {
	err := s.repo.DeleteImage(ctx, imageID)
	if err != nil {
		return fmt.Errorf("ошибка при изменение изображения: %w", err)
	}

	err = s.repo.DeleteImageIdFromProduct(ctx, imageID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении image_id из product: %w", err)
	}

	return nil
}

func (s *ImageService) GetImageByProductId(ctx context.Context, productID uuid.UUID) (model.Image, error) {
	imageId, err := s.repo.GetImageIdByProductId(ctx, productID)
	if err != nil {
		return model.Image{}, fmt.Errorf("ошибка при запросе id из продукта: %w", err)
	}

	image, err := s.repo.GetImageByProductId(ctx, imageId)
	if err != nil {
		return model.Image{}, fmt.Errorf("ошибка при получении изображения: %w", err)
	}

	return image, nil
}

func (s *ImageService) GetImageById(ctx context.Context, imageID uuid.UUID) (model.Image, error) {
	image, err := s.repo.GetImageById(ctx, imageID)
	if err != nil {
		return model.Image{}, fmt.Errorf("ошибка при получении изображения: %w", err)
	}

	return image, nil
}
