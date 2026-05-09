package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/luckxx24/RoomBooking/service"
	"github.com/luckxx24/RoomBooking/store"
)

type dbconfig struct {
	Addr        string
	Maxopencons int
	Maxidlecons int
	Maxidletime string
}

type config struct {
	Addr     string
	DBconfig dbconfig
}

type Application struct {
	Config  config
	Store   store.Storage
	Service service.Service
}

func (app *Application) Mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http//*", "https//*"},
		AllowedMethods:   []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	r.Use(middleware.Timeout(60 * time.Second))

	return r
}

func (app *Application) run(Mux http.Handler) error {
	srv := &http.Server{
		Addr:         app.Config.Addr,
		Handler:      Mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server berjalan di %v", srv)

	return srv.ListenAndServe()
}
