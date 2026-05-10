package main

import (
	"log"

	"github.com/luckxx24/RoomBooking/JWT/auth"
	"github.com/luckxx24/RoomBooking/service"
	"github.com/luckxx24/RoomBooking/settings/db"
	"github.com/luckxx24/RoomBooking/settings/env"
	"github.com/luckxx24/RoomBooking/store"
)

func main() {
	config := config{
		Addr: env.GetString("ADDR", "postgres://Lucky:108099@localhost:5433/BookingRoom?sslmode=disable"),
		DBconfig: dbconfig{
			Addr:        env.GetString("Addr", "8080"),
			Maxopencons: env.GetInt("opencons", 30),
			Maxidlecons: env.GetInt("idlecons", 30),
			Maxidletime: env.GetString("idletime", "1 minute"),
		},
	}

	db, err := db.SettingsDB(config.DBconfig.Addr, config.DBconfig.Maxopencons, config.DBconfig.Maxidlecons, config.DBconfig.Maxidletime)

	if err != nil {
		log.Fatal(err)
	}

	store := store.NewStorage(db)

	Service := service.Service{
		Store: store,
	}

	Token := auth.SecretKey(env.GetString("rahasia", "rahasia"))

	app := Application{
		Config:  config,
		Store:   store,
		Service: Service,
		Token:   Token,
	}

	mux := app.Mount()

	log.Fatal(app.run(mux))
}
