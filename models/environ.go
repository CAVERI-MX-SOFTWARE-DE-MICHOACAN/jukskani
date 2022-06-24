package models

type Environ struct {
	Relays         []Relay          `json:"Relays"`
	SensorDHT      *SensorDHT       `json:"SensorDHT"`
	RelayCronTasks []RelayCronTasks `json:"RelayCronTasks"`
}
type RelayCronTasks struct {
	CronSpec   string `json:"CronSpec"`
	RelayIndex int    `json:"RelayIndex"`
	State      bool   `json:"State"`
}
