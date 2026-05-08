package store

import "github.com/luckxx24/RoomBooking/database"

type storage struct {
	users          Users
	room           Room
	fasilitas      Fasilitas
	room_fasilitas room_fasilitas
	booking        Booking
}

func NewStorage(DB database.DBTX) storage {
	q := database.New(DB)

	return storage{
		users:          &users{q: q},
		room:           &room{q: q},
		fasilitas:      &fasilitas{q: q},
		room_fasilitas: &room_fasilitas{q: q},
		booking:        &booking{q: q},
	}
}
