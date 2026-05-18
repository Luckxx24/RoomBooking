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
	BookingIDstr := chi.URLParam(r, "id_booking")

	if BookingIDstr == " " {
		return uuid.Nil, errors.New("gagal mendapatkan ID dari url param")
	}

	BookingID, errs := uuid.Parse(BookingIDstr)

	if errs != nil {
		return uuid.Nil, errs
	}

	return BookingID, nil
}

func HelperStatusBooking(s string) bool {
	if s == "pending" || s == "approved" || s == "done" {
		return true
	}
	return false
}

func HelperPrice(StartTime, EndTime time.Time, PricePerHour string) (string, error) {
	Selisih := EndTime.Sub(StartTime).Hours()
	priceperHour, errosr := strconv.ParseFloat(PricePerHour, 32)

	if errosr != nil {
		return "", errosr
	}

	Totalprice := Selisih * priceperHour

	totalPrice := decimal.NewFromFloat(Totalprice)

	totalPricestr := totalPrice.StringFixed(2)

	return totalPricestr, nil
}

func (S *Service) CreateBooking(ctx context.Context, StartTime, EndTime time.Time, r *http.Request) (database.Booking, error) {

	if StartTime.IsZero() {
		return database.Booking{}, errors.New("waktu mulai tidak boleh kosong")
	}

	if EndTime.IsZero() {
		return database.Booking{}, errors.New("waktu selesai tidak boleh kosong")
	}
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
	totalPricestr, errro := HelperPrice(StartTime, EndTime, Room.PricePerHour)

	if errro != nil {
		return database.Booking{}, errro
	}

	Booking, err := S.Store.Booking.CreateBooking(
		ctx, database.CreateBookingParams{
			ID:         uuid.New(),
			IDUser:     IDUser,
			IDRooms:    IDRoom,
			StartTime:  StartTime,
			EndTime:    EndTime,
			TotalPrice: totalPricestr,
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

func (S *Service) UpdateBooking(ctx context.Context, StartTime, EndTime time.Time, status string, r *http.Request) (database.Booking, error) {

	ID, errs := HelperGetIDBooking(r)

	if errs != nil {
		return database.Booking{}, errs
	}
	if StartTime.IsZero() {
		return database.Booking{}, errors.New("waktu mulai tidak boleh kosong")
	}

	if EndTime.IsZero() {
		return database.Booking{}, errors.New("waktu selesai tidak boleh kosong")
	}

	if StartTime.After(EndTime) {
		return database.Booking{}, errors.New("tidak boleh mengisi starttime setelah endtime")
	}

	if StartTime.Before(time.Now()) {
		return database.Booking{}, errors.New("tidak boleh mengisi StartTIme di masa lalu")
	}

	IDRoom, errs := HelperGetIDRooms(r)

	if errs != nil {
		return database.Booking{}, errs
	}

	Room, erro := S.Store.Room.GetRoomBYID(ctx, IDRoom)

	if erro == sql.ErrNoRows {
		return database.Booking{}, errors.New("ID room tidak ditemukan")
	}
	if erro != nil {
		return database.Booking{}, erro
	}

	totalpricestr, errrs := HelperPrice(StartTime, EndTime, Room.PricePerHour)

	if errrs != nil {
		return database.Booking{}, errrs
	}

	oke := HelperStatusBooking(status)

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

	if !oke {
		return database.Booking{}, errors.New("format status salah")
	}

	Updated, err := S.Store.Booking.UpdateBooking(ctx, database.UpdateBookingParams{
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     database.Stats(status),
		TotalPrice: totalpricestr,
		ID:         ID,
	})

	if err != nil {
		return database.Booking{}, err
	}
	return Updated, nil
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
