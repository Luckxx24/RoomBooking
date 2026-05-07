package jsonresponse

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func RespondError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		fmt.Printf("errors %v", msg)
	}
	type Error struct {
		Error string `json:"Error"`
	}

	WriteJson(w, code, Error{
		Error: msg,
	})
}

func Succes(w http.ResponseWriter, code int, payload interface{}) {
	err := WriteJson(w, code, payload)

	if err != nil {
		fmt.Printf("errors")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{internal server error 500}"))
	}
}

func WriteJson(w http.ResponseWriter, code int, payload interface{}) error {
	data, errs := json.Marshal(payload)

	if errs != nil {
		return errs
	}

	w.Header().Add("content-type", "application/json")
	w.Write(data)

	return nil
}

func RespondWithBadRequest(w http.ResponseWriter, msg string) {
	RespondError(w, 400, msg)
}

func RespondWithUnauthorized(w http.ResponseWriter, msg string) {
	RespondError(w, 401, msg)
}

func RespondWithNotfound(w http.ResponseWriter, msg string) {
	RespondError(w, 404, msg)
}

func RespondWithForbiden(w http.ResponseWriter, msg string) {
	RespondError(w, 403, msg)
}

func RespondWithConflict(w http.ResponseWriter, msg string) {
	RespondError(w, 409, msg)
}
