package mapper

import (
	"src/internal/api/response"
	"src/internal/repository/model"
	"time"
)

func ToUserModel(req response.CreateUser) model.User {
	birthday, _ := time.Parse("2006-01-02", req.Birthday)

	return model.User{
		ClientName:    req.ClientName,
		ClientSurname: req.ClientSurname,
		Birthday:      birthday,
		Gender:        req.Gender,
	}
}

func ToUserResponse(user model.User) response.UserResponse {
	return response.UserResponse{
		ID:            user.ID.String(),
		ClientName:    user.ClientName,
		ClientSurname: user.ClientSurname,
		Birthday:      user.Birthday.Format("2006-01-02"), // Форматируем дату
		Gender:        user.Gender,
		Registration:  user.RegistrationDate.Format("2006-01-02T15:04:05Z"), // ISO 8601
		Address:       user.AddressID.String(),
	}
}
