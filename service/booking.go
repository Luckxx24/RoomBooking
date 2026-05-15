package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
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

func(S *Service) CreateBooking(ctx context.Context,StartTime,EndTime time.Time,TotalPrice string) (database.Booking,error){

	IDRoom,errs := HelperGetIDRooms(r)

	if errs != inl{
		return database.Booking{},errs
	}

	IDUser,errr :=  HelperGetIDUser(r)

	if errr != inl{
		return database.Booking{},errr
	}

	if StartTime.After(EndTime){
		return database.Booking{},errors.New("tidak boleh mengisi starttime setelah endtime")
	}

	if StartTime.Before(time.Now()){
		return database.Booking{},errors.New("tidak boleh mengisi StartTIme di masa lalu")
	}

	Room,erro := S.Store.Room.GetRoomBYID(ctx,IDRoom)

	if erro != nil{

		return database.Booking{},erro
	}

	if Room == nil{
		return database.Booking{},errors.New("ID Room Tidak Ditemukan")
	}
	S.Store.Booking.CreateBooking(
		ctx,database.CreateBookingParams{
 ID  : uuid.New(),
    IDUser :    
    IDRooms :   
    StartTime : 
    EndTime   : 
    TotalPrice: 
    Status    : 
		}
	)
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
