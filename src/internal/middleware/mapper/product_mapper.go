package mapper

import (
	"fmt"
	"github.com/google/uuid"
	"src/internal/api/response"
	"src/internal/repository/model"
)

func ToProductModel(req response.CreateProduct) model.Product {
	supplierId, err := uuid.Parse(req.SupplierID)
	if err != nil {
		fmt.Printf("Ошибка парсинга supplierID: %s\n", err)
	}
	return model.Product{
		Name:           req.Name,
		Category:       req.Category,
		Price:          req.Price,
		AvailableStock: req.AvailableStock,
		SupplierID:     supplierId,
	}
}

func ToProductResponse(product model.Product) response.ProductResponse {
	imageId := ""

	if product.ImageID != nil {
		str := product.ImageID.String()
		imageId = str
	}
	return response.ProductResponse{
		ID:             product.ID.String(),
		Name:           product.Name,
		Category:       product.Category,
		Price:          product.Price,
		AvailableStock: product.AvailableStock,
		LastUpdateDate: product.LastUpdateDate.String(),
		SupplierID:     product.SupplierID.String(),
		ImageID:        imageId,
	}
}
