package main

import "github.com/kidoman/embd"

type Relay struct {
	Pin_GPIO   uint8 `json:"pin_gpio"`
	State      uint8 `json:"state"`
	DigitalPin embd.DigitalPin
}
type Environment struct {
	Relays [8]Relay `json:"relays"`
}
