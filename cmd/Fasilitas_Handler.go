package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luckxx24/RoomBooking/cmd/jsonresponse"
)

type paramfasilitas struct {
	nama string
}

func (app Application) HandlerCreateFasilitas(w http.ResponseWriter, r http.Request) {

	decode := json.NewDecoder(r.Body)
	param := &paramfasilitas{}
	errs := decode.Decode(&param)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mendecode dari database %v", errs))
		return
	}

	fasilitas, err := app.Service.CreateFasilitas(r.Context(), param.nama)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menyimpan fasilitas %v", err))
		return
	}

	jsonresponse.Succes(w, 201, fasilitas)
}

func (app Application) HandlerGetFasilitas(w http.ResponseWriter, r *http.Request) {
	fasilitas, err := app.Service.GetAllFasilitas(r.Context())

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("error ketika mendapatkan user %v", err))
		return
	}

	jsonresponse.Succes(w, 201, fasilitas)
}

func (app Application) HandleUpdateFasilitas(w http.ResponseWriter, r *http.Request) {

	decode := json.NewDecoder(r.Body)
	param := &paramfasilitas{}
	errs := decode.Decode(&param)

	if errs != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal mendecode dari database %v", errs))
		return
	}

	Updatefasilitas, err := app.Service.Updatefasilitas(r.Context(), param.nama)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal meng-update user %v", Updatefasilitas))
		return
	}

	jsonresponse.Succes(w, 201, Updatefasilitas)
}

func (app Application) HandlerDeleteFasilitas(w http.ResponseWriter, r *http.Request) {
	err := app.Service.DeleteFasilitas(r.Context(), r)

	if err != nil {
		jsonresponse.RespondWithBadRequest(w, fmt.Sprintf("gagal menghapuse fasilitas %v", err))
		return
	}

}
