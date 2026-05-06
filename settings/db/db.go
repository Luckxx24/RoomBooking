package db

import (
	"context"
	"database/sql"
	"time"
)

func SettingsDB(addr string, maxopencons, maxidlecons int, maxidletime string) (*sql.DB, error) {

	db, err := sql.Open("postgres", addr)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxopencons)
	db.SetMaxIdleConns(maxidlecons)

	Idletime, errs := time.ParseDuration(maxidletime)

	if errs != nil {
		return nil, errs
	}

	db.SetConnMaxIdleTime(Idletime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel()

	erro := db.PingContext(ctx)

	if erro != nil {
		return nil, erro
	}

	return db, nil

}
