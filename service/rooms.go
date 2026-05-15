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

func HelperGetIDRooms(r *http.Request) (uuid.UUID, error) {
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

func HelperGetIDfasilitas(r *http.Request) (uuid.UUID, error) {
	FasilitasIDstr := chi.URLParam(r, "id_fasilitas")

	if FasilitasIDstr == " " {
		return uuid.Nil, errors.New("gagal mendapatkan ID dari url param")
	}

	FasilitasID, errs := uuid.Parse(FasilitasIDstr)

	if errs != nil {
		return uuid.Nil, errs
	}

	return FasilitasID, nil
}

func (S *Service) CreateRooms(ctx context.Context) (database.Room, error) {

	type CRoomsFasility struct {
		nama         string
		deskripsi    string
		kapasitas    int32
		PricePerHOur string
		id_fasilitas []uuid.UUID
	}
	var Req CRoomsFasility

	tx, errs := S.Store.DB.BeginTx(ctx, nil)

	if errs != nil {
		return database.Room{}, errs

	}
	defer tx.Rollback()
	Txstore := store.NewTX(tx)

	if Req.nama == " " || Req.deskripsi == " " || Req.kapasitas <= 0 || Req.PricePerHOur == " " {
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

		return database.Room{}, err
	}

	if len(Req.id_fasilitas) <= 0 {
		return database.Room{}, errors.New("mohon masukan fasilitas")
	}

	for _, id_fasilitas := range Req.id_fasilitas {
		_, erros := Txstore.RoomFasilitas.CreateFasilitas_Ruangan(ctx, database.CreateFasilitas_RuanganParams{
			ID:          uuid.New(),
			IDRoom:      Rooms.ID,
			IDFasilitas: id_fasilitas,
		})
		if erros != nil {

			return database.Room{}, erros
		}
	}

	errr := tx.Commit()

	if errr != nil {
		return database.Room{}, errr
	}
	return Rooms, nil

}

func (S *Service) GetRooms(ctx context.Context, Page, Pagesize int) ([]database.GetRoomRow, error) {

	Offset, Limit := HelperPage(Page, Pagesize)

	Rooms, err := S.Store.Room.GetRoom(ctx, database.GetRoomParams{
		Offset: int32(Offset),
		Limit:  int32(Limit),
	})

	if err != nil {
		return []database.GetRoomRow{}, err
	}

	return Rooms, nil
}

func (S *Service) GetRoomsDetail(ctx context.Context, r *http.Request) (database.GetRoomDetailRow, error) {
	id_room, errs := HelperGetIDRooms(r)

	if errs != nil {
		return database.GetRoomDetailRow{}, errs
	}
	room, err := S.Store.Room.GetRoomDetail(ctx, id_room)

	if err != nil {
		return database.GetRoomDetailRow{}, err
	}

	return room, nil
}

func (S *Service) UpdateRooms(ctx context.Context, nama, deskripsi string, kapasitas int32, r *http.Request) (database.Room, error) {

	type URoomsFasility struct {
		nama         string
		deskripsi    string
		kapasitas    int32
		PricePerHOur string
		id_fasilitas []uuid.UUID
	}
	var UReq URoomsFasility

	RoomID, errs := HelperGetIDRooms(r)

	if errs != nil {
		return database.Room{}, errs
	}

	tx, erro := S.Store.DB.BeginTx(ctx, nil)
	defer tx.Rollback()
	if erro != nil {

		return database.Room{}, erro
	}

	TXdb := store.NewTX(tx)

	Errs := TXdb.RoomFasilitas.DeleteFasilitas_Ruangan(ctx, RoomID)

	if Errs != nil {

		return database.Room{}, Errs
	}

	if UReq.nama == " " || UReq.deskripsi == " " || UReq.kapasitas <= 0 || UReq.PricePerHOur == " " {
		return database.Room{}, errors.New("kolom tidak bileh kosong")
	}

	Room, errr := TXdb.Room.UpdateRoom(ctx, database.UpdateRoomParams{
		Nama:         UReq.nama,
		PricePerHour: UReq.PricePerHOur,
		Description:  UReq.deskripsi,
		Kapasitas:    UReq.kapasitas,
		UpdatedAt:    time.Now(),
		ID:           RoomID,
	})

	if errr != nil {

		return database.Room{}, errr
	}

	for _, id_fasilitas := range UReq.id_fasilitas {
		_, erros := TXdb.RoomFasilitas.CreateFasilitas_Ruangan(ctx, database.CreateFasilitas_RuanganParams{
			ID:          uuid.New(),
			IDRoom:      Room.ID,
			IDFasilitas: id_fasilitas,
		})
		if erros != nil {

			return database.Room{}, erros
		}
	}

	errro := tx.Commit()

	if errro != nil {
		return database.Room{}, errro
	}

	return Room, nil
}

func (S *Service) DeletRooms(ctx context.Context, r *http.Request) error {

	RoomsID, errs := HelperGetIDRooms(r)

	if errs != nil {
		return errs
	}
	tx, erro := S.Store.DB.BeginTx(ctx, nil)

	if erro != nil {

		return erro
	}

	TXdb := store.NewTX(tx)

	defer tx.Rollback()

	Errs := TXdb.RoomFasilitas.DeleteFasilitas_Ruangan(ctx, RoomsID)

	if Errs != nil {

		return Errs
	}

	err := TXdb.Room.DeleteRoom(ctx, RoomsID)

	if err != nil {

		return err
	}

	errr := tx.Commit()

	if errr != nil {
		return errr
	}
	return nil
}
