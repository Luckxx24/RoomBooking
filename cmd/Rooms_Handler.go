package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/luckxx24/RoomBooking/cmd/jsonresponse"
	"github.com/luckxx24/RoomBooking/database"
)

type params struct {
	nama         string
	kapasitas    int32
	priceperhour string
	description  string
}

func (app Application) HandlerCreateRoom(w http.ResponseWriter, r http.Request) {

	decode := json.NewDecoder(r.Body)
	param := &params{}
	errs := decode.Decode(&param)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mendecode dari database %v", errs))
		return
	}

	Room, err := app.Store.Room.Createroom(r.Context(), database.CreateroomParams{
		ID:           uuid.New(),
		Nama:         param.nama,
		Kapasitas:    param.kapasitas,
		PricePerHour: param.priceperhour,
		Description:  param.description,
		CreatedAt:    time.Now(),
	})

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menyimpan user %v", err))
		return
	}

	jsonresponse.Succes(w, 201, Room)
}

func (app Application) HandlerGetRoom(w http.ResponseWriter, r *http.Request) {

	page := chi.URLParam(r, "page")
	pagesize := chi.URLParam(r, "pagesize")

	if page == "" || pagesize == "" {
		jsonresponse.RespondWithNotfound(w, fmt.Sprintf("tidak menemukan page dan pagsize di url"))
		return
	}

	Page, errs := strconv.Atoi(page)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal merubah page ke int"))
		return
	}

	Pagesize, errs := strconv.Atoi(page)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal merubah page ke int"))
		return
	}

	User, err := app.Service.GetRooms(r.Context(), Page, Pagesize)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("error ketika mendapatkan user %v", err))
		return
	}

	jsonresponse.Succes(w, 201, User)
}

func (app Application) HandleUpdateRoom(w http.ResponseWriter, r *http.Request) {
	decode := json.NewDecoder(r.Body)
	param := &params{}
	errs := decode.Decode(&param)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mendecode dari database %v", errs))
		return
	}

	RoomIDstr := chi.URLParam(r, "id_room")

	if RoomIDstr == "" {
		jsonresponse.RespondWithNotfound(w, "room id tidak ditemukan di url")
		return
	}

	RoomID, errs := uuid.Parse(RoomIDstr)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal men-parse id room %v", errs))
		return
	}

	Room, err := app.Store.Room.UpdateRoom(r.Context(), database.UpdateRoomParams{
		Nama:         param.nama,
		Kapasitas:    param.kapasitas,
		PricePerHour: param.priceperhour,
		Description:  param.description,
		UpdatedAt:    time.Now(),
		ID:           RoomID,
	},
	)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menyimpan user %v", err))
		return
	}

	jsonresponse.Succes(w, 201, Room)
}

func (app Application) HandlerDeleteRoom(w http.ResponseWriter, r *http.Request) {
	err := app.Service.DeletRooms(r.Context(), r)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menghapuse Rooms %v", err))
		return
	}

}
