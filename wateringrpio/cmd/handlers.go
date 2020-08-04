package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"watering-system-go/wateringrpio"
)

type Command struct {
	Zone  string
	State string
}

type TimedCommand struct {
	Zone          string
	State         string
	TimeInSeconds int64
}

type StatusCommand struct {
	Zone   string
	Status string
}

func getStatus(backyardPin wateringrpio.PinWrapper, frontyardPin wateringrpio.PinWrapper, w http.ResponseWriter, r *http.Request) {
	var statusCommand StatusCommand

	zone := r.URL.Query().Get("zone")

	if zone == "backyard" {
		statusCommand.Status = backyardPin.ReadPin()
	}
	if zone == "frontyard" {
		statusCommand.Status = frontyardPin.ReadPin()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statusCommand)
}

func turnOnWateringSystem(backyardPin wateringrpio.PinWrapper, frontyardPin wateringrpio.PinWrapper, w http.ResponseWriter, r *http.Request) {
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

func timedWateringSystem(backyardPin wateringrpio.PinWrapper, frontyardPin wateringrpio.PinWrapper, w http.ResponseWriter, r *http.Request) {
	var command TimedCommand
	json.NewDecoder(r.Body).Decode(&command)
	log.Printf("Turning on %s\n", command.Zone)
	if command.Zone == "frontyard" && command.State == "on" {
		go func() {
			frontyardPin.TurnOn()
			time.Sleep(time.Second * time.Duration(command.TimeInSeconds))
			frontyardPin.TurnOff()
			log.Printf("Turning off %s", command.Zone)
		}()
	}
	if command.Zone == "backyard" && command.State == "on" {
		go func() {
			log.Println(command.TimeInSeconds)
			backyardPin.TurnOn()
			time.Sleep(time.Second * time.Duration(command.TimeInSeconds))
			backyardPin.TurnOff()
			log.Printf("Turning off %s", command.Zone)
		}()
	}
}
