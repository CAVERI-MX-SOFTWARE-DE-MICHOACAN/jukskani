package models

import "github.com/kidoman/embd"

type Relay struct {
	Pin_GPIO   int `json:"pin_gpio"`
	State      int `json:"state"`
	DigitalPin embd.DigitalPin
}

func (R *Relay) Init(state int) {
	pin, err := embd.NewDigitalPin(R.Pin_GPIO)
	pin.SetDirection(embd.Out)
	if err != nil {
		panic(err)
	}
	R.State = state
	R.DigitalPin = pin
	R.DigitalPin.Write(R.State)
}

func (R *Relay) Write(state int) {
	R.State = state
	R.DigitalPin.Write(state)
}
