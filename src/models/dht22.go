package models

import (
	"fmt"

	"github.com/MichaelS11/go-dht"
)

type SensorDHT struct {
	PinName               string
	Temperature, Humidity float64
	Sensor                *dht.DHT
}

func (DHT *SensorDHT) Init(pin string) {
	DHT.PinName = pin
	err := dht.HostInit()
	if err != nil {
		fmt.Println("HostInit error:", err)
		return
	}

	dht, err := dht.NewDHT("GPIO17", dht.Celsius, "dht22")
	if err != nil {
		fmt.Println("NewDHT error:", err)
		return
	}
	DHT.Sensor = dht

}
func (DHT *SensorDHT) Read() error {
	humidity, temperature, err := DHT.Sensor.Read()
	if err != nil {
		return err
	}
	DHT.Humidity = humidity
	DHT.Temperature = temperature
	return nil
}
