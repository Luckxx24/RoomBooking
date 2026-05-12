package service

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/JWT/middleware"
	"github.com/luckxx24/RoomBooking/database"
	"github.com/luckxx24/RoomBooking/store"
)

func HelperGetID(r *http.Request) (uuid.UUID, error) {
	RoomIDstr := chi.URLParam(r, "id_room")

	if RoomIDstr == " " {
		return uuid.Nil, errors.New("gagal mendapatkan ID dari middleware")
	}

	RoomsID, errs := uuid.Parse(RoomIDstr)

	if errs != nil {
		return uuid.Nil, errs
	}

	return RoomsID, nil
}

type FasilitasCompleted struct {
	nama         string
	deskripsi    string
	kapasitas    int32
	PricePerHOur string
	id_fasilitas []uuid.UUID
}

func (S *Service) CreateRooms(ctx context.Context) (database.Room, error) {
	tx, errs := S.Store.DB.BeginTx(ctx, nil)

	if errs != nil {
		defer tx.Rollback()
		return database.Room{}, errs

	}

	var Req FasilitasCompleted
	Txstore := store.NewTX(tx)

	if Req.nama == " " || Req.deskripsi == " " || Req.kapasitas <= 0 {
		return database.Room{}, errors.New("kolom tidak bileh kosong")
	}
	Role, oke := middleware.GetRoleFromContext(ctx)

	if !oke {
		return database.Room{}, errors.New("format role rusak")
	}
	ok := Helperrole(Role)

	if !ok {
		return database.Room{}, errors.New("Unathorized")
	}
	Rooms, err := Txstore.Room.Createroom(ctx, database.CreateroomParams{
		ID:           uuid.New(),
		Nama:         Req.nama,
		Kapasitas:    Req.kapasitas,
		PricePerHour: Req.PricePerHOur,
		Description:  Req.deskripsi,
		UpdatedAt:    time.Now(),
	})

	if err != nil {
		defer tx.Rollback()
		return database.Room{}, err
	}

	if len(Req.id_fasilitas) < 0 {
		return database.Room{}, errors.New("mohon masukan fasilitas")
	}
	range
	tx.Commit()
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

func (S *Service) UpdateRooms(ctx context.Context, nama, deskripsi string, kapasitas int32, r *http.Request) (database.Room, error) {

	RoomID, errs := HelperGetID(r)

	if errs != nil {
		return database.Room{}, errs
	}

	Rooms, err := S.Store.Room.UpdateRoom(ctx, database.UpdateRoomParams{
		ID:           RoomID,
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

func (S *Service) DeletRooms(ctx context.Context, r *http.Request) error {

	RoomsID, errs := HelperGetID(r)

	if errs != nil {
		return errs
	}
	err := S.Store.Room.DeleteRoom(ctx, RoomsID)

	if err != nil {
		return err
	}
	return nil
}
