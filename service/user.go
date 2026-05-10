package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

func (S *Service) CreateUser(ctx context.Context, nama, email, HashPassword, Role string) (database.User, error) {

	ok := Helperrole(Role)

	if !ok {
		return database.User{}, errors.New("role rusak/tidak ditemukan")
	}
	User, err := S.Store.Users.CreateUsers(ctx, database.CreateUsersParams{
		ID:           uuid.New(),
		Email:        email,
		HashPassword: HashPassword,
		Role:         database.Roles(Role),
	})

	if err != nil {
		return database.User{}, err
	}

	return User, nil
}

func (S *Service) GetUser(ctx context.Context)
