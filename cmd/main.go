package main

import (
	"log"

	"github.com/luckxx24/RoomBooking/settings/db"
	"github.com/luckxx24/RoomBooking/settings/env"
)

func main() {
	config := config{
		Addr: env.GetString("ADDR", "X"),
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
}
