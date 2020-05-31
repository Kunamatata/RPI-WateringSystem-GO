package main

import (
	"encoding/json"
	"log"
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

type Command struct {
	Zone  string
	State string
}

func TurnOnWateringSystem(backyardPin wateringrpio.PinWrapper, frontyardPin wateringrpio.PinWrapper, w http.ResponseWriter, r *http.Request) {
	var command Command
	json.NewDecoder(r.Body).Decode(&command)
	if command.Zone == "frontyard" && command.State == "on" {
		frontyardPin.TurnOn()
	}
	if command.Zone == "frontyard" && command.State == "off" {
		frontyardPin.TurnOff()
	}
	if command.Zone == "backyard" && command.State == "on" {
		backyardPin.TurnOn()
	}
	if command.Zone == "backyard" && command.State == "off" {
		backyardPin.TurnOff()
	}
	log.Println(frontyardPin.ReadPin())
	log.Println(backyardPin.ReadPin())
}

func main() {
	backyardPin, frontyardPin := wateringrpio.InitRPI()
	server := initServer()

	log.Println(backyardPin.ReadPin())

	server.router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		TurnOnWateringSystem(backyardPin, frontyardPin, w, r)
	})
	http.ListenAndServe(":3000", server.router)
}
