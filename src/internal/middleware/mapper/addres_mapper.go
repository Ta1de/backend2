package mapper

import (
	"src/internal/api/response"
	"src/internal/repository/model"
)

func ToAddressModel(dto response.CreateUpdateAddress) model.Address {
	return model.Address{
		Country: dto.Country,
		City:    dto.City,
		Street:  dto.Street,
	}
}
