package main

import (
	"log"
	"wateringsystem/pkg/pi"
	"wateringsystem/pkg/wateringsystem"
)

func main() {
	backyardPin, frontyardPin := pi.InitRPI()

	db, err := wateringsystem.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}

	ws := wateringsystem.NewWateringSystem(backyardPin, frontyardPin, db)

	app := application{
		ws: ws,
	}

	app.RunServer()
}
