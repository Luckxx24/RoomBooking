package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luckxx24/RoomBooking/cmd/jsonresponse"
)

func (app Application) HandlerCreateBooking(w http.ResponseWriter, r http.Request) {
	type params struct {
		nama     string
		Password string
		email    string
		Role     string
	}

	decode := json.NewDecoder(r.Body)
	param := &params{}
	errs := decode.Decode(&param)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mendecode dari database %v", errs))
		return
	}

	User, err := app.Service.CreateUser(r.Context(), param.nama, param.email, param.Password, param.Role)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menyimpan user %v", err))
		return
	}

	jsonresponse.Succes(w, 201, User)
}

func (app Application) HandlerGetBooking(w http.ResponseWriter, r *http.Request) {
	User, err := app.Service.GetUser(r.Context())

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("error ketika mendapatkan user %v", err))
		return
	}

	jsonresponse.Succes(w, 201, User)
}

func (app Application) HandleUpdateBooking(w http.ResponseWriter, r *http.Request) {
	type params struct {
		nama     string
		Password string
		email    string
		Role     string
	}

	decode := json.NewDecoder(r.Body)
	param := &params{}
	errs := decode.Decode(&param)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mendecode dari database %v", errs))
		return
	}

	UpdateUser, err := app.Service.UpdateUser(r.Context(), param.nama, param.email, param.Password)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal meng-update user %v", UpdateUser))
		return
	}

	jsonresponse.Succes(w, 201, UpdateUser)
}

func (app Application) HandlerDeleteBooking(w http.ResponseWriter, r *http.Request) {
	err := app.Service.DeletUser(r.Context())

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menghapuse User %v", err))
		return
	}

}
