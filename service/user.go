package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/JWT/middleware"
	"github.com/luckxx24/RoomBooking/database"
)

func HelperGetIDUser(ctx context.Context) (uuid.UUID, error) {
	UserIDstr, ok := middleware.GetIDFromContext(ctx)

	if !ok {
		return uuid.Nil, errors.New("gagal mendapatkan ID dari middleware")
	}

	UserID, errs := uuid.Parse(UserIDstr)

	if errs != nil {
		return uuid.Nil, errs
	}

	return UserID, nil
}

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

func (S *Service) GetUser(ctx context.Context) (database.GetUsersRow, error) {

	UserID, errs := HelperGetIDUser(ctx)

	if errs != nil {
		return database.GetUsersRow{}, errs
	}

	User, err := S.Store.Users.GetUsers(ctx, UserID)

	if err != nil {
		return database.GetUsersRow{}, err
	}

	return User, nil
}

func (S *Service) GetUserlist(ctx context.Context, Page, PageSize int) ([]database.GetUserslistRow, error) {
	offset, limit := HelperPage(Page, PageSize)
	User, errs := S.Store.Users.GetUserslist(ctx, database.GetUserslistParams{Offset: int32(offset), Limit: int32(limit)})

	if errs != nil {
		return []database.GetUserslistRow{}, errs
	}

	return User, nil
}

func (S *Service) UpdateUser(ctx context.Context, nama, email, haspassword string) (database.User, error) {

	UserID, errs := HelperGetIDUser(ctx)

	if errs != nil {
		return database.User{}, errs
	}
	User, err := S.Store.Users.UpdateUser(ctx, database.UpdateUserParams{
		Nama:         nama,
		Email:        email,
		HashPassword: haspassword,
		ID:           UserID,
	})

	if err != nil {
		return database.User{}, err
	}

	return User, nil
}

func (S *Service) DeletUser(ctx context.Context) error {
	UserID, erro := HelperGetIDUser(ctx)

	if erro != nil {
		return erro
	}
	err := S.Store.Users.DeleteUser(ctx, UserID)

	if err != nil {
		return err
	}
	return nil
}
