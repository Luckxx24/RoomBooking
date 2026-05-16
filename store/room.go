package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

type Room interface {
	DeleteRoom(ctx context.Context, id uuid.UUID) error
	GetRoom(ctx context.Context, arg database.GetRoomParams) ([]database.GetRoomRow, error)
	GetRoomDetail(ctx context.Context, id uuid.UUID) (database.GetRoomDetailRow, error)
	UpdateRoom(ctx context.Context, arg database.UpdateRoomParams) (database.Room, error)
	Createroom(ctx context.Context, arg database.CreateroomParams) (database.Room, error)
	GetRoomBYID(ctx context.Context, id uuid.UUID) (database.GetRoomBYIDRow, error)
}

type room struct {
	q *database.Queries
}

func (r room) DeleteRoom(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteRoom(ctx, id)
}

func (r room) GetRoom(ctx context.Context, arg database.GetRoomParams) ([]database.GetRoomRow, error) {
	return r.q.GetRoom(ctx, arg)
}

func (r room) GetRoomDetail(ctx context.Context, id uuid.UUID) (database.GetRoomDetailRow, error) {
	return r.q.GetRoomDetail(ctx, id)
}

func (r room) UpdateRoom(ctx context.Context, arg database.UpdateRoomParams) (database.Room, error) {
	return r.q.UpdateRoom(ctx, arg)
}

func (r room) Createroom(ctx context.Context, arg database.CreateroomParams) (database.Room, error) {
	return r.q.Createroom(ctx, arg)
}

func (r room) GetRoomBYID(ctx context.Context, id uuid.UUID) (database.GetRoomBYIDRow, error) {
	return r.q.GetRoomBYID(ctx, id)
}
