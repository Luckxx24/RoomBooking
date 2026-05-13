package store

import (
	"database/sql"

	"github.com/luckxx24/RoomBooking/database"
)

type Storage struct {
	Users         Users
	Room          Room
	Fasilitas     Fasilitas
	Booking       Booking
	RoomFasilitas RoomFasilitas
	DB            *sql.DB
}

func NewStorage(DB *sql.DB) Storage {
	q := database.New(DB)

	return Storage{
		Users:         &users{q: q},
		Room:          &room{q: q},
		Fasilitas:     &fasilitas{q: q},
		Booking:       &booking{q: q},
		RoomFasilitas: &roomfasilitas{q: q},
		DB:            DB,
	}
}

func NewTX(tx database.DBTX) Storage {
	q := database.New(tx)

	return Storage{
		Users:         &users{q: q},
		Room:          &room{q: q},
		Fasilitas:     &fasilitas{q: q},
		Booking:       &booking{q: q},
		RoomFasilitas: &roomfasilitas{q: q},
	}
}
