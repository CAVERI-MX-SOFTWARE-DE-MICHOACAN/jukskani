package models

import (
	"fmt"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
)

type Relay struct {
	PinName string `json:"PinName"`
	State   bool   `json:"state"`
}

func (r *Relay) Write(state bool) {

	pin := gpioreg.ByName(r.PinName)

	if pin == nil {
		panic(fmt.Sprintf("pin %s is nill", r.PinName))
	}
	if state {
		pin.Out(gpio.High)
	} else {
		pin.Out(gpio.Low)
	}
	r.State = state

}
