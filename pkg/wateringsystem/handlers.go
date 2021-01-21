package wateringsystem

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"wateringsystem/pkg/wateringsystem/history"
)

// GetStatus gets the status for a given zone
func (ws *WateringSystem) GetStatus(w http.ResponseWriter, r *http.Request) {
	var statusResponses []statusResponse

	statusResponses = append(statusResponses, statusResponse{
		Zone: "Backyard",
		// Status: ws.backyardPin.ReadPin(),
		Status: "High",
	},
		statusResponse{
			Zone: "Frontyard",
			// Status: ws.frontyardPin.ReadPin(),
			Status: "Low",
		})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(statusResponses)
}

// TurnOnWateringSystem turns the watering system on for a given zone
func (ws *WateringSystem) TurnOnWateringSystem(w http.ResponseWriter, r *http.Request) {
	var command waterRequest
	json.NewDecoder(r.Body).Decode(&command)
	if command.Zone == "frontyard" && command.State == "on" {
		ws.frontyardPin.TurnOn()
	}
	if command.Zone == "frontyard" && command.State == "off" {
		ws.frontyardPin.TurnOff()
	}
	if command.Zone == "backyard" && command.State == "on" {
		ws.backyardPin.TurnOn()
	}
	if command.Zone == "backyard" && command.State == "off" {
		ws.backyardPin.TurnOff()
	}
	log.Println(ws.frontyardPin.ReadPin())
	log.Println(ws.backyardPin.ReadPin())
}

// TimedWateringSystem is the handler to turn the watering system for a given amount of time
func (ws *WateringSystem) TimedWateringSystem(w http.ResponseWriter, r *http.Request) {
	var command timedWaterRequest
	json.NewDecoder(r.Body).Decode(&command)
	log.Printf("Turning on %s\n", command.Zone)
	fmt.Fprintf(w, "%s is turned on for: %ds", command.Zone, command.TimeInSeconds)

	if command.Zone == "frontyard" && command.State == "on" {
		go func() {
			ws.frontyardPin.TurnOn()
			time.Sleep(time.Second * time.Duration(command.TimeInSeconds))
			ws.frontyardPin.TurnOff()
			log.Printf("Turning off %s", command.Zone)
		}()
	}
	if command.Zone == "backyard" && command.State == "on" {
		go func() {
			log.Println(command.TimeInSeconds)
			ws.backyardPin.TurnOn()
			time.Sleep(time.Second * time.Duration(command.TimeInSeconds))
			ws.backyardPin.TurnOff()
			log.Printf("Turning off %s", command.Zone)
		}()
	}
}

//GetHistory returns the entire history of our watering system
func (ws *WateringSystem) GetHistory(w http.ResponseWriter, r *http.Request) {
	wateringHistory, err := ws.historyRepo.All()
	var histories []history.History

	for _, h := range wateringHistory {
		histories = append(histories, history.History{
			ID:                   h.ID,
			Startdate:            h.Startdate,
			Enddate:              h.Enddate,
			Area:                 h.Area,
			TimeWateredInSeconds: time.Duration(h.Enddate.Sub(h.Startdate).Seconds()),
		})
	}

	if err != nil {
		http.Error(w, "No history", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(histories)
}
