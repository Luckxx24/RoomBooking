package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

type Booking interface {
	CreateBooking(ctx context.Context, arg database.CreateBookingParams) (database.Booking, error)
	DeleteBooking(ctx context.Context, id uuid.UUID) error
	GetBooking(ctx context.Context, arg database.GetBookingParams) ([]database.GetBookingRow, error)
	UpdateBooking(ctx context.Context, arg database.UpdateBookingParams) (database.Booking, error)
	CheckBookingAvailability(ctx context.Context, arg database.CheckBookingAvailabilityParams) (int64, error)
}

type booking struct {
	q *database.Queries
}

func (b booking) CreateBooking(ctx context.Context, arg database.CreateBookingParams) (database.Booking, error) {
	return b.q.CreateBooking(ctx, arg)
}

func (b booking) DeleteBooking(ctx context.Context, id uuid.UUID) error {
	return b.q.DeleteBooking(ctx, id)
}

func (b booking) GetBooking(ctx context.Context, arg database.GetBookingParams) ([]database.GetBookingRow, error) {
	return b.q.GetBooking(ctx, arg)
}

func (b booking) UpdateBooking(ctx context.Context, arg database.UpdateBookingParams) (database.Booking, error) {
	return b.q.UpdateBooking(ctx, arg)
}

func (b booking) CheckBookingAvailability(ctx context.Context, arg database.CheckBookingAvailabilityParams) (int64, error) {
	return b.q.CheckBookingAvailability(ctx, database.CheckBookingAvailabilityParams{})
}
