package wateringsystem

import (
	"wateringsystem/pkg/pi"

	"github.com/jackc/pgx"
)

//waterRequest to turn on the water
//Zone can either be frontyard or backyard
//State can either be on or off
type waterRequest struct {
	Zone  string `json:"zone,omitempty"`
	State string `json:"state,omitempty"`
}

//timedWaterRequest is the same as command but with a set amount of time
type timedWaterRequest struct {
	Zone          string `json:"zone,omitempty"`
	State         string `json:"state,omitempty"`
	TimeInSeconds int64  `json:"time_in_seconds,omitempty"`
}

//statusRequest returns the watering status for the zone
type statusResponse struct {
	Zone   string `json:"zone,omitempty"`
	Status string `json:"status,omitempty"`
}

//WateringSystem needs to know about the different pins to work with
type WateringSystem struct {
	backyardPin  *pi.PinWrapper
	frontyardPin *pi.PinWrapper
	historyRepo  *HistoryService
}

//NewWateringSystem returns an instance of the watering sytem
func NewWateringSystem(backyardPin *pi.PinWrapper, frontyard *pi.PinWrapper, db *pgx.ConnPool) *WateringSystem {
	return &WateringSystem{
		backyardPin:  backyardPin,
		frontyardPin: frontyard,
		historyRepo:  NewHistoryService(db),
	}
}
