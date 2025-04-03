package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"src/internal/repository"
	"src/internal/repository/model"
)

type UserService struct {
	repoUser    repository.User
	repoAddress repository.Address
}

func NewUserService(repoUser repository.User, repoAddress repository.Address) *UserService {
	return &UserService{
		repoUser:    repoUser,
		repoAddress: repoAddress,
	}
}

func (s *UserService) AddUser(ctx context.Context, user model.User, address model.Address) (uuid.UUID, error) {
	addressID, err := s.repoAddress.CreateAddress(ctx, address)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении адреса: %w", err)
	}

	user.AddressID = addressID

	id, err := s.repoUser.AddUser(ctx, user)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении пользователя: %w", err)
	}

	return id, nil
}

func (s *UserService) RemoveUser(ctx context.Context, userID uuid.UUID) error {
	addressID, err := s.repoUser.GetAddressIDByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("ошибка при получении адреса пользователя: %w", err)
	}

	err = s.repoUser.DeleteUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении пользователя: %w", err)
	}

	err = s.repoAddress.DeleteAddress(ctx, addressID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении адреса: %w", err)
	}

	return nil
}

func (s *UserService) GetUsers(ctx context.Context, name, surname string) ([]model.User, error) {
	users, err := s.repoUser.GetUserNameSurname(ctx, name, surname)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении пользователя: %w", err)
	}

	return users, nil
}

func (s *UserService) GetUsersList(ctx context.Context, limit, offset int) ([]model.User, error) {
	users, err := s.repoUser.GetUserList(ctx, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка пользователей: %w", err)
	}

	return users, nil
}

func (s *UserService) UpdateUserAddress(ctx context.Context, userID uuid.UUID, address model.Address) error {
	addressID, err := s.repoUser.GetAddressIDByUserID(ctx, userID)
	if err != nil {
		return fmt.Errorf("ошибка при получении адреса пользователя: %w", err)
	}

	address.ID = addressID

	err = s.repoAddress.UpdateAddress(ctx, address)
	if err != nil {
		return fmt.Errorf("ошибка при изменении адреса пользователя: %w", err)
	}

	return nil
}
