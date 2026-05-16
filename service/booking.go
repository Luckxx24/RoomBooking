package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
	"github.com/shopspring/decimal"
)

func HelperGetIDBooking(r *http.Request) (uuid.UUID, error) {
	RoomIDstr := chi.URLParam(r, "id_room")

	if RoomIDstr == " " {
		return uuid.Nil, errors.New("gagal mendapatkan ID dari url param")
	}

	RoomsID, errs := uuid.Parse(RoomIDstr)

	if errs != nil {
		return uuid.Nil, errs
	}

	return RoomsID, nil
}

func HelperStatusBooking(s string) bool {
	if s == "pending" || s == "approved" || s == "done" {
		return true
	}
	return false
}

func (S *Service) CreateBooking(ctx context.Context, StartTime, EndTime time.Time, TotalPrice string, r *http.Request, status string) (database.Booking, error) {

	IDRoom, errs := HelperGetIDRooms(r)

	if errs != nil {
		return database.Booking{}, errs
	}

	IDUser, errr := HelperGetIDUser(ctx)

	if errr != nil {
		return database.Booking{}, errr
	}

	if StartTime.After(EndTime) {
		return database.Booking{}, errors.New("tidak boleh mengisi starttime setelah endtime")
	}

	if StartTime.Before(time.Now()) {
		return database.Booking{}, errors.New("tidak boleh mengisi StartTIme di masa lalu")
	}

	Room, erro := S.Store.Room.GetRoomBYID(ctx, IDRoom)

	if erro == sql.ErrNoRows {
		return database.Booking{}, errors.New("ID room tidak ditemukan")
	}
	if erro != nil {
		return database.Booking{}, erro
	}

	Availabler, erros := S.Store.Booking.CheckBookingAvailability(ctx, database.CheckBookingAvailabilityParams{
		IDRooms:   IDRoom,
		StartTime: StartTime,
		EndTime:   EndTime,
	})

	if erros != nil {
		return database.Booking{}, erros
	}

	if Availabler >= 1 {
		return database.Booking{}, errors.New("Jadwal bertabrakan dengan jadwal lainnya")
	}

	Selisih := EndTime.Sub(StartTime).Hours()
	priceperHour, errosr := strconv.ParseFloat(Room.PricePerHour, 32)

	if errosr != nil {
		return database.Booking{}, errosr
	}

	Totalprice := Selisih * priceperHour

	totalPrice := decimal.NewFromFloat(Totalprice)

	Booking, err := S.Store.Booking.CreateBooking(
		ctx, database.CreateBookingParams{
			ID:         uuid.New(),
			IDUser:     IDUser,
			IDRooms:    IDRoom,
			StartTime:  StartTime,
			EndTime:    EndTime,
			TotalPrice: totalPrice,
			Status:     status,
		})

	if err != nil {
		return database.Booking{}, err
	}

	return Booking, nil
}

func (S *Service) GetBooking(ctx context.Context, Page, PageSIze int) ([]database.GetBookingRow, error) {
	offset, limit := HelperPage(Page, PageSIze)
	Booking, err := S.Store.Booking.GetBooking(ctx, database.GetBookingParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})

	if err != nil {
		return []database.GetBookingRow{}, err
	}

	return Booking, nil
}

func ( S *Service) UpdateBooking(ctx context.Context)(){
	S.Store.Booking.UpdateBooking(ctx,database.UpdateBookingParams{
		  StartTime:
    EndTime   : 
    Status    : 
    TotalPrice :
    ID        : 
	})
}

func (S *Service) DeleteBooking(ctx context.Context, r *http.Request) error {

	UserID, errs := HelperGetIDBooking(r)

	if errs != nil {
		return errs
	}
	err := S.Store.Booking.DeleteBooking(ctx, UserID)

	if err != nil {
		return err
	}
	return nil
}
