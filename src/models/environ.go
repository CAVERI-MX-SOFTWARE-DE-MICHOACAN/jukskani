package models

type Environ struct {
	Relays    []Relay `json:"relays"`
	SensorDHT *SensorDHT `json:SensorDHT`
}
