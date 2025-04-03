package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"src/internal/repository/model"

	"github.com/jackc/pgx/v5"
)

type UserPostgres struct {
	db *pgx.Conn
}

func NewUserPostgres(db *pgx.Conn) *UserPostgres {
	return &UserPostgres{db: db}
}

func (r *UserPostgres) AddUser(ctx context.Context, user model.User) (uuid.UUID, error) {
	var id uuid.UUID

	query := `
		INSERT INTO client (client_name, client_surname, birthday, gender, registration_date, address_id) 
		VALUES ($1, $2, $3, $4, CURRENT_TIMESTAMP, $5)
		RETURNING id;
	`

	err := r.db.QueryRow(ctx, query, user.ClientName, user.ClientSurname, user.Birthday, user.Gender, user.AddressID).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при добавлении пользователя: %w", err)
	}

	return id, nil
}

func (r *UserPostgres) GetAddressIDByUserID(ctx context.Context, userID uuid.UUID) (uuid.UUID, error) {
	var addressID uuid.UUID
	query := `SELECT address_id FROM client WHERE id = $1;`
	err := r.db.QueryRow(ctx, query, userID).Scan(&addressID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("ошибка при получении address_id пользователя: %w", err)
	}
	return addressID, nil
}

func (r *UserPostgres) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	query := `DELETE FROM client WHERE id = $1;`
	_, err := r.db.Exec(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("ошибка при удалении пользователя: %w", err)
	}
	return nil
}

func (r *UserPostgres) GetUserNameSurname(ctx context.Context, name, surname string) ([]model.User, error) {
	query := `SELECT * FROM client WHERE client_name = $1 AND client_surname = $2;`

	rows, err := r.db.Query(ctx, query, name, surname)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении пользователей: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID, &user.ClientName, &user.ClientSurname,
			&user.Birthday, &user.Gender, &user.RegistrationDate, &user.AddressID,
		); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return users, nil
}

func (r *UserPostgres) GetUserList(ctx context.Context, limit, offset int) ([]model.User, error) {
	query := `SELECT * FROM client LIMIT $1 OFFSET $2;`

	rows, err := r.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении пользователей: %w", err)
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID, &user.ClientName, &user.ClientSurname,
			&user.Birthday, &user.Gender, &user.RegistrationDate, &user.AddressID,
		); err != nil {
			return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
	}

	return users, nil
}
