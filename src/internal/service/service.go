package service

import (
	"context"
	"github.com/google/uuid"
	"src/internal/repository"
	"src/internal/repository/model"
)

type User interface {
	AddUser(ctx context.Context, user model.User, address model.Address) (uuid.UUID, error)
	RemoveUser(ctx context.Context, userID uuid.UUID) error
	GetUsers(ctx context.Context, name, surname string) ([]model.User, error)
	GetUsersList(ctx context.Context, limit, offset int) ([]model.User, error)
	UpdateUserAddress(ctx context.Context, userID uuid.UUID, address model.Address) error
}

type Supplier interface {
	AddSupplier(ctx context.Context, supplier model.Supplier, address model.Address) (uuid.UUID, error)
	UpdateSupplierAddress(ctx context.Context, SupplierID uuid.UUID, address model.Address) error
	RemoveSupplier(ctx context.Context, SupplierID uuid.UUID) error
	GetSuppliersList(ctx context.Context) ([]model.Supplier, error)
	GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (model.Supplier, error)
}

type Product interface {
	CreateProduct(ctx context.Context, product model.Product) (uuid.UUID, error)
	ReduceStock(ctx context.Context, productID uuid.UUID, quantity int) error
	GetProductById(ctx context.Context, productID uuid.UUID) (model.Product, error)
	GetProductList(ctx context.Context) ([]model.Product, error)
	RemoveProduct(ctx context.Context, productID uuid.UUID) error
}

type Image interface {
	CreateImage(ctx context.Context, image model.Image, productID uuid.UUID) (uuid.UUID, error)
	UpdateImage(ctx context.Context, image model.Image, imageID uuid.UUID) error
	DeleteImage(ctx context.Context, imageID uuid.UUID) error
	GetImageByProductId(ctx context.Context, productID uuid.UUID) (model.Image, error)
	GetImageById(ctx context.Context, imageID uuid.UUID) (model.Image, error)
}

type Service struct {
	User
	Supplier
	Product
	Image
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		User:     NewUserService(repos.User, repos.Address),
		Supplier: NewSupplierService(repos.Supplier, repos.Address),
		Product:  NewProductService(repos.Product),
		Image:    NewImageService(repos.Image),
	}
}
