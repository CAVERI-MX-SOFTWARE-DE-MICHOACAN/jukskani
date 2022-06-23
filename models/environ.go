package models

type Environ struct {
	Relays    []Relay    `json:"Relays"`
	SensorDHT *SensorDHT `json:"SensorDHT"`
}
