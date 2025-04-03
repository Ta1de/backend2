package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"src/internal/repository/model"
)

type User interface {
	AddUser(ctx context.Context, user model.User) (uuid.UUID, error)
	GetAddressIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error)
	DeleteUser(ctx context.Context, userID uuid.UUID) error
	GetUserNameSurname(ctx context.Context, name, surname string) ([]model.User, error)
	GetUserList(ctx context.Context, limit, offset int) ([]model.User, error)
}

type Address interface {
	CreateAddress(ctx context.Context, address model.Address) (uuid.UUID, error)
	DeleteAddress(ctx context.Context, addressID uuid.UUID) error
	UpdateAddress(ctx context.Context, address model.Address) error
}

type Supplier interface {
	AddSupplier(ctx context.Context, supplier model.Supplier) (uuid.UUID, error)
	GetAddressIDBySupplierID(ctx context.Context, supplierID uuid.UUID) (uuid.UUID, error)
	DeleteSupplier(ctx context.Context, userID uuid.UUID) error
	GetSupplierList(ctx context.Context) ([]model.Supplier, error)
	GetSupplierByID(ctx context.Context, supplierID uuid.UUID) (model.Supplier, error)
}

type Product interface {
	CreateProduct(ctx context.Context, product model.Product) (uuid.UUID, error)
	ReduceStock(ctx context.Context, productID uuid.UUID, quantity int) error
	GetProductById(ctx context.Context, productID uuid.UUID) (model.Product, error)
	GetProductList(ctx context.Context) ([]model.Product, error)
	DeleteProduct(ctx context.Context, productID uuid.UUID) error
}

type Image interface {
	AddImage(ctx context.Context, image model.Image) (uuid.UUID, error)
	AddImageToProduct(ctx context.Context, productID uuid.UUID, imageID uuid.UUID) error
	UploadImage(ctx context.Context, image model.Image, imageID uuid.UUID) error
	DeleteImage(ctx context.Context, imageID uuid.UUID) error
	DeleteImageIdFromProduct(ctx context.Context, imageID uuid.UUID) error
	GetImageIdByProductId(ctx context.Context, productId uuid.UUID) (uuid.UUID, error)
	GetImageByProductId(ctx context.Context, productID uuid.UUID) (model.Image, error)
	GetImageById(ctx context.Context, imageID uuid.UUID) (model.Image, error)
}

type Repository struct {
	User
	Address
	Supplier
	Product
	Image
}

func NewRepositore(db *pgx.Conn) *Repository {
	return &Repository{
		User:     NewUserPostgres(db),
		Address:  NewAddressPostgres(db),
		Supplier: NewSupplierPostgres(db),
		Product:  NewProductPostgres(db),
		Image:    NewImagePostgres(db),
	}
}
