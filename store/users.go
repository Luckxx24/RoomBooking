package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

type Users interface {
	CreateUsers(ctx context.Context, arg database.CreateUsersParams) (database.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
	GetUsers(ctx context.Context, id uuid.UUID) (database.GetUsersRow, error)
	GetUserslist(ctx context.Context, arg database.GetUserslistParams) ([]database.GetUserslistRow, error)
	UpdateUser(ctx context.Context, arg database.UpdateUserParams) (database.User, error)
}

type users struct {
	q *database.Queries
}

func (u *users) CreateUsers(ctx context.Context, arg database.CreateUsersParams) (database.User, error) {
	return u.q.CreateUsers(ctx, arg)
}

func (u *users) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return u.q.DeleteUser(ctx, id)
}

func (u *users) GetUsers(ctx context.Context, id uuid.UUID) (database.GetUsersRow, error) {
	return u.q.GetUsers(ctx, id)
}

func (u *users) GetUserslist(ctx context.Context, arg database.GetUserslistParams) ([]database.GetUserslistRow, error) {
	return u.q.GetUserslist(ctx, arg)
}

func (u *users) UpdateUser(ctx context.Context, arg database.UpdateUserParams) (database.User, error) {
	return u.q.UpdateUser(ctx, arg)
}
