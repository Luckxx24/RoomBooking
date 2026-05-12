package service

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/database"
)

func (S *Service) CreateFasilitas(ctx context.Context, nama string) (database.Fasilita, error) {

	if nama == " " {
		return database.Fasilita{}, errors.New("kolom nama tidak boleh kosong")
	}
	fasilitas, err := S.Store.Fasilitas.CreateFasilitas(ctx, database.CreateFasilitasParams{
		ID:   uuid.New(),
		Nama: nama,
	})

	if err != nil {
		return database.Fasilita{}, err
	}

	return fasilitas, nil
}

func (S *Service) GetAllFasilitas(ctx context.Context) (database.Fasilita, error) {

	Fasilitas, err := S.Store.Fasilitas.GetFasilitas(ctx)

	if err != nil {
		return database.Fasilita{}, err
	}

	return Fasilitas, nil

}

func (S *Service) Updatefasilitas(ctx context.Context, nama string) (database.Fasilita, error) {
	if nama == " " {
		return database.Fasilita{}, errors.New("kolom nama tidak boleh kosong")
	}
	fasilitas, err := S.Store.Fasilitas.Updatefasilitas(ctx, database.UpdatefasilitasParams{
		Nama: nama,
	})

	if err != nil {
		return database.Fasilita{}, err
	}

	return fasilitas, nil
}

func (S *Service) DeleteFasilitas(ctx context.Context, r *http.Request) error {
	FasilitasIDstr := chi.URLParam(r, "id_fasilitas")

	if FasilitasIDstr == " " {
		return errors.New("gagagl mendapatkan ID fasilitas")
	}
	FasilitasID, errs := uuid.Parse(FasilitasIDstr)

	if errs != nil {
		return errs
	}

	err := S.Store.Room.DeleteRoom(ctx, FasilitasID)

	if err != nil {
		return err
	}
	return nil
}
