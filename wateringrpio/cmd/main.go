package main

import (
	"net/http"
	"watering-system-go/wateringrpio"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	router *chi.Mux
}

func initServer() Server {
	server := Server{router: chi.NewRouter()}

	server.router.Use(middleware.Logger)
	server.router.Use(middleware.RealIP)
	return server
}

func main() {
	backyardPin, frontyardPin := wateringrpio.InitRPI()
	server := initServer()

	server.router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		turnOnWateringSystem(backyardPin, frontyardPin, w, r)
	})
	server.router.Post("/timed", func(w http.ResponseWriter, r *http.Request) {
		timedWateringSystem(backyardPin, frontyardPin, w, r)
	})
	server.router.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		getStatus(backyardPin, frontyardPin, w, r)
	})
	http.ListenAndServe(":3000", server.router)
}
