package pi

import (
	"log"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/host"
)

//BackyardPin correponds to the backyard pin
const BackyardPin = "GPIO13"

//FrontyardPin coressponds to the frontyard gpio
const FrontyardPin = "GPIO6"

//PinWrapper wraps pins to expose utility methods
type PinWrapper struct {
	pin gpio.PinIO
}

//ReadPin reads the given pin and returns the value as a string
func (pinWrapper PinWrapper) ReadPin() string {
	return pinWrapper.pin.Read().String()
}

//TurnOn turns the pin on Low
func (pinWrapper PinWrapper) TurnOn() {
	pinWrapper.pin.Out(gpio.Low)
}

//TurnOff turns the pin to High
func (pinWrapper PinWrapper) TurnOff() {
	pinWrapper.pin.Out(gpio.High)
}

//InitRPI initiates the raspberry pi
func InitRPI() (PinWrapper, PinWrapper) {
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}
	backyardPin := PinWrapper{pin: gpioreg.ByName(BackyardPin)}
	frontyardPin := PinWrapper{pin: gpioreg.ByName(FrontyardPin)}

	return backyardPin, frontyardPin
}
