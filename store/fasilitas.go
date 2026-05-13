package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

type Fasilitas interface {
	CreateFasilitas(ctx context.Context, arg database.CreateFasilitasParams) (database.Fasilita, error)
	DeleteFasilitas(ctx context.Context, id uuid.UUID) error
	GetFasilitas(ctx context.Context) (database.Fasilita, error)
	Updatefasilitas(ctx context.Context, arg database.UpdatefasilitasParams) (database.Fasilita, error)
}

type fasilitas struct {
	q *database.Queries
}

func (f fasilitas) CreateFasilitas(ctx context.Context, arg database.CreateFasilitasParams) (database.Fasilita, error) {
	return f.q.CreateFasilitas(ctx, arg)
}

func (f fasilitas) DeleteFasilitas(ctx context.Context, id uuid.UUID) error {
	return f.q.DeleteFasilitas(ctx, id)
}

func (f fasilitas) GetFasilitas(ctx context.Context) (database.Fasilita, error) {
	return f.q.GetFasilitas(ctx)
}

func (f fasilitas) Updatefasilitas(ctx context.Context, arg database.UpdatefasilitasParams) (database.Fasilita, error) {
	return f.q.Updatefasilitas(ctx, arg)
}

