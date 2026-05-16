package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

type RoomFasilitas interface {
	CreateFasilitas_Ruangan(ctx context.Context, arg database.CreateFasilitas_RuanganParams) (database.FasilitasRuangan, error)
	DeleteFasilitas_Ruangan(ctx context.Context, id uuid.UUID) error
	GetFasilitas_Ruangan(ctx context.Context) (database.FasilitasRuangan, error)
}

type roomfasilitas struct {
	q *database.Queries
}

func (r roomfasilitas) CreateFasilitas_Ruangan(ctx context.Context, arg database.CreateFasilitas_RuanganParams) (database.FasilitasRuangan, error) {
	return r.q.CreateFasilitas_Ruangan(ctx, arg)
}

func (r roomfasilitas) GetFasilitas_Ruangan(ctx context.Context) (database.FasilitasRuangan, error) {
	return r.q.GetFasilitas_Ruangan(ctx)
}

func (r roomfasilitas) DeleteFasilitas_Ruangan(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteFasilitas_Ruangan(ctx, id)
}
