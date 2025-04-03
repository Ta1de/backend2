package mapper

import (
	"src/internal/api/response"
	"src/internal/repository/model"
)

func ToSupplierModel(req response.CreateSupplier) model.Supplier {
	return model.Supplier{
		Name:        req.Name,
		PhoneNumber: req.Phone,
	}
}

func ToSupplierResponse(supplier model.Supplier) response.SupplierResponse {
	return response.SupplierResponse{
		ID:      supplier.ID.String(),
		Name:    supplier.Name,
		Address: supplier.AddressID.String(),
		Phone:   supplier.PhoneNumber,
	}
}
