package models

import "github.com/kidoman/embd"

type Relay struct {
	Pin_GPIO int `json:"pin_gpio"`
	State    int `json:"state"`
}

func (R *Relay) Write(state int) {

	pin, err := embd.NewDigitalPin(R.Pin_GPIO)
	if err != nil {
		panic(err)
	}

	pin.SetDirection(embd.Out)
	pin.Write(state)

	R.State = state

}
