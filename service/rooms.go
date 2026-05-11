package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/JWT/middleware"
	"github.com/luckxx24/RoomBooking/database"
)

func HelperGetID(ctx context.Context) (uuid.UUID, error) {
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

func (S *Service) CreateRooms(ctx context.Context, nama, deskripsi, Role string, kapasitas int32) (database.Room, error) {

	ok := Helperrole(Role)

	if !ok {
		return database.Room{}, errors.New("role rusak/tidak ditemukan")
	}
	Rooms, err := S.Store.Room.Createroom(ctx, database.CreateroomParams{
		ID:           uuid.New(),
		Nama:         nama,
		Kapasitas:    kapasitas,
		PricePerHour: PricePerHOur,
		Description:  deskripsi,
		UpdatedAt:    time.Now(),
	})

	if err != nil {
		return database.Room{}, err
	}

	return Rooms, nil
}

func (S *Service) GetRooms(ctx context.Context, Page, Pagesize int) ([]database.GetRoomRow, error) {

	HelperPage(Page, Pagesize)

	Rooms, err := S.Store.Room.GetRoom(ctx, database.GetRoomParams{})

	if err != nil {
		return []database.GetRoomRow{}, err
	}

	return Rooms, nil
}

func (S *Service) UpdateRooms(ctx context.Context, nama, deskripsi string, kapasitas int32) (database.Room, error) {

	Rooms, err := S.Store.Room.UpdateRoom(ctx, database.UpdateRoomParams{
		ID:           uuid.New(),
		Nama:         nama,
		Kapasitas:    kapasitas,
		PricePerHour: PricePerHOur,
		Description:  deskripsi,
	})

	if err != nil {
		return database.Room{}, err
	}

	return Rooms, nil
}

func (S *Service) DeletRooms(ctx context.Context) error {

	if erro != nil {
		return erro
	}
	err := S.Store.Room.DeleteRoom(ctx)

	if err != nil {
		return err
	}
	return nil
}
