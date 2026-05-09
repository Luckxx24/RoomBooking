package store

import "github.com/luckxx24/RoomBooking/database"

type Storage struct {
	Users     Users
	Room      Room
	Fasilitas Fasilitas
	Booking   Booking
}

func NewStorage(DB database.DBTX) Storage {
	q := database.New(DB)

	return Storage{
		Users:     &users{q: q},
		Room:      &room{q: q},
		Fasilitas: &fasilitas{q: q},
		Booking:   &booking{q: q},
	}
}
