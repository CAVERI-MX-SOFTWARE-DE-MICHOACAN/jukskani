package models

import (
	"fmt"

	"caveri.mx/jukskani/src/dht22"
)

type SensorDHT struct {
	Pin_GPIO              string
	Temperature, Humidity float32
	Sensor                *dht22.DHT22
}

func (DHT *SensorDHT) Init(pin int) {
	DHT.Pin_GPIO = fmt.Sprintf("GPIO_%d", pin)
	DHT.Sensor = dht22.New(DHT.Pin_GPIO)
}
func (DHT *SensorDHT) Read() error {
	humidity, err := DHT.Sensor.Humidity()
	if err != nil {
		return err
	}
	temperature, err := DHT.Sensor.Temperature()
	if err != nil {
		return err
	}
	DHT.Humidity = humidity
	DHT.Temperature = temperature
	return nil
}
