package main

import (
	"log"
	"net/http"
	"wateringrpio/pkg/handlers"
	"wateringrpio/pkg/pi"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/pkg/profile"
)

type server struct {
	router *chi.Mux
}

func initServer() server {
	server := server{router: chi.NewRouter()}
	server.router.Use(middleware.Logger)
	server.router.Use(middleware.RealIP)
	log.Println("Server running...")
	return server
}

func main() {
	defer profile.Start().Stop()
	backyardPin, frontyardPin := pi.InitRPI()
	server := initServer()
	server.router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.TurnOnWateringSystem(backyardPin, frontyardPin, w, r)
	})
	server.router.Post("/timed", func(w http.ResponseWriter, r *http.Request) {
		handlers.TimedWateringSystem(backyardPin, frontyardPin, w, r)
	})
	server.router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetStatus(backyardPin, frontyardPin, w, r)
	})
	http.ListenAndServe(":3000", server.router)
}
